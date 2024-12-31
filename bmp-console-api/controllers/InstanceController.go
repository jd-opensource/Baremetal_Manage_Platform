package controllers

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-console-api/logic/InstanceLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	response "coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"

	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
)

type InstanceController struct {
	BaseController
}

// swagger:route GET /instance instance getInstanceList
// GetInstanceList 实例列表
//     Responses:
//       200: getInstanceList
//       default: ErrorResponse

func (c *InstanceController) GetInstanceList() {
	defer c.CatchException()
	pr := requestTypes.PagingRequest{
		PageNumber: c.pageable.PageNumber,
		PageSize:   c.pageable.PageSize,
	}
	param := requestTypes.QueryInstancesRequest{
		PagingRequest: pr,
		ProjectID:     c.GetString("projectId"),
		IdcID:         c.GetString("idcId"),
		Status:        c.GetString("status"),       // v2.27.2 状态，逗号分隔；显示所有选项，复选框选择，支持多选。
		DeviceTypeId:  c.GetString("deviceTypeId"), //按照机型搜索
		Name:          c.GetString("name"),         //v2.27.2 实例列表搜索框支持联合搜索| 模糊搜索
		DeviceType:    c.GetString("deviceType"),
		InstanceId:    c.GetString("instanceId"),

		KeypairId:   c.GetString("keypairId"),
		Sn:          c.GetString("sn"),
		PrivateIp:   c.GetString("privateIp"),
		PrivateIpv6: c.GetString("privateIpv6"),
		IloIp:       c.GetString("iloIp"), // 带外管理IP，逗号分隔
		// Cabinets:       c.GetString("cabinets"),       //机柜编码，逗号分隔
		DeviceTypeList: c.GetString("deviceTypeList"), // 规格类型，逗号分隔
		ExportType:     c.GetString("exportType"),
		IsAll:          c.GetString("isAll"),
		IsInstallAgent: c.GetString("isInstallAgent"),
	}
	res, err := InstanceLogic.DescribeInstances(c.logPoints, param) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	if res.Instances == nil {
		res.Instances = make([]response.Instance, 0)
	}
	if c.GetString("exportType") != "" {
		fileName, downloadFileName := InstanceLogic.ExportInstance(c.logPoints, res.Instances)
		//c.Ctx.Input.SetData("download", "yes")
		c.Ctx.Output.Download(fileName, downloadFileName)
		return
	} else {
		c.Res.Result = res
	}
}

// swagger:route GET /instance/{instance_id} instance getInstanceInfo
// GetInstanceInfo 查看实例详情
//     Responses:
//       200: getInstanceInfo
//       default: ErrorResponse

func (c *InstanceController) GetInstanceInfo() {
	defer c.CatchException()
	param := requestTypes.QueryInstanceRequest{
		InstanceId: c.Ctx.Input.Param(":instance_id"),
	}
	res, err := InstanceLogic.DescribeInstance(c.logPoints, param) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}

// swagger:route PUT /instance instance createInstance
// CreateInstance 创建实例
//     Responses:
//       200: createInstance
//       default: ErrorResponse

func (c *InstanceController) CreateInstance() {

	defer c.CatchException()

	req := &sdkModels.CreateInstanceRequest{}
	c.logPoints.Infof("CreateInstance body is:%s", string(c.Ctx.Input.RequestBody))
	//fmt.Println(req, req.Hostname)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	// req.Validate()
	res, err := InstanceLogic.CreateInstanceApi(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	} else {
		c.Res.Result = response.CreateInstanceResponse{
			InstanceIDs: res,
		}
		return
	}

}

// swagger:route POST /instance/startInstance instance startInstance
// StartInstance 开机
//     Responses:
//       200: startInstance
//       default: ErrorResponse

