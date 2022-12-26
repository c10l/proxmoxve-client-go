package ipset

import (
	"testing"
	"time"

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

	assert.Eventually(t, func() bool {
		ipSetList, err := ItemGetRequest{Client: helpers.APITokenTestClient(), Name: req.Name}.Get()
		pass := err == nil &&
			ipSetList != nil &&
			len(*ipSetList) == 0
		return pass
	}, testEventuallyTimeout, 500*time.Millisecond, err)
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

	assert.Eventually(t, func() bool {
		ipSetList, err := ItemGetRequest{Client: helpers.APITokenTestClient(), Name: req.Name}.Get()
		pass := err == nil &&
			ipSetList != nil &&
			len(*ipSetList) == 1
		return pass
	}, testEventuallyTimeout, 500*time.Millisecond, err)
}
