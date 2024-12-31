package response

type VolumeRaid struct {
	// 卷uuid
	VolumeID string `gorm:"column:volume_id" json:"volumeId"`
	// 设备类型uuid
	DeviceTypeID string `gorm:"column:device_type_id" json:"deviceTypeId"`
	// 卷名称
	VolumeName string `gorm:"column:volume_name" json:"volumeName"`
	// 卷类型：系统卷，数据卷
	VolumeType string `gorm:"column:volume_type" json:"volumeType"`
	// 硬盘类型（SSD,HDD）
	DiskType string `gorm:"column:disk_type" json:"diskType"`
	// 接口类型（SATA,SAS,NVME,不限制）
	InterfaceType string `gorm:"column:interface_type" json:"interfaceType"`
	// 单盘大小（最小容量）
	VolumeSize float64 `gorm:"column:volume_size" json:"volumeSize"`
	// 硬盘单位（GB,TB）
	VolumeUnit string `gorm:"column:volume_unit" json:"volumeUnit"`
	// 硬盘数量（最低块数）
	VolumeAmount int `gorm:"column:volume_amount" json:"volumeAmount"`
	//支持的raid模式
	Raids []*Raid `json:"raids"`
}

type Raid struct {
	RaidCan  string `gorm:"column:raid_can" json:"raidCan"`   // RAID配置： (RAID,NO RAID)
	RaidID   string `gorm:"column:raid_id" json:"raidId"`     // RAID模式：RAID1,RIAD10等
	RaidName string `gorm:"column:raid_name" json:"raidName"` // RAID名称
	//容量大小描述
	Detail string `json:"detail"`
}
