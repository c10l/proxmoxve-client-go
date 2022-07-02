package storage

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/test"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemGet(t *testing.T) {
	req := PostRequest{
		Client:      test.APITestClient(),
		Storage:     "pmvetest_" + rand.String(10),
		StorageType: TypeDir,
		DirPath:     func() *string { s := "/foo"; return &s }(),
		Nodes:       &[]string{"pve"},
	}
	_, err := req.Do()
	assert.NoError(t, err)

	resp, err := ItemGetRequest{Client: test.APITestClient(), Storage: req.Storage}.Do()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.Storage, resp.Storage)
	assert.Equal(t, []string{"images", "rootdir"}, resp.Content)
	assert.Equal(t, "/foo", resp.Path)
	assert.Equal(t, TypeDir, resp.Type)
	assert.Equal(t, []string{"pve"}, resp.Nodes)
}

func TestItemNFSGet(t *testing.T) {
	req := PostRequest{
		Client:      test.APITestClient(),
		Storage:     "pmvetest_" + rand.String(10),
		StorageType: TypeNFS,
		NFSExport:   ptrTo("/foo"),
		NFSServer:   ptrTo("1.2.3.4"),
		Disable:     ptrTo(true),
	}
	_, err := req.Do()
	assert.NoError(t, err)

	resp, err := ItemGetRequest{Client: test.APITestClient(), Storage: req.Storage}.Do()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.Storage, resp.Storage)
	assert.Equal(t, []string{"images"}, resp.Content)
	assert.Equal(t, "/foo", resp.NFSExport)
	assert.Equal(t, TypeNFS, resp.Type)
}
