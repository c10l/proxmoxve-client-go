package api2

import (
	"io"
	"net/url"
	"strings"
)

const poolsBasePath = "/pools"

func (c *Client) RetrievePoolList() (io.Reader, error) {
	url := *c.ApiURL
	url.Path += poolsBasePath
	resp, err := doGet(c, &url)
	data := strings.NewReader(string(resp))
	return data, err
}

func (c *Client) CreatePool(poolID, comment string) (io.Reader, error) {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	params := url.Values{}
	params.Add("poolid", poolID)
	params.Add("comment", comment)
	apiURL.RawQuery = params.Encode()
	resp, err := doPost(c, &apiURL)
	data := strings.NewReader(string(resp))
	return data, err
}

func (c *Client) RetrievePool(poolID string) (io.Reader, error) {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	apiURL.Path += "/" + poolID
	resp, err := doGet(c, &apiURL)
	data := strings.NewReader(string(resp))
	return data, err
}

func (c *Client) DeletePool(poolID string) (io.Reader, error) {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	apiURL.Path += "/" + poolID
	resp, err := doDelete(c, &apiURL)
	data := strings.NewReader(string(resp))
	return data, err
}

func (c *Client) UpdatePool(poolID string, comment *string, storage, vms *[]string, delete bool) (io.Reader, error) {
	apiURL := *c.ApiURL
	apiURL.Path += poolsBasePath
	apiURL.Path += "/" + poolID

	params := url.Values{}
	if comment != nil {
		params.Add("comment", *comment)
	}
	if storage != nil {
		params.Add("storage", strings.Join(*storage, ","))
	}
	if vms != nil {
		params.Add("vms", strings.Join(*vms, ","))
	}
	if storage != nil || vms != nil {
		if delete {
			params.Add("delete", "1")
		}
	}
	apiURL.RawQuery = params.Encode()

	resp, err := doPut(c, &apiURL)
	data := strings.NewReader(string(resp))
	return data, err
}
