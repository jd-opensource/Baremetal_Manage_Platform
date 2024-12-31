package api

type RestartInstancesMessage struct {
	ApiMessage
	InstanceIds []string `json:"instance_ids"`
}

func (c RestartInstancesMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.RestartInstancesMessage"
}
