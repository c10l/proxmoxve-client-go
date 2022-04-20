package api2

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestUnmarshalStorage(t *testing.T) {
	expectedStorage := rand.String(10)
	storageJSON := []byte(fmt.Sprintf(`
    {
      "content": "%s,%s",
      "digest": "8391f10ff1f67c76bda33d11a07ca4504cad38be",
      "path": "/foobar",
      "storage": "%s",
      "type": "dir"
    }`, StorageContentImages, StorageContentISO, expectedStorage))

	storage := new(Storage)
	err := json.Unmarshal(storageJSON, storage)
	assert.NoError(t, err)
	assert.Equal(t, expectedStorage, storage.Storage)
	assert.Contains(t, storage.Content, StorageContentImages)
	assert.Contains(t, storage.Content, StorageContentISO)
}

func TestStorageCreateAndRetrieve(t *testing.T) {
	storageStorage := "a" + rand.String(10)
	actualStorage, err := testClient.CreateStorage(storageStorage, StorageTypeDir, map[string]string{"path": "/foo"})
	assert.NoError(t, err)
	assert.Equal(t, storageStorage, actualStorage.Storage)
	assert.Equal(t, StorageTypeDir, actualStorage.Type)

	storage, err := testClient.RetrieveStorage(storageStorage)
	assert.NoError(t, err)
	assert.Equal(t, storageStorage, storage.Storage)
	assert.Contains(t, storage.Content, StorageContentImages)
	assert.Contains(t, storage.Content, StorageContentRootDir)
}

func TestStorageDelete(t *testing.T) {
	storageStorage := "a" + rand.String(10)

	_, err := testClient.CreateStorage(storageStorage, StorageTypeDir, map[string]string{"path": "/foo"})
	assert.NoError(t, err)

	_, err = testClient.RetrieveStorage(storageStorage)
	assert.NoError(t, err)

	err = testClient.DeleteStorage(storageStorage)
	assert.NoError(t, err)

	_, err = testClient.RetrieveStorage(storageStorage)
	assert.ErrorContains(t, err, fmt.Sprintf("500 storage '%s' does not exist", storageStorage))
}
