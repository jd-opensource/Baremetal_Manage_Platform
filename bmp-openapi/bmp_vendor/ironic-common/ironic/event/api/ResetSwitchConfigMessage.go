package api

type ResetSwitchConfigMessage struct {
	ApiMessage
	InstanceIds []string `json:"instance_ids"`
}

func (c ResetSwitchConfigMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.ResetSwitchConfigMessage"
}
