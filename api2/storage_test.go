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
      "content": "vztmpl,images,rootdir,iso",
      "disk": 3812077568,
      "id": "storage/pve/a7ltj8qh7k7",
      "maxdisk": 8087252992,
      "node": "pve",
      "plugintype": "dir",
      "shared": 0,
      "status": "available",
      "storage": "%s",
      "type": "storage"
    }`, expectedStorage))

	storage := new(Storage)
	err := json.Unmarshal(storageJSON, storage)
	assert.NoError(t, err)
	assert.Equal(t, expectedStorage, storage.Storage)
}
