package agent

import "encoding/json"

type WriteImageTar struct {
	Action    string `json:"action"`
	Sn        string `json:"sn"`
	Url       string `json:"url"`
	Format    string `json:"format"`
	Hash      string `json:"hash"`
	Filename  string `json:"filename"`
	OsType    string `json:"os_type"`
	OsVersion string `json:"os_version"`
}

//解决struct不能设置默认值的问题
func (d WriteImageTar) MarshalJSON() ([]byte, error) {
	type Alias WriteImageTar
	if d.Action == "" {
		d.Action = "WriteImageTar"
	}
	return json.Marshal((Alias)(d))
}