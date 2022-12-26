package ipset_cidr

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/cluster/firewall/ipset"
	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/c10l/proxmoxve-client-go/helpers/types"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemPutAddComment(t *testing.T) {
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
		NoMatch:   helpers.PtrTo(types.PVEBool(true)),
	}
	err = req.Put()
	assert.NoError(t, err)

	item, err := ItemGetRequest{Client: helpers.APITokenTestClient(), IPSetName: postReq.IPSetName, CIDR: postReq.CIDR}.Get()
	assert.NoError(t, err)
	assert.Equal(t, req.Comment, item.Comment)
	assert.Equal(t, req.NoMatch, item.NoMatch)
}

func TestItemPutDeleteComment(t *testing.T) {
	ipSetReq := ipset.PostRequest{
		Client: helpers.APITokenTestClient(),
		Name:   testNamePrefix + "delete_comment_" + rand.String(10),
	}
	err := ipSetReq.Post()
	assert.NoError(t, err)

	postReq := PostRequest{
		Client:    helpers.APITokenTestClient(),
		IPSetName: ipSetReq.Name,
		CIDR:      "10.17.0.0/16",
		Comment:   helpers.PtrTo(rand.String(30)),
	}
	err = postReq.Post()
	assert.NoError(t, err)

	item, err := ItemGetRequest{Client: helpers.APITokenTestClient(), IPSetName: ipSetReq.Name, CIDR: postReq.CIDR}.Get()
	assert.NoError(t, err)
	assert.Equal(t, postReq.Comment, item.Comment)

	req := ItemPutRequest{
		Client:    helpers.APITokenTestClient(),
		IPSetName: postReq.IPSetName,
		CIDR:      postReq.CIDR,
		Comment:   nil,
	}
	err = req.Put()
	assert.NoError(t, err)

	item, err = ItemGetRequest{Client: helpers.APITokenTestClient(), IPSetName: postReq.IPSetName, CIDR: postReq.CIDR}.Get()
	assert.NoError(t, err)
	assert.Nil(t, item.Comment)
	assert.Equal(t, req.NoMatch, item.NoMatch)
}
