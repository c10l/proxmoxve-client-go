package ipset

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestPost(t *testing.T) {
	req := PostRequest{
		Client:  helpers.APITokenTestClient(),
		Name:    testNamePrefix + rand.String(10),
		Comment: helpers.PtrTo("foobar"),
	}
	err := req.Post()
	assert.NoError(t, err)

	ipsetList, err := GetRequest{Client: helpers.APITokenTestClient()}.Get()
	assert.NoError(t, err)
	assert.NotNil(t, ipsetList)
	ipSet := ipsetList.FindByName(req.Name)
	assert.NotNil(t, ipSet)
}
