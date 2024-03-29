package api

import (
	"testing"

	"github.com/c10l/proxmoxve-client-go/helpers"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	version, err := GetRequest{Client: helpers.APITokenTestClient()}.Do()
	assert.NoError(t, err)
	assert.NotNil(t, version)
	assert.Regexp(t, `\d\.\d-\d`, version.Version)
	assert.Regexp(t, `.{8}`, version.RepoID)
	assert.Regexp(t, `\d\.\d`, version.Release)
}
