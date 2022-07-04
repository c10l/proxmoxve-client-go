package storage

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/test"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemPut(t *testing.T) {
	req := PostRequest{
		Client:      test.APITokenTestClient(),
		Storage:     "pmvetest_" + rand.String(10),
		StorageType: TypeDir,
		DirPath:     func() *string { s := "/foo"; return &s }(),
	}
	_, err := req.Do()
	assert.NoError(t, err)

	putRest, err := ItemPutRequest{
		Client:  test.APITokenTestClient(),
		Storage: req.Storage,
		Content: &[]string{"images"},
		Nodes:   &[]string{"foo", "bar"},
		Disable: func() *bool { b := true; return &b }(),
		Shared:  func() *bool { b := true; return &b }(),
	}.Do()
	assert.NoError(t, err)
	assert.Equal(t, req.Storage, putRest.Storage)
	assert.Equal(t, TypeDir, putRest.Type)

	resp, err := ItemGetRequest{Client: test.APITokenTestClient(), Storage: req.Storage}.Do()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Content, 1)
	assert.Equal(t, ContentImages, resp.Content[0])
}
