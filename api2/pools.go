package api2

import (
	"encoding/json"
	"io"
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

func (c *Client) RetrievePoolList() (io.Reader, error) {
	url := *c.ApiURL
	url.Path += poolsBasePath
	resp, err := doGet(c, &url)
	data := strings.NewReader(string(resp))
	return data, err
}

func (c *Client) CreatePool(poolID, comment string) (io.Reader, error) {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	params := url.Values{}
	params.Add("poolid", poolID)
	params.Add("comment", comment)
	apiURL.RawQuery = params.Encode()
	resp, err := doPost(c, &apiURL)
	data := strings.NewReader(string(resp))
	return data, err
}

func (c *Client) RetrievePool(poolID string) (io.Reader, error) {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	apiURL.Path += "/" + poolID
	resp, err := doGet(c, &apiURL)
	data := strings.NewReader(string(resp))
	return data, err
}

func (c *Client) DeletePool(poolID string) (io.Reader, error) {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	apiURL.Path += "/" + poolID
	resp, err := doDelete(c, &apiURL)
	data := strings.NewReader(string(resp))
	return data, err
}

func (c *Client) UpdatePool(poolID string, comment *string, storage, vms *[]string, delete bool) (io.Reader, error) {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	apiURL.Path += "/" + poolID

	params := url.Values{}
	if comment != nil {
		params.Add("comment", *comment)
	}
	if storage != nil {
		params.Add("storage", strings.Join(*storage, ","))
	}
	if vms != nil {
		params.Add("vms", strings.Join(*vms, ","))
	}
	if storage != nil || vms != nil {
		if delete {
			params.Add("delete", "1")
		}
	}
	apiURL.RawQuery = params.Encode()

	resp, err := doPut(c, &apiURL)
	data := strings.NewReader(string(resp))
	return data, err
}
