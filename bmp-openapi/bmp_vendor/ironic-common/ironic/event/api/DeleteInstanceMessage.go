package api

type DeleteInstanceMessage struct {
	ApiMessage
	InstanceId string `json:"instance_id"`
}

func (c DeleteInstanceMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.DeleteInstanceMessage"
}
