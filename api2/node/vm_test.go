package node

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestUnmarshalQemu(t *testing.T) {
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

	qemu := new(Qemu)
	err := json.Unmarshal(qemuJSON, qemu)
	assert.NoError(t, err)
	assert.Equal(t, expectedName, qemu.Name)
}

func TestUnmarshalLXC(t *testing.T) {
	expectedName := rand.String(10)
	lxcJSON := []byte(fmt.Sprintf(`
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
      "name": "%s",
      "netin": 0,
      "netout": 0,
      "node": "pve",
      "status": "stopped",
      "template": 0,
      "type": "lxc",
      "uptime": 0,
      "vmid": 101
    }`, expectedName))

	lxc := new(LXC)
	err := json.Unmarshal(lxcJSON, lxc)
	assert.NoError(t, err)
	assert.Equal(t, expectedName, lxc.Name)
}
