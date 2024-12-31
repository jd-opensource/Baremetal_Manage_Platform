package dao

type GetDeviceStatusRequest struct {
	IdcId string `json:"idc_id"`
	Sn    string `json:"sn"`
	PageRequest
	// ExportType 非空要导出
	ExportType string `json:"exportType"`
}

type GetDeviceStatusResponse struct {
	Detail []*DeviceStatus `json:"detail"`
	PageResult
}

type DeviceStatus struct {
	DeviceItem
	StatusItem
}

type DeviceItem struct {
	IdcId    string `json:"idc_id"`
	IdcName  string `json:"idc_name"`
	Sn       string `json:"sn"`
	DeviceId string `json:"device_id"`
	Brand    string `json:"brand"`
	// Manufacturer string `json:"manufacturer"`
	Model string `json:"model"`
	// IsSale int    `json:"is_sale"`
	ManageStatus string `json:"manage_status"`
}

type StatusItem struct {
	CpuStatus   int `json:"cpu_status"`
	MemStatus   int `json:"mem_status"`
	DiskStatus  int `json:"disk_status"`
	NicStatus   int `json:"nic_status"`
	PowerStatus int `json:"power_status"`
	OtherStatus int `json:"other_status"`
}

type GetPowerStatusDiffRequest struct {
	IdcId      string `json:"idc_id"`
	InstanceId string `json:"instanceId"`
}

type PowerStatusSyncRequest struct {
	IdcId      string `json:"idc_id"`
	InstanceId string `json:"instanceId"`
}
