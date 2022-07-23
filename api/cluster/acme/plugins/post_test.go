package plugins

import (
	"testing"
	"time"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestPost(t *testing.T) {
	req := PostRequest{
		Client: helpers.TicketTestClient(),
		ID:     "pmvetest_acme_" + rand.String(10),
		Type:   "standalone",
	}
	err := req.Post()
	assert.NoError(t, err)

	assert.Eventually(t, func() bool {
		pluginsList, err := GetRequest{Client: helpers.APITokenTestClient()}.Get()
		for _, i := range pluginsList {
			if i.Plugin == req.ID {
				assert.NoError(t, err)
				assert.NotNil(t, pluginsList)
				return true
			}
		}
		return false
	}, 5*time.Second, 500*time.Millisecond, err)
}
