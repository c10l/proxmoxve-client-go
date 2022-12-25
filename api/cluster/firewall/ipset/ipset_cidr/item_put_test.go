package ipset_cidr

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/cluster/firewall/ipset"
	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemPutComment(t *testing.T) {
	ipSetReq := ipset.PostRequest{
		Client: helpers.APITokenTestClient(),
		Name:   testNamePrefix + rand.String(10),
	}
	err := ipSetReq.Post()
	assert.NoError(t, err)

	postReq := PostRequest{
		Client:    helpers.APITokenTestClient(),
		IPSetName: ipSetReq.Name,
		CIDR:      "10.17.0.0/16",
	}
	err = postReq.Post()
	assert.NoError(t, err)

	_, err = ItemGetRequest{Client: helpers.APITokenTestClient(), IPSetName: ipSetReq.Name, CIDR: postReq.CIDR}.Get()
	assert.NoError(t, err)

	req := ItemPutRequest{
		Client:    helpers.APITokenTestClient(),
		IPSetName: postReq.IPSetName,
		CIDR:      postReq.CIDR,
		Comment:   helpers.PtrTo(rand.String(30)),
	}
	err = req.Put()
	assert.NoError(t, err)

	item, err := ItemGetRequest{Client: helpers.APITokenTestClient(), IPSetName: postReq.IPSetName, CIDR: postReq.CIDR}.Get()
	assert.NoError(t, err)
	assert.Equal(t, req.Comment, item.Comment)
}
