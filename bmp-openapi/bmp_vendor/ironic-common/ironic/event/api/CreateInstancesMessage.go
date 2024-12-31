package api

type CreateInstancesMessage struct {
	ApiMessage
	InstanceIds       []string    `json:"instance_ids"`
	Password          string      `json:"password"`
	UserData          string      `json:"user_data"`
	AliasIps          interface{} `json:"alias_ips"`           //本来是[]AliasIP，依赖太深，直接定义成interface{}
	ExtensionAliasIps interface{} `json:"extension_alias_ips"` //本来是[]AliasIP，依赖太深，直接定义成interface{}
}

func (c CreateInstancesMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.CreateInstancesMessage"
}
