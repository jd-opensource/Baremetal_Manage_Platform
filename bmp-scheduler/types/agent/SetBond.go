package agent

import "encoding/json"

type SetBond struct {
	Action                 string   `json:"action"`
	Sn                     string   `json:"sn"`
	PrimaryInterfaceName   string   `json:"primary_interface_name"`
	PrimaryMacAddr         string   `json:"primary_mac_addr"`
	SecondaryInterfaceName string   `json:"secondary_interface_name"`
	SecondaryMacAddr       string   `json:"secondary_mac_addr"`
	IpAddr                 string   `json:"ip_addr"`
	Netmask                string   `json:"netmask"`
	Gateway                string   `json:"gateway"`
	Nameservers            []string `json:"nameservers"`
	Mode                   string   `json:"mode"`
}

//解决struct不能设置默认值的问题
func (d SetBond) MarshalJSON() ([]byte, error) {
	type Alias SetBond
	if d.Action == "" {
		d.Action = "SetBond"
	}
	return json.Marshal((Alias)(d))
}
