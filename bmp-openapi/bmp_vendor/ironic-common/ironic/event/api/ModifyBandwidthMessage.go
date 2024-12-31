package api

type ModifyBandwidthMessage struct {
	ApiMessage
	InstanceId     string `json:"instance_id"`
	Bandwidth      int32  `json:"bandwidth"`
	InstanceStatus string `json:"instance_status"`
}

func (c ModifyBandwidthMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.ModifyBandwidthMessage"
}
