package agent

import "encoding/json"

type CleanRaid struct {
	Action    string     `json:"action"`
	Sn        string     `json:"sn"`
	RaidDatas []RaidData `json:"raid_datas"`
}

type RaidData struct {
	AdapterId  int    `json:"adapter_id"`
	RaidDriver string `json:"raid_driver"`
}

//解决struct不能设置默认值的问题
func (d CleanRaid) MarshalJSON() ([]byte, error) {
	type Alias CleanRaid
	if d.Action == "" {
		d.Action = "CleanRaid"
	}
	return json.Marshal((Alias)(d))
}
