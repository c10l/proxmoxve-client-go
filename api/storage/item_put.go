package storage

import (
	"encoding/json"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type ItemPutRequest struct {
	client  *api.Client
	Storage string
	Content []Content
}

type ItemPutResponse struct {
	Storage string `json:"storage"`
	Type    Type   `json:"type"`
	Config  string `json:"config,omitempty"`
}

func (g ItemPutRequest) Do() (*ItemPutResponse, error) {
	var r ItemPutResponse
	apiURL := g.client.ApiURL
	apiURL.Path += basePath + "/" + g.Storage
	params := url.Values{}
	params.Add("content", contentList(&g.Content))
	apiURL.RawQuery = params.Encode()
	resp, err := g.client.Put(apiURL)
	if err != nil {
		return nil, err
	}
	return &r, json.Unmarshal(resp, &r)
}
