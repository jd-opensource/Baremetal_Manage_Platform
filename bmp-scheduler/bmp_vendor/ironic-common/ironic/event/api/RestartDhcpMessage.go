package api

type RestartDhcpMessage struct {
	ApiMessage
	Az string `json:"az"`
}

func (c RestartDhcpMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.RestartDhcpMessage"
}
