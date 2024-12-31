package driver

import "encoding/json"

type PingHost struct {
	Action string `json:"action"`
	Sn     string `json:"sn"`
	Ip     string `json:"ip"`
}

//解决struct不能设置默认值的问题
func (d PingHost) MarshalJSON() ([]byte, error) {
	type Alias PingHost
	if d.Action == "" {
		d.Action = "PingHost"
	}
	return json.Marshal((Alias)(d))
}
