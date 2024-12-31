package agent

import "encoding/json"

type CleanBlockDevice struct {
	Action string `json:"action"`
	Sn     string `json:"sn"`
}

//解决struct不能设置默认值的问题
func (d CleanBlockDevice) MarshalJSON() ([]byte, error) {
	type Alias CleanBlockDevice
	if d.Action == "" {
		d.Action = "CleanBlockDevice"
	}
	return json.Marshal((Alias)(d))
}
