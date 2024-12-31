package agent

import (
	"encoding/json"

	domainTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/domain"
)

type MonutPartitions struct {
	Action       string  `json:"action"`
	Sn           string  `json:"sn"`
	Version      string  `json:"version"`
	AutoMountEfi bool    `json:"auto_mount_efi"`
	Mounts       []Mount `json:"mounts"`
}

type Mount struct {
	IsRootDevice bool                         `json:"is_root_device"`
	IsDataDevice bool                         `json:"is_data_device"`
	DiskHints    *domainTypes.RootDeviceHints `json:"disk_hints,omitempty"` //经测试，缺省值给到agent，agent会报错,所以omitempty属性
	Label        string                       `json:"label"`                //label只在每个盘符下唯一，所以同时需要确定盘符的相应信息
	Options      string                       `json:"options"`              //"options": "errors=remount-ro",
	Mountpoint   string                       `json:"mountpoint"`           // "mountpoint": "/"
}

//解决struct不能设置默认值的问题
func (d MonutPartitions) MarshalJSON() ([]byte, error) {
	type Alias MonutPartitions
	if d.Action == "" {
		d.Action = "MountPartitions"
	}
	return json.Marshal((Alias)(d))
}
