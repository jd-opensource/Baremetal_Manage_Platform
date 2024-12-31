package driver

import "encoding/json"

type PowerOff struct {
	Action   string `json:"action"`
	Sn       string `json:"sn"`
	IloIp    string `json:"ilo_ip"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//解决struct不能设置默认值的问题
func (d PowerOff) MarshalJSON() ([]byte, error) {
	type Alias PowerOff
	if d.Action == "" {
		d.Action = "PowerOff"
	}
	return json.Marshal((Alias)(d))
}
