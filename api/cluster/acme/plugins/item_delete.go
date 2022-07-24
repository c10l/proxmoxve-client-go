package plugins

import "github.com/c10l/proxmoxve-client-go/api"

type ItemDeleteRequest struct {
	Client *api.Client
	ID     string
}

func (g ItemDeleteRequest) Delete() error {
	return g.Client.DeleteItem(g, basePath, g.ID)
}
