package storage

import (
	"encoding/json"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type PostRequest struct {
	Client *api.Client

	// Required fields
	Storage     string
	StorageType string

	// Optional fields
	Path          *string
	Content       *[]string
	Nodes         *[]string
	Disable       *bool
	Shared        *bool
	Preallocation *string
}

type PostResponse struct {
	Storage string `json:"storage"`
	Type    string `json:"type"`
	Config  string `json:"config,omitempty"`
}

func (p PostRequest) Do() (*PostResponse, error) {
	var s PostResponse
	apiURL := *p.Client.ApiURL
	apiURL.Path += basePath
	params := url.Values{}
	params.Add("storage", p.Storage)
	params.Add("type", string(p.StorageType))
	params.Add("path", *p.Path)
	if p.Content != nil {
		params.Add("content", listJoin(p.Content, ","))
	}
	if p.Nodes != nil {
		params.Add("nodes", listJoin(p.Nodes, ","))
	}
	if p.Disable != nil {
		params.Add("disable", boolToInt(*p.Disable))
	}
	if p.Shared != nil {
		params.Add("shared", boolToInt(*p.Shared))
	}
	if p.Preallocation != nil {
		params.Add("preallocation", string(*p.Preallocation))
	}
	apiURL.RawQuery = params.Encode()
	resp, err := p.Client.Post(&apiURL)
	if err != nil {
		return nil, err
	}
	return &s, json.Unmarshal(resp, &s)
}
