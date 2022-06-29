package storage

import (
	"encoding/json"
	"fmt"

	"github.com/c10l/proxmoxve-client-go/api"
)

type ItemGetRequest struct {
	Client *api.Client

	Storage string
}

type ItemGetResponse struct {
	Storage       string
	Type          string
	Content       []string
	Digest        string
	Nodes         []string
	Disable       bool
	Shared        bool
	Preallocation string
	Path          string
	PruneBackups  string
}

func (r *ItemGetResponse) UnmarshalJSON(b []byte) error {
	var helper struct {
		Storage       string          `json:"storage"`
		Type          string          `json:"type"`
		Content       json.RawMessage `json:"content"`
		Digest        string          `json:"digest"`
		Nodes         json.RawMessage `json:"nodes"`
		Disable       bool            `json:"disable"`
		Shared        bool            `json:"shared"`
		Preallocation string          `json:"preallocation"`
		Path          string          `json:"path"`
		PruneBackups  string          `json:"prune-backups"`
	}
	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}
	r.Storage = helper.Storage
	r.Type = helper.Type
	r.Content = rawListSplitAndSort(helper.Content)
	r.Digest = helper.Digest
	r.Disable = helper.Disable
	r.Shared = helper.Shared
	r.Preallocation = helper.Preallocation
	r.Path = helper.Path
	r.PruneBackups = helper.PruneBackups
	r.Nodes = rawListSplitAndSort(helper.Nodes)
	return nil
}

func (g ItemGetRequest) Do() (*ItemGetResponse, error) {
	if g.Storage == "" {
		return nil, fmt.Errorf("storage is required")
	}

	var r ItemGetResponse
	apiURL := *g.Client.ApiURL
	apiURL.Path += basePath + "/" + g.Storage
	resp, err := g.Client.Get(&apiURL)
	if err != nil {
		return nil, err
	}
	return &r, json.Unmarshal(resp, &r)
}
