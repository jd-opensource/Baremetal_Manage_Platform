package BaseLogic

const (
	SUCCESS_CODE_OPENAPI = 0 //openapi返回的code，成功返回0，如果不是0，说明有错误，需要抛出错误信息
	ERROR_CODE           = 400
	EOPEN_CODE           = 500
	ERROR_REDIS_CODE     = 401
	ERROR_MYSQL_CODE     = 402

	LOCAL_FILE_PATH          = "./log/batch/"
	UPLOAD_FILE_FORMAT_ERROR = "20201"
	UPLOAD_FILE_SIZE_ERROR   = "20202"
	UPLOAD_FILE_BASE64       = "20203"
	UPLOAD_FILE_NOT_BASE64   = "20204"

	//uc的调用路径配置

	USERCENTER_GET_PRIVILEGENEW = "/usercenter/getPrivilegeNew"
	ORDER_SUBMITRENEW           = "/operate/submitRenew"
	//以下错误码返给前端特定的错误提示，并且支持国际化
	ERROR_NO                    = 400001 //通用返回错误码
	ERROR_NO_INSTANCE_NOT_FOUND = "400002"
	ERROR_NO_STRING_TO_STRING   = 400003

	COMMON       = "common"       //内置镜像
	USER_DEFINED = "user_defined" //自定义

	EN_US = "en_US"
	CN_ZH = "zh_CN"
)

