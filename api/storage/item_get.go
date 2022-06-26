package storage

import (
	"encoding/json"

	"github.com/c10l/proxmoxve-client-go/api"
)

type ItemGetRequest struct {
	Client *api.Client

	Storage string
}

type ItemGetResponse struct {
	Storage       string        `json:"storage"`
	Type          Type          `json:"type"`
	Content       ContentList   `json:"content"`
	Digest        string        `json:"digest"`
	Nodes         string        `json:"nodes"`
	Disable       bool          `json:"disable"`
	Shared        bool          `json:"shared"`
	PreAllocation PreAllocation `json:"preallocation"`
	Path          string        `json:"path"`
	PruneBackups  string        `json:"prune-backups"`
}

func (g ItemGetRequest) Do() (*ItemGetResponse, error) {
	var r ItemGetResponse
	apiURL := g.Client.ApiURL
	apiURL.Path += basePath + "/" + g.Storage
	resp, err := g.Client.Get(apiURL)
	if err != nil {
		return nil, err
	}
	return &r, json.Unmarshal(resp, &r)
}
