package response

type LicenseToken struct {
	Msg string `json:"msg"`
}

type LicenseContent struct {
	HardWare     string       `json:"hardwareinfo"`
	StartTime    int64        `json:"start_time"`
	EndTime      int64        `json:"end_time"`
	Models       []string     `json:"models"`
	NodesNum     int          `json:"nodes_num"`
	UsedNodesNum int          `json:"used_nodes_num"`
	Version      string       `json:"version"`
	LicenseType  string       `json:"license_type"`
	LicenseName  string       `json:"license_name"`
	ModelList    []ModuleItem `json:"modelList"`
}

type ModuleItem struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	NameEn        string `json:"nameEn"`
	Description   string `json:"description"`
	DescriptionEn string `json:"descriptionEn"`
	Effective     bool   `json:"effective"`
}
