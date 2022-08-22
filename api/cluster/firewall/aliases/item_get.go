package aliases

import (
	"encoding/json"

	"github.com/c10l/proxmoxve-client-go/api"
)

type ItemGetRequest struct {
	Client *api.Client
	Name   string
}

type ItemGetResponse struct {
	Name      string  `json:"name"`
	CIDR      string  `json:"cidr"`
	Comment   *string `json:"comment,omitempty"`
	IPVersion int     `json:"ipversion"`
}

// GetItem satisfies the ItemGetter interface
// Not to be used directly. Use Get() instead.
func (g ItemGetRequest) GetItem() ([]byte, error) {
	return g.Client.GetItem(g, basePath, g.Name)
}

func (g ItemGetRequest) Get() (*ItemGetResponse, error) {
	item, err := g.GetItem()
	if err != nil {
		return nil, err
	}
	resp := new(ItemGetResponse)
	return resp, json.Unmarshal(item, resp)
}
