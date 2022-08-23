package aliases

import (
	"testing"
	"time"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemPutCIDR(t *testing.T) {
	req := PostRequest{
		Client: helpers.APITokenTestClient(),
		Name:   "pmvetest_fw_aliases_" + rand.String(10),
		CIDR:   "1.1.1.0/24",
	}
	err := req.Post()
	assert.NoError(t, err)
	assert.Eventually(t, func() bool {
		_, err = ItemGetRequest{Client: helpers.APITokenTestClient(), Name: req.Name}.Get()
		return err == nil
	}, testEventuallyTimeout, 500*time.Millisecond)

	err = ItemPutRequest{Client: helpers.APITokenTestClient(), Name: req.Name, CIDR: "2.2.0.0/16"}.Put()
	assert.NoError(t, err)
	assert.Eventually(t, func() bool {
		item, _ := ItemGetRequest{Client: helpers.APITokenTestClient(), Name: req.Name}.Get()
		return assert.Equal(t, item.CIDR, "2.2.0.0/16")
	}, testEventuallyTimeout, 500*time.Millisecond)
}

func TestItemPutRename(t *testing.T) {
	req := PostRequest{
		Client: helpers.APITokenTestClient(),
		Name:   testNamePrefix + rand.String(10),
		CIDR:   "1.1.1.0/24",
	}
	err := req.Post()
	assert.NoError(t, err)
	assert.Eventually(t, func() bool {
		_, err = ItemGetRequest{Client: helpers.APITokenTestClient(), Name: req.Name}.Get()
		return err == nil
	}, testEventuallyTimeout, 500*time.Millisecond)

	expectedName := testNamePrefix + rand.String(10)
	err = ItemPutRequest{Client: helpers.APITokenTestClient(), Name: req.Name, CIDR: "1.1.1.0/24", Rename: &expectedName}.Put()
	assert.NoError(t, err)
	assert.Eventually(t, func() bool {
		item, _ := ItemGetRequest{Client: helpers.APITokenTestClient(), Name: req.Name}.Get()
		return item == nil
	}, testEventuallyTimeout, 500*time.Millisecond)
	assert.Eventually(t, func() bool {
		item, _ := ItemGetRequest{Client: helpers.APITokenTestClient(), Name: expectedName}.Get()
		return item != nil &&
			item.CIDR == "1.1.1.0/24"
	}, testEventuallyTimeout, 500*time.Millisecond)
}
