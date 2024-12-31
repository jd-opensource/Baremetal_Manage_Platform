package agent

import "encoding/json"

type SetMetaData struct {
	Action   string `json:"action"`
	Sn       string `json:"sn"`
	MetaData struct {
		HostName string `json:"hostname,omitempty"`
		Uuid     string `json:"uuid,omitempty"`
	} `json:"meta_data"`
}

//解决struct不能设置默认值的问题
func (d SetMetaData) MarshalJSON() ([]byte, error) {
	type Alias SetMetaData
	if d.Action == "" {
		d.Action = "SetMetaData"
	}
	return json.Marshal((Alias)(d))
}
