package api2

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestUnmarshalPoolMemberStorage(t *testing.T) {
	expectedStorage := rand.String(10)
	storageJSON := []byte(fmt.Sprintf(`
    [{
      "content": "vztmpl,images,rootdir,iso",
      "disk": 3812077568,
      "id": "storage/pve/%s",
      "maxdisk": 8087252992,
      "node": "pve",
      "plugintype": "dir",
      "shared": 0,
      "status": "available",
      "storage": "%s",
      "type": "storage"
    }]`, expectedStorage, expectedStorage))

	var poolMembers PoolMembers
	err := json.Unmarshal(storageJSON, &poolMembers)
	assert.NoError(t, err)
	assert.Len(t, poolMembers.Storage, 1)
	assert.Len(t, poolMembers.Qemu, 0)
	assert.Len(t, poolMembers.LXC, 0)
	assert.Equal(t, PoolMemberTypeStorage, poolMembers.Storage[0].Type)
	assert.Equal(t, StorageTypeDir, poolMembers.Storage[0].PluginType)
	assert.Equal(t, expectedStorage, poolMembers.Storage[0].Storage)
}

func TestUnmarshalPoolMemberQemu(t *testing.T) {
	expectedName := rand.String(10)
	qemuJSON := []byte(fmt.Sprintf(`
    [{
      "cpu": 0,
      "disk": 0,
      "diskread": 0,
      "diskwrite": 0,
      "id": "qemu/100",
      "maxcpu": 1,
      "maxdisk": 34359738368,
      "maxmem": 2147483648,
      "mem": 0,
      "name": "%s",
      "netin": 0,
      "netout": 0,
      "node": "pve",
      "status": "stopped",
      "template": 0,
      "type": "qemu",
      "uptime": 0,
      "vmid": 100
    }]`, expectedName))

	var poolMembers PoolMembers
	err := json.Unmarshal(qemuJSON, &poolMembers)
	assert.NoError(t, err)
	assert.Len(t, poolMembers.Storage, 0)
	assert.Len(t, poolMembers.Qemu, 1)
	assert.Len(t, poolMembers.LXC, 0)
	assert.Equal(t, string(PoolMemberTypeQemu), poolMembers.Qemu[0].Type)
	assert.Equal(t, expectedName, poolMembers.Qemu[0].Name)
}

func TestUnmarshalPoolMemberLXC(t *testing.T) {
	expectedName := rand.String(10)
	lxcJSON := []byte(fmt.Sprintf(`
    [{
      "cpu": 0,
      "disk": 0,
      "diskread": 0,
      "diskwrite": 0,
      "id": "lxc/100",
      "maxcpu": 1,
      "maxdisk": 34359738368,
      "maxmem": 2147483648,
      "mem": 0,
      "name": "%s",
      "netin": 0,
      "netout": 0,
      "node": "pve",
      "status": "stopped",
      "template": 0,
      "type": "lxc",
      "uptime": 0,
      "vmid": 100
    }]`, expectedName))

	var poolMembers PoolMembers
	err := json.Unmarshal(lxcJSON, &poolMembers)
	assert.NoError(t, err)
	assert.Len(t, poolMembers.Storage, 0)
	assert.Len(t, poolMembers.Qemu, 0)
	assert.Len(t, poolMembers.LXC, 1)
	assert.Equal(t, string(PoolMemberTypeLXC), poolMembers.LXC[0].Type)
	assert.Equal(t, "lxc", poolMembers.LXC[0].Type)
	assert.Equal(t, expectedName, poolMembers.LXC[0].Name)
}

func TestUnmarshalPoolMembers(t *testing.T) {
	poolJSON := []byte(`{
  "comment": "mn4fhr2sk8nvwwl8vfmg",
  "members": [
    {
      "cpu": 0,
      "disk": 0,
      "diskread": 0,
      "diskwrite": 0,
      "id": "qemu/100",
      "maxcpu": 1,
      "maxdisk": 34359738368,
      "maxmem": 2147483648,
      "mem": 0,
      "name": "foobar",
      "netin": 0,
      "netout": 0,
      "node": "pve",
      "status": "stopped",
      "template": 0,
      "type": "qemu",
      "uptime": 0,
      "vmid": 100
    },
    {
      "cpu": 0,
      "disk": 0,
      "diskread": 0,
      "diskwrite": 0,
      "id": "lxc/101",
      "maxcpu": 1,
      "maxdisk": 8589934592,
      "maxmem": 536870912,
      "mem": 0,
      "name": "lexi",
      "netin": 0,
      "netout": 0,
      "node": "pve",
      "status": "stopped",
      "template": 0,
      "type": "lxc",
      "uptime": 0,
      "vmid": 101
    },
    {
      "content": "vztmpl,images,rootdir,iso",
      "disk": 3812077568,
      "id": "storage/pve/a7ltj8qh7k7",
      "maxdisk": 8087252992,
      "node": "pve",
      "plugintype": "dir",
      "shared": 0,
      "status": "available",
      "storage": "a7ltj8qh7k7",
      "type": "storage"
    }
  ]
}`)

	var pool Pool
	err := json.Unmarshal(poolJSON, &pool)
	assert.NoError(t, err)
	assert.Equal(t, string(PoolMemberTypeQemu), pool.Members.Qemu[0].Type)
	assert.Equal(t, string(PoolMemberTypeLXC), pool.Members.LXC[0].Type)
	assert.Equal(t, PoolMemberTypeStorage, pool.Members.Storage[0].Type)
}

func TestPoolCreateAndRetrieve(t *testing.T) {
	poolID := rand.String(10)
	expectedComment := rand.String(20)
	err := testClient.CreatePool(poolID, expectedComment)
	assert.NoError(t, err)

	pool, err := testClient.RetrievePool(poolID)
	assert.NoError(t, err)
	assert.Equal(t, poolID, pool.PoolID)
	assert.Equal(t, expectedComment, pool.Comment)
}

func TestPoolUpdateComment(t *testing.T) {
	poolID := rand.String(10)
	expectedComment := rand.String(20)

	err := testClient.CreatePool(poolID, expectedComment)
	assert.NoError(t, err)

	err = testClient.UpdatePool(poolID, &expectedComment, nil, nil, false)
	assert.NoError(t, err)

	pool, err := testClient.RetrievePool(poolID)
	assert.NoError(t, err)
	assert.Equal(t, expectedComment, pool.Comment)
}

func TestPoolDelete(t *testing.T) {
	poolID := rand.String(10)

	err := testClient.CreatePool(poolID, "")
	assert.NoError(t, err)

	_, err = testClient.RetrievePool(poolID)
	assert.NoError(t, err)

	err = testClient.DeletePool(poolID)
	assert.NoError(t, err)

	_, err = testClient.RetrievePool(poolID)
	assert.ErrorContains(t, err, fmt.Sprintf("500 pool '%s' does not exist", poolID))
}

func TestPoolListRetrieve(t *testing.T) {
	poolID1 := rand.String(10)
	poolID2 := rand.String(10)

	err := testClient.CreatePool(poolID1, "")
	assert.NoError(t, err)
	err = testClient.CreatePool(poolID2, "")
	assert.NoError(t, err)

	poolList, err := testClient.RetrievePoolList()
	assert.NoError(t, err)
	assert.Contains(t, poolList, Pool{PoolID: poolID1, Comment: ""})
	assert.Contains(t, poolList, Pool{PoolID: poolID2, Comment: ""})
}
