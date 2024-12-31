package command

type DefaultCommandMessage struct {
	CommandMessage
	CommandId  int64  `json:"command_id"`
	RequestId  string `json:"request_id"`
	InstanceId string `json:"instance_id"`
	Type       string `json:"type"`
}

func (c DefaultCommandMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.command.DefaultCommandMessage"
}
