package account

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type PostRequest struct {
	Client    *api.Client
	Contact   string  `json:"contact"`
	Directory *string `json:"directory"`
	Name      *string `json:"name"`

	// TOSurl is the URL of CA TermsOfService - setting this indicates agreement.
	TOSurl *string `json:"tos_url"`
}

type PostResponse string

func (g PostRequest) Do() (*PostResponse, error) {
	if g.Client.UserPass == nil {
		return nil, fmt.Errorf("POST to /cluster/acme/account requires username=root@pam and password")
	}
	var s PostResponse
	apiURL := *g.Client.ApiURL
	apiURL.Path += basePath
	params := url.Values{}
	params.Add("contact", g.Contact)
	if g.Directory != nil {
		params.Add("directory", *g.Directory)
	}
	if g.Name != nil {
		params.Add("name", *g.Name)
	}
	if g.TOSurl != nil {
		params.Add("tos_url", *g.TOSurl)
	}
	apiURL.RawQuery = params.Encode()
	resp, err := g.Client.Post(&apiURL)
	if err != nil {
		return nil, err
	}
	return &s, json.Unmarshal(resp, &s)
}
