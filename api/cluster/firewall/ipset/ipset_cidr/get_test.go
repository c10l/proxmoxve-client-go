package ipset_cidr

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/cluster/firewall/ipset"
	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestGet(t *testing.T) {
	ipSetReq := ipset.PostRequest{
		Client: helpers.APITokenTestClient(),
		Name:   testNamePrefix + rand.String(10),
	}
	err := ipSetReq.Post()
	assert.NoError(t, err)

	cidrList, err := GetRequest{Client: helpers.APITokenTestClient(), IPSetName: ipSetReq.Name}.Get()
	assert.NoError(t, err)
	assert.NotNil(t, cidrList)
}
