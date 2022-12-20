package api

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAPITokenClient(t *testing.T) {
	expectedBaseURL := os.Getenv("PROXMOXVE_TEST_BASE_URL")
	expectedTokenID := os.Getenv("PROXMOXVE_TEST_TOKEN_ID")
	expectedSecret := os.Getenv("PROXMOXVE_TEST_SECRET")
	expectedInsecure, err := strconv.ParseBool(os.Getenv("PROXMOXVE_TEST_TLS_INSECURE"))
	assert.NoError(t, err)
	client, err := NewAPITokenClient(expectedBaseURL, expectedTokenID, expectedSecret, expectedInsecure)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%s/api2/json", expectedBaseURL), client.APIurl.String())
	assert.Equal(t, expectedTokenID, client.APIToken.TokenID)
	assert.Equal(t, expectedSecret, client.APIToken.Secret)
	assert.Equal(t, expectedInsecure, client.TLSInsecure)
}
