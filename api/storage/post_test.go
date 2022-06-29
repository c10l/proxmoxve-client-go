package storage

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/test"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestPost(t *testing.T) {
	req := PostRequest{Client: test.APITestClient()}
	req.Storage = "a" + rand.String(10)
	req.StorageType = TypeDir
	path := "/foo"
	req.Path = &path
	req.Nodes = &[]string{"pve"}
	response, err := req.Do()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, req.Storage, response.Storage)
	assert.Equal(t, *req.Nodes, []string{"pve"})
}

func TestPostSharedAndDisable(t *testing.T) {
	req := PostRequest{Client: test.APITestClient(), Shared: ptrTo(true), Disable: ptrTo(true)}
	req.Storage = "a" + rand.String(10)
	req.StorageType = TypeDir
	path := "/foo"
	req.Path = &path
	req.Nodes = &[]string{"pve"}
	response, err := req.Do()
	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func ptrTo[T any](t T) *T {
	return &t
}
