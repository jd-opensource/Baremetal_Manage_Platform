package api

type ReinstallInstanceMessage struct {
	ApiMessage
	InstanceId string `json:"instance_id"`
	KeepData   bool   `json:"keep_data"`
	Password   string `json:"password"`
	UserData   string `json:"user_data"`
}

func (c ReinstallInstanceMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.ReinstallInstanceMessage"
}
