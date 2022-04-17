package api2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient("http://foobar/", "test-token-id", "test-secret", false)
	assert.Equal(t, client.BaseURL, "http://foobar/api2/json")
	assert.Equal(t, client.TokenID, "test-token-id")
	assert.Equal(t, client.Secret, "test-secret")
	assert.Equal(t, client.TLSInsecure, false)
}
