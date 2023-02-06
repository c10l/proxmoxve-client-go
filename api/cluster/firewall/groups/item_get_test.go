package groups

import (
	"testing"
	"time"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemGet(t *testing.T) {
	req := PostRequest{
		Client:  helpers.APITokenTestClient(),
		Group:   testNamePrefix + rand.String(3),
		Comment: helpers.PtrTo(rand.String(10)),
	}
	err := req.Post()
	assert.NoError(t, err)

	assert.Eventually(t, func() bool {
		group, err := ItemGetRequest{Client: helpers.APITokenTestClient(), Group: req.Group}.Get()
		return err == nil &&
			len(*group) == 0
	}, testEventuallyTimeout, 500*time.Millisecond, err)
}
