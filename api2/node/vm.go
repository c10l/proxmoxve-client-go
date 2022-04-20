package node

type VM struct {
	Cpu       int    `json:"cpu"`
	Disk      int    `json:"disk"`
	DiskRead  int    `json:"diskread"`
	DiskWrite int    `json:"diskwrite"`
	ID        string `json:"id"`
	MaxCPU    int    `json:"maxcpu"`
	MaxDisk   int    `json:"maxdisk"`
	MaxMem    int    `json:"maxmem"`
	Mem       int    `json:"mem"`
	Name      string `json:"name"`
	NetIn     int    `json:"netin"`
	NetOut    int    `json:"netout"`
	Node      string `json:"node"`
	Status    string `json:"status"`
	Template  int    `json:"template"`
	Type      string `json:"type"`
	Uptime    int    `json:"uptime"`
	VMID      int    `json:"vmid"`
}
type Qemu VM
type LXC VM
