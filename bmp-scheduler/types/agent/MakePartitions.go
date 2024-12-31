package agent

import (
	"encoding/json"

	domainTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/domain"
)

type MakePartitions struct {
	Action string `json:"action"`
	Sn     string `json:"sn"`
	// OsType    string `json:"os_type"`
	// OsVersion string `json:"os_version"`
	Version  string `json:"version"` // "2.0"
	BootMode string `json:"boot_mode"`

	AutoCreateEfi      bool `json:"auto_create_efi"`       ////可选，默认:true。uefi模式下自动创建efi分区
	AutoCreateBiosGrub bool `json:"auto_create_bios_grub"` //可选，默认:true。bios模式下自动创建bios_grub分区

	Disks []Disk `json:"volumes"` //其实叫disks更合适，老板非要叫volumes

	// DiskLabel        string                       `json:"disk_label"` //系统盘label
	// DataDiskLabel    string                       `json:"data_disk_label"`
	// PartitionRoot    *domainTypes.Partition       `json:"partition_root,omitempty"`
	// PartitionBoot    *domainTypes.Partition       `json:"partition_boot,omitempty"`
	// SystemPartitions []domainTypes.Partition      `json:"system_partitions"`
	// DataPartitions   []domainTypes.Partition      `json:"data_partitions"`
	// RootDeviceHints  *domainTypes.RootDeviceHints `json:"root_device_hints,omitempty"` //经测试，缺省值给到agent，agent会报错,所以omitempty属性
}

type Disk struct {
	KeepData     bool                         `json:"keep_data"`
	IsRootDevice bool                         `json:"is_root_device"`
	IsDataDevice bool                         `json:"is_data_device"`
	DiskHints    *domainTypes.RootDeviceHints `json:"disk_hints,omitempty"` //经测试，缺省值给到agent，agent会报错,所以omitempty属性
	DiskLabel    string                       `json:"disk_label"`           //gpt
	Partitions   []MPartition                 `json:"partitions"`
}

type MPartition struct {
	ID       string `json:"id"`        //instance_partiton表uuid ipt-haxcsxs570fhqew36zuuilks5w6v
	PartType string `json:"part_type"` // "esp",//可选
	FsType   string `json:"fs_type"`   //"fat32",//可选
	BootFlag string `json:"boot_flag"` //"boot",//可选。bios模式:引导分区需要设置此参数(有boot分区时是)
	Number   int    `json:"number"`    //1
	Size     int    `json:"size"`      //512,单位Mb
}

//解决struct不能设置默认值的问题
func (d MakePartitions) MarshalJSON() ([]byte, error) {
	type Alias MakePartitions
	if d.Action == "" {
		d.Action = "MakePartitions"
	}
	return json.Marshal((Alias)(d))
}
