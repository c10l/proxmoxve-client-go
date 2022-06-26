package storage

import (
	"bytes"
	"encoding/json"
	"net/url"
	"sort"
	"strings"

	"github.com/c10l/proxmoxve-client-go/api"
)

type GetRequest struct {
	client      *api.Client
	storageType Type
}

type GetResponse []GetResponseStorage
type GetResponseStorage struct {
	Content      GetResponseContentList `json:"content,omitempty"`
	Digest       string                 `json:"digest,omitempty"`
	Path         string                 `json:"path,omitempty"`
	PruneBackups string                 `json:"prune-backups,omitempty"`
	Shared       int                    `json:"shared,omitempty"`
	Storage      string                 `json:"storage,omitempty"`
	Type         Type                   `json:"type,omitempty"`
}

type GetResponseContentList []Content

func (l *GetResponseContentList) UnmarshalJSON(b []byte) error {
	parts := strings.Split(string(bytes.Trim(b, `"`)), ",")
	sort.Strings(parts)
	for _, item := range parts {
		*l = append(*l, Content(item))
	}
	return nil
}

func (g GetRequest) Do() (*GetResponse, error) {
	var s GetResponse
	apiURL := g.client.ApiURL
	apiURL.Path += basePath
	if g.storageType != "" {
		params := url.Values{}
		params.Add("type", string(g.storageType))
		apiURL.RawQuery = params.Encode()
	}
	resp, err := g.client.Get(apiURL)
	if err != nil {
		return nil, err
	}
	return &s, json.Unmarshal(resp, &s)
}
