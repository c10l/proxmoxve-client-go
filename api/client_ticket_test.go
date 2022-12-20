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
	user := os.Getenv("PROXMOXVE_TEST_USER")
	pass := os.Getenv("PROXMOXVE_TEST_PASS")
	totpSeed := os.Getenv("PROXMOXVE_TEST_TOTPSEED")
	expectedInsecure, err := strconv.ParseBool(os.Getenv("PROXMOXVE_TEST_TLS_INSECURE"))
	assert.NoError(t, err)
	client, err := NewTicketClient(expectedBaseURL, user, pass, totpSeed, expectedInsecure)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%s/api2/json", expectedBaseURL), client.APIurl.String())
	assert.Equal(t, expectedInsecure, client.TLSInsecure)
	assert.Contains(t, client.Ticket.Ticket, "PVE:root@pam:")
	assert.Regexp(t, regexp.MustCompile(`[A-Z0-9]{8}:[A-Za-z0-9\+/]{43}`), client.Ticket.CSRFPreventionToken)
}
