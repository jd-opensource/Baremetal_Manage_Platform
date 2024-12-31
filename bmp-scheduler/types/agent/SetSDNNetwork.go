package agent

import "encoding/json"

type SetSDNNetwork struct {
	Action                 string           `json:"action"`
	Sn                     string           `json:"sn"`
	PrimaryInterfaceName   string           `json:"primary_interface_name"`
	PrimaryMacAddr         string           `json:"primary_mac_addr"`
	PrimaryIpAddr          string           `json:"primary_ip_addr"`
	PrimaryNetmask         string           `json:"primary_netmask"`
	PrimaryGateway         string           `json:"primary_gateway"`
	SecondaryInterfaceName string           `json:"secondary_interface_name"`
	SecondaryMacAddr       string           `json:"secondary_mac_addr"`
	SecondaryIpAddr        string           `json:"secondary_ip_addr"`
	SecondaryNetmask       string           `json:"secondary_netmask"`
	SecondaryGateway       string           `json:"secondary_gateway"`
	Nameservers            []string         `json:"nameservers"`
	Routes                 []InterfaceRoute `json:"routes"`
}

//解决struct不能设置默认值的问题
func (d SetSDNNetwork) MarshalJSON() ([]byte, error) {
	type Alias SetSDNNetwork
	if d.Action == "" {
		d.Action = "SetSDNNetwork"
	}
	return json.Marshal((Alias)(d))
}
