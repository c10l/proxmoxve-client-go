package api2

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveVersion(t *testing.T) {
	var v []byte
	resp, err := testClient.RetrieveVersion()
	assert.NoError(t, err)
	v, err = io.ReadAll(resp)
	assert.NoError(t, err)
	// Because the API doesn't guarantee the order in which the data keys are returned,
	// we test each of them individually.
	assert.Regexp(t, `{"data":{.*"version":"\d\.\d-\d".*}`, string(v))
	assert.Regexp(t, `{"data":{.*"repoid":".{8}".*}`, string(v))
	assert.Regexp(t, `{"data":{.*"release":"\d\.\d".*}`, string(v))
}
