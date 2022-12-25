package ipset_cidr

import "fmt"

func basePath(name string) string {
	return fmt.Sprintf("/cluster/firewall/ipset/%s", name)
}
