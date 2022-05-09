package api2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestStorageCreateAndRetrieve(t *testing.T) {
	storageStorage := "a" + rand.String(10)
	actualStorage, err := testClient.PostStorage(storageStorage, StorageTypeDir, map[string]string{"path": "/foo"})
	assert.NoError(t, err)
	assert.Equal(t, storageStorage, actualStorage.Storage)
	assert.Equal(t, StorageTypeDir, actualStorage.Type)

	storageList, err := testClient.GetStorageList()
	assert.NoError(t, err)

	var storage Storage
	for _, i := range *storageList {
		if i.Storage == storageStorage {
			storage = i
		}
	}
	assert.Equal(t, storageStorage, storage.Storage)
	assert.Contains(t, storage.Content, StorageContentImages)
	assert.Contains(t, storage.Content, StorageContentRootDir)
}

func TestStorageListRetrieve(t *testing.T) {
	storageStorage1 := "a" + rand.String(10)
	storageStorage2 := "a" + rand.String(10)

	_, err := testClient.PostStorage(storageStorage1, StorageTypeDir, map[string]string{"path": "/foo"})
	assert.NoError(t, err)
	_, err = testClient.PostStorage(storageStorage2, StorageTypeDir, map[string]string{"path": "/bar"})
	assert.NoError(t, err)

	storageList, err := testClient.GetStorageList(WithStorageTypeFilter(StorageTypeDir))
	assert.NoError(t, err)
	var storageStorage1Path string
	var storageStorage2Path string
	for _, i := range *storageList {
		if i.Storage == storageStorage1 {
			storageStorage1Path = i.Path
		}
		if i.Storage == storageStorage2 {
			storageStorage2Path = i.Path
		}
	}
	assert.Equal(t, "/foo", storageStorage1Path)
	assert.Equal(t, "/bar", storageStorage2Path)
}
