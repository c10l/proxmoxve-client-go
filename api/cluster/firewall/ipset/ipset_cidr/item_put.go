package ipset_cidr

import (
	"fmt"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
	"github.com/c10l/proxmoxve-client-go/helpers/types"
)

type ItemPutRequest struct {
	Client    *api.Client
	IPSetName string
	CIDR      string

	// Optional arguments
	Comment *string
	Digest  *string
	NoMatch *types.PVEBool
}

// PutItem satisfies the ItemPutter interface.
// Not to be used directly. Use Put() instead.
func (g ItemPutRequest) PutItem() ([]byte, error) {
	return g.Client.PutItem(g, basePath(g.IPSetName), g.CIDR)
}

func (g ItemPutRequest) Put() error {
	_, err := g.PutItem()
	return err
}

// ParseParams satisfies the ItemPutter interface.
// Not to be used directly. Use Put() instead.
func (g ItemPutRequest) ParseParams(apiURL *url.URL) error {
	params := url.Values{}
	params.Add("cidr", g.CIDR)
	if g.Comment != nil {
		params.Add("comment", *g.Comment)
	}
	if g.Digest != nil {
		params.Add("digest", *g.Digest)
	}
	if g.NoMatch != nil {
		params.Add("nomatch", g.NoMatch.ToAPIRequestParam())
	}
	if len(params) == 0 {
		return fmt.Errorf("no params")
	}
	apiURL.RawQuery = params.Encode()
	return nil
}