func (c *InstanceController) StartInstance() {
	defer c.CatchException()

	req := &requestTypes.StartInstanceRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	// req.Validate()
	res, err := InstanceLogic.StartInstance(c.logPoints, *req) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route POST /instance/stopInstance instance stopInstance
// StopInstance 关机
//     Responses:
//       200: stopInstance
//       default: ErrorResponse

func (c *InstanceController) StopInstance() {
	defer c.CatchException()
	req := &requestTypes.StopInstanceRequest{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	// req.Validate()
	res, err := InstanceLogic.StopInstance(c.logPoints, *req) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route POST /instance/restartInstance instance restartInstance
// RestartInstance 重启实例
//     Responses:
//       200: restartInstance
//       default: ErrorResponse

func (c *InstanceController) RestartInstance() {
	defer c.CatchException()
	req := &requestTypes.RestartInstanceRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
	}
	// req.Validate()
	res, err := InstanceLogic.RestartInstance(c.logPoints, *req) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route POST /instance/modifyInstance instance modifyInstance
// ModifyInstance 修改实例信息
//     Responses:
//       200: modifyInstance
//       default: ErrorResponse

func (c *InstanceController) ModifyInstance() {
	defer c.CatchException()
	req := &requestTypes.ModifyInstanceRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
	}
	// req.Validate()
	res, err := InstanceLogic.ModifyInstance(c.logPoints, *req) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route DELETE /instance/{instance_id} instance deleteInstance
// DeleteInstance 删除实例
//     Responses:
//       200: deleteInstance
//       default: ErrorResponse

func (c *InstanceController) DeleteInstance() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")

	// req.Validate()
	res, err := InstanceLogic.DeleteInstance(c.logPoints, instanceId) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route POST /instance/lockInstance instance lockInstance
// LockInstance 锁定实例
//     Responses:
//       200: lockInstance
//       default: ErrorResponse

func (c *InstanceController) LockInstance() {
	defer c.CatchException()
	req := &requestTypes.LockInstanceRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
	}
	// req.Validate()
	res, err := InstanceLogic.LockInstance(c.logPoints, *req) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route POST /instance/unlockInstance instance unlockInstance
// UnlockInstance 解锁实例
//     Responses:
//       200: unlockInstance
//       default: ErrorResponse

func (c *InstanceController) UnlockInstance() {
	defer c.CatchException()
	req := &requestTypes.UnLockInstanceRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
	}
	// req.Validate()
	res, err := InstanceLogic.UnLockInstance(c.logPoints, *req) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route POST /instance/startInstances instance startInstances
// StartInstances 批量开机
//     Responses:
//       200: startInstances
//       default: ErrorResponse
// StartInstances 批量开机
func (c *InstanceController) StartInstances() {
	defer c.CatchException()
	req := &requestTypes.StartInstancesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warnf("StartInstances unmarshal error:%s", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	// req.Validate()
	res, err := InstanceLogic.StartInstances(c.logPoints, req.InstanceIDs) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route POST /instance/stopInstances instance stopInstances
// StopInstances 批量关机
//     Responses:
//       200: stopInstances
//       default: ErrorResponse
func (c *InstanceController) StopInstances() {
	defer c.CatchException()
	req := &requestTypes.StartInstancesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warnf("StopInstances unmarshal error:%s", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	// req.Validate()
	res, err := InstanceLogic.StopInstances(c.logPoints, req.InstanceIDs) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route POST /instance/restartInstances instance restartInstances
// RestartInstances 批量重启
//     Responses:
//       200: restartInstances
//       default: ErrorResponse
func (c *InstanceController) RestartInstances() {
	defer c.CatchException()
	req := &requestTypes.StartInstancesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warnf("RestartInstances unmarshal error:%s", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	// req.Validate()
	res, err := InstanceLogic.RestartInstances(c.logPoints, req.InstanceIDs) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route POST /instance/modifyInstances instance modifyInstances
// ModifyInstances 批量编辑实例名称
//     Responses:
//       200: modifyInstances
//       default: ErrorResponse
func (c *InstanceController) ModifyInstances() {
	defer c.CatchException()
	req := &requestTypes.ModifyInstancesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warnf("ModifyInstancesRequest unmarshal error:%s", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	// req.Validate()
	res, err := InstanceLogic.ModifyInstances(c.logPoints, req.InstanceIDs, req.InstanceName) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route DELETE /instance/deleteInstances instance deleteInstances
// DeleteInstances 批量删除实例
//     Responses:
//       200: deleteInstances
//       default: ErrorResponse
func (c *InstanceController) DeleteInstances() {
	defer c.CatchException()
	req := &requestTypes.StartInstancesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warnf("DeleteInstances unmarshal error:%s", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	// req.Validate()
	res, err := InstanceLogic.DeleteInstances(c.logPoints, req.InstanceIDs) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route POST /instance/resetPasswd instance resetPasswd
// resetPasswd 重置密码
//     Responses:
//       200: resetPasswd
//       default: ErrorResponse

func (c *InstanceController) ResetPasswd() {
	defer c.CatchException()

	req := &requestTypes.ResetPasswdRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	// req.Validate()
	res, err := InstanceLogic.ResetInstancePassword(c.logPoints, req.InstanceId, req.Password) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route POST /instance/batchResetPasswd instance resetInstancesPasswd
// resetInstancesPasswd 批量重置密码
//     Responses:
//       200: resetInstancesPasswd
//       default: ErrorResponse

func (c *InstanceController) ResetInstancesPasswd() {
	defer c.CatchException()

	req := &requestTypes.ResetInstancesPasswdRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	// req.Validate()
	res, err := InstanceLogic.ResetInstancesPassword(c.logPoints, req.InstanceIds, req.Password) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route POST /instance/reinstallInstance instance reinstallInstance
// ReinstallInstance 重装实例
//     Responses:
//       200: reinstallInstance
//       default: ErrorResponse

func (c *InstanceController) ReinstallInstance() {
	defer c.CatchException()

	req := &requestTypes.ReinstallInstanceRequest{}

	//fmt.Println(req, req.Hostname)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	// req.Validate()
	res, err := InstanceLogic.ReinstallInstance(c.logPoints, *req)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	} else {
		c.Res.Result = response.CommonResponse{
			Success: res,
		}
		return
	}
}

// swagger:route GET /instance/forshare/list instance getInstanceListForShare
// GetInstanceListForShare 实例列表
//     Responses:
//       200: getInstanceListForShare
//       default: ErrorResponse

func (c *InstanceController) GetInstanceListForShare() {
	defer c.CatchException()

	param := requestTypes.QueryInstancesForShareRequest{
		ProjectID:  c.GetString("projectId"),
		OwnerName:  c.GetString("ownerName"),
		SharerName: c.GetString("sharerName"),
		//支持如下的模糊搜索
		InstanceName: c.GetString("instanceName"),
		InstanceId:   c.GetString("instanceId"),
	}
	res, err := InstanceLogic.DescribeInstancesForShare(c.logPoints, param) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}
