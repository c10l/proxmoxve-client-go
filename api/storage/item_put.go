package storage

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
	"github.com/c10l/proxmoxve-client-go/helpers/types"
)

type ItemPutRequest struct {
	Client *api.Client

	Storage         string
	Content         *[]string
	Nodes           *[]string
	Disable         *types.PVEBool
	Shared          *types.PVEBool
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
		params.Add("disable", g.Disable.ToAPIRequestParam())
	}
	if g.Shared != nil {
		params.Add("shared", g.Shared.ToAPIRequestParam())
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
