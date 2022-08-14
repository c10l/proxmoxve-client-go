package plugins

import (
	"encoding/base64"
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
		Type:   "dns",
		API:    helpers.PtrTo("lua"),
		Data:   helpers.PtrTo(base64.StdEncoding.EncodeToString([]byte("test"))),
	}
	err := postReq.Post()
	assert.NoError(t, err)

	assert.Eventually(t, func() bool {
		_, err := ItemGetRequest{Client: helpers.APITokenTestClient(), ID: postReq.ID}.Get()
		return err == nil
	}, eventuallyTimeout, 500*time.Millisecond, err)

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
	assert.Equal(t, "node1,node2", *item.Nodes)
	assert.Equal(t, "lua", *item.API)
	assert.Equal(t, "test", *item.Data)
}

func TestItemPutDelete(t *testing.T) {
	postReq := PostRequest{
		Client: helpers.TicketTestClient(),
		ID:     "pmvetest_acme_" + rand.String(10),
		Type:   "dns",
		API:    helpers.PtrTo("lua"),
		Data:   helpers.PtrTo(base64.StdEncoding.EncodeToString([]byte("test"))),
		Nodes:  helpers.PtrTo([]string{"node1", "node2"}),
	}
	err := postReq.Post()
	assert.NoError(t, err)

	assert.Eventually(t, func() bool {
		_, err := ItemGetRequest{Client: helpers.APITokenTestClient(), ID: postReq.ID}.Get()
		return err == nil
	}, eventuallyTimeout, 500*time.Millisecond, err)

	itemPutReq := ItemPutRequest{
		Client:  helpers.TicketTestClient(),
		ID:      postReq.ID,
		Disable: helpers.PtrTo(types.PVEBool(true)),
		Delete:  helpers.PtrTo("nodes,data"),
	}
	err = itemPutReq.Put()
	assert.NoError(t, err)
	item, err := ItemGetRequest{Client: helpers.APITokenTestClient(), ID: postReq.ID}.Get()
	assert.NoError(t, err)
	assert.Equal(t, true, bool(item.Disable))
	assert.Nil(t, item.Nodes)
	assert.Nil(t, item.Data)
}
