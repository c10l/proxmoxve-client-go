package api2

import (
	"bytes"
	"encoding/json"
	"net/url"
	"sort"
	"strings"
)

const storageBasePath = "/storage"

type StorageList []Storage

type Storage struct {
	Content StorageContentList `json:"content"`
	Digest  string             `json:"digest"`
	Path    string             `json:"path"`
	Storage string             `json:"storage"`
	Type    StorageType        `json:"type"`
	Config  json.RawMessage    `json:"config,omitempty"`
}

type StorageType string

const (
	StorageTypeBTRFS       StorageType = "btrfs"
	StorageTypeCephFS      StorageType = "cephfs"
	StorageTypeCIFS        StorageType = "cifs"
	StorageTypeDir         StorageType = "dir"
	StorageTypeGlusterFS   StorageType = "glusterfs"
	StorageTypeiSCSI       StorageType = "iscsi"
	StorageTypeiSCSIDirect StorageType = "iscsidirect"
	StorageTypeLVM         StorageType = "lvm"
	StorageTypeLVMThin     StorageType = "lvmthin"
	StorageTypeNFS         StorageType = "nfs"
	StorageTypePBS         StorageType = "pbs"
	StorageTypeRBD         StorageType = "rbd"
	StorageTypeZFS         StorageType = "zfs"
	StorageTypeZFSPool     StorageType = "zfspool"
)

type StorageContentList []StorageContent
type StorageContent string

const (
	StorageContentVZTMPL  StorageContent = "vztmpl"
	StorageContentImages  StorageContent = "images"
	StorageContentRootDir StorageContent = "rootdir"
	StorageContentISO     StorageContent = "iso"
)

func (scl *StorageContentList) UnmarshalJSON(b []byte) error {
	parts := strings.Split(string(bytes.Trim(b, `"`)), ",")
	sort.Strings(parts)
	for _, item := range parts {
		*scl = append(*scl, StorageContent(item))
	}
	return nil
}

type retrieveStorageListOption func(*url.URL)

func WithTypeFilter(storageType StorageType) retrieveStorageListOption {
	return func(apiURL *url.URL) {
		params := url.Values{}
		params.Add("type", string(storageType))
		apiURL.RawQuery = params.Encode()
	}
}

func (c *Client) RetrieveStorageList(opts ...retrieveStorageListOption) (*StorageList, error) {
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

func (c *Client) RetrieveStorage(storage string) (*Storage, error) {
	apiURL := *c.ApiURL
	apiURL.Path += storageBasePath
	apiURL.Path += "/" + storage
	resp, err := doGet(c, &apiURL)
	if err != nil {
		return nil, err
	}
	var data Storage
	err = json.Unmarshal(resp, &data)
	return &data, err
}

func (c *Client) CreateStorage(storage string, storageType StorageType, options map[string]string) (*Storage, error) {
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

func (c *Client) DeleteStorage(storage string) error {
	apiURL := *c.ApiURL
	apiURL.Path += storageBasePath
	apiURL.Path += "/" + storage
	_, err := doDelete(c, &apiURL)
	return err
}

func (c *Client) UpdateStorage(storage string, options map[string]string) (*Storage, error) {
	apiURL := *c.ApiURL
	apiURL.Path += storageBasePath
	apiURL.Path += "/" + storage
	params := url.Values{}
	for k, v := range options {
		params.Add(k, v)
	}
	apiURL.RawQuery = params.Encode()
	resp, err := doPut(c, &apiURL)
	if err != nil {
		return nil, err
	}
	var data Storage
	if err := json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	return &data, err
}
