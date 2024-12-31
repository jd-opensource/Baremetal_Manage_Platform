package agent

import "encoding/json"

type SetKeypairs struct {
	Action    string `json:"action"`
	Sn        string `json:"sn"`
	PublicKey string `json:"public_key"`
}

//解决struct不能设置默认值的问题
func (d SetKeypairs) MarshalJSON() ([]byte, error) {
	type Alias SetKeypairs
	if d.Action == "" {
		d.Action = "SetKeypairs"
	}
	return json.Marshal((Alias)(d))
}
