package baseLogic

var (
	PartitionType = map[string]string{
		"/":    "root",
		"swap": "system",
	}
	PartitionLabel = map[string]string{
		"/":    "l_root",
		"swap": "l_swap",
	}
	SystemDiskLabel = map[string]string{
		"/":    "gpt",
		"swap": "gpt",
	}
	DataDiskLabel = map[string]string{
		"/":    "gpt",
		"swap": "gpt",
	}
	DeviceManageStatus = map[string]string{
		IN:           "已入库",
		PUTAWAY:      "已上架",
		UNPUTAWAY:    "已下架",
		CREATED:      "已创建",
		PUTAWAYING:   "上架中",
		PUTAWAYFAIL:  "上架失败",
		UNPUTAWAYING: "下架中",
		REMOVED:      "已移除",
	}
	DeviceManageStatusEn = map[string]string{
		IN:           "Stocked",
		PUTAWAY:      "Mounted",
		UNPUTAWAY:    "Unmounted",
		CREATED:      "created",
		PUTAWAYING:   "Mounting",
		PUTAWAYFAIL:  "Mount failed",
		UNPUTAWAYING: "Unmounting",
		REMOVED:      "removed",
	}
	DeviceReason = map[string]string{
		DEFAULT_BMP_ERROR_CODE:     "未知错误",
		PING_ILOIP_ERROR:           "ping带外不通",
		USERNAME_OR_PASSWORD_ERROR: "带外用户名或密码错误",
		USER_PRIVILEGE_ERROR:       "用户权限错误",
		PING_SUBNET_GATEWAY_ERROR:  "ping子网网关错误",
		MAC1_ERROR:                 "mac1地址有误",
		MAC2_ERROR:                 "mac2地址有误",
	}
	DeviceReasonEn = map[string]string{
		DEFAULT_BMP_ERROR_CODE:     "unknown error message",
		PING_ILOIP_ERROR:           "ping OOB fails",
		USERNAME_OR_PASSWORD_ERROR: "OOB Username does not exist or password is wrong",
		USER_PRIVILEGE_ERROR:       "User permission error",
		PING_SUBNET_GATEWAY_ERROR:  "ping subnet gateway error",
		MAC1_ERROR:                 "mac1 address is wrong",
		MAC2_ERROR:                 "mac2 address is wrong",
	}
	InstanceReason = map[string]string{
		DEFAULT_BMP_ERROR_CODE: "未知错误",
		MakeRaidError:          "raid失败",
		MakePartitionsError:    "分区失败",
		WriteImageTarError:     "写入镜像失败",
		SetPasswordError:       "设置密码失败",
		SetCloudinitConfError:  "cloudinit初始化失败",

		SetNocloudUserDataError: "设置用户自定义指令失败",
		SetNocloudMetaDataError: "设置metadata失败",
		SetNocloudNetworkError:  "设置网络失败",
		PingError:               "PXE启动失败",

		DHCPConfigAddHostError: "设置DHCP配置失败",
		DHCPConfigDelHostError: "删除DHCP配置失败",
		SetDISKBootError:       "设置硬盘启动失败",
		PowerCycleError:        "重启失败",
		PowerOffError:          "关机失败",
		PowerOnError:           "开机失败",
		SetPXEBootError:        "设置PXE启动失败",

		InitRootDeviceError: "初始化设备信息失败",
		//windows专用指令
		WriteImageError:     "写入镜像失败",
		SetMetaDataError:    "设置metadata失败",
		SetNetworkDataError: "设置网络失败",
	}
	InstanceReasonEn = map[string]string{
		DEFAULT_BMP_ERROR_CODE: "unknown error message",
		MakeRaidError:          "raid failed",
		MakePartitionsError:    "Partition failed",
		WriteImageTarError:     "Failed to import image",
		SetPasswordError:       "Failed to set password",
		SetCloudinitConfError:  "cloudinit initialization failed",

		SetNocloudUserDataError: "Failed to set user-defined directive",
		SetNocloudMetaDataError: "Failed to set metadata",
		SetNocloudNetworkError:  "Network setup failed",
		PingError:               "PXE start failed",

		DHCPConfigAddHostError: "Add dhcp failed",
		DHCPConfigDelHostError: "Del dhcp failed",
		SetDISKBootError:       "Set disk boot failed",
		PowerCycleError:        "PowerCycle failed",
		PowerOffError:          "PowerOff failed",
		PowerOnError:           "PowerOn failed",
		SetPXEBootError:        "SetPXEBoot failed",

		InitRootDeviceError: "InitRootDevice failed",
		//windows专用指令
		WriteImageError:     "Failed to import image",
		SetMetaDataError:    "Failed to set metadata",
		SetNetworkDataError: "Network setup failed",
	}
	InstanceLock = map[string]string{
		LOCKED:   "锁定",
		UNLOCKED: "解锁",
	}
	InstanceLockEn = map[string]string{
		LOCKED:   "Locked",
		UNLOCKED: "Unlocked",
	}
	InstanceStatus = map[string]string{
		CREATING:           "创建中",
		STARTING:           "开机中",
		RUNNING:            "运行",
		STOPPING:           "关机中",
		STOPPED:            "已关机",
		RESTARTING:         "重启中",
		RESETTING_PASSWORD: "重置密码中",
		DESTROYING:         "删除中",
		DESTROYED:          "已删除",
		ERROR:              "错误",
		// 创建失败
		CREATE_ERROR: "创建失败",
		// 开机失败
		START_ERROR: "开机失败",
		// 关机失败
		STOP_ERROR: "关机失败",
		// 重启失败
		RESTART_ERROR: "重启失败",
		//重装失败
		REINSTALL_ERROR: "重装失败",
		//重置密码失败
		RESETPASSWD_ERROR: "重置密码失败",
		// 销毁失败
		DESTROY_ERROR: "删除失败",
		UPGRADING:     "调整配置中",
		REINSTALLING:  "重装系统中",
	}
	InstanceStatusEn = map[string]string{
		CREATING:           "Creating",
		STARTING:           "Starting up",
		RUNNING:            "running",
		STOPPING:           "Shutting Down",
		STOPPED:            "Shut Down",
		RESTARTING:         "Rebooting",
		RESETTING_PASSWORD: "Resetting password",
		DESTROYING:         "Destroying",
		DESTROYED:          "destroyed",
		ERROR:              "Error",
		// 创建失败
		CREATE_ERROR: "Creation failed",
		// 开机失败
		START_ERROR: "Startup failed",
		// 关机失败
		STOP_ERROR: "Shutdown failed",
		// 重启失败
		RESTART_ERROR: "Reboot failed",
		//重装失败
		REINSTALL_ERROR: "Reinstall failed",
		//重置密码失败
		RESETPASSWD_ERROR: "Resetpasswd failed",
		// 销毁失败
		DESTROY_ERROR: "Delete failed",
		UPGRADING:     "Upgrading",
		REINSTALLING:  "Reinstalling System",
	}
	DeviceTypeSeries = map[string]string{
		COMPUTER: "计算型",
		STORAGE:  "存储型",
		GPU:      "GPU",
		OTHER:    "其它",
	}
	DeviceTypeSeriesEn = map[string]string{
		COMPUTER: "Computer Class",
		STORAGE:  "Storage Class",
		GPU:      "GPU",
		OTHER:    "Other Classes",
	}
	Source = map[string]string{
		COMMON:       "预置镜像",
		USER_DEFINED: "自定义镜像",
	}
	SourceEn = map[string]string{
		COMMON:       "Preset image",
		USER_DEFINED: "Custom image",
	}
	VolumeType = map[string]string{
		VOLUME_TYPE_SYSTEM: "系统卷",
		VOLUME_TYPE_DATA:   "数据卷",
	}
	VolumeTypeEn = map[string]string{
		VOLUME_TYPE_SYSTEM: "system",
		VOLUME_TYPE_DATA:   "data",
	}
	RaidCan = map[string]string{
		RAID_CAN_NO_RAID:     "NO RAID",
		RAID_CAN_RAID:        "RAID",
		RAID_CAN_SINGLE_RAID: "单盘RIAD0",
	}
	RaidCanEn = map[string]string{
		RAID_CAN_NO_RAID:        "NO RAID",
		RAID_CAN_RAID:           "RAID",
		RAID_CAN_SINGLE_RAID_EN: RAID_CAN_SINGLE_RAID_EN,
	}
)

