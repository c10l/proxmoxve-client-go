package api2

import (
	"fmt"
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
	url := fmt.Sprintf(c.BaseURL + poolsBasePath)
	return doGet(c, data, url)
}

func (c *Client) PostPool(poolID, comment string) error {
	url := fmt.Sprintf(c.BaseURL + poolsBasePath + "?poolid=" + poolID + "&comment=" + comment)
	_, err := doPost(c, new(Pools), url, nil)
	return err
}

func (c *Client) GetPool(poolID string) (*Pool, error) {
	url := fmt.Sprintf(c.BaseURL + poolsBasePath + "/" + poolID)
	pool := &Pool{PoolID: poolID}
	return doGet(c, pool, url)
}

func (c *Client) DeletePool(poolID string) error {
	url := fmt.Sprintf(c.BaseURL + poolsBasePath + "/" + poolID)
	return doDelete(c, url)
}
