package types

type InstanceExtra struct {
	InstanceId string `json:"instance_id"`

	Password string `json:"password"`

	UserData string `json:"user_data"`

	KeepData bool `json:"keep_data"`

	AliasIps []AliasIP `json:"alias_ips"`

	ExtensionAliasIps []AliasIP `json:"extension_alias_ips"`
}

type AliasIP struct {
	SubnetId string `json:"subnet_id"`
	Cidr     string `json:"cidr"`
}
