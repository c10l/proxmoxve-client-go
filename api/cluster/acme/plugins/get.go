package plugins

import (
	"encoding/json"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type GetRequest struct {
	Client *api.Client

	Type string
}

type GetResponse []struct {
	Digest string `json:"digest"`
	Plugin string `json:"plugin"`
	Type   string `json:"type"`
}

func (g GetRequest) Get() (GetResponse, error) {
	items, err := g.GetAll()
	if err != nil {
		return nil, err
	}
	var s GetResponse
	return s, json.Unmarshal(items, &s)
}

func (g GetRequest) GetAll() ([]byte, error) {
	return g.Client.GetAll(g, basePath)
}

func (g GetRequest) ParseParams(apiURL *url.URL) error {
	return nil
}
