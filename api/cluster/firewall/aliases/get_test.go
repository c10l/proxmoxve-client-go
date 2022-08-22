package aliases

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	aliasesList, err := GetRequest{Client: helpers.APITokenTestClient()}.Get()
	assert.NoError(t, err)
	assert.NotNil(t, aliasesList)
}
