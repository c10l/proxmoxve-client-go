package account

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type ItemPutRequest struct {
	Client *api.Client

	Name    string
	Contact string
}

type ItemPutResponse string

// PutItem satisfies the ItemPutter interface.
// Not to be used directly. Use Put() instead.
func (g ItemPutRequest) PutItem() ([]byte, error) {
	return g.Client.PutItem(g, basePath, g.Name)
}

func (g ItemPutRequest) Put() (*ItemPutResponse, error) {
	item, err := g.PutItem()
	if err != nil {
		return nil, err
	}
	resp := new(ItemPutResponse)
	return resp, json.Unmarshal(item, resp)
}

// ParseParams satisfies the ItemPutter interface.
// Not to be used directly. Use Put() instead.
func (g ItemPutRequest) ParseParams(apiURL *url.URL) error {
	params := url.Values{}
	if g.Contact != "" {
		params.Add("contact", g.Contact)
	}
	if len(params) == 0 {
		return fmt.Errorf("no params")
	}
	apiURL.RawQuery = params.Encode()
	return nil
}
