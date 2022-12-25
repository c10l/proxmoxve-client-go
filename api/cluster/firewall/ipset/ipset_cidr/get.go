package ipset_cidr

import (
	"encoding/json"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type GetRequest struct {
	Client    *api.Client
	IPSetName string
}

type GetResponse []*struct {
	CIDR    string  `json:"cidr"`
	Digest  string  `json:"digest"`
	Comment *string `json:"comment,omitempty"`
	NoMatch *string `json:"nomatch,omitempty"`
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
	return g.Client.GetAll(g, basePath(g.IPSetName))
}

func (g GetRequest) ParseParams(apiURL *url.URL) error {
	return nil
}
