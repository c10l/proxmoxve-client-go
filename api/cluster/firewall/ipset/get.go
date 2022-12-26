package ipset

import (
	"encoding/json"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type GetRequest struct {
	Client *api.Client
}

type GetResponseList []GetResponse

type GetResponse struct {
	Name    string  `json:"name"`
	Digest  string  `json:"digest"`
	Comment *string `json:"comment,omitempty"`
}

func (l *GetResponseList) FindByName(name string) *GetResponse {
	for _, i := range *l {
		if i.Name == name {
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
	return g.Client.GetAll(g, basePath)
}

func (g GetRequest) ParseParams(apiURL *url.URL) error {
	return nil
}
