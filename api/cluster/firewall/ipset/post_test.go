package ipset

import (
	"testing"
	"time"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestPost(t *testing.T) {
	req := PostRequest{
		Client:  helpers.APITokenTestClient(),
		Name:    testNamePrefix + rand.String(10),
		Comment: helpers.PtrTo("foobar"),
	}
	err := req.Post()
	assert.NoError(t, err)

	assert.Eventually(t, func() bool {
		ipsetList, err := GetRequest{Client: helpers.APITokenTestClient()}.Get()
		for _, i := range ipsetList {
			if i.Name == req.Name {
				assert.NoError(t, err)
				assert.NotNil(t, ipsetList)
				return true
			}
		}
		return false
	}, testEventuallyTimeout, 500*time.Millisecond, err)
}