var (
	code      int
	chMessage string
	enMessage string
	status    string
)
var (
	FileError = map[string]string{
		UPLOAD_FILE_FORMAT_ERROR: "uploadFileFormatError",
		UPLOAD_FILE_SIZE_ERROR:   "uploadFileSizeError",
		UPLOAD_FILE_BASE64:       "uploadFileBase64",
		UPLOAD_FILE_NOT_BASE64:   "uploadFileNotBase64",
	}
	FormatTime  = "2006-01-02 15:04:05"
	FormatDate  = "2006-01-02"
	MountPoints = map[string]map[string][]string{
		"points": map[string][]string{
			"data":   []string{"/var", "/export", "/data0"},
			"system": []string{"/", "/boot", "swap", "/var", "/export", "/data0"},
		},
	}
	OSFileSystemFormats = map[string][]string{
		"formats": []string{"XFS", "EXT3", "EXT4", "SWAP"},
	}
	Cpulist = []string{
		"Intel,E5-2620V4,2,8,2.1",
		"Intel,E5-2650V4,2,12,2.2",
		"Intel,E5-2683V4,2,16,2.1",
		"Intel,Silver 4116,2,12,2.1",
		"Intel,Silver 4210,2,10,2.2",
		"Intel,Gold 5218,2,16,2.3",
		"Intel,Gold 6146,2,12,3.2",

		"Intel,Gold 6148,2,20,2.4",
		"Intel,GOLD 6267,2,24,2.6",
		"Intel,Platinum 8338,2,32,2.6",
		"AMD,EPYC-7H12,2,64,2.6",
		"AMD,EPYC-7642,2,48,2.3",

		"Ampere®,Altra® Q80-30,2,80,3.0",
		"Kunpeng,920-7260,2,64,2.60",
	}
	Memlist = []string{
		"128 16 8 DDR3 1600",
		"192 16 12 DDR4 2400",
		"256 32 8 DDR4 2666",
		"384 32 12 DDR4 2933",
		"512 32 16 DDR4 2933",
		"768 64 12 DDR4 3200",
		"1024 64 16 DDR4 3200",
	}

	HeaderList = map[string]map[string]map[string]map[string]bool{
		//镜像自定义列表
		"imageList": {
			"customInfo": {
				"imageName": {
					"selected": true, //默认选中
					"required": true, //true代表灰色按钮，直接就是选中，不可更改选中状态，默认就是选中
				},
				"imageId": {
					"selected": true,  //默认选中
					"required": false, //true代表灰色按钮，直接就是选中，不可更改选中状态，默认就是选中
				},
				"source": { //内置镜像
					"selected": true,
					"required": false,
				},
				"architecture": { //体系架构
					"selected": true,
					"required": false,
				},
				"osType": {
					"selected": true,
					"required": false,
				},
				"osVersion": {
					"selected": true,
					"required": false,
				},
				"systemPartition": {
					"required": false,
					"selected": true,
				},
				"description": {
					"selected": true,
					"required": false,
				},
				"deviceTypeNum": {
					"selected": true,
					"required": false,
				},
				"createdTime": {
					"selected": true,
					"required": false,
				},
				"operate": {
					"selected": true,
					"required": true,
				},
				"bootMode": {
					"required": false,
					"selected": true,
				},
				"format": {
					"required": false,
					"selected": true,
				},
			},
		},
		//设备列自定义表头
		"deviceList": {
			"customInfo": {
				"sn": {
					"selected": true,
					"required": true,
				},
				"deviceTypeName": { // 机型
					"selected": true,
					"required": false,
				},
				"architecture": { // 体系架构
					"selected": true,
					"required": false,
				},
				"description": { // 描述
					"selected": true,
					"required": false,
				},
				"deviceSeries": { // 机型类型
					"selected": true,
					"required": false,
				},

				"manageStatusName": { // 管理状态
					"selected": true,
					"required": false,
				},
				"idcName": { //机房名称
					"selected": true,
					"required": false,
				},
				"cabinet": { // 机柜编码
					"selected": true,
					"required": false,
				},
				"uPosition": { // 所在U位
					"selected": true,
					"required": false,
				},
				"brand": { // 品牌
					"selected": true,
					"required": false,
				},

				"model": { // 型号
					"selected": true,
					"required": false,
				},
				"imageName": { // 实例镜像
					"selected": true,
					"required": false,
				},
				"instanceName": { // 实例名称
					"selected": true,
					"required": false,
				},
				"instanceId": { // 实例ID
					"selected": true,
					"required": false,
				},
				"instanceStatus": { // 实例状态
					"selected": true,
					"required": false,
				},
				"instanceLockedStatus": { // 实例锁定状态
					"selected": true,
					"required": false,
				},

				"instanceOwer": { // 实例属主
					"selected": true,
					"required": false,
				},
				"cpuInfo": { // CPU 英特尔 I7 (10*8 物理核, 2.5GHz)
					"selected": true,
					"required": false,
				},
				"memInfo": { // 内存 256GB DDR4 2933MHz
					"selected": true,
					"required": false,
				},
				"svInfo": { // 系统盘	2880GB (240GB SAS SSD*12,RAID1)
					"selected": true,
					"required": false,
				},
				"dvInfo": { // 数据盘 10GB (1GB SATA SSD*10,RAID1)
					"selected": true,
					"required": false,
				},

				"gpuInfo": { // GPU INVIDA I8*10
					"selected": true,
					"required": false,
				},
				"switchIp1": { // 网口1交换机IP
					"selected": true,
					"required": false,
				},
				"switchIp2": { // 网口2交换机IP
					"selected": true,
					"required": false,
				},
				"switchPort1": { // 网口1（eth0）上联端口
					"selected": true,
					"required": false,
				},
				"switchPort2": { // 网口2（eth1）上联端口
					"selected": true,
					"required": false,
				},

				"nicInfo": { // 网络设置 单网卡 or 双网口
					"selected": true,
					"required": false,
				},
				"iloIp": { // 带外IP
					"selected": true,
					"required": false,
				},
				"privateIpv4": { // 内网IPv4
					"selected": true,
					"required": false,
				},
				"privateEth1Ipv4": { // eth1内网IPv4
					"selected": true,
					"required": false,
				},
				"privateIpv6": { // 内网IPv6
					"selected": true,
					"required": false,
				},
				"privateEth1Ipv6": { // eth1内网IPv6
					"selected": true,
					"required": false,
				},
				"createdTime": { // 创建时间
					"selected": true,
					"required": false,
				},
				"operate": {
					"selected": true,
					"required": true,
				},
				"collectStatus": {
					"selected": true,
					"required": false,
				},
			},
		},
		//设备类型列表自定义表头
		"deviceTypeList": {
			"customInfo": {
				"name": {
					"selected": true,
					"required": true,
				},
				"deviceSeries": {
					"selected": true,
					"required": false,
				},
				"idcName": {
					"selected": true,
					"required": false,
				},
				"idcRegion": {
					"selected": true,
					"required": false,
				},
				"architecture": {
					"selected": true,
					"required": false,
				},
				"deviceType": {
					"selected": true,
					"required": false,
				},

				"image": {
					"selected": true,
					"required": false,
				},
				"cpu": {
					"selected": true,
					"required": false,
				},
				"memory": {
					"selected": true,
					"required": false,
				},
				"systemVolume": {
					"selected": true,
					"required": false,
				},
				"systemVolumeRaid": {
					"selected": true,
					"required": false,
				},
				"raidCan": {
					"selected": true,
					"required": false,
				},
				"dataVolume": {
					"selected": true,
					"required": false,
				},
				"nicInfo": {
					"selected": true,
					"required": false,
				},
				"nic": {
					"selected": true,
					"required": false,
				},
				"gpu": {
					"selected": true,
					"required": false,
				},

				"remark": {
					"selected": true,
					"required": false,
				},
				"operate": {
					"selected": true,
					"required": true,
				},
				"deviceCount": {
					"selected": true,
					"required": false,
				},
				"bootMode": {
					"required": false,
					"selected": true,
				},
			},
		},
		//机房列表自定义表头
		"idcList": {
			"customInfo": {
				"address": {
					"selected": true,
					"required": false,
				},

				"level": {
					"selected": true,
					"required": false,
				},
				"name": {
					"selected": true,
					"required": true,
				},
				"nameEn": {
					"selected": true,
					"required": false,
				},

				"updateTime": {
					"selected": true,
					"required": false,
				},
				"createTime": {
					"selected": true,
					"required": false,
				},

				"createdBy": {
					"selected": true,
					"required": false,
				},
				"updatedBy": {
					"selected": true,
					"required": false,
				},
				"operate": {
					"selected": true,
					"required": true,
				},
			},
		},
		// 用户列表自定义表头
		"userList": {
			"customInfo": {
				"userName": {
					"selected": true,
					"required": true,
				},
				"userId": {
					"selected": true,
					"required": false,
				},
				"roleName": {
					"selected": true,
					"required": false,
				},
				"phone": {
					"selected": true,
					"required": false,
				},
				"email": {
					"selected": true,
					"required": false,
				},
				"remark": {
					"selected": true,
					"required": false,
				},
				"createTime": {
					"selected": true,
					"required": false,
				},
				"projectCount": {
					"selected": true,
					"required": false,
				},
				"operate": {
					"selected": true,
					"required": true,
				},
			},
		},
		// 角色列表自定义表头
		"roleList": {
			"customInfo": {
				"roleName": {
					"selected": true,
					"required": true,
				},
				"userName": {
					"selected": true,
					"required": false,
				},
				"permission": {
					"selected": true,
					"required": false,
				},
				"operate": {
					"selected": true,
					"required": true,
				},
			},
		},
		//消息自定义表头
		"messageList": {
			"customInfo": {
				"detail": {
					"selected": true,
					"required": true,
				},
				"finish_time": { //
					"selected": true,
					"required": false,
				},
				"message_type": { //
					"selected": true,
					"required": false,
				},
				"message_sub_type": { //
					"selected": true,
					"required": false,
				},
			},
		},
		"auditLogsList": {
			"customInfo": {
				"logid": {
					"selected": true,
					"required": true,
				},
				"operationName": { //
					"selected": true,
					"required": true,
				},
				"userName": { //
					"selected": true,
					"required": false,
				},
				"operateTime": { //
					"selected": true,
					"required": false,
				},
				"result": { //
					"selected": true,
					"required": false,
				},
				"failReason": { //
					"selected": true,
					"required": false,
				},
			},
		},
		"monitorRulesList": {
			"customInfo": {
				"ruleName": {
					"selected": true,
					"required": false,
				},
				"ruleId": { //
					"selected": true,
					"required": false,
				},
				"resourceName": { //
					"selected": true,
					"required": false,
				},
				"instanceCount": { //
					"selected": true,
					"required": false,
				},
				"triggerDescription": { //
					"selected": true,
					"required": false,
				},
				"statusName": { //
					"selected": true,
					"required": false,
				},
				"operate": {
					"required": true,
					"selected": true,
				},
				"userName": {
					"required": true,
					"selected": true,
				},
			},
		},
		"monitorAlertsList": {
			"customInfo": {
				"alertTime": {
					"selected": true,
					"required": true,
				},
				"resourceName": { //
					"selected": true,
					"required": false,
				},
				"triggerDescription": { //
					"selected": true,
					"required": false,
				},
				"alertValue": { //
					"selected": true,
					"required": false,
				},
				"alertLevelDescription": { //
					"selected": true,
					"required": false,
				},
				"alertPeriod": { //
					"selected": true,
					"required": false,
				},
				"userName": { //
					"selected": true,
					"required": false,
				},
			},
		},
	}
	ExportImageHeader = []string{
		"镜像名称",
		"镜像ID",
		"镜像类型",
		"体系架构",
		"操作系统平台",
		"操作系统版本",
		"镜像格式",
		"引导模式",
		"系统盘分区",
		"镜像描述",
		"已关联机型数",
		"创建时间",
	}
	ExportImageHeaderEn = []string{
		"Image Name",
		"Image ID",
		"Image Type",
		"Architecture",
		"OS Platform",
		"OS Version",
		"Image Format",
		"Boot Mode",
		"System Volume Partition",
		"Image Desp",
		"Associated ITs",
		"Created Time",
	}

	NetworkType = map[string]string{
		"basic":  "基础网络",
		"vpc":    "私有网络",
		"retail": "定制网络",
	}

	NetworkTypeEn = map[string]string{
		"basic":  "basic",
		"vpc":    "vpc",
		"retail": "retail",
	}

	InterfaceMode = map[string]string{
		"single": "单网口",
		"bond":   "单网口bond",
		"dual":   "双网口",
	}
	InterfaceModeEn = map[string]string{
		"single": "Single Network Interface",
		"bond":   "Single Network Interface bond",
		"dual":   "Double Network Interfaces",
	}
	InstanceStatus = map[string]string{
		"creating":           "创建中",
		"starting":           "开机中",
		"running":            "运行",
		"stopping":           "关机中",
		"stopped":            "关机",
		"restarting":         "重启中",
		"resetting_password": "重置密码中",
		"destroying":         "销毁中",
		"destroyed":          "已销毁",
		"error":              "错误",
		"upgrading":          "调整配置中",
		"reinstalling":       "重装系统中",
	}

	InstanceStatusEn = map[string]string{
		"creating":           "Creating",
		"starting":           "Starting up",
		"running":            "running",
		"stopping":           "Shutting Down",
		"stopped":            "Shut Down",
		"restarting":         "Rebooting",
		"resetting_password": "Resetting password",
		"destroying":         "Destroying",
		"destroyed":          "destroyed",
		"error":              "Error",
		"upgrading":          "Upgrading",
		"reinstalling":       "Reinstalling System",
	}

	FilterInstanceStatus = map[string]string{ //给前端用的实例列表的fiter筛选接口用
		"running": "运行",
		"stopped": "关机",
		"error":   "错误",
	}

	FilterInstanceStatusEn = map[string]string{ //给前端用的实例列表的fiter筛选接口用
		"running": "running",
		"stopped": "stopped",
		"error":   "error",
	}
	Instance = map[string]map[string]string{
		ERROR_NO_INSTANCE_NOT_FOUND: map[string]string{
			"cn": "资源实例未找到",
			"en": "instance not found",
		},
		//todo
	}
	InstanceLockedName = map[string]string{ //给前端用的设备列表的实力锁定状态
		"locked":   "已锁定",
		"unlocked": "已解锁",
	}

	InstanceLockedNameEn = map[string]string{
		"locked":   "locked",
		"unlocked": "unlocked",
	}

	ExportDeviceHeader = []string{
		"SN",
		"机型名称",
		"机型类型",
		"管理状态",
		"机房名称",
		"机柜编码",
		"所在U位",
		"品牌",
		"型号",
		"体系架构",
		"镜像",

		"实例名称",
		"实例ID",
		"实例状态",
		"实例锁定状态",
		"实例属主",
		"CPU",
		"内存",
		"系统盘",
		"数据盘",
		"GPU",
		"网口1交换机IP",

		"网口2交换机IP",
		"网口1(eth0)上联端口",
		"网口2(eth1)上联端口",
		"网络设置",
		"带外IP",
		"内网IPv4(eth0)",
		"内网IPv4(eth1)",
		"内网IPv6(eth0)",
		"内网IPv6(eth1)",
		"创建时间",
		"备注",
	}
	ExportDeviceHeaderEn = []string{
		"SN",
		"Type Name",
		"Instance Type",
		"Management Status",
		"Name",
		"Cabinet ID",
		"Unit Location",
		"Brand",
		"Model",
		"Architecture",
		"Image",

		"Instance Name",
		"Instance ID",
		"Instance Status",
		"Locked Status",
		"Instance Owner",
		"CPU",
		"Memory",
		"System Volume",
		"Data Volume",
		"GPU",
		"Uplink SW IP（Port #1）",

		"Uplink SW IP（Port #2）",
		"Uplink SW Port #1",
		"Uplink SW Port #2",
		"NIC settings",
		"OOB IP",
		"Private IPv4(eth0)",
		"Private IPv4(eth1)",
		"Private IPv6(eth0)",
		"Private IPv6(eth1)",
		"Created Time",
		"Remark",
	}

	ExportDeviceTypeHeader = []string{
		"机型名称",
		"机型类型",
		"机房名称",
		"机房英文名称",
		"体系架构",
		"引导模式",
		"机型规格",

		"镜像",
		"设备",
		"CPU",
		"内存",
		//"系统盘",
		//"RAID配置",
		//"系统盘RAID",
		//
		//"数据盘",
		//"数据盘 RAID",
		"网卡",
		"网络设置",
		"GPU",

		"描述",
	}

	ExportDeviceTypeHeaderEn = []string{
		"Type Name",
		"Instance Type",
		"Name",         // TODO 机房名称：显示所属的机房
		"English Name", // TODO 机房编码： cn-north-1
		"Architecture",
		"BootMode",
		"Instance  Spec (IS)",

		"Image",
		"Devices",
		"CPU",
		"Memory",
		"System Volume",
		"RAID Config",
		"System Volume RAID Type",
		"Data Volume",
		//"Data Volume Raid Type",
		"NIC",
		"NIC settings",
		"GPU",

		"Description",
	}

	ExportIdcHeader = []string{
		"机房名称",
		"机房英文名称",
		"机房等级",
		"机房地址",
		"创建时间",
		"创建人",
		"更新时间",
		"更新人",
	}

	ExportIdcHeaderEn = []string{
		"Name",
		"English name",
		"Level",
		"Address",
		"Created Time",
		"Creator",
		"Update time",
		"Updater",
	}

	ExportUserHeader = []string{
		"用户名",
		"用户ID",
		"角色",
		"手机",
		"邮箱",
		"描述",
		"创建时间",
	}
	ExportUserHeaderEn = []string{
		"User Name",
		"User ID",
		"Role",
		"Mobile Phone",
		"Email",
		"Description",
		"Create Time",
	}

	ExportMessageHeader = []string{
		"消息内容",
		"接收时间",
		"消息类型",
		"消息子类型",
	}

	ExportMessageHeaderEn = []string{
		"message content",
		"time",
		"message type",
		"message subtype",
	}

	ExportAuditLogsHeader = []string{
		"LogID",
		"操作名称",
		"操作人",
		"操作时间",
		"操作反馈",
		"错误码",
	}

	ExportAuditLogsHeaderEn = []string{
		"LogID",
		"Operation Name",
		"Operator",
		"Operation Time",
		"Operation Feedback",
		"Fault Info",
	}

	OssFilter = map[string]interface{}{
		"architecture": map[string]interface{}{
			"i386": map[string]interface{}{
				"osType": map[string]interface{}{
					"其它": "",
				},
			},
			"x86_64": map[string]interface{}{
				"osType": map[string]interface{}{
					"CentOS":    []string{"7.1", "7.2", "7.4", "7.5", "7.6", "7.8", "7.9", "8.0", "8.2"},
					"Ubuntu":    []string{"14.04", "16.04", "18.04", "20.04"},
					"Windows":   []string{"2012", "2016", "2019", "2022"},
					"Debian":    []string{"10.9"},
					"OpenEuler": []string{"20.03"},
					"其它":        "",
				},
			},
			"ARM64(aarch64)": map[string]interface{}{
				"osType": map[string]interface{}{
					"CentOS":    []string{"8.2"},
					"OpenEuler": []string{"20.03"},
					"其它":        "",
				},
			},
			"LoongArch™": map[string]interface{}{
				"osType": map[string]interface{}{
					"loongnix-server": []string{"8.4"},
				},
			},
			"format": []string{
				"tar",
				"qcow2",
			},
			"bootMode": []string{
				"UEFI",
				"Legacy/BIOS",
			},
		},
	}
	OssFilterEn = map[string]interface{}{
		"architecture": map[string]interface{}{
			"i386": map[string]interface{}{
				"osType": map[string]interface{}{
					"Other": "",
				},
			},
			"x86_64": map[string]interface{}{
				"osType": map[string]interface{}{
					"CentOS":    []string{"6.6", "7.1", "7.2", "7.4", "7.5", "7.6", "7.8", "7.9", "8.0", "8.2"},
					"Ubuntu":    []string{"14.04", "16.04", "18.04", "20.04"},
					"Windows":   []string{"2012", "2016", "2019", "2022"},
					"Debian":    []string{"10.9"},
					"OpenEuler": []string{"20.03"},
					"Other":     "",
				},
			},
			"ARM64(aarch64)": map[string]interface{}{
				"osType": map[string]interface{}{
					"CentOS":    []string{"8.2"},
					"OpenEuler": []string{"20.03"},
					"Other":     "",
				},
			},
			"LoongArch™": map[string]interface{}{
				"osType": map[string]interface{}{
					"loongnix-server": []string{"8.4"},
				},
			},
		},
	}
	ExportRulesHeader = []string{
		"用户名",
		"规则名称",
		"规则ID",
		"资源类型",
		"实例关联数量",
		"触发条件",
		"状态",
	}

	ExportRulesHeaderEn = []string{
		"User Name",
		"Rule Name",
		"Rule ID",
		"Resource Type",
		"Instance associated number",
		"Triggering conditions",
		"Status",
	}
	ExportAlertsHeader = []string{
		"用户名",
		"报警时间",
		"资源类型",
		"报警资源ID",
		"触发条件",
		"报警值",
		"报警级别",
		"持续时间",
		"通知对象",
	}

	ExportAlertsHeaderEn = []string{
		"User Name",
		"Alarm time",
		"Resource Type",
		"Alarm Resources ID",
		"Trigger Conditions",
		"Alarm Value",
		"Alarm Level",
		"Duration",
		"Notification Object",
	}
)
