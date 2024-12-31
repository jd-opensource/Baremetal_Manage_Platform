package command

import "time"

type MonitorMessage struct {
	RequestId     string    `json:"request_id"`
	Sn            string    `json:"sn"`
	InstanceId    string    `json:"instance_id"`
	Action        string    `json:"action"`
	TimeoutPolicy string    `json:"timeout_policy"`
	Now           time.Time `json:"now"`
}

func (c MonitorMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.command.MonitorMessage"
}
