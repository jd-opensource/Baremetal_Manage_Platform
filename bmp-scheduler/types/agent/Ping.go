package agent

import "encoding/json"

type Ping struct {
	Action string `json:"action"`
	Sn     string `json:"sn"`
}

//解决struct不能设置默认值的问题
func (d Ping) MarshalJSON() ([]byte, error) {
	type Alias Ping
	if d.Action == "" {
		d.Action = "Ping"
	}
	return json.Marshal((Alias)(d))
}
