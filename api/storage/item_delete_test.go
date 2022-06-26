package storage

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/test"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemDelete(t *testing.T) {
	req := PostRequest{client: test.APITestClient()}
	req.Storage = "a" + rand.String(10)
	req.StorageType = TypeDir
	path := "/foo"
	req.Path = &path
	_, err := req.Do()
	assert.NoError(t, err)

	err = ItemDeleteRequest{client: test.APITestClient(), Storage: req.Storage}.Do()
	assert.NoError(t, err)

	_, err = ItemGetRequest{client: test.APITestClient(), Storage: req.Storage}.Do()
	assert.ErrorContains(t, err, "does not exist")
}