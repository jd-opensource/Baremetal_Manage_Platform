package agent

import "encoding/json"

type SetNocloudUserData struct {
	Action   string `json:"action"`
	Sn       string `json:"sn"`
	Version  string `json:"version"`
	UserData string `json:"user_data"`
}

//解决struct不能设置默认值的问题
func (d SetNocloudUserData) MarshalJSON() ([]byte, error) {
	type Alias SetNocloudUserData
	if d.Action == "" {
		d.Action = "SetUserData"
	}
	return json.Marshal((Alias)(d))
}
