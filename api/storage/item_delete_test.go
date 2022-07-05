package storage

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/test"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemDelete(t *testing.T) {
	req := PostRequest{Client: test.APITokenTestClient()}
	req.Storage = "pmvetest_" + rand.String(10)
	req.StorageType = TypeDir
	path := "/foo"
	req.DirPath = &path
	_, err := req.Do()
	assert.NoError(t, err)

	err = ItemDeleteRequest{Client: test.APITokenTestClient(), Storage: req.Storage}.Delete()
	assert.NoError(t, err)

	_, err = ItemGetRequest{Client: test.APITokenTestClient(), Storage: req.Storage}.Do()
	assert.ErrorContains(t, err, "does not exist")
}
