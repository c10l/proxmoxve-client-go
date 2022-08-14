package account

import (
	"testing"
	"time"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemPut(t *testing.T) {
	postReq := PostRequest{
		Client:    helpers.TicketTestClient(),
		Name:      "pmvetest_acme_" + rand.String(10),
		Contact:   "foobar@baz.com",
		Directory: helpers.PtrTo("https://127.0.0.1:14000/dir"),
		TOSurl:    helpers.PtrTo("https://letsencrypt.org/documents/LE-SA-v1.2-November-15-2017.pdf"),
	}
	postResp, err := postReq.Post()
	assert.NoError(t, err)
	assert.NotNil(t, postResp)
	assert.Contains(t, *postResp, "UPID:pve:")

	assert.Eventually(t, func() bool {
		accountList, err := GetRequest{Client: helpers.APITokenTestClient()}.Get()
		for _, i := range accountList {
			if i.Name == postReq.Name {
				assert.NoError(t, err)
				assert.NotNil(t, accountList)
				return true
			}
		}
		return false
	}, eventuallyTimeout, 500*time.Millisecond, err)

	itemPutReq := ItemPutRequest{
		Client:  helpers.TicketTestClient(),
		Name:    postReq.Name,
		Contact: "foo@barbaz.com",
	}
	itemPutResp, err := itemPutReq.Put()
	assert.NoError(t, err)
	assert.NotNil(t, itemPutResp)
	assert.Contains(t, *itemPutResp, "UPID:pve:")
	assert.Equal(t, "foo@barbaz.com", itemPutReq.Contact)
}
