package driver

import "encoding/json"

type DHCPRestart struct {
	Action string `json:"action"`
	Sn     string `json:"sn"`
}

//解决struct不能设置默认值的问题
func (d DHCPRestart) MarshalJSON() ([]byte, error) {
	type Alias DHCPRestart
	if d.Action == "" {
		d.Action = "DHCPRestart"
	}
	return json.Marshal((Alias)(d))
}
