package groups

import (
	"testing"
	"time"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemDelete(t *testing.T) {
	req := PostRequest{
		Client: helpers.TicketTestClient(),
		Group:  testNamePrefix + rand.String(3),
	}
	err := req.Post()
	assert.NoError(t, err)
	assert.Eventually(t, func() bool {
		_, err = ItemGetRequest{Client: helpers.APITokenTestClient(), Group: req.Group}.Get()
		return err == nil
	}, testEventuallyTimeout, 500*time.Millisecond)

	err = ItemDeleteRequest{Client: helpers.APITokenTestClient(), Group: req.Group}.Delete()
	assert.NoError(t, err)
	assert.Eventually(t, func() bool {
		_, err = ItemGetRequest{Client: helpers.APITokenTestClient(), Group: req.Group}.Get()
		return assert.Contains(t, err.Error(), "no such security group")
	}, testEventuallyTimeout, 500*time.Millisecond)
}
