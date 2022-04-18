package api2

import (
	"fmt"
	"os"
)

var testClient *Client

func init() {
	baseURL := os.Getenv("PROXMOXVE_TEST_BASE_URL")
	tokenID := os.Getenv("PROXMOXVE_TEST_TOKEN_ID")
	secret := os.Getenv("PROXMOXVE_TEST_SECRET")
	if baseURL == "" || tokenID == "" || secret == "" {
		fmt.Println("test environment not setup")
		fmt.Println("set environment variables:")
		fmt.Println("  - PROXMOXVE_TEST_BASE_URL")
		fmt.Println("  - PROXMOXVE_TEST_TOKEN_ID")
		fmt.Println("  - PROXMOXVE_TEST_SECRET")
		os.Exit(1)
	}
	var err error
	testClient, err = NewClient(
		baseURL,
		tokenID,
		secret,
		true,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
