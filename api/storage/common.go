package storage

import (
	"encoding/json"
	"sort"
	"strings"
)

const basePath = "/storage"

const (
	TypeBTRFS       string = "btrfs"
	TypeCephFS      string = "cephfs"
	TypeCIFS        string = "cifs"
	TypeDir         string = "dir"
	TypeGlusterFS   string = "glusterfs"
	TypeISCSI       string = "iscsi"
	TypeISCSIDirect string = "iscsidirect"
	TypeLVM         string = "lvm"
	TypeLVMThin     string = "lvmthin"
	TypeNFS         string = "nfs"
	TypePBS         string = "pbs"
	TypeRBD         string = "rbd"
	TypeZFS         string = "zfs"
	TypeZFSPool     string = "zfspool"
)

const (
	ContentBackup   string = "backup"
	ContentImages   string = "images"
	ContentISO      string = "iso"
	ContentRootDir  string = "rootdir"
	ContentSnippets string = "snippets"
	ContentVZTmpl   string = "vztmpl"
)

const (
	PreAllocationOff       string = "off"
	PreAllocationMetadata  string = "metadata"
	PreAllocationFallocate string = "fallocate"
	PreAllocationFull      string = "full"
)

func listJoin(l *[]string, separator string) string {
	contentList := ""
	for i, c := range *l {
		if i == len(*l)-1 {
			contentList += string(c)
		} else {
			contentList += string(c) + separator
		}
	}
	return contentList
}

func rawListSplitAndSort(s json.RawMessage) []string {
	slice := strings.Split(strings.Trim(string(s), `"`), ",")

	// remove empty items
	clearedSlice := []string{}
	for _, v := range slice {
		if v != "" {
			clearedSlice = append(clearedSlice, v)
		}
	}

	sort.Strings(clearedSlice)
	return clearedSlice
}
