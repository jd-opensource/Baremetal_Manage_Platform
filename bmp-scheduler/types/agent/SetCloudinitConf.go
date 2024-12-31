package agent

import "encoding/json"

type SetCloudinitConf struct {
	Action    string `json:"action"`
	Sn        string `json:"sn"`
	Version   string `json:"version"`
	SshPwauth string `json:"ssh_pwauth"`
}

//解决struct不能设置默认值的问题
func (d SetCloudinitConf) MarshalJSON() ([]byte, error) {
	type Alias SetCloudinitConf
	if d.Action == "" {
		d.Action = "SetCloudinitConf"
	}
	return json.Marshal((Alias)(d))
}
