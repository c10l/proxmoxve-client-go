package storage

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type PostRequest struct {
	client        *api.Client
	Storage       string
	StorageType   Type
	Path          *string
	Content       *[]Content
	Nodes         *string
	Disable       *bool
	Shared        *bool
	Preallocation *PreAllocation
}

type PreAllocation string

const (
	PreAllocationOff       PreAllocation = "off"
	PreAllocationMetadata  PreAllocation = "metadata"
	PreAllocationFallocate PreAllocation = "fallocate"
	PreAllocationFull      PreAllocation = "full"
)

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
	params.Add("path", *p.Path)
	if p.Content != nil {
		params.Add("content", contentList(p.Content))
	}
	if p.Nodes != nil {
		params.Add("nodes", *p.Nodes)
	}
	if p.Disable != nil {
		params.Add("disable", fmt.Sprintf("%v", p.Disable))
	}
	if p.Shared != nil {
		params.Add("shared", fmt.Sprintf("%v", p.Shared))
	}
	if p.Preallocation != nil {
		params.Add("preallocation", string(*p.Preallocation))
	}
	apiURL.RawQuery = params.Encode()
	resp, err := p.client.Post(apiURL)
	if err != nil {
		return nil, err
	}
	return &s, json.Unmarshal(resp, &s)
}

func contentList(l *[]Content) string {
	contentList := ""
	for i, c := range *l {
		if i == len(*l) {
			contentList += string(c)
		} else {
			contentList += string(c) + ","
		}
	}
	return contentList
}
