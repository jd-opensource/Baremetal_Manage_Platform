package response

type Raid struct {
	// raid uuid
	RaidID string `json:"raidId"`
	// raid 名称
	Name string `json:"name"`
	//设备类型uuid
	DeviceTypeId string `json:"deviceTypeId"`
	//磁盘类型
	VolumeType string `json:"volumeType"`
	//磁盘详细信息
	VolumeDetail string `json:"volumeDetail"`
	//中文描述
	DescriptionEn string `json:"descriptionEn"`
	// description
	DescriptionZh string `json:"descriptionZh"`
	// 系统盘分区数量
	SystemPartitionCount int `json:"systemPartitionCount"`
	// 可用分区值，单位GB
	AvailableValue int `json:"availableValue"`
	// 磁盘类型：SAS/SATA/SSD/NVME
	DiskType string `json:"diskType"`
}
type RaidList struct {
	//raid实体列表
	Raids []*Raid `json:"raids"`
	//PageNumber int64   `json:"pageNumber"`
	//PageSize   int64   `json:"pageSize"`
	//TotalCount int64   `json:"totalCount"`
}
type RaidInfo struct {
	//raid实体
	Raid *Raid `json:"raid"`
}
type RaidId struct {
	//raid uuid
	RaidId string `json:"raidId"`
}
