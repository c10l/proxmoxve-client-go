package account

import (
	"github.com/c10l/proxmoxve-client-go/api"
)

type ItemDeleteRequest struct {
	Client *api.Client

	Name string
}

func (r ItemDeleteRequest) Delete() error {
	return r.Client.DeleteItem(r, basePath, r.Name)
}
