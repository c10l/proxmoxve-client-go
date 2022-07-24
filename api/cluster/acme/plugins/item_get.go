package plugins

import (
	"encoding/json"

	"github.com/c10l/proxmoxve-client-go/api"
	"github.com/c10l/proxmoxve-client-go/helpers"
)

type ItemGetRequest struct {
	Client *api.Client
	ID     string
}

type ItemGetResponse struct {
	Digest          string          `json:"digest"`
	Plugin          string          `json:"plugin"`
	Type            string          `json:"type"`
	Disable         helpers.IntBool `json:"disable"`
	ValidationDelay int             `json:"validation-delay"`
	Nodes           string          `json:"nodes"`
}

// GetItem satisfies the ItemGetter interface
// Not to be used directly. Use Get() instead.
func (g ItemGetRequest) GetItem() ([]byte, error) {
	return g.Client.GetItem(g, basePath, g.ID)
}

func (g ItemGetRequest) Get() (*ItemGetResponse, error) {
	item, err := g.GetItem()
	if err != nil {
		return nil, err
	}
	resp := new(ItemGetResponse)
	return resp, json.Unmarshal(item, resp)
}