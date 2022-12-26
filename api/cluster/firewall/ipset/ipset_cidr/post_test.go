package ipset_cidr

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/cluster/firewall/ipset"
	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/c10l/proxmoxve-client-go/helpers/types"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestPost(t *testing.T) {
	client := helpers.APITokenTestClient()
	ipSetReq := ipset.PostRequest{
		Client: client,
		Name:   testNamePrefix + rand.String(10),
	}
	err := ipSetReq.Post()
	assert.NoError(t, err)

	ipSetCIDRReq := PostRequest{
		Client:    client,
		IPSetName: ipSetReq.Name,
		CIDR:      "192.168.0.0/16",
		Comment:   helpers.PtrTo("foobar"),
		NoMatch:   helpers.PtrTo(types.PVEBool(true)),
	}
	err = ipSetCIDRReq.Post()
	assert.NoError(t, err)

	ipSetCIDR, err := GetRequest{Client: helpers.APITokenTestClient(), IPSetName: ipSetReq.Name}.Get()
	assert.NoError(t, err)
	assert.NotNil(t, ipSetCIDR)
	assert.Len(t, ipSetCIDR, 1)
	assert.Equal(t, ipSetCIDRReq.CIDR, ipSetCIDR[0].CIDR)
	assert.Equal(t, ipSetCIDRReq.Comment, ipSetCIDR[0].Comment)
	assert.Equal(t, ipSetCIDRReq.NoMatch, ipSetCIDR[0].NoMatch)
}
