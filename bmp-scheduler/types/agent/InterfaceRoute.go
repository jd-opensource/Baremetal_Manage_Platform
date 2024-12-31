package agent

type InterfaceRoute struct {
	AddressCidr   string `json:"address_cidr"`
	Gateway       string `json:"gateway"`
	InterfaceName string `json:"interface_name"`
}
