package api2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVer(t *testing.T) {
	version, _ := testClient.GetVersion()
	assert.GreaterOrEqual(t, version.Release, "7")
}
