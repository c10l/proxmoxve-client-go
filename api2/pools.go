package api2

import (
	"net/url"
	"strings"
)

type PoolList []Pool

type Pool struct {
	PoolID  string `json:"poolid"`
	Comment string `json:"comment,omitempty"`
	Members []any  `json:"members,omitempty"`
}

const poolsBasePath = "/pools"

func (c *Client) RetrievePoolList() (*PoolList, error) {
	data := new(PoolList)
	url := *c.ApiURL
	url.Path += poolsBasePath
	err := doGet(c, data, &url)
	return data, err
}

func (c *Client) CreatePool(poolID, comment string) error {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	params := url.Values{}
	params.Add("poolid", poolID)
	params.Add("comment", comment)
	apiURL.RawQuery = params.Encode()
	err := doPost(c, new(PoolList), &apiURL)
	return err
}

func (c *Client) RetrievePool(poolID string) (*Pool, error) {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	apiURL.Path += "/" + poolID
	data := &Pool{PoolID: poolID}
	err := doGet(c, data, &apiURL)
	return data, err
}

func (c *Client) DeletePool(poolID string) error {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	apiURL.Path += "/" + poolID
	return doDelete(c, &apiURL)
}

func (c *Client) UpdatePool(poolID string, comment *string, storage, vms *[]string, delete bool) error {
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

	return doPut(c, &apiURL)
}
