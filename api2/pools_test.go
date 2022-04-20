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

	var poolMember PoolMember
	err := json.Unmarshal(storageJSON, &poolMember)
	assert.NoError(t, err)
	assert.Equal(t, PoolMemberTypeStorage, poolMember.Type)
	assert.Equal(t, "storage", poolMember.Storage.Type)
	assert.Equal(t, expectedStorage, poolMember.Storage.Storage)
}

func TestUnmarshalPoolMemberQemu(t *testing.T) {
	expectedName := rand.String(10)
	qemuJSON := []byte(fmt.Sprintf(`
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
      "name": "%s",
      "netin": 0,
      "netout": 0,
      "node": "pve",
      "status": "stopped",
      "template": 0,
      "type": "qemu",
      "uptime": 0,
      "vmid": 100
    }`, expectedName))

	var poolMember PoolMember
	err := json.Unmarshal(qemuJSON, &poolMember)
	assert.NoError(t, err)
	assert.Equal(t, PoolMemberTypeQemu, poolMember.Type)
	assert.Equal(t, "qemu", poolMember.Qemu.Type)
	assert.Equal(t, expectedName, poolMember.Qemu.Name)
}

func TestUnmarshalPoolMemberLXC(t *testing.T) {
	expectedName := rand.String(10)
	lxcJSON := []byte(fmt.Sprintf(`
    {
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
    }`, expectedName))

	var poolMember PoolMember
	err := json.Unmarshal(lxcJSON, &poolMember)
	assert.NoError(t, err)
	assert.Equal(t, PoolMemberTypeLXC, poolMember.Type)
	assert.Equal(t, "lxc", poolMember.LXC.Type)
	assert.Equal(t, expectedName, poolMember.LXC.Name)
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
	assert.Equal(t, PoolMemberTypeQemu, pool.Members[0].Type)
	assert.Equal(t, PoolMemberTypeLXC, pool.Members[1].Type)
	assert.Equal(t, PoolMemberTypeStorage, pool.Members[2].Type)
}

func TestPoolCreateAndRetrieve(t *testing.T) {
	poolID := rand.String(10)
	expectedComment := rand.String(20)
	err := testClient.CreatePool(poolID, expectedComment)
	assert.NoError(t, err)

	pool, err := testClient.RetrievePool(poolID)
	assert.NoError(t, err)
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
