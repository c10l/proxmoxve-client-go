package storage

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type ItemPutRequest struct {
	Client *api.Client

	Storage         string
	Content         *[]string
	Nodes           *[]string
	Disable         *bool
	Shared          *bool
	Preallocation   *string
	NFSMountOptions *string
}

type ItemPutResponse struct {
	Storage string `json:"storage"`
	Type    string `json:"type"`
	Config  string `json:"config,omitempty"`
}

func (g ItemPutRequest) Do() (*ItemPutResponse, error) {
	if g.Storage == "" {
		return nil, fmt.Errorf("storage is required")
	}

	var r ItemPutResponse
	apiURL := *g.Client.ApiURL
	apiURL.Path += basePath + "/" + g.Storage
	params := url.Values{}
	if g.Content != nil {
		params.Add("content", listJoin(g.Content, ","))
	}
	if g.Nodes != nil {
		params.Add("nodes", listJoin(g.Nodes, ","))
	}
	if g.Disable != nil {
		params.Add("disable", boolToInt(*g.Disable))
	}
	if g.Shared != nil {
		params.Add("shared", boolToInt(*g.Shared))
	}
	if g.Preallocation != nil {
		params.Add("preallocation", string(*g.Preallocation))
	}
	if g.NFSMountOptions != nil {
		params.Add("options", string(*g.NFSMountOptions))
	}
	apiURL.RawQuery = params.Encode()
	resp, err := g.Client.Put(&apiURL)
	if err != nil {
		return nil, err
	}
	return &r, json.Unmarshal(resp, &r)
}
