package response

type RDeviceTypeRaid struct {
	ID uint64 `json:"-"` // ID
	// RAID uuid
	RaidID string `json:"raidId"`
	// 设备类型id
	DeviceTypeID string `json:"deviceTypeId"`
	// 磁盘类型 system/data
	VolumeType string `json:"volumeType"`
	// 磁盘详细信息
	VolumeDetail string `json:"volumeDetail"`
	// 可用分区值，单位GB
	AvailableValue int `json:"availableValue"`
	// 系统盘noraid模式真实数量;此模式多块盘只能用一块
	SystemPartitionCount int `json:"systemPartitionCount"`
	// 磁盘类型：SAS/SATA/SSD/NVME
	DiskType string `json:"diskType"`
	// 磁盘类型：SAS/SATA/SSD/NVME
	DiskInterfaceType string `json:"DiskInterfaceType"`
	// 创建者
	CreatedBy string `json:"createdBy"`
	// 更新者
	UpdatedBy string `json:"updatedBy"`
	// 创建时间
	CreatedTime string `json:"createdTime"`
	// 更新时间
	UpdatedTime string `json:"updatedTime"`
}
type RDeviceTypeRaidList struct {
	RDeviceTypeRaids []*RDeviceTypeRaid `json:"rDeviceTypeRaids"`
	PageNumber       int64              `json:"pageNumber"`
	PageSize         int64              `json:"pageSize"`
	TotalCount       int64              `json:"totalCount"`
}
type RDeviceTypeRaidInfo struct {
	RDeviceTypeRaid *RDeviceTypeRaid `json:"rDeviceTypeRaid"`
}
