package agent

import "encoding/json"

type CollectHardwareInfo struct {
	Action string `json:"action"`
	Sn     string `json:"sn"`
}

//解决struct不能设置默认值的问题
func (d CollectHardwareInfo) MarshalJSON() ([]byte, error) {
	type Alias CollectHardwareInfo
	if d.Action == "" {
		d.Action = "CollectHardwareInfo"
	}
	return json.Marshal((Alias)(d))
}
