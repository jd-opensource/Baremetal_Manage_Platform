package BaseLogic

import (
	"fmt"
	"os"
	"runtime/debug"
)

const EN_US = "en_US"
const ZH_CN = "zh_CN"

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
)

func init() {
	_, e := os.Stat(LOCAL_FILE_PATH)
	if e != nil {
		if os.IsNotExist(e) {
			if e := os.MkdirAll(LOCAL_FILE_PATH, os.ModePerm); e != nil {
				fmt.Println(fmt.Sprintf("create LOCAL_FILE_PATH error, %v\n%s", e, debug.Stack()))
				return
			}
		}
	}

}

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
	HeaderList = map[string]map[string]map[string]map[string]bool{
		//实例自定义表头
		"instanceList": {
			"customInfo": {
				"instanceName": {
					"selected": true, //默认选中
					"required": true, //true代表灰色按钮，直接就是选中，不可更改选中状态，默认就是选中
				},
				"status": {
					"selected": true,  //默认选中
					"required": false, //true代表灰色按钮，直接就是选中，不可更改选中状态，默认就是选中
				},
				"deviceTypeName": { //机型
					"selected": true,
					"required": false,
				},
				"imageName": { //镜像
					"selected": true,
					"required": false,
				},
				"iloIp": {
					"selected": true,
					"required": false,
				},
				"privateIpv4": {
					"selected": true,
					"required": false,
				},
				"privateEth1Ipv4": {
					"selected": true,
					"required": false,
				},
				"privateIpv6": {
					"selected": true,
					"required": false,
				},
				"privateEth1Ipv6": {
					"selected": true,
					"required": false,
				},
				"config": {
					"selected": true,
					"required": false,
				},
				"createdTime": {
					"selected": true,
					"required": false,
				},
			},
		},
		//项目自定义表头
		"projectList": {
			"customInfo": {
				"projectName": {
					"selected": true,
					"required": true,
				},
				"projectId": { //
					"selected": true,
					"required": false,
				},
				"instanceCount": { //
					"selected": true,
					"required": false,
				},
				"createdTime": { //
					"selected": true,
					"required": false,
				},
				"updatedTime": { //
					"selected": true,
					"required": false,
				},
				"owned": { //
					"selected": true,
					"required": false,
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
					"required": true,
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
	ExportInstanceHeader = []string{
		"实例名称/ID",
		"状态",
		"机型",
		"镜像",
		"带外IP",
		"内网IPv4(eth0)",
		"内网IPv4(eth1)",
		"内网IPV6(eth0)",
		"内网IPV6(eth1)",
		"配置",
		"创建时间",
	}

	ExportInstanceHeaderEn = []string{
		"Instance Name/ID",
		"Status",
		"Instance Type",
		"Image",
		"Out-of-band IP",
		"Private IPv4(eth0)",
		"Private IPv4(eth1)",
		"Private IPv6(eth0)",
		"Private IPv6(eth1)",
		"Configuration",
		"Created Time",
	}

	ExportProjectHeader = []string{
		"项目名称",
		"项目ID",
		"项目协作状态",
		"实例数",
		"创建时间",
		"更新时间",
	}

	ExportProjectHeaderEn = []string{
		"ProjectName",
		"ProjectID",
		"sharedStatus",
		"Instances",
		"CreateTime",
		"UpdateTime",
	}

	ExportIdcHeader = []string{
		"机房名称",
		"机房编码",
		"机房等级",
		"机房地址",
		"创建时间",
		"创建人",
		"更新时间",
		"更新人",
	}

	ExportIdcHeaderEn = []string{
		"IdcName",
		"Cabinet",
		"Level",
		"Address",
		"CreateTime",
		"CreatedBy",
		"UpdatedTime",
		"UpdatedBy",
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
		"bond": "单网口",
		"dual": "双网口",
	}
	InterfaceModeEn = map[string]string{
		"bond": "bond",
		"dual": "dual",
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
		"Creation failed":    "创建失败",
		"Startup failed":     "开机失败",
		"Shutdown failed":    "关机失败",
		"Reboot failed":      "重启失败",
		"Delete failed":      "删除失败",
	}

	InstanceStatusEn = map[string]string{
		"creating":           "Creating",
		"starting":           "Starting up",
		"running":            "Running",
		"stopping":           "Shutting Down",
		"stopped":            "Stopped",
		"restarting":         "Restarting",
		"resetting_password": "Reset Password",
		"destroying":         "Destructing",
		"destroyed":          "Destroyed",
		"error":              "Error",
		"upgrading":          "Upgrading",
		"reinstalling":       "Reinstalling System",
		"Creation failed":    "Creation failed",
		"Startup failed":     "Startup failed",
		"Shutdown failed":    "Shutdown failed",
		"Reboot failed":      "Reboot failed",
		"Delete failed":      "Delete failed",
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

	ExportRulesHeader = []string{
		"规则名称",
		"规则ID",
		"资源类型",
		"实例关联数量",
		"触发条件",
		"状态",
	}

	ExportRulesHeaderEn = []string{
		"Rule Name",
		"Rule ID",
		"Resource Type",
		"Instance associated number",
		"Triggering conditions",
		"Status",
	}
	ExportAlertsHeader = []string{
		"报警时间",
		"报警资源ID",
		"触发条件",
		"报警值",
		"报警级别",
		"持续时间",
		"通知对象",
	}

	ExportAlertsHeaderEn = []string{
		"Alarm time",
		"Alarm Resources ID",
		"Trigger conditions",
		"Alarm value",
		"Alarm level",
		"duration",
		"Notification object",
	}
)
