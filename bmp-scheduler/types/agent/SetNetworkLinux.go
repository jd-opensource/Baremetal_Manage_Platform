package agent

import "encoding/json"

type SetNocloudNetwork struct {
	Action  string `json:"action"`
	Version string `json:"version"`
	Sn      string `json:"sn"`
	Network struct {
		Ethernets struct {
			Eth0 struct {
				Match struct {
					Macaddress string `json:"macaddress,omitempty"`
				} `json:"match,omitempty"`
				SetName string `json:"set-name,omitempty"`
			} `json:"eth0,omitempty"`
			Eth1 struct {
				Match struct {
					Macaddress string `json:"macaddress,omitempty"`
				} `json:"match,omitempty"`
				SetName string `json:"set-name,omitempty"`
			} `json:"eth1,omitempty"`
		} `json:"ethernets"`
		Bonds struct {
			Bond0 struct {
				Macaddress  string      `json:"macaddress,omitempty"`
				Addresses   []string    `json:"addresses,omitempty"`
				Gateway4    string      `json:"gateway4,omitempty"`
				Gateway6    string      `json:"gateway6,omitempty"`
				Nameservers Nameservers `json:"nameservers"`
				Interfaces  []string    `json:"interfaces,omitempty"`
				Parameters  struct {
					Mode               string `json:"mode"`
					MiiMonitorInterval int    `json:"mii-monitor-interval"`
				} `json:"parameters,omitempty"`
			} `json:"bond0,omitempty"`
		} `json:"bonds,omitempty"`
		Version int `json:"version"`
	} `json:"network"`
}

type Nameservers struct {
	Addresses []string `json:"addresses,omitempty"`
}

type RouteItem struct {
	To  string `json:"to,omitempty"`
	Via string `json:"via,omitempty"`
}

//解决struct不能设置默认值的问题
func (d SetNocloudNetwork) MarshalJSON() ([]byte, error) {
	type Alias SetNocloudNetwork
	if d.Action == "" {
		d.Action = "SetNetwork"
	}
	return json.Marshal((Alias)(d))
}

type SetNocloudNetworkNoBond struct {
	Action  string `json:"action"`
	Version string `json:"version"`
	Sn      string `json:"sn"`
	Network struct {
		Ethernets struct {
			Eth0 struct {
				Match struct {
					Macaddress string `json:"macaddress,omitempty"`
				} `json:"match,omitempty"`
				SetName     string      `json:"set-name,omitempty"`
				Addresses   []string    `json:"addresses,omitempty"`
				Netmask     string      `json:"netmask,omitempty"`
				Gateway4    string      `json:"gateway4,omitempty"`
				Gateway6    string      `json:"gateway6,omitempty"`
				Nameservers Nameservers `json:"nameservers,omitempty"`
				Routes      []RouteItem `json:"routes,omitempty"`
			} `json:"eth0,omitempty"`
			Eth1 struct {
				Match struct {
					Macaddress string `json:"macaddress,omitempty"`
				} `json:"match,omitempty"`
				SetName     string      `json:"set-name,omitempty"`
				Addresses   []string    `json:"addresses,omitempty"`
				Netmask     string      `json:"netmask,omitempty"`
				Gateway4    string      `json:"gateway4,omitempty"`
				Gateway6    string      `json:"gateway6,omitempty"`
				Nameservers Nameservers `json:"nameservers,omitempty"`
				Routes      []RouteItem `json:"routes,omitempty"`
			} `json:"eth1,omitempty"` //暂时先设置一个网口
		} `json:"ethernets"`

		Version int `json:"version"`
	} `json:"network"`
}

//解决struct不能设置默认值的问题
func (d SetNocloudNetworkNoBond) MarshalJSON() ([]byte, error) {
	type Alias SetNocloudNetworkNoBond
	if d.Action == "" {
		d.Action = "SetNetwork"
	}
	return json.Marshal((Alias)(d))
}
