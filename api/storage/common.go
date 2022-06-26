package storage

const basePath = "/storage"

type Type string

const (
	TypeBTRFS       Type = "btrfs"
	TypeCephFS      Type = "cephfs"
	TypeCIFS        Type = "cifs"
	TypeDir         Type = "dir"
	TypeGlusterFS   Type = "glusterfs"
	TypeISCSI       Type = "iscsi"
	TypeISCSIDirect Type = "iscsidirect"
	TypeLVM         Type = "lvm"
	TypeLVMThin     Type = "lvmthin"
	TypeNFS         Type = "nfs"
	TypePBS         Type = "pbs"
	TypeRBD         Type = "rbd"
	TypeZFS         Type = "zfs"
	TypeZFSPool     Type = "zfspool"
)

type Content string

const (
	ContentVZTMPL   Content = "vztmpl"
	ContentImages   Content = "images"
	ContentRootDir  Content = "rootdir"
	ContentISO      Content = "iso"
	ContentSnippets Content = "snippets"
)
