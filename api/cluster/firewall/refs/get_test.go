package refs

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	refs, err := GetRequest{Client: helpers.APITokenTestClient()}.Do()
	assert.NoError(t, err)
	assert.NotNil(t, refs)
}
