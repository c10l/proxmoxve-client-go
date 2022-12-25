package ipset_cidr

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/cluster/firewall/ipset"
	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemDelete(t *testing.T) {
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
	}
	err = req.Post()
	assert.NoError(t, err)

	_, err = ItemGetRequest{Client: helpers.APITokenTestClient(), IPSetName: ipSetReq.Name, CIDR: req.CIDR}.Get()
	assert.NoError(t, err)

	err = ItemDeleteRequest{Client: helpers.APITokenTestClient(), IPSetName: req.IPSetName, CIDR: req.CIDR}.Delete()
	assert.NoError(t, err)

	_, err = ItemGetRequest{Client: helpers.APITokenTestClient(), IPSetName: ipSetReq.Name, CIDR: req.CIDR}.Get()
	assert.Contains(t, err.Error(), "no such IP/Network")
}
