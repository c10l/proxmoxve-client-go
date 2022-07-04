package account

import (
	"fmt"

	"github.com/c10l/proxmoxve-client-go/api"
)

type ItemDeleteRequest struct {
	Client *api.Client

	Name string
}

func (r ItemDeleteRequest) Do() error {
	if r.Name == "" {
		return fmt.Errorf("account is required")
	}

	apiURL := *r.Client.ApiURL
	apiURL.Path += basePath + "/" + r.Name
	_, err := r.Client.Delete(&apiURL)
	return err
}
