package storage

import (
	"encoding/json"
	"net/url"

	"github.com/c10l/proxmoxve-client-go/api"
)

type PostRequest struct {
	Client *api.Client

	// Required fields
	Storage     string
	StorageType string

	// Dir fields
	DirPath   *string
	DirShared *bool

	// NFS fields
	NFSMountOptions *string
	NFSServer       *string
	NFSExport       *string

	// Global optional fields
	Content       *[]string
	Nodes         *[]string
	Disable       *bool
	Preallocation *string
}

type PostResponse struct {
	Storage string `json:"storage"`
	Type    string `json:"type"`
	Config  string `json:"config,omitempty"`
}

func (p PostRequest) Post() (*PostResponse, error) {
	item, err := p.PostItem()
	if err != nil {
		return nil, err
	}
	resp := new(PostResponse)
	return resp, json.Unmarshal(item, resp)
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
	params.Add("storage", p.Storage)
	params.Add("type", string(p.StorageType))
	if p.DirPath != nil {
		params.Add("path", string(*p.DirPath))
	}
	if p.Content != nil {
		params.Add("content", stringSliceJoin(p.Content, ","))
	}
	if p.Nodes != nil {
		params.Add("nodes", stringSliceJoin(p.Nodes, ","))
	}
	if p.Disable != nil {
		params.Add("disable", boolToInt(*p.Disable))
	}
	if p.DirShared != nil {
		params.Add("shared", boolToInt(*p.DirShared))
	}
	if p.Preallocation != nil {
		params.Add("preallocation", string(*p.Preallocation))
	}
	if p.NFSMountOptions != nil {
		params.Add("options", string(*p.NFSMountOptions))
	}
	if p.NFSServer != nil {
		params.Add("server", string(*p.NFSServer))
	}
	if p.NFSExport != nil {
		params.Add("export", string(*p.NFSExport))
	}
	apiURL.RawQuery = params.Encode()
	return nil
}
