package agent

import "encoding/json"

type SetNocloudMetaData struct {
	Action   string `json:"action"`
	Version  string `json:"version"`
	Sn       string `json:"sn"`
	MetaData struct {
		InstanceID    string   `json:"instance-id"`
		LocalHostname string   `json:"local-hostname,omitempty"`
		PublicKeys    []string `json:"public-keys,omitempty"`
	} `json:"meta_data"`
}

//解决struct不能设置默认值的问题
func (d SetNocloudMetaData) MarshalJSON() ([]byte, error) {
	type Alias SetNocloudMetaData
	if d.Action == "" {
		d.Action = "SetMetaData"
	}
	return json.Marshal((Alias)(d))
}
