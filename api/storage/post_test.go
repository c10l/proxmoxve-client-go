package storage

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/test"
	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestPostDir(t *testing.T) {
	req := PostRequest{Client: test.APITokenTestClient()}
	req.Storage = "pmvetest_dir_" + rand.String(10)
	req.StorageType = TypeDir
	path := "/foo"
	req.DirPath = &path
	req.Nodes = &[]string{"pve"}
	response, err := req.Do()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, req.Storage, response.Storage)
	assert.Equal(t, *req.Nodes, []string{"pve"})
}

func TestPostDirSharedAndDisable(t *testing.T) {
	req := PostRequest{Client: test.APITokenTestClient(), DirShared: helpers.PtrTo(true), Disable: helpers.PtrTo(true)}
	req.Storage = "pmvetest_dir_" + rand.String(10)
	req.StorageType = TypeDir
	path := "/foo"
	req.DirPath = &path
	req.Nodes = &[]string{"pve"}
	response, err := req.Do()
	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestPostNFS(t *testing.T) {
	req := PostRequest{Client: test.APITokenTestClient()}
	req.Storage = "pmvetest_nfs_" + rand.String(10)
	req.StorageType = TypeNFS
	req.NFSMountOptions = helpers.PtrTo("vers=4.2")
	req.NFSServer = helpers.PtrTo("100.100.100.100")
	req.NFSExport = helpers.PtrTo("/mnt/nfs_export/path")
	req.Nodes = &[]string{"pve"}
	req.Disable = helpers.PtrTo(true)
	resp, err := req.Do()
	assert.NoError(t, err)
	assert.Contains(t, resp.Storage, "pmvetest_nfs_")
}
