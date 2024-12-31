package response

type LicenseContent struct {
	HardWare     string   `json:"hardwareinfo"`
	StartTime    int64    `json:"start_time"`
	EndTime      int64    `json:"end_time"`
	Models       []string `json:"models"`
	NodesNum     int      `json:"nodes_num"`
	UsedNodesNum int      `json:"used_nodes_num"`
	Version      string   `json:"version"`
	LicenseType  string   `json:"license_type"`
	LicenseName  string   `json:"license_name"`
}
