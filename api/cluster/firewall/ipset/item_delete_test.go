package ipset

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemDelete(t *testing.T) {
	req := PostRequest{
		Client: helpers.APITokenTestClient(),
		Name:   testNamePrefix + rand.String(10),
	}
	err := req.Post()
	assert.NoError(t, err)
	_, err = ItemGetRequest{Client: helpers.APITokenTestClient(), Name: req.Name}.Get()
	assert.NoError(t, err)

	err = ItemDeleteRequest{Client: helpers.APITokenTestClient(), Name: req.Name}.Delete()
	assert.NoError(t, err)
	_, err = ItemGetRequest{Client: helpers.APITokenTestClient(), Name: req.Name}.Get()
	assert.Contains(t, err.Error(), "no such IPSet")
}
