package storage

import "github.com/c10l/proxmoxve-client-go/api"

type ItemDeleteRequest struct {
	client  *api.Client
	Storage string
}

func (r ItemDeleteRequest) Do() error {
	apiURL := r.client.ApiURL
	apiURL.Path += basePath + "/" + r.Storage
	_, err := r.client.Delete(apiURL)
	return err
}
