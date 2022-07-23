package storage

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestUnmarshalGetResponseStorageList(t *testing.T) {
	storageListJSON := []byte(`[
		{
			"content": "snippets,rootdir,backup,vztmpl,images,iso",
			"digest": "21d306286ca5d843b74e18f8cfb98dc705e975e6",
			"path": "/var/lib/vz",
			"prune-backups": "keep-all=1",
			"storage": "local",
			"type": "dir"
		},
		{
			"content": "images",
			"digest": "21d306286ca5d843b74e18f8cfb98dc705e975e6",
			"path": "/foo",
			"prune-backups": "keep-all=1",
			"shared": 0,
			"storage": "foo",
			"type": "dir"
		}
	]`)
	var storageList []GetResponseStorage
	err := json.Unmarshal(storageListJSON, &storageList)
	assert.NoError(t, err)
	assert.Equal(t, "local", (storageList)[0].Storage)
	assert.Contains(t, (storageList)[0].Content, ContentImages)
	assert.Contains(t, (storageList)[0].Content, ContentISO)
	assert.Equal(t, "foo", (storageList)[1].Storage)
	assert.Contains(t, (storageList)[1].Content, ContentImages)
}

func TestUnmarshalGetResponseStorage(t *testing.T) {
	expectedStorage := rand.String(10)
	storageJSON := []byte(fmt.Sprintf(`
    {
      "content": "%s,%s",
      "digest": "8391f10ff1f67c76bda33d11a07ca4504cad38be",
      "path": "/foobar",
      "storage": "%s",
      "type": "dir"
    }`, ContentImages, ContentISO, expectedStorage))

	storage := new(GetResponseStorage)
	err := json.Unmarshal(storageJSON, storage)
	assert.NoError(t, err)
	assert.Equal(t, expectedStorage, storage.Storage)
	assert.Contains(t, storage.Content, ContentImages)
	assert.Contains(t, storage.Content, ContentISO)
}

func TestGet(t *testing.T) {
	storageList, err := GetRequest{Client: helpers.APITokenTestClient()}.Get()
	assert.NoError(t, err)
	assert.NotNil(t, storageList)
}
