package ipset

import (
	"github.com/c10l/proxmoxve-client-go/api"
	"github.com/c10l/proxmoxve-client-go/helpers/types"
)

type ItemDeleteRequest struct {
	Client *api.Client
	Name   string
	Force  types.PVEBool
}

func (g ItemDeleteRequest) Delete() error {
	return g.Client.DeleteItem(g, basePath, g.Name, "")
}

func (g ItemDeleteRequest) ForceDelete() error {
	return g.Client.DeleteItem(g, basePath, g.Name, "", api.URLParam{Key: "force", Value: types.PVEBool(true)})
}
