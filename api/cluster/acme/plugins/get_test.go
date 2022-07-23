package plugins

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	accountList, err := GetRequest{Client: helpers.APITokenTestClient()}.Get()
	assert.NoError(t, err)
	assert.NotNil(t, accountList)
	assert.Equal(t, 1, len(accountList))
	assert.Equal(t, "standalone", accountList[0].Plugin)
}
