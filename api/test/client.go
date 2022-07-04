package test

import (
	"fmt"
	"os"

	"github.com/c10l/proxmoxve-client-go/api"
)

func APITokenTestClient() *api.Client {
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
	testClient, err := api.NewAPITokenClient(
		baseURL,
		tokenID,
		secret,
		true,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return testClient
}

func TicketTestClient() *api.Client {
	baseURL := os.Getenv("PROXMOXVE_TEST_BASE_URL")
	user := os.Getenv("PROXMOXVE_TEST_USER")
	pass := os.Getenv("PROXMOXVE_TEST_PASS")
	if baseURL == "" || user == "" || pass == "" {
		fmt.Println("test environment not setup")
		fmt.Println("set environment variables:")
		fmt.Println("  - PROXMOXVE_TEST_BASE_URL")
		fmt.Println("  - PROXMOXVE_TEST_USER")
		fmt.Println("  - PROXMOXVE_TEST_PASS")
		os.Exit(1)
	}
	var err error
	testClient, err := api.NewTicketClient(
		baseURL,
		user,
		pass,
		true,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return testClient
}
