package aliases

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
		Name:   testNamePrefix + rand.String(10),
		CIDR:   "1.1.1.0/24",
	}
	err := req.Post()
	assert.NoError(t, err)
	assert.Eventually(t, func() bool {
		_, err = ItemGetRequest{Client: helpers.APITokenTestClient(), Name: req.Name}.Get()
		return err == nil
	}, testEventuallyTimeout, 500*time.Millisecond)

	err = ItemDeleteRequest{Client: helpers.APITokenTestClient(), Name: req.Name}.Delete()
	assert.NoError(t, err)
	assert.Eventually(t, func() bool {
		_, err = ItemGetRequest{Client: helpers.APITokenTestClient(), Name: req.Name}.Get()
		return assert.Contains(t, err.Error(), "no such alias")
	}, testEventuallyTimeout, 500*time.Millisecond)
}
