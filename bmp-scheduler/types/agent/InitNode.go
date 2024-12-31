package agent

import (
	"encoding/json"

	domainTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/domain"
)

type InitNode struct {
	Action   string   `json:"action"`
	Sn       string   `json:"sn"`
	NodeData NodeData `json:"node_data"`
}

type NodeData struct {
	RootDeviceHints domainTypes.RootDeviceHints `json:"root_device_hints,omitempty"`
	OsType          string                      `json:"os_type"`
	OsVersion       string                      `json:"os_version"`
}

//解决struct不能设置默认值的问题
func (d InitNode) MarshalJSON() ([]byte, error) {
	type Alias InitNode
	if d.Action == "" {
		d.Action = "InitNode"
	}
	return json.Marshal((Alias)(d))
}
