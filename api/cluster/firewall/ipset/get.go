package ipset

import (
	"encoding/json"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type GetRequest struct {
	Client *api.Client
}

type GetResponsePlugins struct {
	Name    string `json:"name"`
	Digest  string `json:"digest"`
	Comment string `json:"comment"`
}

func (g GetRequest) Get() ([]GetResponsePlugins, error) {
	items, err := g.GetAll()
	if err != nil {
		return nil, err
	}
	var s []GetResponsePlugins
	return s, json.Unmarshal(items, &s)
}

func (g GetRequest) GetAll() ([]byte, error) {
	return g.Client.GetAll(g, basePath)
}

func (g GetRequest) ParseParams(apiURL *url.URL) error {
	return nil
}
