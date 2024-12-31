package request

import (
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
)

type DeleteInstanceRequest struct {
	HeaderRequest
}

type GetInstanceStatusRequest struct {
	InstanceID string `json:"instanceId"`
}

type LockInstanceRequest struct {
	HeaderRequest
	InstanceId string `json:"instanceId"`
}

type UnLockInstanceRequest struct {
	HeaderRequest
	InstanceId string `json:"instanceId"`
}

type StopInstanceRequest struct {
	HeaderRequest
	InstanceId  string `json:"instanceId"`
	ClientToken int    `json:"clientToken"`
}

type StartInstanceRequest struct {
	HeaderRequest
	InstanceId  string `json:"instanceId"`
	ClientToken int    `json:"clientToken"`
}

type RestartInstanceRequest struct {
	HeaderRequest
	InstanceId  string `json:"instanceId"`
	ClientToken int    `json:"clientToken"`
}

type ModifyInstanceRequest struct {
	HeaderRequest

	// required: true
	InstanceID string `json:"instanceId"`

	Name string `json:"name"`
	// required: true
	Description *string `json:"description"`
}

type QueryInstanceRequest struct {
	HeaderRequest
	InstanceId string `json:"instanceId" valid:"Required"`
}

type QueryInstancesRequest struct {
	PagingRequest
	ProjectID    string `json:"projectId"`
	IdcID        string `json:"idcId"`
	Status       string `json:"status"` // v2.27.2 状态，逗号分隔；显示所有选项，复选框选择，支持多选。旧代码：private List<String> status string `json:""`
	DeviceTypeId string `json:"deviceTypeId"`
	Name         string `json:"name"` // v2.27.2 实例列表搜索框支持联合搜索| 模糊搜索，单个条件值搜索

	DeviceType string `json:"deviceType"`

	InstanceId string `json:"instanceId"`

	PrivateIp string `json:"privateIp"`

	PrivateIpv6 string `json:"privateIpv6"`

	KeypairId string `json:"keypairId"`

	Sn string `json:"sn"` // v2.18.1 新增

	IloIp string `json:"iloIp"` // v2.27.2 带外管理IP，逗号分隔； 精确搜索，多个条件值搜索

	//Cabinets string `json:"cabinets"` // v2.27.2 机柜编码，逗号分隔；精确搜索，多个条件值搜索；如：T4-L3-152-Test

	DeviceTypeList string `json:"deviceTypeList"` // v2.27.2 规格类型，逗号分隔；显示所有选项，复选框选择，支持多选；精确搜索，多个条件值搜索；如：cps.c2.per2.highdensity

	ExportType string `json:"exportType"` // v2.27.2   0：导出全部实力；1：导出选中实例；2：导出搜索结果
	// "1" 表示全部
	IsAll string `json:"isAll"`

	//是否安装了agent，"0"表示未安装，"1"表示已安装,不传表示全部
	IsInstallAgent string `json:"isInstallAgent"`
}

type RemoveInstancesRequest struct {
	HeaderRequest
	InstanceIds []string `json:"instanceIds"`
}

type StartInstancesRequest struct {
	InstanceIDs []string `json:"instanceIds"`
}

type ModifyInstancesRequest struct {
	InstanceIDs  []string `json:"instanceIds"`
	InstanceName string   `json:"instanceName"`
}

type ResetPasswdRequest struct {
	HeaderRequest
	InstanceId string `json:"instanceId" valid:"Required"`
	Password   string `json:"password"`
}

type ResetInstancesPasswdRequest struct {
	HeaderRequest
	InstanceIds []string `json:"instanceIds" valid:"Required"`
	Password    string   `json:"password"`
}

type ReinstallInstanceRequest struct {
	HeaderRequest
	InstanceId string `json:"instanceId" valid:"Required"`
	sdkModels.ReinstallInstanceRequest
}

type QueryInstancesForShareRequest struct {
	ProjectID  string `json:"projectId"`
	OwnerName  string `json:"ownerName"`
	SharerName string `json:"sharerName"`

	//支持搜索
	InstanceName string `json:"instanceName"`
	InstanceId   string `json:"instanceId"`
}
