package CommandAction

import (
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/TimeoutPolicy"
)

type CommandAction struct {
	Name        string
	CommandType string
	Timeout     int
	Policy      string
}

var (
	Start               = CommandAction{"Start", "", 0, TimeoutPolicy.WARNING}
	Heart               = CommandAction{"Heart", "", 0, TimeoutPolicy.WARNING}
	SystemErrorCommand  = CommandAction{"SystemErrorCommand", "", 0, TimeoutPolicy.WARNING}
	DHCPConfigAddSubnet = CommandAction{"DHCPConfigAddSubnet", CommandType.DRIVER, 3 * 60, TimeoutPolicy.WARNING}
	CheckInitConfig     = CommandAction{"CheckInitConfig", CommandType.DRIVER, 3 * 60, TimeoutPolicy.WARNING}
	DHCPConfigAddHost   = CommandAction{"DHCPConfigAddHost", CommandType.DRIVER, 3 * 60, TimeoutPolicy.WARNING}
	DHCPConfigDelHost   = CommandAction{"DHCPConfigDelHost", CommandType.DRIVER, 3 * 60, TimeoutPolicy.WARNING}
	SetPXEBoot          = CommandAction{"SetPXEBoot", CommandType.DRIVER, 3 * 60, TimeoutPolicy.WARNING}
	SetDISKBoot         = CommandAction{"SetDISKBoot", CommandType.DRIVER, 3 * 60, TimeoutPolicy.WARNING}
	PowerOn             = CommandAction{"PowerOn", CommandType.DRIVER, 3 * 60, TimeoutPolicy.WARNING}
	PowerOff            = CommandAction{"PowerOff", CommandType.DRIVER, 3 * 60, TimeoutPolicy.WARNING}
	PowerReset          = CommandAction{"PowerReset", CommandType.DRIVER, 3 * 60, TimeoutPolicy.WARNING}
	PowerCycle          = CommandAction{"PowerCycle", CommandType.DRIVER, 3 * 60, TimeoutPolicy.WARNING}
	DHCPRestart         = CommandAction{"DHCPRestart", CommandType.DRIVER, 3 * 60, TimeoutPolicy.SKIP}

	CollectHardwareInfo  = CommandAction{"CollectHardwareInfo", CommandType.AGENT, 10 * 60, TimeoutPolicy.SKIP}
	CollectDiskLocations = CommandAction{"CollectDiskLocations", CommandType.AGENT, 10 * 60, TimeoutPolicy.SKIP}
	UploadSystemLog      = CommandAction{"UploadSystemLog", CommandType.AGENT, 3 * 60, TimeoutPolicy.SKIP}
	CleanRaid            = CommandAction{"CleanRaid", CommandType.AGENT, 6 * 60, TimeoutPolicy.WARNING}
	MakeRaid             = CommandAction{"MakeRaid", CommandType.AGENT, 6 * 60, TimeoutPolicy.WARNING}
	WriteImage           = CommandAction{"WriteImage", CommandType.AGENT, 25 * 60, TimeoutPolicy.WARNING}
	WriteImageTar        = CommandAction{"WriteImageTar", CommandType.AGENT, 25 * 60, TimeoutPolicy.WARNING}
	Qcow2MakePartitions  = CommandAction{"Qcow2MakePartitions", CommandType.AGENT, 6 * 60, TimeoutPolicy.WARNING}
	MakePartitions       = CommandAction{"MakePartitions", CommandType.AGENT, 25 * 60, TimeoutPolicy.WARNING}
	SetHostname          = CommandAction{"SetHostname", CommandType.AGENT, 3 * 60, TimeoutPolicy.WARNING}
	SetPassword          = CommandAction{"SetPassword", CommandType.AGENT, 3 * 60, TimeoutPolicy.WARNING}
	SetNetwork           = CommandAction{"SetNetwork", CommandType.AGENT, 3 * 60, TimeoutPolicy.WARNING}
	SetVpcNetwork        = CommandAction{"SetVpcNetwork", CommandType.AGENT, 3 * 60, TimeoutPolicy.WARNING}
	SetBond              = CommandAction{"SetBond", CommandType.AGENT, 3 * 60, TimeoutPolicy.WARNING}
	SetUserCmd           = CommandAction{"SetUserCmd", CommandType.AGENT, 3 * 60, TimeoutPolicy.WARNING}
	SetKeypairs          = CommandAction{"SetKeypairs", CommandType.AGENT, 3 * 60, TimeoutPolicy.WARNING}
	//帮忙Agent重试mq链接
	Ping             = CommandAction{"Ping", CommandType.AGENT, 30 * 60, TimeoutPolicy.WARNING}
	InitRootDevice   = CommandAction{"InitRootDevice", CommandType.AGENT, 6 * 60, TimeoutPolicy.WARNING}
	CleanBlockDevice = CommandAction{"CleanBlockDevice", CommandType.AGENT, 25 * 60, TimeoutPolicy.WARNING}

	CreateVRF             = CommandAction{"CreateVRF", CommandType.NETWORK, 60, TimeoutPolicy.WARNING}
	CreateVRFBalance      = CommandAction{"CreateVRFBalance", CommandType.NETWORK, 60, TimeoutPolicy.WARNING}
	CreateVSIInterface    = CommandAction{"CreateVSIInterface", CommandType.NETWORK, 60, TimeoutPolicy.WARNING}
	CreateVSI             = CommandAction{"CreateVSI", CommandType.NETWORK, 60, TimeoutPolicy.WARNING}
	BindingVSI            = CommandAction{"BindingVSI", CommandType.NETWORK, 60, TimeoutPolicy.WARNING}
	UnBindingVSI          = CommandAction{"UnBindingVSI", CommandType.NETWORK, 60, TimeoutPolicy.WARNING}
	SetRetailBond         = CommandAction{"SetRetailBond", CommandType.NETWORK, 60, TimeoutPolicy.WARNING}
	CreateVNI             = CommandAction{"CreateVNI", CommandType.NETWORK, 60, TimeoutPolicy.WARNING}
	AddArpStatic          = CommandAction{"AddArpStatic", CommandType.NETWORK, 60, TimeoutPolicy.WARNING}
	CleanArpStatic        = CommandAction{"CleanArpStatic", CommandType.NETWORK, 60, TimeoutPolicy.WARNING}
	SetBandwidth          = CommandAction{"SetBandwidth", CommandType.NETWORK, 60, TimeoutPolicy.WARNING}
	CleanBandwidth        = CommandAction{"CleanBandwidth", CommandType.NETWORK, 60, TimeoutPolicy.WARNING}
	SaveSwitchConfig      = CommandAction{"SaveSwitchConfig", CommandType.NETWORK, 60, TimeoutPolicy.WARNING}
	SaveConfigToFtpServer = CommandAction{"SaveConfigToFtpServer", CommandType.NETWORK, 60, TimeoutPolicy.WARNING}

	SDNRegister        = CommandAction{"SDNRegister", CommandType.SDN, 60, TimeoutPolicy.WARNING}
	SDNSubnetBinding   = CommandAction{"SDNSubnetBinding", CommandType.SDN, 60, TimeoutPolicy.WARNING}
	SDNSubnetUnBinding = CommandAction{"SDNSubnetUnBinding", CommandType.SDN, 60, TimeoutPolicy.WARNING}
	SDNEipBinding      = CommandAction{"SDNEipBinding", CommandType.SDN, 60, TimeoutPolicy.WARNING}
	SDNEipUnBinding    = CommandAction{"SDNEipUnBinding", CommandType.SDN, 60, TimeoutPolicy.WARNING}
	SDNAddAliasIP      = CommandAction{"SDNAddAliasIP", CommandType.SDN, 60, TimeoutPolicy.WARNING}
	SDNDeleteAliasIP   = CommandAction{"SDNDeleteAliasIP", CommandType.SDN, 60, TimeoutPolicy.WARNING}
	SDNClean           = CommandAction{"SDNClean", CommandType.SDN, 60, TimeoutPolicy.WARNING}
)

