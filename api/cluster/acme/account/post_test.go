package account

import (
	"testing"
	"time"

	"github.com/c10l/proxmoxve-client-go/api/test"
	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestPost(t *testing.T) {
	req := PostRequest{
		Client:    test.TicketTestClient(),
		Name:      helpers.PtrTo("pmvetest_acme_" + rand.String(10)),
		Contact:   "foobar@baz.com",
		Directory: helpers.PtrTo("https://acme-staging-v02.api.letsencrypt.org/directory"),
		TOSurl:    helpers.PtrTo("https://letsencrypt.org/documents/LE-SA-v1.2-November-15-2017.pdf"),
	}
	resp, err := req.Post()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Contains(t, *resp, "UPID:pve:")

	assert.Eventually(t, func() bool {
		accountList, err := GetRequest{Client: test.APITokenTestClient()}.Get()
		for _, i := range *accountList {
			if i.Name == *req.Name {
				assert.NoError(t, err)
				assert.NotNil(t, accountList)
				return true
			}
		}
		return false
	}, 5*time.Second, 500*time.Millisecond, err)
}
