package agent

import "encoding/json"

type SetUserCmd struct {
	Action  string `json:"action"`
	Sn      string `json:"sn"`
	Content string `json:"content"`
}

//解决struct不能设置默认值的问题
func (d SetUserCmd) MarshalJSON() ([]byte, error) {
	type Alias SetUserCmd
	if d.Action == "" {
		d.Action = "SetUserCmd"
	}
	return json.Marshal((Alias)(d))
}
