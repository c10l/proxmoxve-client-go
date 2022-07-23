package account

import (
	"encoding/json"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type GetRequest struct {
	Client *api.Client
}

type GetResponseAccount struct {
	Name string `json:"name"`
}

func (g GetRequest) Get() ([]GetResponseAccount, error) {
	items, err := g.GetAll()
	if err != nil {
		return nil, err
	}
	var s []GetResponseAccount
	return s, json.Unmarshal(items, &s)
}

func (g GetRequest) GetAll() ([]byte, error) {
	return g.Client.GetAll(g, basePath)
}

func (g GetRequest) ParseParams(apiURL *url.URL) error {
	return nil
}
