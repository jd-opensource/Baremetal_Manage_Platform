package driver

import "encoding/json"

type PowerOn struct {
	Action   string `json:"action"`
	Sn       string `json:"sn"`
	IloIp    string `json:"ilo_ip"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//解决struct不能设置默认值的问题
func (d PowerOn) MarshalJSON() ([]byte, error) {
	type Alias PowerOn
	if d.Action == "" {
		d.Action = "PowerOn"
	}
	return json.Marshal((Alias)(d))
}
