package response

import (
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
)

type Raid struct {
	RaidID string `json:"raidId"`
	Name   string `json:"name"`
	//DeviceTypeId         string `json:"deviceTypeId"` //设备类型
	//VolumeType           string `json:"volumeType"`   //磁盘类型
	//VolumeDetail         string `json:"volumeDetail"` //磁盘详细信息
	DescriptionEn string `json:"descriptionEn"`
	DescriptionZh string `json:"descriptionZh"`
	//SystemPartitionCount int    `json:"systemPartitionCount"`
	//AvailableValue       int    `json:"availableValue"` // 可用分区值，单位GB
	//DiskType             string `json:"diskType"`       // 磁盘类型：SAS/SATA/SSD/NVME
}
type RaidList struct {
	Raids []*Raid `json:"raids"`
	//PageNumber int64   `json:"pageNumber"`
	//PageSize   int64   `json:"pageSize"`
	//TotalCount int64   `json:"totalCount"`
}
type RaidInfo struct {
	Raid *Raid `json:"raid"`
}
type RaidId struct {
	RaidId string `json:"raidId"`
}

type DeviceTypeRaid struct {
	sdkModels.RDeviceTypeRaid
	RaidName string `json:"raidName"`
}
