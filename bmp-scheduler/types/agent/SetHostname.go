package agent

import "encoding/json"

type SetHostname struct {
	Action   string `json:"action"`
	Sn       string `json:"sn"`
	Hostname string `json:"hostname"`
}

//解决struct不能设置默认值的问题
func (d SetHostname) MarshalJSON() ([]byte, error) {
	type Alias SetHostname
	if d.Action == "" {
		d.Action = "SetHostname"
	}
	return json.Marshal((Alias)(d))
}
