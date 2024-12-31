package agent

import "encoding/json"

type InitRootDevice struct {
	Action    string `json:"action"`
	Sn        string `json:"sn"`
	OsType    string `json:"os_type"`
	OsVersion string `json:"os_version"`
}

//解决struct不能设置默认值的问题
func (d InitRootDevice) MarshalJSON() ([]byte, error) {
	type Alias InitRootDevice
	if d.Action == "" {
		d.Action = "InitRootDevice"
	}
	return json.Marshal((Alias)(d))
}
