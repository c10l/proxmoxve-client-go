package api2

import (
	"bytes"
	"net/url"
	"sort"
	"strings"
)

type StorageList []Storage

type StorageBackend string

const (
	StorageBackendBTRFS       = "btrfs"
	StorageBackendCeph        = "cephfs"
	StorageBackendCIFS        = "cifs"
	StorageBackendDirectory   = "dir"
	StorageBackendGlusterFS   = "glusterfs"
	StorageBackendiSCSI       = "iscsi"
	StorageBackendiSCSIDirect = "iscsidirect"
	StorageBackendLVM         = "lvm"
	StorageBackendLVMThin     = "lvmthin"
	StorageBackendNFS         = "nfs"
	StorageBackendPBS         = "pbs"
	StorageBackendRBD         = "rbd"
	StorageBackendZFS         = "zfs"
	StorageBackendZFSPool     = "zfspool"
)

type StoragePreAllocationType string

const (
	StoragePreAllocationTypeOff      = "off"
	StoragePreAllocationTypeMetadata = "metadata"
	StoragePreAllocationTypeFalloc   = "falloc"
	StoragePreAllocationTypeFull     = "full"
)

type StorageSMBVersionType string

const (
	StorageSMBVersionTypeDefault = "default"
	StorageSMBVersionType2_0     = "2.0"
	StorageSMBVersionType2_1     = "2.1"
	StorageSMBVersionType3       = "3"
	StorageSMBVersionType3_0     = "3.0"
	StorageSMBVersionType3_11    = "3.11"
)

type StorageGlusterFSTransportType string

const (
	StorageGlusterFSTransportTypeTCP  = "tcp"
	StorageGlusterFSTransportTypeRDMA = "rdma"
	StorageGlusterFSTransportTypeUnix = "unix"
)

type StorageContentList []string

func (cl *StorageContentList) UnmarshalJSON(b []byte) error {
	parts := strings.Split(string(bytes.Trim(b, `"`)), ",")
	sort.Strings(parts)
	*cl = parts
	return nil
}

type StorageContentType string

const (
	StorageContentTypeVZTMPL  = "vztmpl"
	StorageContentTypeISO     = "iso"
	StorageContentTypeBackup  = "backup"
	StorageContentTypeImages  = "images"
	StorageContentTypeRootDir = "rootdir"
)

