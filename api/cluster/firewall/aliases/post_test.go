package aliases

import (
	"testing"
	"time"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestPost(t *testing.T) {
	req := PostRequest{
		Client: helpers.APITokenTestClient(),
		Name:   testNamePrefix + rand.String(10),
		CIDR:   "10.10.0.0/16",
	}
	err := req.Post()
	assert.NoError(t, err)

	assert.Eventually(t, func() bool {
		aliasesList, err := GetRequest{Client: helpers.APITokenTestClient()}.Get()
		for _, i := range aliasesList {
			if i.Name == req.Name {
				assert.NoError(t, err)
				assert.NotNil(t, aliasesList)
				return true
			}
		}
		return false
	}, testEventuallyTimeout, 500*time.Millisecond, err)
}
