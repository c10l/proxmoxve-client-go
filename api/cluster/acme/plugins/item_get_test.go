package plugins

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
		ID:      "pmvetest_acme_" + rand.String(10),
		Type:    "standalone",
		Disable: helpers.PtrTo(helpers.IntBool(true)),
	}
	err := req.Post()
	assert.NoError(t, err)

	assert.Eventually(t, func() bool {
		plugin, err := ItemGetRequest{Client: helpers.APITokenTestClient(), ID: req.ID}.Get()
		return err == nil &&
			plugin.Plugin == req.ID &&
			plugin.Type == req.Type
	}, 5*time.Second, 500*time.Millisecond, err)
}
