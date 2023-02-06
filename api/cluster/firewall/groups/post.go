package groups

import (
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type PostRequest struct {
	Client *api.Client

	// Required fields
	Group string

	// Optional fields
	Comment *string
	Digest  *string
	Rename  *string
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
	params.Add("group", p.Group)
	if p.Comment != nil {
		params.Add("comment", string(*p.Comment))
	}
	if p.Digest != nil {
		params.Add("digest", string(*p.Digest))
	}
	if p.Rename != nil {
		params.Add("rename", string(*p.Rename))
	}
	apiURL.RawQuery = params.Encode()
	return nil
}
