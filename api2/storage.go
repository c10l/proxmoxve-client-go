package api2

import (
	"bytes"
	"io"
	"net/url"
	"sort"
	"strings"
)

const storageBasePath = "/storage"

type Storage struct {
	Content    StorageContent `json:"content"`
	Disk       int            `json:"disk"`
	ID         string         `json:"id"`
	MaxDisk    int            `json:"maxdisk"`
	Node       string         `json:"node"`
	PluginType string         `json:"plugintype"`
	Shared     int            `json:"shared"`
	Status     string         `json:"status"`
	Storage    string         `json:"storage"`
	Type       string         `json:"type"`
}

type StorageContent []string

const (
	StorageContentVZTMPL  = "vztmpl"
	StorageContentImages  = "images"
	StorageContentRootDir = "rootdir"
	StorageContentISO     = "iso"
)

func (sc *StorageContent) UnmarshalJSON(b []byte) error {
	parts := strings.Split(string(bytes.Trim(b, `"`)), ",")
	sort.Strings(parts)
	*sc = parts
	return nil
}

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