var CommandActionDict = map[string]CommandAction{
	"Start":                 Start,
	"Heart":                 Heart,
	"SystemErrorCommand":    SystemErrorCommand,
	"DHCPConfigAddSubnet":   DHCPConfigAddSubnet,
	"CheckInitConfig":       CheckInitConfig,
	"DHCPConfigAddHost":     DHCPConfigAddHost,
	"DHCPConfigDelHost":     DHCPConfigDelHost,
	"SetPXEBoot":            SetPXEBoot,
	"SetDISKBoot":           SetDISKBoot,
	"PowerOn":               PowerOn,
	"PowerOff":              PowerOff,
	"PowerReset":            PowerReset,
	"PowerCycle":            PowerCycle,
	"DHCPRestart":           DHCPRestart,
	"CollectHardwareInfo":   CollectHardwareInfo,
	"UploadSystemLog":       UploadSystemLog,
	"CleanRaid":             CleanRaid,
	"MakeRaid":              MakeRaid,
	"WriteImage":            WriteImage,
	"WriteImageTar":         WriteImageTar,
	"Qcow2MakePartitions":   Qcow2MakePartitions,
	"MakePartitions":        MakePartitions,
	"SetHostname":           SetHostname,
	"SetPassword":           SetPassword,
	"SetNetwork":            SetNetwork,
	"SetVpcNetwork":         SetVpcNetwork,
	"SetBond":               SetBond,
	"SetUserCmd":            SetUserCmd,
	"SetKeypairs":           SetKeypairs,
	"Ping":                  Ping,
	"InitRootDevice":        InitRootDevice,
	"CleanBlockDevice":      CleanBlockDevice,
	"CreateVRF":             CreateVRF,
	"CreateVRFBalance":      CreateVRFBalance,
	"CreateVSIInterface":    CreateVSIInterface,
	"CreateVSI":             CreateVSI,
	"BindingVSI":            BindingVSI,
	"UnBindingVSI":          UnBindingVSI,
	"SetRetailBond":         SetRetailBond,
	"CreateVNI":             CreateVNI,
	"AddArpStatic":          AddArpStatic,
	"CleanArpStatic":        CleanArpStatic,
	"SetBandwidth":          SetBandwidth,
	"CleanBandwidth":        CleanBandwidth,
	"SaveSwitchConfig":      SaveSwitchConfig,
	"SaveConfigToFtpServer": SaveConfigToFtpServer,
	"SDNRegister":           SDNRegister,
	"SDNSubnetBinding":      SDNSubnetBinding,
	"SDNSubnetUnBinding":    SDNSubnetUnBinding,
	"SDNEipBinding":         SDNEipBinding,
	"SDNEipUnBinding":       SDNEipUnBinding,
	"SDNAddAliasIP":         SDNAddAliasIP,
	"SDNDeleteAliasIP":      SDNDeleteAliasIP,
	"SDNClean":              SDNClean,
}
