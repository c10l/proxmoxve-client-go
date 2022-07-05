package account

import (
	"testing"
	"time"

	"github.com/c10l/proxmoxve-client-go/api/test"
	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemDelete(t *testing.T) {
	req := PostRequest{
		Client:    test.TicketTestClient(),
		Name:      helpers.PtrTo("pmvetest_acme_" + rand.String(10)),
		Contact:   "foobar@baz.com",
		Directory: helpers.PtrTo("https://acme-staging-v02.api.letsencrypt.org/directory"),
		TOSurl:    helpers.PtrTo("https://letsencrypt.org/documents/LE-SA-v1.2-November-15-2017.pdf"),
	}
	_, err := req.Do()
	assert.NoError(t, err)

	assert.Eventually(t, func() bool {
		err = ItemDeleteRequest{Client: test.TicketTestClient(), Name: *req.Name}.Do()
		return err == nil
	}, 5*time.Second, 500*time.Millisecond, err)

	// _, err = ItemGetRequest{Client: test.APITokenTestClient(), Account: req.Account}.Do()
	// assert.ErrorContains(t, err, "does not exist")
}
