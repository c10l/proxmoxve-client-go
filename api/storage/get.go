package storage

import (
	"bytes"
	"encoding/json"
	"net/url"
	"sort"
	"strings"

	"github.com/c10l/proxmoxve-client-go/api"
)

const basePath = "/storage"

type GetRequest struct {
	client      *api.Client
	storageType GetRequestStorageType
}

type GetRequestStorageType string

const (
	GetRequestStorageTypeBTRFS       GetRequestStorageType = "btrfs"
	GetRequestStorageTypeCephFS      GetRequestStorageType = "cephfs"
	GetRequestStorageTypeCIFS        GetRequestStorageType = "cifs"
	GetRequestStorageTypeDir         GetRequestStorageType = "dir"
	GetRequestStorageTypeGlusterFS   GetRequestStorageType = "glusterfs"
	GetRequestStorageTypeISCSI       GetRequestStorageType = "iscsi"
	GetRequestStorageTypeISCSIDirect GetRequestStorageType = "iscsidirect"
	GetRequestStorageTypeLVM         GetRequestStorageType = "lvm"
	GetRequestStorageTypeLVMThin     GetRequestStorageType = "lvmthin"
	GetRequestStorageTypeNFS         GetRequestStorageType = "nfs"
	GetRequestStorageTypePBS         GetRequestStorageType = "pbs"
	GetRequestStorageTypeRBD         GetRequestStorageType = "rbd"
	GetRequestStorageTypeZFS         GetRequestStorageType = "zfs"
	GetRequestStorageTypeZFSPool     GetRequestStorageType = "zfspool"
)

type GetResponse []GetResponseStorage
type GetResponseStorage struct {
	Content      GetResponseContentList `json:"content,omitempty"`
	Digest       string                 `json:"digest,omitempty"`
	Path         string                 `json:"path,omitempty"`
	PruneBackups string                 `json:"prune-backups,omitempty"`
	Shared       int                    `json:"shared,omitempty"`
	Storage      string                 `json:"storage,omitempty"`
	Type         GetRequestStorageType  `json:"type,omitempty"`
}

type GetResponseContentList []GetResponseContent
type GetResponseContent string

const (
	GetResponseContentVZTMPL  GetResponseContent = "vztmpl"
	GetResponseContentImages  GetResponseContent = "images"
	GetResponseContentRootDir GetResponseContent = "rootdir"
	GetResponseContentISO     GetResponseContent = "iso"
)

func (l *GetResponseContentList) UnmarshalJSON(b []byte) error {
	parts := strings.Split(string(bytes.Trim(b, `"`)), ",")
	sort.Strings(parts)
	for _, item := range parts {
		*l = append(*l, GetResponseContent(item))
	}
	return nil
}

func (g GetRequest) Do() (*GetResponse, error) {
	var s GetResponse
	apiURL := g.client.ApiURL
	apiURL.Path += basePath
	if g.storageType != "" {
		params := url.Values{}
		params.Add("type", string(g.storageType))
		apiURL.RawQuery = params.Encode()
	}
	resp, err := g.client.Get(apiURL)
	if err != nil {
		return nil, err
	}
	return &s, json.Unmarshal(resp, &s)
}
