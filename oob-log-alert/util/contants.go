package util

const (
	EN_US = "en_US"
	CN_ZH = "zh_CN"
)

//设备是否有带外故障
var statusDesc = map[int]string{
	0: "正常",
	1: "异常",
}
var statusDescEn = map[int]string{
	0: "normal",
	1: "unexpected",
}

//带外故障是否已处理
var logDealStatusDesc = map[int]string{
	0: "未处理",
	1: "已处理",
}
var logDealStatusDescEn = map[int]string{
	0: "Unhandled",
	1: "Handled",
}

var pushStatusDesc = map[int]string{
	0: "未启用",
	1: "已启用",
}

var pushStatusDescEn = map[int]string{
	0: "disabled",
	1: "enabled",
}

func StatusIntDesc(status int, language string) string {
	if language == EN_US {
		return statusDescEn[status]
	} else {
		return statusDesc[status]
	}
}

func LogDealStatusIntDesc(status int, language string) string {
	if language == EN_US {
		return logDealStatusDescEn[status]
	} else {
		return logDealStatusDesc[status]
	}
}

func PushStatusIntDesc(status int, language string) string {
	if language == EN_US {
		return pushStatusDescEn[status]
	} else {
		return pushStatusDesc[status]
	}
}

var (
	FormatTime = "2006-01-02 15:04:05"
	FormatDate = "2006-01-02"

	ExportDeviceStatusHeader = []string{
		"SN",
		"机房",
		"品牌",
		"型号",
		"管理状态",
		"CPU",
		"内存",
		"硬盘",
		"网卡",
		"电源",
		"其他",
	}
	ExportDeviceStatusHeaderEn = []string{
		"SN",
		"Idc Name",
		"Brand",
		"Model",
		"Manage Status",
		"CPU",
		"Memory",
		"Harddisk",
		"NIC",
		"POWER Supply",
		"Other",
	}

	ExportLogCollectionsHeader = []string{
		"LogID",
		"SN",
		"机房",
		"故障等级",
		"故障类型",
		"故障配件",
		"iloIP",
		"故障描述",
		"故障原始日志",
		"状态",
		"报警次数",
		"带外故障发现时间",
		"首次故障报警时间",
		"更新时间",
	}

	ExportLogCollectionsHeaderEn = []string{
		"LogID",
		"SN",
		"Idc Name",
		"Failure Level",
		"Fault Type",
		"Faulty accessories",
		"iloIP",
		"Fault Desc",
		"Fault Original log",
		"Status",
		"Number of alarms",
		"Out-of-band fault detection time",
		"First fault alarm time",
		"Update Time",
	}

	ExportLogCollectionBySnHeader = []string{
		"序号",
		"故障类型",
		"故障描述",
		"故障报警时间",
		"更新时间",
		"报警次数",
		"状态",
	}

	ExportLogCollectionBySnHeaderEn = []string{
		"LogID",
		"Fault Type",
		"Fault Desc",
		"Out-of-band fault detection time",
		"Update Time",
		"Number of alarms",
		"Status",
	}

	ExportRulesHeader = []string{
		"故障名称",
		"故障配件",
		"故障类型",
		"判定条件",
		"判定阈值",
		"故障等级",
		"故障描述",
		"故障推送",
		"启用状态",
	}

	ExportRulesHeaderEn = []string{
		"Fault Name",
		"Faulty accessories",
		"Fault Type",
		"Judgment Conditions",
		"Decision Threshold",
		"Failure Level",
		"Fault Description",
		"Failure Push",
		"Enabled Status",
	}
)

