package account

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/api/test"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	accountList, err := GetRequest{Client: test.APITokenTestClient()}.Get()
	assert.NoError(t, err)
	assert.NotNil(t, accountList)
}