type Storage struct {
	ID                 string              `json:"storage"`
	Backend            StorageBackend      `json:"type"`
	AuthSupported      *string             `json:"authsupported,omitempty"`
	Base               *string             `json:"base,omitempty"`
	BlockSize          *string             `json:"blocksize,omitempty"`
	BandwidthLimit     *string             `json:"bwlimit,omitempty"`
	ComstarHostGroup   *string             `json:"comstarhg,omitempty"`
	ComstarTargetGroup *string             `json:"comstartg,omitempty"`
	Content            *StorageContentList `json:"content,omitempty"`
	Datastore          *string             `json:"datastore,omitempty"`
	Disable            *bool               `json:"disable,omitempty"`
	CIFSDomain         *string             `json:"domain,omitempty"`

	// EncryptionKey Encryption key. Use 'autogen' to generate one automatically without passphrase.
	EncryptionKey *string `json:"encryption-key,omitempty"`

	NFSExportPath          *string `json:"export,omitempty"`
	CertificateFingerprint *string `json:"fingerprint,omitempty"`
	DefaultImageFormat     *string `json:"format,omitempty"`
	CephFilesystemName     *string `json:"fs-name,omitempty"`
	MountCephFSThroughFUSE *bool   `json:"fuse,omitempty"`
	IsMountpoint           *string `json:"is-mountpoint,omitempty"`
	IscsiProvider          *string `json:"iscsiprovider,omitempty"`
	ClientKeyring          *string `json:"keyring,omitempty"`
	Krbd                   *bool   `json:"krbd,omitempty"`
	LIOTargetPortalGroup   *string `json:"liotpg,omitempty"`

	// MasterPubKey Base64-encoded, PEM-formatted public RSA key.
	// Used to encrypt a copy of the encryption-key which will be added to each encrypted backup.
	MasterPubKey *string `json:"master-pubkey,omitempty"`

	// MkDir Default: true
	MkDir *bool `json:"mkdir,omitempty"`

	// MonitorHost IP addresses of monitors (for external clusters).
	MonitorHost *string `json:"monhost,omitempty"`

	MountPoint   *string `json:"mountpoint,omitempty"`
	RBDNamespace *string `json:"namespace,omitempty"`

	// NoCOW Set the NOCOW flag on files.
	// Disables data checksumming and causes data errors to be unrecoverable from while allowing direct I/O.
	// Only use this if data does not need to be any more safe than on a single ext4 formatted disk with no underlying raid system.
	NoCOW *bool `json:"nocow,omitempty"`

	Nodes           *string `json:"nodes,omitempty"`
	NoWriteCache    *bool   `json:"nowritecache,omitempty"`
	NFSMountOptions *string `json:"options,omitempty"`
	Password        *string `json:"password,omitempty"`
	FilesystemPath  *string `json:"path,omitempty"`
	Pool            *string `json:"pool,omitempty"`

	// Port Default: 8007
	Port *int `json:"port,omitempty"`

	IscsiPortal *string `json:"portal,omitempty"`

	// PreAllocation Preallocation mode for raw and qcow2 images.
	// Using 'metadata' on raw images results in preallocation=off.
	PreAllocation *StoragePreAllocationType `json:"preallocation,omitempty"`

	PruneBackup *string `json:"prune-backup,omitempty"`

	// SafeRemove Zero-out data when removing LVs.
	SafeRemove *bool `json:"saferemove,omitempty"`

	// Server Server IP or DNS name.
	Server *string `json:"server,omitempty"`

	// Server2 Backup volfile server IP or DNS name.
	Server2 *string `json:"server2,omitempty"`

	CIFSShare  *string                `json:"share,omitempty"`
	Shared     *int                   `json:"shared,omitempty"`
	SMBVersion *StorageSMBVersionType `json:"smbversion,omitempty"`
	Sparse     *bool                  `json:"sparse,omitempty"`
	Subdir     *string                `json:"subdir,omitempty"`

	// TaggedOnly Only use logical volumes tagged with 'pve-vm-ID'.
	TaggedOnly *bool `json:"tagged_only,omitempty"`

	IscsiTarget        *string                        `json:"target,omitempty"`
	LVMThinPool        *string                        `json:"thinpool,omitempty"`
	GlusterFSTransport *StorageGlusterFSTransportType `json:"transport,omitempty"`
	RBDUsername        *string                        `json:"username,omitempty"`
	VGName             *string                        `json:"vgname,omitempty"`
	GlusterFSVolume    *string                        `json:"volume,omitempty"`
}

const storageBasePath = "/storage"

func (c *Client) RetrieveStorageList() (*StorageList, error) {
	data := new(StorageList)
	apiURL := *c.ApiURL
	apiURL.Path += storageBasePath
	err := doGet(c, data, &apiURL)
	return data, err
}

func (c *Client) RetrieveStorage(storageID string) (*Storage, error) {
	data := new(Storage)
	apiURL := *c.ApiURL
	apiURL.Path += storageBasePath
	apiURL.Path += "/" + storageID
	err := doGet(c, data, &apiURL)
	return data, err
}

func (c *Client) CreateStorage(storageID, backend string, options map[string]string) error {
	apiURL := *c.ApiURL
	apiURL.Path += storageBasePath
	params := url.Values{}
	params.Add(getTagByFieldName("ID", Storage{}), storageID)
	params.Add(getTagByFieldName("Backend", Storage{}), backend)
	for k, v := range options {
		params.Add(k, v)
	}
	apiURL.RawQuery = params.Encode()
	err := doPost(c, new(Storage), &apiURL)
	return err
}

func (c *Client) DeleteStorage(storageID string) error {
	apiURL := *c.ApiURL
	apiURL.Path += storageBasePath
	apiURL.Path += "/" + storageID
	return doDelete(c, &apiURL)
}

func (c *Client) UpdateStorage(storageID string, options map[string]string) error {
	apiURL := *c.ApiURL
	apiURL.Path += storageBasePath
	apiURL.Path += "/" + storageID
	params := url.Values{}
	for k, v := range options {
		params.Add(k, v)
	}
	apiURL.RawQuery = params.Encode()
	return doPut(c, &apiURL)
}
