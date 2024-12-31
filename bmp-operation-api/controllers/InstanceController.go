package controllers

import (
	// "encoding/json"

	"encoding/json"

	InstanceLogic "coding.jd.com/aidc-bmp/bmp-operation-api/logic/InstanceLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	response "coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

type InstanceController struct {
	BaseController
}

// StartInstance【设备管理】【开机】
// swagger:route POST /instances/start/{instance_id} instance startInstance
// StartInstance【设备管理】【开机】
//     Responses:
//       200: startInstance
//       default: ErrorResponse

func (c *InstanceController) StartInstance() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")
	req := &requestTypes.StartInstanceRequest{
		InstanceID: instanceId,
	}
	result, err := InstanceLogic.StartInstance(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// RestartInstance【设备管理】【重启】
// swagger:route POST /instances/restart/{instance_id} instance restartInstance
// RestartInstance【设备管理】【重启】
//     Responses:
//       200: restartInstance
//       default: ErrorResponse

func (c *InstanceController) RestartInstance() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")
	req := &requestTypes.RestartInstanceRequest{
		InstanceID: instanceId,
	}
	result, err := InstanceLogic.RestartInstance(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// StopInstance【设备管理】【关机】
// swagger:route POST /instances/stop/{instance_id} instance stopInstance
// StopInstance【设备管理】【关机】
//     Responses:
//       200: stopInstance
//       default: ErrorResponse

func (c *InstanceController) StopInstance() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")
	req := &requestTypes.StopInstanceRequest{
		InstanceID: instanceId,
	}
	result, err := InstanceLogic.StopInstance(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// swagger:route POST /instances/resetStatus/{instance_id} instance resetInstanceStatus
// ResetInstanceStatus【设备管理】【重置实例状态】
//     Responses:
//       200: resetInstanceStatus
//       default: ErrorResponse
func (c *InstanceController) ResetInstanceStatus() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")

	result, err := InstanceLogic.ResetInstanceStatus(c.logPoints, instanceId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// swagger:route DELETE /instances/{instance_id} instance deleteInstance
// DeleteInstance【设备管理】【回收实例】
//     Responses:
//       200: deleteInstance
//       default: ErrorResponse
func (c *InstanceController) DeleteInstance() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")

	result, err := InstanceLogic.DeleteInstance(c.logPoints, instanceId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// swagger:route POST /instances/lock/{instance_id} instance lockInstance
// LockInstance 锁定实例
//     Responses:
//       200: lockInstance
//       default: ErrorResponse

func (c *InstanceController) LockInstance() {
	defer c.CatchException()

	instanceId := c.Ctx.Input.Param(":instance_id")

	res, err := InstanceLogic.LockInstance(c.logPoints, instanceId) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route POST /instances/unlock/{instance_id} instance unlockInstance
// UnlockInstance 解锁实例
//     Responses:
//       200: unlockInstance
//       default: ErrorResponse

func (c *InstanceController) UnlockInstance() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")
	// req.Validate()
	res, err := InstanceLogic.UnLockInstance(c.logPoints, instanceId) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

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

// StartInstances 批量关机
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

// StartInstances 批量重启
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

// ModifyInstances 批量编辑实例名称
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

// StartInstances 批量删除实例
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

// swagger:route POST /instances/batchResetPasswd instance resetInstancesPasswd
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

// swagger:route POST /instances/reinstallInstance instance reinstallInstance
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

// swagger:route GET /instances/{instance_id} instance getInstanceInfo
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
