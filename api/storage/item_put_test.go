package storage

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemPut(t *testing.T) {
	req := PostRequest{
		Client:      helpers.APITokenTestClient(),
		Storage:     "pmvetest_" + rand.String(10),
		StorageType: TypeDir,
		DirPath:     func() *string { s := "/foo"; return &s }(),
	}
	_, err := req.Post()
	assert.NoError(t, err)

	putRest, err := ItemPutRequest{
		Client:  helpers.APITokenTestClient(),
		Storage: req.Storage,
		Content: &[]string{"images"},
		Nodes:   &[]string{"foo", "bar"},
		Disable: helpers.PtrTo(helpers.IntBool(true)),
		Shared:  helpers.PtrTo(helpers.IntBool(false)),
	}.Put()
	assert.NoError(t, err)
	assert.Equal(t, req.Storage, putRest.Storage)
	assert.Equal(t, TypeDir, putRest.Type)

	resp, err := ItemGetRequest{Client: helpers.APITokenTestClient(), Storage: req.Storage}.Get()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Content, 1)
	assert.Equal(t, ContentImages, resp.Content[0])
}
