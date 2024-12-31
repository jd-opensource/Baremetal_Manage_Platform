package driver

import "encoding/json"

type DelDHCPConfigHost struct {
	Action string `json:"action"`
	Sn     string `json:"sn"`
	Mac    string `json:"mac"`
	Ip     string `json:"ip"`
}

//解决struct不能设置默认值的问题
func (d DelDHCPConfigHost) MarshalJSON() ([]byte, error) {
	type Alias DelDHCPConfigHost
	if d.Action == "" {
		d.Action = "DHCPConfigDelHost"
	}
	return json.Marshal((Alias)(d))
}
