package api2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveVersion(t *testing.T) {
	version, err := testClient.RetrieveVersion()
	assert.NoError(t, err)
	assert.Regexp(t, `\d\.\d-\d`, version.Version)
	assert.Regexp(t, `.{8}`, version.RepoID)
	assert.Regexp(t, `\d\.\d`, version.Release)
}
