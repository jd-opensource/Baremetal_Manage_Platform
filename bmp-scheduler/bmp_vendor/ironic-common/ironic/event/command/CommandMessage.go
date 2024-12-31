package command

type CommandMessage struct {
	Sn     string `json:"sn,omitempty"`
	Action string `json:"action,omitempty"`
}

func (c CommandMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.command.CommandMessage"
}
