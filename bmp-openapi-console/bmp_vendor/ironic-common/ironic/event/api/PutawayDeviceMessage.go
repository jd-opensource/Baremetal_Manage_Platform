package api

//设备上架
type PutawayDeviceMessage struct {
	ApiMessage
	Sn string `json:"sn"`
}

func (c PutawayDeviceMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.PutawayDeviceMessage"
}
