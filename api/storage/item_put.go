package storage

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/c10l/proxmoxve-client-go/api"
)

type ItemPutRequest struct {
	Client *api.Client

	Storage       string
	Content       *[]string
	Nodes         *[]string
	Disable       *bool
	Shared        *bool
	Preallocation *string
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
		disableParam := "0"
		if *g.Disable {
			disableParam = "1"
		}
		params.Add("disable", disableParam)
	}
	if g.Shared != nil {
		params.Add("shared", strconv.FormatBool(*g.Shared))
	}
	if g.Preallocation != nil {
		params.Add("preallocation", string(*g.Preallocation))
	}
	apiURL.RawQuery = params.Encode()
	resp, err := g.Client.Put(&apiURL)
	if err != nil {
		return nil, err
	}
	return &r, json.Unmarshal(resp, &r)
}
