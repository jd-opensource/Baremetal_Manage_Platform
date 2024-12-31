package agent

import "encoding/json"

type SetNetwork struct {
	Action               string   `json:"action"`
	Sn                   string   `json:"sn"`
	PublicInterfaceName  string   `json:"public_interface_name"`
	PublicMacAddr        string   `json:"public_mac_addr"`
	PublicIpAddr         string   `json:"public_ip_addr"`
	PublicNetmask        string   `json:"public_netmask"`
	PublicGateway        string   `json:"public_gateway"`
	PrivateInterfaceName string   `json:"private_interface_name"`
	PrivateMacAddr       string   `json:"private_mac_addr"`
	PrivateIpAddr        string   `json:"private_ip_addr"`
	PrivateNetmask       string   `json:"private_netmask"`
	PrivateGateway       string   `json:"private_gateway"`
	EnableNet            bool     `json:"enable_net"`
	IsSetDns             bool     `json:"is_set_dns"`
	Nameservers          []string `json:"nameservers"`
}

//解决struct不能设置默认值的问题
func (d SetNetwork) MarshalJSON() ([]byte, error) {
	type Alias SetNetwork
	if d.Action == "" {
		d.Action = "SetNetwork"
	}
	return json.Marshal((Alias)(d))
}
