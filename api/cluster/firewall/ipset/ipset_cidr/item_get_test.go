package ipset_cidr

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/cluster/firewall/ipset"
	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/c10l/proxmoxve-client-go/helpers/types"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemGet(t *testing.T) {
	ipSetReq := ipset.PostRequest{
		Client: helpers.APITokenTestClient(),
		Name:   testNamePrefix + rand.String(10),
	}
	err := ipSetReq.Post()
	assert.NoError(t, err)

	req := PostRequest{
		Client:    helpers.APITokenTestClient(),
		IPSetName: ipSetReq.Name,
		CIDR:      "192.168.0.0/17",
		NoMatch:   helpers.PtrTo(types.PVEBool(true)),
	}
	err = req.Post()
	assert.NoError(t, err)

	ipSetCIDR, err := ItemGetRequest{Client: helpers.APITokenTestClient(), IPSetName: ipSetReq.Name, CIDR: req.CIDR}.Get()
	assert.NoError(t, err)
	assert.Equal(t, req.CIDR, ipSetCIDR.CIDR)
	assert.Equal(t, req.NoMatch, ipSetCIDR.NoMatch)
}
