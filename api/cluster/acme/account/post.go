package account

import (
	"encoding/json"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type PostRequest struct {
	Client    *api.Client
	Name      string  `json:"name"`
	Contact   string  `json:"contact"`
	Directory *string `json:"directory"`

	// TOSurl is the URL of CA TermsOfService - setting this indicates agreement.
	TOSurl *string `json:"tos_url"`
}

type PostResponse string

func (g PostRequest) Post() (*PostResponse, error) {
	item, err := g.PostItem()
	if err != nil {
		return nil, err
	}
	resp := new(PostResponse)
	return resp, json.Unmarshal(item, resp)
}

// PostItem satisfies the ItemPutter interface.
// Not to be used directly. Use Post() instead.
func (g PostRequest) PostItem() ([]byte, error) {
	return g.Client.PostItem(g, basePath)
}

// ParseParams satisfies the ItemPutter interface.
// Not to be used directly. Use Post() instead.
func (g PostRequest) ParseParams(apiURL *url.URL) error {
	params := url.Values{}
	params.Add("name", g.Name)
	params.Add("contact", g.Contact)
	if g.Directory != nil {
		params.Add("directory", *g.Directory)
	}
	if g.TOSurl != nil {
		params.Add("tos_url", *g.TOSurl)
	}
	apiURL.RawQuery = params.Encode()
	return nil
}
