package groups

import (
	"encoding/json"

	"github.com/c10l/proxmoxve-client-go/api"
	"github.com/c10l/proxmoxve-client-go/helpers/types"
)

type ItemGetRequest struct {
	Client *api.Client
	Group  string
}

type ItemGetResponse []struct {
	Action string `json:"action"`
	Pos    int    `json:"pos"`
	Type   string `json:"type"`

	// Optional
	Comment   string             `json:"comment"`
	Dest      string             `json:"dest"`
	DPort     string             `json:"dport"`
	Enable    types.PVEBool      `json:"enable"`
	ICMPType  string             `json:"icmp-type"`
	Iface     string             `json:"iface"`
	IPVersion int                `json:"ipversion"`
	Log       ItemGetResponseLog `json:"log"`
	Macro     string             `json:"macro"`
	Proto     string             `json:"proto"`
	Source    string             `json:"source"`
	SPort     string             `json:"sport"`
}

type ItemGetResponseLog string

const (
	ItemGetResponseLogEmerg   ItemGetResponseLog = "emerg"
	ItemGetResponseLogAlert   ItemGetResponseLog = "alert"
	ItemGetResponseLogCrit    ItemGetResponseLog = "crit"
	ItemGetResponseLogErr     ItemGetResponseLog = "err"
	ItemGetResponseLogWarning ItemGetResponseLog = "warning"
	ItemGetResponseLogNotice  ItemGetResponseLog = "notice"
	ItemGetResponseLogInfo    ItemGetResponseLog = "info"
	ItemGetResponseLogDebug   ItemGetResponseLog = "debug"
	ItemGetResponseLogNoLog   ItemGetResponseLog = "nolog"
)

// GetItem satisfies the ItemGetter interface
// Not to be used directly. Use Get() instead.
func (g ItemGetRequest) GetItem() ([]byte, error) {
	return g.Client.GetItem(g, basePath, g.Group)
}

func (g ItemGetRequest) Get() (*ItemGetResponse, error) {
	item, err := g.GetItem()
	if err != nil {
		return nil, err
	}
	resp := new(ItemGetResponse)
	return resp, json.Unmarshal(item, resp)
}
