package ipset

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
