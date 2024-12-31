package agent

import "encoding/json"

type SetPassword struct {
	Action   string `json:"action"`
	Sn       string `json:"sn"`
	Password string `json:"password"`
}

//解决struct不能设置默认值的问题
func (d SetPassword) MarshalJSON() ([]byte, error) {
	type Alias SetPassword
	if d.Action == "" {
		d.Action = "SetPassword"
	}
	return json.Marshal((Alias)(d))
}
