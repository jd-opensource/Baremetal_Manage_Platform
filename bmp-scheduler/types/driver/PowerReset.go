package driver

import "encoding/json"

type PowerReset struct {
	Action   string `json:"action"`
	Sn       string `json:"sn"`
	IloIp    string `json:"ilo_ip"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//解决struct不能设置默认值的问题
func (d PowerReset) MarshalJSON() ([]byte, error) {
	type Alias PowerReset
	if d.Action == "" {
		d.Action = "PowerReset"
	}
	return json.Marshal((Alias)(d))
}
