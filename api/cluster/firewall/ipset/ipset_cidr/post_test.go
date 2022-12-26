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
	ipSetName := testNamePrefix + rand.String(10)

	assert.NoError(t, ipset.PostRequest{
		Client: client,
		Name:   ipSetName,
	}.Post())

	req := PostRequest{
		Client:    client,
		IPSetName: ipSetName,
		CIDR:      "192.168.0.0/16",
		Comment:   helpers.PtrTo("foobar"),
		NoMatch:   helpers.PtrTo(types.PVEBool(true)),
	}
	assert.NoError(t, req.Post())

	ipSetCIDRList, err := GetRequest{Client: helpers.APITokenTestClient(), IPSetName: ipSetName}.Get()
	assert.NoError(t, err)
	assert.NotNil(t, ipSetCIDRList)
	assert.Len(t, ipSetCIDRList, 1)

	ipSetCIDR := ipSetCIDRList.FindByCIDR(req.CIDR)
	assert.Equal(t, req.CIDR, ipSetCIDR.CIDR)
	assert.Equal(t, req.Comment, ipSetCIDR.Comment)
	assert.Equal(t, req.NoMatch, ipSetCIDR.NoMatch)
}
