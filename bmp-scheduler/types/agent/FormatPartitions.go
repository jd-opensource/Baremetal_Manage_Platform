package agent

import (
	"encoding/json"
)

type FormatPartitions struct {
	Action  string `json:"action"`
	Sn      string `json:"sn"`
	Version string `json:"version"`

	Partitions []FPartition `json:"partitions,omitempty"`
}

type FPartition struct {
	Volume string `json:"volume"` // 此参数来自，MakePartitions的 "id": "ipt-1bt811r0feavi4sfln5nyyz8qgvg"
	FsType string `json:"fs_type"`
	Label  string `json:"label"`
	IsRoot bool   `json:"is_root"`
}

//解决struct不能设置默认值的问题
func (d FormatPartitions) MarshalJSON() ([]byte, error) {
	type Alias FormatPartitions
	if d.Action == "" {
		d.Action = "FormatPartitions"
	}
	return json.Marshal((Alias)(d))
}
