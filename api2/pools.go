package api2

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/c10l/proxmoxve-client-go/api2/node"
)

const poolsBasePath = "/pools"

type PoolMemberType string

const (
	PoolMemberTypeStorage PoolMemberType = "storage"
	PoolMemberTypeQemu    PoolMemberType = "qemu"
	PoolMemberTypeLXC     PoolMemberType = "lxc"
)

type PoolMember struct {
	Type    PoolMemberType
	Storage *Storage
	Qemu    *node.Qemu
	LXC     *node.LXC
}

func (pm *PoolMember) UnmarshalJSON(b []byte) error {
	type poolMemberType struct {
		Type string `json:"type"`
	}
	var pmType poolMemberType
	if err := json.Unmarshal(b, &pmType); err != nil {
		return err
	}
	switch pmType.Type {
	case "storage":
		var storage Storage
		if err := json.Unmarshal(b, &storage); err != nil {
			return err
		}
		pm.Type = PoolMemberTypeStorage
		pm.Storage = &storage
	case "qemu":
		var qemu node.Qemu
		if err := json.Unmarshal(b, &qemu); err != nil {
			return err
		}
		pm.Type = PoolMemberTypeQemu
		pm.Qemu = &qemu
	case "lxc":
		var lxc node.LXC
		if err := json.Unmarshal(b, &lxc); err != nil {
			return err
		}
		pm.Type = PoolMemberTypeLXC
		pm.LXC = &lxc
	}
	return nil
}

type Pool struct {
	PoolID  string
	Comment string       `json:"comment,omitempty"`
	Members []PoolMember `json:"members"`
}

// func (c *Client) RetrievePoolList() (io.Reader, error) {
// 	url := *c.ApiURL
// 	url.Path += poolsBasePath
// 	resp, err := doGet(c, &url)
// 	data := strings.NewReader(string(resp))
// 	return data, err
// }

func (c *Client) CreatePool(poolID, comment string) error {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	params := url.Values{}
	params.Add("poolid", poolID)
	params.Add("comment", comment)
	apiURL.RawQuery = params.Encode()
	_, err := doPost(c, &apiURL) // POST /pool has no response
	return err
}

func (c *Client) RetrievePool(poolID string) (*Pool, error) {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	apiURL.Path += "/" + poolID
	resp, err := doGet(c, &apiURL)
	if err != nil {
		return nil, err
	}
	var data Pool
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

func (c *Client) UpdatePool(poolID string, comment *string, storageStorages, vmNames []string, delete bool) error {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	apiURL.Path += "/" + poolID

	params := url.Values{}
	params.Add("poolid", poolID)
	if comment != nil {
		params.Add("comment", *comment)
	}
	if len(storageStorages) > 0 {
		for _, item := range storageStorages {
			if _, err := (*c).RetrieveStorage(item); err != nil {
				return err
			}
		}
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
