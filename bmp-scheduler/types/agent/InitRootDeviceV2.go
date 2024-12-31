package agent

import (
	"encoding/json"

	domainTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/domain"
)

type InitRootDeviceV2 struct {
	Action          string                      `json:"action"`
	Sn              string                      `json:"sn"`
	OsType          string                      `json:"os_type"`
	OsVersion       string                      `json:"os_version"`
	RootPartition   string                      `json:"root_partition"`
	RootDeviceHints domainTypes.RootDeviceHints `json:"root_device_hints"`
}

//解决struct不能设置默认值的问题
func (d InitRootDeviceV2) MarshalJSON() ([]byte, error) {
	type Alias InitRootDeviceV2
	if d.Action == "" {
		d.Action = "InitRootDeviceV2"
	}
	return json.Marshal((Alias)(d))
}
