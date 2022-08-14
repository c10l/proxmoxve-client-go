package account

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemDelete(t *testing.T) {
	req := PostRequest{
		Client:    helpers.TicketTestClient(),
		Name:      "pmvetest_acme_" + rand.String(10),
		Contact:   "foobar@baz.com",
		Directory: helpers.PtrTo("https://127.0.0.1:14000/dir"),
		TOSurl:    helpers.PtrTo("https://letsencrypt.org/documents/LE-SA-v1.2-November-15-2017.pdf"),
	}
	_, err := req.Post()
	assert.NoError(t, err)
	assert.Eventually(t, func() bool {
		_, err = ItemGetRequest{Client: helpers.TicketTestClient(), Name: req.Name}.Get()
		return err == nil
	}, eventuallyTimeout, 500*time.Millisecond)

	err = ItemDeleteRequest{Client: helpers.TicketTestClient(), Name: req.Name}.Delete()
	assert.NoError(t, err)
	assert.Eventually(t, func() bool {
		_, err = ItemGetRequest{Client: helpers.TicketTestClient(), Name: req.Name}.Get()
		return strings.Contains(err.Error(), fmt.Sprintf("ACME account config file '%s' does not exist.", req.Name))
	}, eventuallyTimeout, 500*time.Millisecond)
}
