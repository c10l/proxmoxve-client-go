package api2

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	"github.com/c10l/proxmoxve-client-go/api2/node"
)

const poolsBasePath = "/pools"

type PoolsList []Pool

type Pool struct {
	PoolID  string
	Comment string       `json:"comment,omitempty"`
	Members *PoolMembers `json:"members,omitempty"`
}

type PoolMembers struct {
	Storage []PoolMemberStorage
	Qemu    []node.Qemu
	LXC     []node.LXC
}

type PoolMemberStorage struct {
	Content    StorageContentList `json:"content"`
	Disk       int                `json:"disk"`
	ID         string             `json:"id"`
	MaxDisk    int                `json:"maxdisk"`
	Node       string             `json:"node"`
	PluginType StorageType        `json:"plugintype"`
	Shared     int                `json:"shared"`
	Status     string             `json:"status"`
	Storage    string             `json:"storage"`
	Type       PoolMemberType     `json:"type"`
}

type PoolMemberType string

const (
	PoolMemberTypeStorage PoolMemberType = "storage"
	PoolMemberTypeQemu    PoolMemberType = "qemu"
	PoolMemberTypeLXC     PoolMemberType = "lxc"
)

func (pm *PoolMembers) UnmarshalJSON(b []byte) error {
	var l []json.RawMessage
	if err := json.Unmarshal(b, &l); err != nil {
		return err
	}
	for _, item := range l {
		type poolMember struct {
			Type string `json:"type"`
		}
		var m poolMember
		if err := json.Unmarshal(item, &m); err != nil {
			return err
		}
		switch m.Type {
		case "storage":
			var storage PoolMemberStorage
			if err := json.Unmarshal(item, &storage); err != nil {
				return err
			}
			pm.Storage = append(pm.Storage, storage)
		case "qemu":
			var qemu node.Qemu
			if err := json.Unmarshal(item, &qemu); err != nil {
				return err
			}
			pm.Qemu = append(pm.Qemu, qemu)
		case "lxc":
			var lxc node.LXC
			if err := json.Unmarshal(item, &lxc); err != nil {
				return err
			}
			pm.LXC = append(pm.LXC, lxc)
		}
	}
	return nil
}

// RetrievePoolList Retrieves a list of pools.
// **Note that the response DOES NOT include Members!**
func (c *Client) GetPoolsList() (PoolsList, error) {
	url := *c.ApiURL
	url.Path += poolsBasePath
	resp, err := doGet(c, &url)
	if err != nil {
		return nil, err
	}
	var data PoolsList
	err = json.Unmarshal(resp, &data)
	return data, err
}

func (c *Client) PostPool(poolID, comment string) error {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	params := url.Values{}
	params.Add("poolid", poolID)
	params.Add("comment", comment)
	apiURL.RawQuery = params.Encode()
	_, err := doPost(c, &apiURL) // POST /pools has no response
	return err
}

func (c *Client) GetPool(poolID string) (*Pool, error) {
	if poolID == "" {
		return nil, errors.New("poolID cannot be empty")
	}
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	apiURL.Path += "/" + poolID
	resp, err := doGet(c, &apiURL)
	if err != nil {
		return nil, err
	}
	data := Pool{PoolID: poolID}
	err = json.Unmarshal(resp, &data)
	return &data, err
}

func (c *Client) DeletePool(poolID string) error {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	apiURL.Path += "/" + poolID
	_, err := doDelete(c, &apiURL)
	return err
}

func (c *Client) PutPool(poolID string, comment *string, storageStorages, vmNames []string, delete bool) error {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	apiURL.Path += "/" + poolID

	params := url.Values{}
	params.Add("poolid", poolID)
	if comment != nil {
		params.Add("comment", *comment)
	}
	if len(storageStorages) > 0 {
		// TODO: Implement validation after we have RetrieveStorage()
		// for _, item := range storageStorages {
		// 	if _, err := (*c).RetrieveStorage(item); err != nil {
		// 		return err
		// 	}
		// }
		storages := strings.Join(storageStorages, ",")
		params.Add("storage", storages)
	}
	if len(vmNames) > 0 {
		// TODO: Implement validation after we have RetrieveVMs()
		// for _, item := range vmNames {
		// 	if _, err := (*c).RetrieveVM(item); err != nil {
		// 		return err
		// 	}
		// }
		vmNames := strings.Join(vmNames, ",")
		params.Add("vms", vmNames)
	}
	if len(storageStorages) > 0 || len(vmNames) > 0 {
		if delete {
			params.Add("delete", "1")
		}
	}
	apiURL.RawQuery = params.Encode()

	_, err := doPut(c, &apiURL)
	return err
}
