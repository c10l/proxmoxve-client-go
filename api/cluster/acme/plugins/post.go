package plugins

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/c10l/proxmoxve-client-go/api"
	"github.com/c10l/proxmoxve-client-go/helpers/types"
)

type PostRequest struct {
	Client *api.Client
	ID     string
	Type   string

	// Optional attributes
	API             *string
	Data            *string
	Disable         *types.PVEBool
	Nodes           *[]string
	ValidationDelay *int
}

func (g PostRequest) Post() error {
	_, err := g.PostItem()
	return err
}

// PostItem satisfies the ItemPoster interface.
// Not to be used directly. Use Post() instead.
func (g PostRequest) PostItem() ([]byte, error) {
	return g.Client.PostItem(g, basePath)
}

// ParseParams satisfies the ItemPutter interface.
// Not to be used directly. Use Post() instead.
func (g PostRequest) ParseParams(apiURL *url.URL) error {
	params := url.Values{}
	params.Add("id", g.ID)
	params.Add("type", g.Type)
	if g.API != nil {
		params.Add("api", *g.API)
	}
	if g.Data != nil {
		params.Add("data", *g.Data)
	}
	if g.Disable != nil {
		params.Add("disable", g.Disable.ToAPIRequestParam())
	}
	if g.Nodes != nil {
		params.Add("nodes", strings.Join(*g.Nodes, ","))
	}
	if g.ValidationDelay != nil {
		params.Add("validation_delay", fmt.Sprintf("%d", *g.ValidationDelay))
	}
	apiURL.RawQuery = params.Encode()
	return nil
}
