package storage

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/test"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestPostDir(t *testing.T) {
	req := PostRequest{Client: test.APITestClient()}
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
	req := PostRequest{Client: test.APITestClient(), DirShared: ptrTo(true), Disable: ptrTo(true)}
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
	req := PostRequest{Client: test.APITestClient()}
	req.Storage = "pmvetest_nfs_" + rand.String(10)
	req.StorageType = TypeNFS
	req.NFSMountOptions = ptrTo("vers=4.2")
	req.NFSServer = ptrTo("100.100.100.100")
	req.NFSExport = ptrTo("/mnt/nfs_export/path")
	req.Nodes = &[]string{"pve"}
	req.Disable = ptrTo(true)
	resp, err := req.Do()
	assert.NoError(t, err)
	assert.Contains(t, resp.Storage, "pmvetest_nfs_")
}

func ptrTo[T any](t T) *T {
	return &t
}
