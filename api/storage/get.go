package storage

import (
	"encoding/json"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type GetRequest struct {
	Client *api.Client

	StorageType Type
}

type GetResponse []GetResponseStorage
type GetResponseStorage struct {
	Content      ContentList `json:"content,omitempty"`
	Digest       string      `json:"digest,omitempty"`
	Path         string      `json:"path,omitempty"`
	PruneBackups string      `json:"prune-backups,omitempty"`
	Shared       int         `json:"shared,omitempty"`
	Storage      string      `json:"storage,omitempty"`
	Type         Type        `json:"type,omitempty"`
}

func (g GetRequest) Do() (*GetResponse, error) {
	var s GetResponse
	apiURL := *g.Client.ApiURL
	apiURL.Path += basePath
	if g.StorageType != "" {
		params := url.Values{}
		params.Add("type", string(g.StorageType))
		apiURL.RawQuery = params.Encode()
	}
	resp, err := g.Client.Get(&apiURL)
	if err != nil {
		return nil, err
	}
	return &s, json.Unmarshal(resp, &s)
}
