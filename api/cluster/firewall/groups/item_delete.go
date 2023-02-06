package groups

import "github.com/c10l/proxmoxve-client-go/api"

type ItemDeleteRequest struct {
	Client *api.Client
	Group  string
}

func (g ItemDeleteRequest) Delete() error {
	return g.Client.DeleteItem(g, basePath, g.Group, "")
}
