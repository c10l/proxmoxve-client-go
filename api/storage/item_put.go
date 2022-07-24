package storage

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
	"github.com/c10l/proxmoxve-client-go/helpers"
)

type ItemPutRequest struct {
	Client *api.Client

	Storage         string
	Content         *[]string
	Nodes           *[]string
	Disable         *helpers.IntBool
	Shared          *helpers.IntBool
	Preallocation   *string
	NFSMountOptions *string
}

type ItemPutResponse struct {
	Storage string `json:"storage"`
	Type    string `json:"type"`
	Config  string `json:"config,omitempty"`
}

// PutItem satisfies the ItemPutter interface.
// Not to be used directly. Use Put() instead.
func (g ItemPutRequest) PutItem() ([]byte, error) {
	return g.Client.PutItem(g, basePath, g.Storage)
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
	if g.Content != nil {
		params.Add("content", stringSliceJoin(g.Content, ","))
	}
	if g.Nodes != nil {
		params.Add("nodes", stringSliceJoin(g.Nodes, ","))
	}
	if g.Disable != nil {
		params.Add("disable", fmt.Sprintf("%d", g.Disable.Int()))
	}
	if g.Shared != nil {
		params.Add("shared", fmt.Sprintf("%d", g.Shared.Int()))
	}
	if g.Preallocation != nil {
		params.Add("preallocation", string(*g.Preallocation))
	}
	if g.NFSMountOptions != nil {
		params.Add("options", string(*g.NFSMountOptions))
	}
	if len(params) == 0 {
		return fmt.Errorf("no params")
	}
	apiURL.RawQuery = params.Encode()
	return nil
}
