package api

type CutDeviceStockMessage struct {
	ApiMessage
	Sns []string `json:"sns"`
}

func (c CutDeviceStockMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.CutDeviceStockMessage"
}
