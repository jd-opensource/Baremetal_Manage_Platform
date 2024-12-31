package api

type SDNCallbackMessage struct {
	ApiMessage
	InstanceId string      `json:"instance_id"`
	Sn         string      `json:"sn"`
	Action     string      `json:"action"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func (c SDNCallbackMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.SDNCallbackMessage"
}
