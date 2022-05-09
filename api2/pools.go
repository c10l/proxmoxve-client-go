package api2

import (
	"encoding/json"
	"net/url"
)

const poolsBasePath = "/pools"

type PoolsList []Pool

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
