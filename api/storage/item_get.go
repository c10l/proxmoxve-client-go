package storage

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/c10l/proxmoxve-client-go/api"
)

type ItemGetRequest struct {
	Client *api.Client

	Storage string
}

type ItemGetResponse struct {
	Storage         string
	Type            string
	Content         []string
	Digest          string
	Nodes           []string
	Disable         bool
	Shared          bool
	Preallocation   string
	Path            string
	PruneBackups    string
	NFSMountOptions string
	NFSServer       string
	NFSExport       string
}

func (r *ItemGetResponse) UnmarshalJSON(b []byte) error {
	var helper struct {
		Storage         string          `json:"storage"`
		Type            string          `json:"type"`
		Content         json.RawMessage `json:"content"`
		Digest          string          `json:"digest"`
		Nodes           json.RawMessage `json:"nodes"`
		Disable         int             `json:"disable"`
		Shared          int             `json:"shared"`
		Preallocation   string          `json:"preallocation"`
		Path            string          `json:"path"`
		PruneBackups    string          `json:"prune-backups"`
		NFSMountOptions string          `json:"options"`
		NFSServer       string          `json:"server"`
		NFSExport       string          `json:"export"`
	}
	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}
	disable, err := strconv.ParseBool(fmt.Sprint(helper.Disable))
	if err != nil {
		return err
	}
	shared, err := strconv.ParseBool(fmt.Sprint(helper.Shared))
	if err != nil {
		return err
	}

	r.Storage = helper.Storage
	r.Type = helper.Type
	r.Content = rawListSplitAndSort(helper.Content)
	r.Digest = helper.Digest
	r.Disable = disable
	r.Shared = shared
	r.Preallocation = helper.Preallocation
	r.Path = helper.Path
	r.PruneBackups = helper.PruneBackups
	r.Nodes = rawListSplitAndSort(helper.Nodes)
	r.NFSMountOptions = helper.NFSMountOptions
	r.NFSServer = helper.NFSServer
	r.NFSExport = helper.NFSExport
	return nil
}

// GetItem satisfies the ItemGetter interface
// Not to be used directly. Use Get() instead.
func (g ItemGetRequest) GetItem() ([]byte, error) {
	return g.Client.GetItem(g, basePath, g.Storage)
}

func (g ItemGetRequest) Get() (*ItemGetResponse, error) {
	item, err := g.GetItem()
	if err != nil {
		return nil, err
	}
	resp := new(ItemGetResponse)
	return resp, json.Unmarshal(item, resp)
}
