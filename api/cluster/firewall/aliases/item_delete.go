package aliases

import "github.com/c10l/proxmoxve-client-go/api"

type ItemDeleteRequest struct {
	Client *api.Client
	Name   string
	Digest string
}

func (g ItemDeleteRequest) Delete() error {
	return g.Client.DeleteItem(g, basePath, g.Name, "")
}
