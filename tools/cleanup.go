package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c10l/proxmoxve-client-go/api"
	"github.com/c10l/proxmoxve-client-go/api/storage"
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
	c, err := api.NewAPITokenClient(
		baseURL,
		tokenID,
		secret,
		true,
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	storageList, err := storage.GetRequest{Client: c}.Do()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, item := range *storageList {
		if !strings.HasPrefix(item.Storage, "pmvetest_") {
			continue
		}
		fmt.Printf("Deleting Storage %s\n", item.Storage)
		delReq := storage.ItemDeleteRequest{Client: c, Storage: item.Storage}
		if delReq.Do() != nil {
			fmt.Println(err)
		}
	}
	// clusterAcmeAccountList, err := account.GetRequest{Client: c}.Do()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }
	// for _, item := range *clusterAcmeAccountList {
	// 	if !strings.HasPrefix(item.Name, "pmvetest_") {
	// 		continue
	// 	}
	// 	fmt.Printf("Deleting Cluster ACME Account %s\n", item.Name)
	// 	delReq := account.ItemDeleteRequest{Client: c, Name: item.Name}
	// 	if delReq.Do() != nil {
	// 		fmt.Println(err)
	// 	}
	// }
}
