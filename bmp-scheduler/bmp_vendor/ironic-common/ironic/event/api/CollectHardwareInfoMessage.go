package api

type CollectHardwareInfoMessage struct {
	ApiMessage
	Sns []string `json:"sns"`
	//是否清空raid
	AllowOverride bool   `json:"allowOverride"`
	NetworkType   string `json:"network_type"`
}

func (c CollectHardwareInfoMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.CollectHardwareInfoMessage"
}
