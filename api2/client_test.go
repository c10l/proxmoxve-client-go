package api2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client, _ := NewClient("http://foobar/", "test-token-id", "test-secret", false)
	assert.Equal(t, "http://foobar/api2/json", client.ApiURL.String())
	assert.Equal(t, "test-token-id", client.TokenID)
	assert.Equal(t, "test-secret", client.Secret)
	assert.Equal(t, false, client.TLSInsecure)
}
