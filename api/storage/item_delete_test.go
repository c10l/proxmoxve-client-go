package storage

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemDelete(t *testing.T) {
	req := PostRequest{Client: helpers.APITokenTestClient()}
	req.Storage = "pmvetest_" + rand.String(10)
	req.StorageType = TypeDir
	path := "/foo"
	req.DirPath = &path
	_, err := req.Post()
	assert.NoError(t, err)

	err = ItemDeleteRequest{Client: helpers.APITokenTestClient(), Storage: req.Storage}.Delete()
	assert.NoError(t, err)

	_, err = ItemGetRequest{Client: helpers.APITokenTestClient(), Storage: req.Storage}.Get()
	assert.ErrorContains(t, err, "does not exist")
}
