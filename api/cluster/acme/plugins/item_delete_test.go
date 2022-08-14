package plugins

import (
	"fmt"
	"testing"
	"time"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemDelete(t *testing.T) {
	req := PostRequest{
		Client: helpers.TicketTestClient(),
		ID:     "pmvetest_acme_" + rand.String(10),
		Type:   "standalone",
	}
	err := req.Post()
	assert.NoError(t, err)
	assert.Eventually(t, func() bool {
		_, err = ItemGetRequest{Client: helpers.TicketTestClient(), ID: req.ID}.Get()
		return err == nil
	}, eventuallyTimeout, 500*time.Millisecond)

	err = ItemDeleteRequest{Client: helpers.TicketTestClient(), ID: req.ID}.Delete()
	assert.NoError(t, err)
	assert.Eventually(t, func() bool {
		_, err = ItemGetRequest{Client: helpers.TicketTestClient(), ID: req.ID}.Get()
		return assert.Contains(t, err.Error(), fmt.Sprintf("ACME plugin '%s' not defined", req.ID))
	}, eventuallyTimeout, 500*time.Millisecond)
}
