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

func TestStorageListRetrieve(t *testing.T) {
	storageStorage1 := "a" + rand.String(10)
	storageStorage2 := "a" + rand.String(10)

	_, err := testClient.CreateStorage(storageStorage1, StorageTypeDir, map[string]string{"path": "/foo"})
	assert.NoError(t, err)
	_, err = testClient.CreateStorage(storageStorage2, StorageTypeDir, map[string]string{"path": "/bar"})
	assert.NoError(t, err)

	storageList, err := testClient.RetrieveStorageListOfType(StorageTypeDir)
	assert.NoError(t, err)
	var storageStorage1Path string
	var storageStorage2Path string
	for _, item := range *storageList {
		if item.Storage == storageStorage1 {
			storageStorage1Path = item.Path
		}
		if item.Storage == storageStorage2 {
			storageStorage2Path = item.Path
		}
	}
	assert.Equal(t, "/foo", storageStorage1Path)
	assert.Equal(t, "/bar", storageStorage2Path)
}

func TestStorageUpdate(t *testing.T) {
	storageStorage := "a" + rand.String(10)

	createOptions := map[string]string{"path": "/foo", "content": fmt.Sprintf("%s,%s", StorageContentISO, StorageContentRootDir)}
	_, err := testClient.CreateStorage(storageStorage, StorageTypeDir, createOptions)
	assert.NoError(t, err)

	_, err = testClient.UpdateStorage(storageStorage, map[string]string{"content": string(StorageContentVZTMPL)})
	assert.NoError(t, err)

	actualStorage, err := testClient.RetrieveStorage(storageStorage)
	assert.NoError(t, err)
	assert.Len(t, actualStorage.Content, 1)
	assert.Contains(t, actualStorage.Content, StorageContentVZTMPL)
}
