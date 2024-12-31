package sdn_api

const (
	CREATE_SDN_INSTANCE    = "http://%s/cps/createbm"
	DELETE_SDN_INSTANCE    = "http://%s/cps/delbm"
	BINDING_SUBNET         = "http://%s/cps/bindsub"
	UNBINDING_SUBNET       = "http://%s/cps/unbindsub"
	BINDING_EIP            = "http://%s/cps/bindeip"
	UNBINDING_EIP          = "http://%s/cps/unbindeip"
	GET_SUBNET             = "http://%s/search/getvnet"
	GET_SUBNET_IP_STOCK    = "http://%s/search/getvnetcidr"
	GET_EIP                = "http://%s/search/eipinfo"
	QUERY_ALIAS_IP         = "http://%s/plugin/getalias"
	QUERY_INSTANCE_EIPS    = "http://%s/plugin/getalog"
	BATCH_ADD_ALIAS_IP     = "http://%s/plugin/batchalias"
	DELETE_ALIAS_IP        = "http://%s/plugin/delalias"
	ASSIGN_BM_IPV6_ADDRESS = "http://%s/ipv6/assignbmipv6address"
	GET_VPC                = "http://%s/search/vpcinfo"
)

type Cfg struct {
	Ip          string
	Nameservers []string
}

var SdnProperties = map[string]Cfg{}

func InitSdnApiConfig(c map[string]Cfg) error {
	for k, v := range c {
		SdnProperties[k] = v
	}
	return nil
}
