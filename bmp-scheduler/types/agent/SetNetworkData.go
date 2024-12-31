package agent

import "encoding/json"

type SetNetworkData struct {
	Sn          string `json:"sn"`
	Action      string `json:"action"`
	NetworkData struct {
		Services []Service `json:"services"`
		Networks []NetWork `json:"networks"`
		Links    []Link    `json:"links"`
	} `json:"network"`
}
type Service struct {
	Type    string `json:"type"`
	Address string `json:"address"`
}
type NetWork struct {
	IPAddress string  `json:"ip_address"`
	Type      string  `json:"type"`
	Netmask   string  `json:"netmask"`
	Link      string  `json:"link"`
	Routes    []Route `json:"routes"`
}
type Route struct {
	Netmask string `json:"netmask"`
	Network string `json:"network"`
	Gateway string `json:"gateway"`
}
type Link struct {
	EthernetMacAddress string   `json:"ethernet_mac_address"`
	Type               string   `json:"type"`
	ID                 string   `json:"id"`
	Mtu                int      `json:"mtu,omitempty"`
	BondMiimon         int      `json:"bond_miimon,omitempty"`
	BondMode           string   `json:"bond_mode,omitempty"`
	BondLinks          []string `json:"bond_links,omitempty"`
}

//解决struct不能设置默认值的问题
func (d SetNetworkData) MarshalJSON() ([]byte, error) {
	type Alias SetNetworkData
	if d.Action == "" {
		d.Action = "SetNetworkWindows"
	}
	return json.Marshal((Alias)(d))
}
