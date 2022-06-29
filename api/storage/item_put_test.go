package storage

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/test"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemPut(t *testing.T) {
	req := PostRequest{
		Client:      test.APITestClient(),
		Storage:     "a" + rand.String(10),
		StorageType: TypeDir,
		Path:        func() *string { s := "/foo"; return &s }(),
	}
	_, err := req.Do()
	assert.NoError(t, err)

	putRest, err := ItemPutRequest{Client: test.APITestClient(), Storage: req.Storage, Content: []string{"images"}}.Do()
	assert.NoError(t, err)
	assert.Equal(t, req.Storage, putRest.Storage)
	assert.Equal(t, TypeDir, putRest.Type)

	resp, err := ItemGetRequest{Client: test.APITestClient(), Storage: req.Storage}.Do()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Content, 1)
	assert.Equal(t, ContentImages, resp.Content[0])
}
