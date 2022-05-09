package api2

import (
	"encoding/json"
	"net/url"
)

const storageBasePath = "/storage"

type StorageList []Storage

type getStorageListOption func(*url.URL)

func WithStorageTypeFilter(storageType StorageType) getStorageListOption {
	return func(apiURL *url.URL) {
		params := url.Values{}
		params.Add("type", string(storageType))
		apiURL.RawQuery = params.Encode()
	}
}

func (c *Client) GetStorageList(opts ...getStorageListOption) (*StorageList, error) {
	apiURL := *c.ApiURL
	apiURL.Path += storageBasePath

	for _, opt := range opts {
		opt(&apiURL)
	}

	resp, err := doGet(c, &apiURL)
	if err != nil {
		return nil, err
	}

	var data StorageList
	err = json.Unmarshal(resp, &data)
	return &data, err
}

func (c *Client) PostStorage(storage string, storageType StorageType, options map[string]string) (*Storage, error) {
	apiURL := *c.ApiURL
	apiURL.Path += storageBasePath
	params := url.Values{}
	params.Add("storage", storage)
	params.Add("type", string(storageType))
	for k, v := range options {
		params.Add(k, v)
	}
	apiURL.RawQuery = params.Encode()
	resp, err := doPost(c, &apiURL)
	if err != nil {
		return nil, err
	}
	var data Storage
	if err := json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	return &data, err
}
