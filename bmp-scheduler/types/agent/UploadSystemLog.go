package agent

import "encoding/json"

type UploadSystemLog struct {
	Action string `json:"action"`
	Sn     string `json:"sn"`
}

//解决struct不能设置默认值的问题
func (d UploadSystemLog) MarshalJSON() ([]byte, error) {
	type Alias UploadSystemLog
	if d.Action == "" {
		d.Action = "UploadSystemLog"
	}
	return json.Marshal((Alias)(d))
}
