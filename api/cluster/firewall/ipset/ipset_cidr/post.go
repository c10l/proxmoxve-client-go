package ipset_cidr

import (
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
	"github.com/c10l/proxmoxve-client-go/helpers/types"
)

type PostRequest struct {
	Client *api.Client

	// Required fields
	IPSetName string
	CIDR      string

	// Optional fields
	Comment *string
	NoMatch *types.PVEBool
}

func (p PostRequest) Post() error {
	_, err := p.PostItem()
	return err
}

// PostItem satisfies the ItemPutter interface.
// Not to be used directly. Use Post() instead.
func (p PostRequest) PostItem() ([]byte, error) {
	return p.Client.PostItem(p, basePath(p.IPSetName))
}

// ParseParams satisfies the ItemPutter interface.
// Not to be used directly. Use Post() instead.
func (p PostRequest) ParseParams(apiURL *url.URL) error {
	params := url.Values{}
	params.Add("cidr", string(p.CIDR))
	if p.Comment != nil {
		params.Add("comment", string(*p.Comment))
	}
	if p.NoMatch != nil {
		params.Add("nomatch", p.NoMatch.ToAPIRequestParam())
	}
	apiURL.RawQuery = params.Encode()
	return nil
}
