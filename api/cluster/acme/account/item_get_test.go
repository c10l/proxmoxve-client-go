package account

import (
	"testing"
	"time"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemGet(t *testing.T) {
	req := PostRequest{
		Client:    helpers.TicketTestClient(),
		Name:      "pmvetest_acme_" + rand.String(10),
		Contact:   "foobar@baz.com",
		Directory: helpers.PtrTo("https://127.0.0.1:14000/dir"),
		TOSurl:    helpers.PtrTo("https://letsencrypt.org/documents/LE-SA-v1.2-November-15-2017.pdf"),
	}
	_, err := req.Post()
	assert.NoError(t, err)

	var account *ItemGetResponse
	assert.Eventually(t, func() bool {
		account, err = ItemGetRequest{Client: helpers.TicketTestClient(), Name: req.Name}.Get()
		return err == nil &&
			*req.Directory == account.Directory &&
			len(account.Account.Contact) > 0 &&
			"mailto:"+req.Contact == account.Account.Contact[0]
	}, eventuallyTimeout, 500*time.Millisecond)
	assert.NoError(t, err)
	assert.Equal(t, *req.Directory, account.Directory)
	assert.Greater(t, len(account.Account.Contact), 0)
	assert.Equal(t, "mailto:"+req.Contact, account.Account.Contact[0])
}
