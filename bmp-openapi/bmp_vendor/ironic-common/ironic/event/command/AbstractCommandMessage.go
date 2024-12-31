package command

type AbstractCommandMessage struct {
	Sn     string `json:"sn"`
	Action string `json:"action"`
}

func (c AbstractCommandMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.command.AbstractCommandMessage"
}
