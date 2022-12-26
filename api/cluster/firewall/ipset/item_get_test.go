package ipset

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/cluster/firewall/ipset/ipset_cidr"
	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemGet(t *testing.T) {
	req := PostRequest{
		Client: helpers.APITokenTestClient(),
		Name:   testNamePrefix + rand.String(10),
	}
	err := req.Post()
	assert.NoError(t, err)

	ipSetList, err := ItemGetRequest{Client: helpers.APITokenTestClient(), Name: req.Name}.Get()
	assert.NoError(t, err)
	assert.NotNil(t, ipSetList)
	assert.Len(t, *ipSetList, 0)
}

func TestItemGetWithCIDR(t *testing.T) {
	req := PostRequest{
		Client: helpers.APITokenTestClient(),
		Name:   testNamePrefix + rand.String(10),
	}
	err := req.Post()
	assert.NoError(t, err)

	cidrReq := ipset_cidr.PostRequest{
		Client:    req.Client,
		IPSetName: req.Name,
		CIDR:      "10.0.0.0/16",
	}
	assert.NoError(t, cidrReq.Post())

	ipSetCIDRList, err := ItemGetRequest{Client: helpers.APITokenTestClient(), Name: req.Name}.Get()
	assert.NoError(t, err)
	assert.Len(t, *ipSetCIDRList, 1)
	assert.Equal(t, (*ipSetCIDRList)[0].CIDR, cidrReq.CIDR)
	assert.Nil(t, (*ipSetCIDRList)[0].Comment)
}

func TestItemGetDeleteComment(t *testing.T) {
	req := PostRequest{
		Client: helpers.APITokenTestClient(),
		Name:   testNamePrefix + rand.String(10),
	}
	err := req.Post()
	assert.NoError(t, err)

	cidrReq := ipset_cidr.PostRequest{
		Client:    req.Client,
		IPSetName: req.Name,
		CIDR:      "10.0.0.0/16",
		Comment:   helpers.PtrTo("foobar"),
	}
	assert.NoError(t, cidrReq.Post())

	ipSetList, err := GetRequest{Client: helpers.APITokenTestClient()}.Get()
	assert.NoError(t, err)
	ipSet := ipSetList.FindByName(req.Name)
	assert.NotNil(t, ipSet)
}
