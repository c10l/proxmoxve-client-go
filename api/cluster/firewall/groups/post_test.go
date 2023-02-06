package groups

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
		Group:  testNamePrefix + rand.String(3),
	}
	err := req.Post()
	assert.NoError(t, err)

	assert.Eventually(t, func() bool {
		groupsList, err := GetRequest{Client: helpers.APITokenTestClient()}.Get()
		for _, i := range groupsList {
			if i.Group == req.Group {
				assert.NoError(t, err)
				assert.NotNil(t, groupsList)
				return true
			}
		}
		return false
	}, testEventuallyTimeout, 500*time.Millisecond, err)
}
