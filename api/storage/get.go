package storage

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/c10l/proxmoxve-client-go/api"
)

type GetRequest struct {
	Client *api.Client
	Type   string
}

type GetResponse []GetResponseStorage
type GetResponseStorage struct {
	Content      []string
	Digest       string
	Path         string
	PruneBackups string
	Shared       bool
	Storage      string
	Type         string
}

func (r *GetResponseStorage) UnmarshalJSON(b []byte) error {
	var helper struct {
		Content      json.RawMessage `json:"content"`
		Digest       string          `json:"digest"`
		Path         string          `json:"path"`
		PruneBackups string          `json:"prune-backups"`
		Shared       int             `json:"shared"`
		Storage      string          `json:"storage"`
		Type         string          `json:"type"`
	}
	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}
	shared, err := strconv.ParseBool(fmt.Sprint(helper.Shared))
	if err != nil {
		return err
	}

	r.Content = strings.Split(strings.Trim(string(helper.Content), `"`), ",")
	r.Digest = helper.Digest
	r.Path = helper.Path
	r.PruneBackups = helper.PruneBackups
	r.Shared = shared
	r.Storage = helper.Storage
	r.Type = helper.Type
	return nil
}

func (g GetRequest) Do() (*GetResponse, error) {
	var s GetResponse
	apiURL := *g.Client.ApiURL
	apiURL.Path += basePath
	if g.Type != "" {
		params := url.Values{}
		params.Add("type", string(g.Type))
		apiURL.RawQuery = params.Encode()
	}
	resp, err := g.Client.Get(&apiURL)
	if err != nil {
		return nil, err
	}
	return &s, json.Unmarshal(resp, &s)
}
