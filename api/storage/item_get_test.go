package storage

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/c10l/proxmoxve-client-go/helpers/types"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestItemGet(t *testing.T) {
	req := PostRequest{
		Client:      helpers.APITokenTestClient(),
		Storage:     "pmvetest_" + rand.String(10),
		StorageType: TypeDir,
		DirPath:     func() *string { s := "/foo"; return &s }(),
		Nodes:       &[]string{"pve"},
	}
	_, err := req.Post()
	assert.NoError(t, err)

	resp, err := ItemGetRequest{Client: helpers.APITokenTestClient(), Storage: req.Storage}.Get()
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
		Client:      helpers.APITokenTestClient(),
		Storage:     "pmvetest_" + rand.String(10),
		StorageType: TypeNFS,
		NFSExport:   helpers.PtrTo("/foo"),
		NFSServer:   helpers.PtrTo("1.2.3.4"),
		Disable:     helpers.PtrTo(types.PVEBool(true)),
	}
	_, err := req.Post()
	assert.NoError(t, err)

	resp, err := ItemGetRequest{Client: helpers.APITokenTestClient(), Storage: req.Storage}.Get()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.Storage, resp.Storage)
	assert.Equal(t, []string{"images"}, resp.Content)
	assert.Equal(t, "/foo", resp.NFSExport)
	assert.Equal(t, TypeNFS, resp.Type)
}
