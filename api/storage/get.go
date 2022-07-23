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

func (g GetRequest) Get() ([]GetResponseStorage, error) {
	items, err := g.GetAll()
	if err != nil {
		return nil, err
	}
	var resp []GetResponseStorage
	return resp, json.Unmarshal(items, &resp)
}

// GetAll implements the Getter interface.
// Not to be used directly. Use Get() instead.
func (g GetRequest) GetAll() ([]byte, error) {
	return g.Client.GetAll(g, basePath)
}

// ParsePath implements the Getter interface.
// Not to be used directly. Use Get() instead.
func (g GetRequest) ParseParams(apiURL *url.URL) error {
	if g.Type != "" {
		params := url.Values{}
		params.Add("type", string(g.Type))
		apiURL.RawQuery = params.Encode()
	}
	return nil
}
