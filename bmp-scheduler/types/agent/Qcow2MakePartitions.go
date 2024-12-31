package agent

import (
	"encoding/json"

	domainTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/domain"
)

type Qcow2MakePartitions struct {
	Action           string                  `json:"action"`
	Sn               string                  `json:"sn"`
	OsType           string                  `json:"os_type"`
	OsVersion        string                  `json:"os_version"`
	KeepData         bool                    `json:"keep_data"`
	DataDiskLabel    string                  `json:"data_disk_label"`
	SystemPartitions []domainTypes.Partition `json:"system_partitions"`
	DataPartitions   []domainTypes.Partition `json:"data_partitions"`
}

//解决struct不能设置默认值的问题
func (d Qcow2MakePartitions) MarshalJSON() ([]byte, error) {
	type Alias Qcow2MakePartitions
	if d.Action == "" {
		d.Action = "Qcow2MakePartitions"
	}
	return json.Marshal((Alias)(d))
}
