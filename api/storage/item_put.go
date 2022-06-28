package storage

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type ItemPutRequest struct {
	Client *api.Client

	Storage string
	Content []Content
}

type ItemPutResponse struct {
	Storage string `json:"storage"`
	Type    Type   `json:"type"`
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
	params.Add("content", contentList(&g.Content))
	apiURL.RawQuery = params.Encode()
	resp, err := g.Client.Put(&apiURL)
	if err != nil {
		return nil, err
	}
	return &r, json.Unmarshal(resp, &r)
}
