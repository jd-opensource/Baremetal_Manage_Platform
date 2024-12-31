package api

type CreateDevicesMessage struct {
	ApiMessage
	Action     string `json:"action"`
	Sn         string `json:"sn"`
	Subnet     string `json:"subnet"`
	SubnetMask string `json:"subnet_mask"`
	Routes     string `json:"routes"`
}

func (c CreateDevicesMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.CreateDevicesMessage"
}
