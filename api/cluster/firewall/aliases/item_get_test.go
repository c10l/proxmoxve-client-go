package aliases

import (
	"testing"
	"time"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemGet(t *testing.T) {
	req := PostRequest{
		Client: helpers.APITokenTestClient(),
		Name:   "pmvetest_fw_alias_" + rand.String(10),
		CIDR:   "1.1.1.0/24",
	}
	err := req.Post()
	assert.NoError(t, err)

	assert.Eventually(t, func() bool {
		alias, err := ItemGetRequest{Client: helpers.APITokenTestClient(), Name: req.Name}.Get()
		return err == nil &&
			alias.CIDR == req.CIDR &&
			alias.IPVersion == 4
	}, eventuallyTimeout, 500*time.Millisecond, err)
}
