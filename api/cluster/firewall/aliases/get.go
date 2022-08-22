package aliases

import (
	"encoding/json"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type GetRequest struct {
	Client *api.Client
}

type GetResponsePlugins struct {
	CIDR    string `json:"cidr"`
	Digest  string `json:"digest"`
	Name    string `json:"name"`
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
