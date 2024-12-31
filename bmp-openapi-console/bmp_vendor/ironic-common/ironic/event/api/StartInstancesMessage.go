package api

type StartInstancesMessage struct {
	ApiMessage
	InstanceIds []string `json:"instance_ids"`
}

func (c StartInstancesMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.StartInstancesMessage"
}
