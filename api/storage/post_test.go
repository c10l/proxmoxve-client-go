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
	response, err := req.Do()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, req.Storage, response.Storage)
}
