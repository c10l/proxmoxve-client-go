package plugins

import (
	"testing"
	"time"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/c10l/proxmoxve-client-go/helpers/types"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemPut(t *testing.T) {
	postReq := PostRequest{
		Client: helpers.TicketTestClient(),
		ID:     "pmvetest_acme_" + rand.String(10),
		Type:   "standalone",
	}
	err := postReq.Post()
	assert.NoError(t, err)

	assert.Eventually(t, func() bool {
		_, err := ItemGetRequest{Client: helpers.APITokenTestClient(), ID: postReq.ID}.Get()
		return err == nil
	}, 5*time.Second, 500*time.Millisecond, err)

	itemPutReq := ItemPutRequest{
		Client:  helpers.TicketTestClient(),
		ID:      postReq.ID,
		Nodes:   &[]string{"node1", "node2"},
		Disable: helpers.PtrTo(types.PVEBool(true)),
	}
	err = itemPutReq.Put()
	assert.NoError(t, err)
	item, err := ItemGetRequest{Client: helpers.APITokenTestClient(), ID: postReq.ID}.Get()
	assert.NoError(t, err)
	assert.Equal(t, true, bool(item.Disable))
	assert.Equal(t, "node1,node2", item.Nodes)
}
