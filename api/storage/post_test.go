package storage

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/test"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestPost(t *testing.T) {
	req := PostRequest{client: test.APITestClient()}
	req.Storage = "a" + rand.String(10)
	req.StorageType = TypeDir
	req.Path = "/foo"
	response, err := req.Do()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, req.Storage, response.Storage)
}