var HeaderList = map[string]map[string]map[string]map[string]bool{
	//运营平台故障日志自定义列表
	"alert_log_list": {
		"custom_info": {
			"logid": {
				"selected": true, //默认选中
				"required": true, //true代表灰色按钮，直接就是选中，不可更改选中状态，默认就是选中
			},
			"sn": {
				"selected": true, //默认选中
				"required": true, //true代表灰色按钮，直接就是选中，不可更改选中状态，默认就是选中
			},
			"idc_name": { //
				"selected": true,
				"required": false,
			},
			"alert_level": { //体系架构
				"selected": true,
				"required": false,
			},
			"alert_type": {
				"selected": true,
				"required": false,
			},
			"alert_part": {
				"selected": true,
				"required": false,
			},
			"ilo_ip": {
				"required": false,
				"selected": true,
			},
			"alert_desc": {
				"selected": true,
				"required": false,
			},
			"alert_content": {
				"selected": true,
				"required": false,
			},
			"status": {
				"selected": true,
				"required": false,
			},
			"count": {
				"selected": true,
				"required": false,
			},
			"collect_time": {
				"selected": true,
				"required": false,
			},
			"event_time": {
				"selected": true,
				"required": false,
			},
			"update_time": {
				"selected": true,
				"required": false,
			},
			"operate": {
				"selected": true,
				"required": true,
			},
		},
	},
	//控制台故障日志自定义列表
	"alert_log_by_sn_list": {
		"custom_info": {
			"logid": {
				"selected": true, //默认选中
				"required": true, //true代表灰色按钮，直接就是选中，不可更改选中状态，默认就是选中
			},
			"alert_type": {
				"selected": true,
				"required": false,
			},
			"alert_desc": {
				"selected": true,
				"required": false,
			},
			"event_time": {
				"selected": true,
				"required": true,
			},
			"update_time": {
				"selected": true,
				"required": true,
			},
			"count": {
				"selected": true,
				"required": true,
			},
			"status": {
				"selected": true,
				"required": false,
			},
			"operate": {
				"selected": true,
				"required": true,
			},
		},
	},
	//运营平台故障规则自定义列表
	"alert_rule_list": {
		"custom_info": {
			"alert_name": {
				"selected": true,
				"required": false,
			},
			"alert_spec": {
				"selected": true,
				"required": false,
			},
			"alert_type": {
				"selected": true,
				"required": false,
			},
			"alert_condition": {
				"selected": true,
				"required": false,
			},
			"alert_threshold": {
				"selected": true,
				"required": false,
			},
			"alert_level": {
				"selected": true,
				"required": false,
			},
			"alert_desc": {
				"selected": true,
				"required": false,
			},
			"push_status": {
				"selected": true,
				"required": true,
			},
			"use_status": {
				"selected": true,
				"required": true,
			},
			"operate": {
				"selected": true,
				"required": true,
			},
		},
	},
	//运营平台硬件设备状态
	"device_status_list": {
		"custom_info": {
			"sn": {
				"selected": true,
				"required": true,
			},
			"idc_name": {
				"selected": true,
				"required": false,
			},
			"brand": {
				"selected": true,
				"required": false,
			},
			"model": {
				"selected": true,
				"required": false,
			},
			"manage_status": {
				"selected": true,
				"required": false,
			},
			"cpu_status": {
				"selected": true,
				"required": false,
			},
			"mem_status": {
				"selected": true,
				"required": false,
			},
			"disk_status": {
				"selected": true,
				"required": false,
			},
			"nic_status": {
				"selected": true,
				"required": false,
			},
			"power_status": {
				"selected": true,
				"required": false,
			},
			"other_status": {
				"selected": true,
				"required": false,
			},
			"operate": {
				"selected": true,
				"required": true,
			},
		},
	},
	//运营平台设备管理硬件监控状态
	"device_alert_log_list": {
		"custom_info": {
			"logid": {
				"selected": true, //默认选中
				"required": true, //true代表灰色按钮，直接就是选中，不可更改选中状态，默认就是选中
			},
			"alert_type": {
				"selected": true,
				"required": false,
			},
			"alert_desc": {
				"selected": true,
				"required": false,
			},
			"event_time": {
				"selected": true,
				"required": false,
			},
			"update_time": {
				"selected": true,
				"required": false,
			},
			"count": {
				"selected": true,
				"required": false,
			},
			"status": {
				"selected": true,
				"required": true,
			},
			"operate": {
				"selected": true,
				"required": true,
			},
		},
	},
}
