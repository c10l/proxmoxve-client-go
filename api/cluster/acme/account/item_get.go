package account

import (
	"encoding/json"

	"github.com/c10l/proxmoxve-client-go/api"
)

type ItemGetRequest struct {
	Client *api.Client

	Name string
}

type ItemGetResponse struct {
	Account   ItemGetResponseAccount `json:"account"`
	Directory string                 `json:"directory"`
	Location  string                 `json:"location"`
	TOS       string                 `json:"tos"`
}

type ItemGetResponseAccount struct {
	Contact   []string                  `json:"contact"`
	InitialIP string                    `json:"initialIp"`
	Status    string                    `json:"status"`
	CreatedAt string                    `json:"createdAt"`
	Key       ItemGetResponseAccountKey `json:"key"`
}

type ItemGetResponseAccountKey struct {
	E   string `json:"e"`
	N   string `json:"n"`
	Use string `json:"use"`
	Kty string `json:"kty"`
}

// GetItem satisfies the ItemGetter interface
// Not to be used directly. Use Get() instead.
func (g ItemGetRequest) GetItem() ([]byte, error) {
	var name string
	if g.Name == "" {
		name = "default"
	} else {
		name = g.Name
	}
	return g.Client.GetItem(g, basePath, name)
}

func (g ItemGetRequest) Get() (*ItemGetResponse, error) {
	item, err := g.GetItem()
	if err != nil {
		return nil, err
	}
	resp := new(ItemGetResponse)
	return resp, json.Unmarshal(item, resp)
}
