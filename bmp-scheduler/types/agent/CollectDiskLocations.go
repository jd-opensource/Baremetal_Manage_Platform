package agent

import "encoding/json"

type CollectDiskLocations struct {
	Action  string `json:"action"`
	Sn      string `json:"sn"`
	Version string `json:"version"`
	// AdapterIds    []string `json:"adapter_ids"`
	AllowOverride bool   `json:"allow_override"`
	RaidDriver    string `json:"raid_driver"`
}

//解决struct不能设置默认值的问题
func (d CollectDiskLocations) MarshalJSON() ([]byte, error) {
	type Alias CollectDiskLocations
	if d.Action == "" {
		d.Action = "CollectDiskLocations"
	}
	return json.Marshal((Alias)(d))
}
