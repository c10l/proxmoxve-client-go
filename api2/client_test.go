package api2

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	expectedBaseURL := os.Getenv("PROXMOXVE_TEST_BASE_URL")
	expectedTokenID := os.Getenv("PROXMOXVE_TEST_TOKEN_ID")
	expectedSecret := os.Getenv("PROXMOXVE_TEST_SECRET")
	expectedInsecure, err := strconv.ParseBool(os.Getenv("PROXMOXVE_TEST_TLS_INSECURE"))
	assert.NoError(t, err)
	client, err := NewClient(expectedBaseURL, expectedTokenID, expectedSecret, expectedInsecure)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%s/api2/json", expectedBaseURL), client.ApiURL.String())
	assert.Equal(t, expectedTokenID, client.TokenID)
	assert.Equal(t, expectedSecret, client.Secret)
	assert.Equal(t, expectedInsecure, client.TLSInsecure)
}