const (
	IS_DEL          = 1
	IS_NOT_DEL      = 0
	DATE_FORMAT     = "2006-01-02 15:04:05"
	DATE_UTC_FORMAT = "2006-01-02T15:04:05Z"
	IS_ALL          = "1"
	IS_NOT_ALL      = ""

	EN_US = "en_US"
	ZH_CN = "zh_CN"

	IN           = "in"           //已入库
	PUTAWAY      = "putaway"      //已上架
	CREATED      = "created"      //已创建
	PUTAWAYING   = "putawaying"   //上架中
	PUTAWAYFAIL  = "putawayfail"  //上架失败
	UNPUTAWAY    = "unputaway"    //下架
	UNPUTAWAYING = "unputawaying" //下架中，将“已上架”的设备执行“下架”操作的中间过程状态
	REMOVED      = "removed"      //已移除

	DEFAULT_BMP_ERROR_CODE     = "ERROR"
	USERNAME_OR_PASSWORD_ERROR = "USERNAME_OR_PASSWORD_ERROR"
	USER_PRIVILEGE_ERROR       = "USER_PRIVILEGE_ERROR"
	PING_ILOIP_ERROR           = "PING_ILOIP_ERROR"
	PING_SUBNET_GATEWAY_ERROR  = "PING_SUBNET_GATEWAY_ERROR"
	MAC1_ERROR                 = "MAC1_ERROR"
	MAC2_ERROR                 = "MAC2_ERROR"

	//instance error
	DHCPConfigAddHostError = "DHCPConfigAddHostError"
	DHCPConfigDelHostError = "DHCPConfigDelHostError"
	SetDISKBootError       = "SetDISKBootError"
	PowerCycleError        = "PowerCycleError"
	PowerOffError          = "PowerOffError"
	PowerOnError           = "PowerOnError"
	SetPXEBootError        = "SetPXEBootError"

	InitRootDeviceError = "InitRootDeviceError"
	//windows专用指令
	WriteImageError     = "WriteImageError"
	SetMetaDataError    = "SetMetaDataError"
	SetNetworkDataError = "SetNetworkDataError"

	PingError     = "PingError"
	MakeRaidError = "MakeRaidError"

	MakePartitionsError   = "MakePartitionsError"
	WriteImageTarError    = "WriteImageTarError"
	SetPasswordError      = "SetPasswordError"
	SetCloudinitConfError = "SetCloudinitConfError"

	SetNocloudUserDataError = "SetNocloudUserDataError"
	SetNocloudMetaDataError = "SetNocloudMetaDataError"
	SetNocloudNetworkError  = "SetNocloudNetworkError"

	LOCKED                = "locked"
	UNLOCKED              = "unlocked"
	DEFAULTPROJECTNAME    = "默认项目"
	DEFAULTPROJECTNAME_EN = "default project"
	//实例状态相关
	CREATING           = "creating"
	STARTING           = "starting"
	RUNNING            = "running"
	STOPPING           = "stopping"
	STOPPED            = "stopped"
	RESTARTING         = "restarting"
	RESETTING_PASSWORD = "resetting_password"
	DESTROYING         = "destroying"
	DESTROYED          = "destroyed"
	ERROR              = "error"
	// 创建失败
	CREATE_ERROR = "Creation failed"
	// 开机失败
	START_ERROR = "Startup failed"
	// 关机失败
	STOP_ERROR = "Shutdown failed"
	// 重启失败
	RESTART_ERROR = "Reboot failed"
	// 重装失败
	REINSTALL_ERROR = "Reinstall failed"
	//重置密码失败
	RESETPASSWD_ERROR = "Resetpasswd failed"
	// 销毁失败
	DESTROY_ERROR = "Delete failed"
	UPGRADING     = "upgrading"
	REINSTALLING  = "reinstalling"
	//机型相关
	COMPUTER = "computer"
	STORAGE  = "storage"
	GPU      = "gpu"
	OTHER    = "other"

	COMMON       = "common"       //内置镜像
	USER_DEFINED = "user_defined" //自定义

	ROLE_ADMIN_UUID    = "role-admin-uuid-01"    //平台拥有着
	ROLE_OPERATOR_UUID = "role-operator-uuid-01" //运营平台
	ROLE_USER_UUID     = "role-user-uuid-01"     //控制台

	VOLUME_TYPE_SYSTEM = "system" //卷类型
	VOLUME_TYPE_DATA   = "data"

	RAID_CAN_NO_RAID        = "NO RAID"
	RAID_CAN_RAID           = "RAID"
	RAID_CAN_SINGLE_RAID    = "单盘RAID0"
	RAID_CAN_SINGLE_RAID_EN = "RAID0-stripping"

	NOT_LIMITED = "notLimited"
)
