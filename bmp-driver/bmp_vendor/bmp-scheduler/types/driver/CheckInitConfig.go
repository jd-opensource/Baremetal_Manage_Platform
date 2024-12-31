package driver

import "encoding/json"

type CheckInitConfig struct {
	Action        string `json:"action"`
	Sn            string `json:"sn"`
	IloIp         string `json:"ilo_ip"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Privilege     string `json:"privilege"` //用户权限:Administrator, User, Operator
	SubnetGateway string `json:"subnet_gateway"`
	Mac1          string `json:"mac1"`
	Mac2          string `json:"mac2"`
}

//解决struct不能设置默认值的问题
func (d CheckInitConfig) MarshalJSON() ([]byte, error) {
	type Alias CheckInitConfig
	if d.Action == "" {
		d.Action = "CheckInitConfig"
	}
	return json.Marshal((Alias)(d))
}
