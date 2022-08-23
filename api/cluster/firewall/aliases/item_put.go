package aliases

import (
	"fmt"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type ItemPutRequest struct {
	Client *api.Client
	Name   string
	CIDR   string

	// Optional arguments
	Comment *string
	Digest  *string
	Rename  *string
}

// PutItem satisfies the ItemPutter interface.
// Not to be used directly. Use Put() instead.
func (g ItemPutRequest) PutItem() ([]byte, error) {
	return g.Client.PutItem(g, basePath, g.Name)
}

func (g ItemPutRequest) Put() error {
	_, err := g.PutItem()
	return err
}

// ParseParams satisfies the ItemPutter interface.
// Not to be used directly. Use Put() instead.
func (g ItemPutRequest) ParseParams(apiURL *url.URL) error {
	params := url.Values{}
	params.Add("name", g.Name)
	params.Add("cidr", g.CIDR)
	if g.Comment != nil {
		params.Add("comment", *g.Comment)
	}
	if g.Digest != nil {
		params.Add("digest", *g.Digest)
	}
	if g.Rename != nil {
		params.Add("rename", *g.Rename)
	}
	if len(params) == 0 {
		return fmt.Errorf("no params")
	}
	apiURL.RawQuery = params.Encode()
	return nil
}
