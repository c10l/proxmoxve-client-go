package api

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTicketClient(t *testing.T) {
	expectedBaseURL := os.Getenv("PROXMOXVE_TEST_BASE_URL")
	expectedUser := os.Getenv("PROXMOXVE_TEST_USER")
	expectedPass := os.Getenv("PROXMOXVE_TEST_PASS")
	expectedTOTPSeed := os.Getenv("PROXMOXVE_TEST_TOTPSEED")
	expectedInsecure, err := strconv.ParseBool(os.Getenv("PROXMOXVE_TEST_TLS_INSECURE"))
	assert.NoError(t, err)
	client, err := NewTicketClient(expectedBaseURL, expectedUser, expectedPass, "", expectedInsecure)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%s/api2/json", expectedBaseURL), client.APIurl.String())
	assert.Equal(t, expectedInsecure, client.TLSInsecure)
	assert.Equal(t, expectedUser, client.UserPass.User)
	assert.Equal(t, expectedPass, client.UserPass.Pass)
	assert.Equal(t, expectedTOTPSeed, client.UserPass.TOTPSeed)
	assert.Equal(t, "root@pam", client.UserPass.Ticket.Username)
	assert.Contains(t, client.UserPass.Ticket.Ticket, "PVE:root@pam:")
	assert.Regexp(t, regexp.MustCompile(`[A-Z0-9]{8}:[A-Za-z0-9\+/]{43}`), client.UserPass.Ticket.CSRFPreventionToken)
}
