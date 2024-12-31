package driver

import "encoding/json"

type DHCPConfigAddSubnet struct {
	Action     string `json:"action"`
	Sn         string `json:"sn"`
	Subnet     string `json:"subnet"`
	SubnetMask string `json:"subnet_mask"`
	Routes     string `json:"routes"`
}

//解决struct不能设置默认值的问题
func (d DHCPConfigAddSubnet) MarshalJSON() ([]byte, error) {
	type Alias DHCPConfigAddSubnet
	if d.Action == "" {
		d.Action = "DHCPConfigAddSubnet"
	}
	return json.Marshal((Alias)(d))
}
