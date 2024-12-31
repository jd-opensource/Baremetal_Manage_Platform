package response

type FilterResponse struct {
	Azs            map[string]string `json:"azs"`
	DeviceTypes    map[string]string `json:"deviceTypes"`
	InterfaceModes map[string]string `json:"interfaceModes"`
	Status         map[string]string `json:"status"`
}
