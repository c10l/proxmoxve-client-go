package ipset_cidr

import "github.com/c10l/proxmoxve-client-go/api"

type ItemDeleteRequest struct {
	Client    *api.Client
	IPSetName string
	CIDR      string
}

func (g ItemDeleteRequest) Delete() error {
	return g.Client.DeleteItem(g, basePath(g.IPSetName), g.CIDR, "")
}
