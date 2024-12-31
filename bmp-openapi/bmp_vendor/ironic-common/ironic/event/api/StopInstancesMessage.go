package api

type StopInstancesMessage struct {
	ApiMessage
	InstanceIds []string `json:"instance_ids"`
}

func (c StopInstancesMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.StopInstancesMessage"
}
