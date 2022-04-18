package api2

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestRetrieveStorage(t *testing.T) {
	actualStorage, retrieveStorageError := testClient.RetrieveStorage("local")
	assert.NoError(t, retrieveStorageError)
	actualStorageContent := *actualStorage.Content
	expectedStorageContent := StorageContentList{StorageContentTypeBackup, StorageContentTypeISO, StorageContentTypeVZTMPL}
	assert.Equal(t, expectedStorageContent, actualStorageContent)
}

func TestCreateStorageAndRetrieveStorage(t *testing.T) {
	storageID := "a" + rand.String(10) // This must begin with a letter
	expectedStorageBackend := StorageBackendDirectory
	expectedStoragePath := "/" + rand.String(10)
	expectedStorageContent := StorageContentList{StorageContentTypeBackup, StorageContentTypeImages}
	options := map[string]string{
		getTagByFieldName("FilesystemPath", Storage{}): expectedStoragePath,
		getTagByFieldName("Content", Storage{}):        strings.Join(expectedStorageContent, ","),
	}
	assert.NoError(t, testClient.CreateStorage(storageID, expectedStorageBackend, options))

	actualStorage, retrieveStorageError := testClient.RetrieveStorage(storageID)
	assert.NoError(t, retrieveStorageError)
	actualStorageContent := *actualStorage.Content
	assert.Equal(t, expectedStorageContent, actualStorageContent)
}

func TestRetrieveStorageList(t *testing.T) {
	actualStorageList, retrieveStorageListError := testClient.RetrieveStorageList()
	assert.NoError(t, retrieveStorageListError)
	assert.NotEmpty(t, actualStorageList)
}

func TestDeleteStorage(t *testing.T) {
	storageID := "a" + rand.String(10)
	options := map[string]string{
		getTagByFieldName("FilesystemPath", Storage{}): "/" + rand.String(10),
	}
	assert.NoError(t, testClient.CreateStorage(storageID, StorageBackendDirectory, options))
	assert.NoError(t, testClient.DeleteStorage(storageID))
	_, retrieveStorageError := testClient.RetrieveStorage(storageID)
	assert.ErrorContains(t, retrieveStorageError, fmt.Sprintf("500 storage '%s' does not exist", storageID))
}

func TestUpdateStoragePath(t *testing.T) {
	storageID := "a" + rand.String(10)
	expectedContent := StorageContentList{StorageContentTypeImages, StorageContentTypeISO}
	options := map[string]string{
		getTagByFieldName("FilesystemPath", Storage{}): "/" + rand.String(10),
		getTagByFieldName("Content", Storage{}):        strings.Join(StorageContentList{StorageContentTypeBackup}, ","),
	}
	assert.NoError(t, testClient.CreateStorage(storageID, StorageBackendDirectory, options))

	updatedOptions := map[string]string{
		getTagByFieldName("Content", Storage{}): strings.Join(expectedContent, ","),
	}
	assert.NoError(t, testClient.UpdateStorage(storageID, updatedOptions))
	actualStorage, _ := testClient.RetrieveStorage(storageID)
	assert.Equal(t, expectedContent, *actualStorage.Content)
}
