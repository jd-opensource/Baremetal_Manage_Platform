package api

type AllocRetailPrivateIpMessage struct {
	ApiMessage
}

func (c AllocRetailPrivateIpMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.AllocRetailPrivateIpMessage"
}
