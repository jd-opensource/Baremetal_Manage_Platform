package driver

import "encoding/json"

type SetDISKBoot struct {
	Action   string `json:"action"`
	Sn       string `json:"sn"`
	IloIp    string `json:"ilo_ip"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//解决struct不能设置默认值的问题
func (d SetDISKBoot) MarshalJSON() ([]byte, error) {
	type Alias SetDISKBoot
	if d.Action == "" {
		d.Action = "SetDISKBoot"
	}
	return json.Marshal((Alias)(d))
}
