package refs

import (
	"encoding/json"

	"github.com/c10l/proxmoxve-client-go/api"
)

const basePath = "/cluster/firewall/refs"

type GetRequest struct {
	Client *api.Client
}

type GetResponse []GetResponseRef

type GetResponseRef struct {
	Name    string `json:"name"`
	Ref     string `json:"ref"`
	Type    string `json:"type"`
	Comment string `json:"comment,omitempty"`
}

type GetResponseRefType string

const (
	GetResponseRefTypeAlias GetResponseRefType = "alias"
	GetResponseRefTypeIPSet GetResponseRefType = "ipset"
)

func (g GetRequest) Do() (*GetResponse, error) {
	var v GetResponse
	apiURL := g.Client.APIurl
	apiURL.Path += basePath
	resp, err := g.Client.Get(&apiURL)
	if err != nil {
		return nil, err
	}
	return &v, json.Unmarshal(resp, &v)
}
