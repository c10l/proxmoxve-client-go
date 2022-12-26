package ipset_cidr

import (
	"encoding/json"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
	"github.com/c10l/proxmoxve-client-go/helpers/types"
)

type GetRequest struct {
	Client    *api.Client
	IPSetName string
}

type GetResponseList []GetResponse

type GetResponse struct {
	CIDR    string         `json:"cidr"`
	Digest  string         `json:"digest"`
	Comment *string        `json:"comment,omitempty"`
	NoMatch *types.PVEBool `json:"nomatch,omitempty"`
}

func (l *GetResponseList) FindByCIDR(cidr string) *GetResponse {
	for _, i := range *l {
		if i.CIDR == cidr {
			return &i
		}
	}
	return nil
}

func (g GetRequest) Get() (GetResponseList, error) {
	items, err := g.GetAll()
	if err != nil {
		return nil, err
	}
	var s GetResponseList
	return s, json.Unmarshal(items, &s)
}

func (g GetRequest) GetAll() ([]byte, error) {
	return g.Client.GetAll(g, basePath(g.IPSetName))
}

func (g GetRequest) ParseParams(apiURL *url.URL) error {
	return nil
}
