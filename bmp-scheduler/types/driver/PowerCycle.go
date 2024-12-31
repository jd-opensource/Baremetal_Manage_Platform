package driver

import "encoding/json"

type PowerCycle struct {
	Action   string `json:"action"`
	Sn       string `json:"sn"`
	IloIp    string `json:"ilo_ip"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//解决struct不能设置默认值的问题
func (d PowerCycle) MarshalJSON() ([]byte, error) {
	type Alias PowerCycle
	if d.Action == "" {
		d.Action = "PowerCycle"
	}
	return json.Marshal((Alias)(d))
}
