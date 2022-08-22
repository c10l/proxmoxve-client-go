package aliases

import (
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type PostRequest struct {
	Client *api.Client

	// Required fields
	CIDR string
	Name string

	// Optional fields
	Comment *string
}

func (p PostRequest) Post() error {
	_, err := p.PostItem()
	return err
}

// PostItem satisfies the ItemPutter interface.
// Not to be used directly. Use Post() instead.
func (p PostRequest) PostItem() ([]byte, error) {
	return p.Client.PostItem(p, basePath)
}

// ParseParams satisfies the ItemPutter interface.
// Not to be used directly. Use Post() instead.
func (p PostRequest) ParseParams(apiURL *url.URL) error {
	params := url.Values{}
	params.Add("cidr", p.CIDR)
	params.Add("name", string(p.Name))
	if p.Comment != nil {
		params.Add("comment", string(*p.Comment))
	}
	apiURL.RawQuery = params.Encode()
	return nil
}
