package command

type CallbackCommandMessage struct {
	CommandMessage
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (c CallbackCommandMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.command.CallbackCommandMessage"
}
