package ipset_cidr

import (
	"encoding/json"

	"github.com/c10l/proxmoxve-client-go/api"
)

type ItemGetRequest struct {
	Client    *api.Client
	IPSetName string
	CIDR      string
}

type ItemGetResponse struct {
	CIDR    string  `json:"cidr"`
	Digest  string  `json:"digest"`
	Comment *string `json:"comment,omitempty"`
}

// GetItem satisfies the ItemGetter interface
// Not to be used directly. Use Get() instead.
func (g ItemGetRequest) GetItem() ([]byte, error) {
	return g.Client.GetItem(g, basePath(g.IPSetName), g.CIDR)
}

func (g ItemGetRequest) Get() (*ItemGetResponse, error) {
	item, err := g.GetItem()
	if err != nil {
		return nil, err
	}
	resp := new(ItemGetResponse)
	return resp, json.Unmarshal(item, resp)
}
