package api2

import (
	"io"
	"net/url"
	"strings"
)

const storageBasePath = "/storage"

func (c *Client) RetrieveStorageList() (io.Reader, error) {
	apiURL := *c.ApiURL
	apiURL.Path += storageBasePath
	resp, err := doGet(c, &apiURL)
	data := strings.NewReader(string(resp))
	return data, err
}

func (c *Client) RetrieveStorage(storage string) (io.Reader, error) {
	apiURL := *c.ApiURL
	apiURL.Path += storageBasePath
	apiURL.Path += "/" + storage
	resp, err := doGet(c, &apiURL)
	data := strings.NewReader(string(resp))
	return data, err
}

func (c *Client) CreateStorage(storage, storageType string, options map[string]string) (io.Reader, error) {
	apiURL := *c.ApiURL
	apiURL.Path += storageBasePath
	params := url.Values{}
	params.Add("storage", storage)
	params.Add("type", storageType)
	for k, v := range options {
		params.Add(k, v)
	}
	apiURL.RawQuery = params.Encode()
	resp, err := doPost(c, &apiURL)
	data := strings.NewReader(string(resp))
	return data, err
}

func (c *Client) DeleteStorage(storage string) (io.Reader, error) {
	apiURL := *c.ApiURL
	apiURL.Path += storageBasePath
	apiURL.Path += "/" + storage
	resp, err := doDelete(c, &apiURL)
	data := strings.NewReader(string(resp))
	return data, err
}

func (c *Client) UpdateStorage(storage string, options map[string]string) (io.Reader, error) {
	apiURL := *c.ApiURL
	apiURL.Path += storageBasePath
	apiURL.Path += "/" + storage
	params := url.Values{}
	for k, v := range options {
		params.Add(k, v)
	}
	apiURL.RawQuery = params.Encode()
	resp, err := doPut(c, &apiURL)
	data := strings.NewReader(string(resp))
	return data, err
}
