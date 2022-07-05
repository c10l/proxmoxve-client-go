package storage

import (
	"github.com/c10l/proxmoxve-client-go/api"
)

type ItemDeleteRequest struct {
	Client *api.Client

	Storage string
}

func (r ItemDeleteRequest) Delete() error {
	return r.Client.DeleteItem(r, basePath, r.Storage)
}
