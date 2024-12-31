package response

type Interface struct {
	Sn            string `json:"sn"`
	InterfaceName string `json:"interface_name"`
	InterfaceType string `json:"interface_type"`
	Mac           string `json:"mac"`
	SwitchIp      string `json:"switch_ip"`
	SwitchPort    string `json:"switch_port"`
}
