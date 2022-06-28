package storage

import (
	"fmt"

	"github.com/c10l/proxmoxve-client-go/api"
)

type ItemDeleteRequest struct {
	Client *api.Client

	Storage string
}

func (r ItemDeleteRequest) Do() error {
	if r.Storage == "" {
		return fmt.Errorf("storage is required")
	}

	apiURL := *r.Client.ApiURL
	apiURL.Path += basePath + "/" + r.Storage
	_, err := r.Client.Delete(&apiURL)
	return err
}
