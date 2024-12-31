package Constants

const (
	SUFFIX = "%s_command_processor"

	INSTANCE_EXTRA_KEY = "ironic_instance_extra_%s"

	RESTART_DHCP_EXTRA_KEY = "restart_dhcp_extra"

	COLLECT_EXTRA_KEY = "collect_sn_extra_%s"

	INSTANCE_ALIAS_IP_KEY = "ironic_instance_alias_ip_%s_%s"

	COMMAND_RESULT = "ironic_command_result_%s"

	REDIS_KEY_CLUSTER = "ironic_cluster"

	REDIS_KEY_CHECK = "ironic_check"

	REDIS_KEY_INSTANCE = "ironic_instance"

	REDIS_KEY_STATUS = "ironic_instance_status"

	DHCP_RESTART_KEY = "dhcp_restart_key"

	REDIS_KEY_MAIL_RECEIVER_FORMAT = "ironic_mail_receiver_%s"

	REDIS_KEY_MONITOR_COMMAND_FORMAT = "ironic_monitor_command_%s_%s"

	TIMEOUT_INSTANCE_START int = 2 * 60

	TIMEOUT_INSTANCE_STOP int = 30

	TIMEOUT_INSTANCE_RESTART int = 3 * 60

	TIMEOUT_INSTANCE_RESTART_WINDOWS int = 8 * 60

	TIMEOUT_INSTANCE_STATUS int = 60 * 60

	TIMEOUT_MONITOR_COMMAND int = 60 * 60

	TAR_IMAGE = "tar"

	SET_PXE_BOOT_PROCESSOR_NAME = "SetPXEBoot_command_processor"

	SET_DISK_BOOT_PROCESSOR_NAME = "SetDISKBoot_command_processor"

	POWER_ON_PROCESSOR_NAME = "PowerOn_command_processor"

	POWER_OFF_PROCESSOR_NAME = "PowerOff_command_processor"

	DHCP_RESTART_PROCESSOR_NAME = "DHCPRestart_command_processor"

	POWER_RESET_PROCESSOR_NAME = "PowerReset_command_processor"

	POWER_CYCLE_PROCESSOR_NAME = "PowerCycle_command_processor"

	INIT_ROOT_DEVICE_PROCESSOR_NAME = "InitRootDevice_command_processor"

	CLEAN_RAID_PROCESSOR_NAME = "CleanRaid_command_processor"

	MAKE_RAID_PROCESSOR_NAME = "MakeRaid_command_processor"

	PING_PROCESSOR_NAME = "Ping_command_processor"

	ADD_DHCP_CONFIG_HOST_PROCESSOR_NAME = "DHCPConfigAddHost_command_processor"

	DEL_DHCP_CONFIG_HOST_PROCESSOR_NAME = "DHCPConfigDelHost_command_processor"

	WRITE_IMAGE_RAID_PROCESSOR_NAME = "WriteImage_command_processor"

	WRITE_IMAGE_TAR_PROCESSOR_NAME = "WriteImageTar_command_processor"

	QCOW2_MAKE_PARTITIONS_PROCESSOR_NAME = "Qcow2MakePartitions_command_processor"

	MAKE_PARTITIONS_PROCESSOR_NAME = "MakePartitions_command_processor"

	SET_HOSTNAME_PROCESSOR_NAME = "SetHostname_command_processor"

	SET_PASSWORD_PROCESSOR_NAME = "SetPassword_command_processor"

	SET_NETWORK_PROCESSOR_NAME = "SetNetwork_command_processor"

	SET_BOND_PROCESSOR_NAME = "SetBond_command_processor"

	SET_RETAIL_BOND_PROCESSOR_NAME = "SetRetailBond_command_processor"

	SDN_SET_NETWORK_PROCESSOR_NAME = "SetVpcNetwork_command_processor"

	SET_USER_CMD_PROCESSOR_NAME = "SetUserCmd_command_processor"

	SET_KEYPAIRS_PROCESSOR_NAME = "SetKeypairs_command_processor"

	CLEAN_BLOCK_DEVICE_PROCESSOR_NAME = "CleanBlockDevice_command_processor"

	COLLECT_HARDWARE_INFO_PROCESSOR_NAME = "CollectHardwareInfo_command_processor"

	UPLOAD_SYSTEM_LOG_PROCESSOR_NAME = "UploadSystemLog_command_processor"

	CREATE_VRF_PROCESSOR_NAME = "CreateVRF_command_processor"

	CREATE_VRF_BALANCE_PROCESSOR_NAME = "CreateVRFBalance_command_processor"

	CREATE_VSI_INTERFACE_PROCESSOR_NAME = "CreateVSIInterface_command_processor"

	CREATE_VSI_PROCESSOR_NAME = "CreateVSI_command_processor"

	BINDING_VSI_PROCESSOR_NAME = "BindingVSI_command_processor"

	UN_BINDING_VSI_PROCESSOR_NAME = "UnBindingVSI_command_processor"

	BINDING_IPV6_PROCESSOR_NAME = "BindingIPv6_command_processor"

	CREATE_VNI_PROCESSOR_NAME = "CreateVNI_command_processor"

	ADD_ARP_STATIC_PROCESSOR_NAME = "AddArpStatic_command_processor"

	CLEAN_ARP_STATIC_PROCESSOR_NAME = "CleanArpStatic_command_processor"

	SET_BANDWIDTH_PROCESSOR_NAME = "SetBandwidth_command_processor"

	CLEAN_BANDWIDTH__PROCESSOR_NAME = "CleanBandwidth_command_processor"

	SAVE_SWITCH_CONFIG_PROCESSOR_NAME = "SaveSwitchConfig_command_processor"

	SAVE_CONFIG_TO_FTP_SERVER_PROCESSOR_NAME = "SaveConfigToFtpServer_command_processor"

	SYSTEM_ERROR_COMMAND_PROCESSOR_NAME = "SystemErrorCommand_command_processor"

	SDN_REGISTER_PROCESSOR_NAME = "SDNRegister_command_processor"

	SDN_SUBNET_BINDING_PROCESSOR_NAME = "SDNSubnetBinding_command_processor"

	SDN_SUBNET_UNBINDING_PROCESSOR_NAME = "SDNSubnetUnBinding_command_processor"

	SDN_EIP_BINDING_PROCESSOR_NAME = "SDNEipBinding_command_processor"

	SDN_EIP_UNBINDING_PROCESSOR_NAME = "SDNEipUnBinding_command_processor"

	SDN_CLEAN_PROCESSOR_NAME = "SDNClean_command_processor"

	SDN_ADD_ALIAS_IP_PROCESSOR_NAME = "SDNAddAliasIP_command_processor"

	SDN_DELETE_ALIAS_IP_PROCESSOR_NAME = "SDNDeleteAliasIP_command_processor"
)
