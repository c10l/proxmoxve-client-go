package api2

import (
	"fmt"
)

type Pools struct {
	Data []Pool `json:"data"`
}

type Pool struct {
	PoolID  string `json:"poolid"`
	Comment string `json:"comment,omitempty"`
}

const poolsBasePath = "/pools"

func (c *Client) GetPools() (*Pools, error) {
	var data *Pools
	url := fmt.Sprintf(c.BaseURL + poolsBasePath)

	return doGet(c, data, url)
}

func (c *Client) PostPools(poolID, comment string) error {
	url := fmt.Sprintf(c.BaseURL + poolsBasePath + "?poolid=" + poolID + "&comment=" + comment)
	fmt.Println(url)
	_, err := doPost(c, new(Pools), url, nil)
	return err
}
