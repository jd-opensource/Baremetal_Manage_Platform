package api

type InstanceResetPasswordMessage struct {
	ApiMessage
	InstanceId string `json:"instance_id"`
	Password   string `json:"password"`
	Action     string `json:"action"`
}

func (c InstanceResetPasswordMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.InstanceResetPasswordMessage"
}
