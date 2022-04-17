package api2

import (
	"fmt"
)

type Pools struct {
	Data []string
}

const poolsBasePath = "/pools"

func (c *Client) GetPools() (*Pools, error) {
	var data *Pools
	url := fmt.Sprintf(c.BaseURL + poolsBasePath)

	return doGet(c, data, url)
}
