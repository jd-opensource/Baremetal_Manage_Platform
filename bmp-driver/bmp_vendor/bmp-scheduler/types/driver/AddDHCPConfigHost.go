package driver

import "encoding/json"

type AddDHCPConfigHost struct {
	Action string `json:"action"`
	Sn     string `json:"sn"`
	Mac    string `json:"mac"`
	Ip     string `json:"ip"`
}

//解决struct不能设置默认值的问题
func (d AddDHCPConfigHost) MarshalJSON() ([]byte, error) {
	type Alias AddDHCPConfigHost
	if d.Action == "" {
		d.Action = "DHCPConfigAddHost"
	}
	return json.Marshal((Alias)(d))
}
