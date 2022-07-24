package plugins

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/c10l/proxmoxve-client-go/api"
	"github.com/c10l/proxmoxve-client-go/helpers/types"
)

type ItemPutRequest struct {
	Client *api.Client
	ID     string

	// Optional arguments
	API             *string
	Data            *string
	Delete          *string
	Digest          *string
	Disable         *types.PVEBool
	Nodes           *[]string
	ValidationDelay *int
}

// PutItem satisfies the ItemPutter interface.
// Not to be used directly. Use Put() instead.
func (g ItemPutRequest) PutItem() ([]byte, error) {
	return g.Client.PutItem(g, basePath, g.ID)
}

func (g ItemPutRequest) Put() error {
	_, err := g.PutItem()
	return err
}

// ParseParams satisfies the ItemPutter interface.
// Not to be used directly. Use Put() instead.
func (g ItemPutRequest) ParseParams(apiURL *url.URL) error {
	params := url.Values{}
	if g.API != nil {
		params.Add("api", *g.API)
	}
	if g.Data != nil {
		params.Add("data", *g.Data)
	}
	if g.Delete != nil {
		params.Add("delete", *g.Delete)
	}
	if g.Digest != nil {
		params.Add("digest", *g.Digest)
	}
	if g.Disable != nil {
		params.Add("disable", g.Disable.ToAPIRequestParam())
	}
	if g.Nodes != nil {
		params.Add("nodes", strings.Join(*g.Nodes, ","))
	}
	if g.ValidationDelay != nil {
		params.Add("validation-delay", fmt.Sprintf("%d", *g.ValidationDelay))
	}
	if len(params) == 0 {
		return fmt.Errorf("no params")
	}
	apiURL.RawQuery = params.Encode()
	return nil
}
