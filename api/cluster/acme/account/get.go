package account

import (
	"encoding/json"

	"github.com/c10l/proxmoxve-client-go/api"
)

type GetRequest struct {
	Client *api.Client
}

type GetResponse []GetResponseAccount
type GetResponseAccount struct {
	Name string `json:"name"`
}

func (g GetRequest) Do() (*GetResponse, error) {
	var s GetResponse
	apiURL := *g.Client.ApiURL
	apiURL.Path += basePath
	resp, err := g.Client.Get(&apiURL)
	if err != nil {
		return nil, err
	}
	return &s, json.Unmarshal(resp, &s)
}
