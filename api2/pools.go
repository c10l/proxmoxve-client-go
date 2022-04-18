package api2

import (
	"fmt"
	"net/url"
	"strings"
)

type Pools []Pool

type Pool struct {
	PoolID  string `json:"poolid"`
	Comment string `json:"comment,omitempty"`
	Members []any  `json:"members,omitempty"`
}

const poolsBasePath = "/pools"

func (c *Client) GetPools() (*Pools, error) {
	var data *Pools
	url := *c.ApiURL
	url.Path += poolsBasePath
	return doGet(c, data, &url)
}

func (c *Client) PostPool(poolID, comment string) error {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	params := url.Values{}
	params.Add("poolid", poolID)
	params.Add("comment", comment)
	apiURL.RawQuery = params.Encode()
	_, err := doPost(c, new(Pools), &apiURL)
	return err
}

func (c *Client) GetPool(poolID string) (*Pool, error) {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	apiURL.Path += "/" + poolID
	pool := &Pool{PoolID: poolID}
	return doGet(c, pool, &apiURL)
}

func (c *Client) DeletePool(poolID string) error {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	apiURL.Path += "/" + poolID
	return doDelete(c, &apiURL)
}

func (c *Client) PutPool(poolID string, comment *string, storage, vms *[]string, delete bool) error {
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

	fmt.Println(apiURL.String())
	return doPut(c, &apiURL)
}
