package storage

import (
	"encoding/json"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type PostRequest struct {
	client      *api.Client
	Storage     string
	StorageType Type
	Path        string
}

type PostResponse struct {
	Storage string `json:"storage"`
	Type    Type   `json:"type"`
	Config  string `json:"config,omitempty"`
}

func (p PostRequest) Do() (*PostResponse, error) {
	var s PostResponse
	apiURL := p.client.ApiURL
	apiURL.Path += basePath
	params := url.Values{}
	params.Add("storage", p.Storage)
	params.Add("type", string(p.StorageType))
	params.Add("path", p.Path)
	apiURL.RawQuery = params.Encode()
	resp, err := p.client.Post(apiURL)
	if err != nil {
		return nil, err
	}
	return &s, json.Unmarshal(resp, &s)
}
