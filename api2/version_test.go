package api2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveVersion(t *testing.T) {
	actualVersion, _ := testClient.RetrieveVersion()
	assert.GreaterOrEqual(t, actualVersion.Release, "7")
}
