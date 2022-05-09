package main

import (
	"fmt"
	"os"

	"github.com/c10l/proxmoxve-client-go/api2"
)

func main() {
	baseURL := os.Getenv("PROXMOXVE_TEST_BASE_URL")
	confirmCleanup := os.Getenv("PROXMOXVE_TEST_URL_CLEANUP")
	if confirmCleanup != baseURL {
		fmt.Printf(`
!!! WARNING !!!

This is a dangerous operation.
For safety, set the environment variable PROXMOXVE_TEST_URL_CLEANUP to the base URL of the ProxMox VE server you want to clean up.

Current environment:
PROXMOXVE_TEST_BASE_URL = %s
PROXMOXVE_TEST_URL_CLEANUP = %s
`, baseURL, confirmCleanup)
		os.Exit(1)
	}

	fmt.Printf("Cleaning up %s\n\n", baseURL)

	tokenID := os.Getenv("PROXMOXVE_TEST_TOKEN_ID")
	secret := os.Getenv("PROXMOXVE_TEST_SECRET")
	c, err := api2.NewClient(
		baseURL,
		tokenID,
		secret,
		true,
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	poolList, err := c.GetPoolsList()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, item := range poolList {
		fmt.Printf("Deleting Pool %s\n", item.PoolID)
		if c.DeletePool(item.PoolID) != nil {
			fmt.Println(err)
		}
	}

	storageList, err := c.GetStorageList()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, item := range *storageList {
		fmt.Printf("Deleting Storage %s\n", item.Storage)
		if c.DeleteStorage(item.Storage) != nil {
			fmt.Println(err)
		}
	}
}
