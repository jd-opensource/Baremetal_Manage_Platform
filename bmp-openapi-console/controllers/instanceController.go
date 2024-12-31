package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/instanceLogic"
	log "coding.jd.com/aidc-bmp/bmp_log"

	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	validation "git.jd.com/cps-golang/ironic-common/ironic/common/Validation"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceTypeController operations for Device
type InstanceController struct {
	BaseController
}

func ValidateLocalUserProjectInstances(logger *log.Logger, projectId string, InstanceIds []string) error {
	userId := logger.GetPoint("userId").(string)
	query := map[string]interface{}{
		"user_id":    userId,
		"project_id": projectId,
		"is_del":     0,
	}
	if len(InstanceIds) > 0 {
		query["instance_id.in"] = InstanceIds
	}
	c, err := instanceLogic.GetInstanceCount(logger, query)
	if err != nil {
		return err
	}
	if c == 0 || c != int64(len(InstanceIds)) {
		return errors.New("权限不够")
	}
	return nil
}

// Instance ...
// swagger:route POST /project/instances instance createProjectInstance
//
// CreateProjectInstance 创建实例
//     Responses:
//       200: createProjectInstance
//       default: ErrorResponse

func (c *InstanceController) CreateProjectInstance() {
	defer c.CatchException()
	c.logPoints.Infof("CreateProjectInstance body:%s", string(c.Ctx.Input.RequestBody))
	req := &requestTypes.CreateInstanceRequest{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("CreateInstance parse pInstancet body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	instanceIds, err := instanceLogic.CreateInstance(c.logPoints, *req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.InstanceIds{
		InstanceIds: instanceIds,
	}
}

// GetInstanceList ...
// swagger:route GET /project/instances instance describeProjectInstances
//
// DescribeProjectInstances 获取实例列表
//     Responses:
//       200: describeProjectInstances
//       default: ErrorResponse

func (c *InstanceController) DescribeProjectInstances() {
	defer c.CatchException()
	p := util.Pageable{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	req := &requestTypes.QueryInstancesRequest{
		InstanceName:   c.GetString("instanceName"),
		InstanceID:     c.GetString("instanceId"),
		Sn:             c.GetString("sn"),
		ProjectID:      c.GetString("projectId"),
		IdcID:          c.GetString("idcId"),
		DeviceTypeID:   c.GetString("deviceTypeId"),
		DeviceID:       c.GetString("deviceId"),
		Status:         c.GetString("status"),
		IloIP:          c.GetString("ilo_ip"),
		IPV4:           c.GetString("ipv4"),
		IPV6:           c.GetString("ipv6"),
		IsAll:          c.GetString("isAll"),
		IsInstallAgent: c.GetString("isInstallAgent"),
	}
	req.Validate(c.logPoints)
	res, count, err := instanceLogic.GetInstanceList(c.logPoints, req, p)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.InstanceList{
		Instances:  res,
		PageNumber: c.pageable.PageNumber,
		PageSize:   c.pageable.PageSize,
		TotalCount: count,
	}

}

// DescribeInstance ...
// swagger:route GET /project/instances/{instance_id} instance describeProjectInstance
// DescribeProjectInstance 获取实例详情
//
//     Responses:
//       200: describeProjectInstance
//       default: ErrorResponse

func (c *InstanceController) DescribeProjectInstance() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")
	res, err := instanceLogic.GetInstanceById(c.logPoints, instanceId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.InstanceInfo{
		Instance: res,
	}
}

// ModifyInstance ...
// swagger:route PUT /project/instances/{instance_id} instance modifyProjectInstance
// ModifyProjectInstance 修改实例信息
//
//     Responses:
//       200: modifyProjectInstance
//       default: ErrorResponse

func (c *InstanceController) ModifyProjectInstance() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")
	req := &requestTypes.ModifyInstanceRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyInstance parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := instanceLogic.ModifyProjectInstance(c.logPoints, req, instanceId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// LockInstance ...
// swagger:route PUT /project/instances/{instance_id}/lock instance lockProjectInstance
//
// LockProjectInstance 锁定实例
//     Responses:
//       200: lockProjectInstance
//       default: ErrorResponse

func (c *InstanceController) LockProjectInstance() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")
	if err := instanceLogic.LockInstance(c.logPoints, instanceId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// LockInstance ...
// swagger:route PUT /project/instances/{instance_id}/unlock instance unLockProjectInstance
//
// UnLockProjectInstance 解锁实例
//     Responses:
//       200: unLockProjectInstance
//       default: ErrorResponse

func (c *InstanceController) UnLockProjectInstance() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")
	if err := instanceLogic.UnLockInstance(c.logPoints, instanceId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// StartInstance ...
// swagger:route PUT /project/instances/{instance_id}/start instance startProjectInstance
//
// StartProjectInstance 实例开机
//     Responses:
//       200: startProjectInstance
//       default: ErrorResponse

func (c *InstanceController) StartProjectInstance() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")
	if err := instanceLogic.StartInstance(c.logPoints, instanceId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// StopInstance ...
// swagger:route PUT /project/instances/{instance_id}/stop instance stopProjectInstance
//
// StopProjectInstance 实例关机
//     Responses:
//       200: stopProjectInstance
//       default: ErrorResponse

func (c *InstanceController) StopProjectInstance() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")
	if err := instanceLogic.StopInstance(c.logPoints, instanceId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// RestartInstance ...
// swagger:route PUT /project/instances/{instance_id}/restart instance restartProjectInstance
//
// RestartProjectInstance 实例重启
//     Responses:
//       200: restartProjectInstance
//       default: ErrorResponse

func (c *InstanceController) RestartProjectInstance() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")
	if err := instanceLogic.RestartInstance(c.logPoints, instanceId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// DeleteInstance ...
// swagger:route DELETE /project/instances/{instance_id} instance deleteProjectInstance
//
// DeleteProjectInstance 删除实例
//     Responses:
//       200: deleteProjectInstance
//       default: ErrorResponse

func (c *InstanceController) DeleteProjectInstance() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")
	if err := instanceLogic.DeleteInstance(c.logPoints, instanceId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// ResetInstanceStatus ...
// swagger:route PUT /project/instances/{instance_id}/resetStatus instance resetInstanceStatus
//
// ResetInstanceStatus 重置实例状态
//     Responses:
//       200: resetInstanceStatus
//       default: ErrorResponse
func (c *InstanceController) ResetInstanceStatus() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")
	if err := instanceLogic.ResetInstanceStatus(c.logPoints, instanceId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route PUT /project/instances/batch/instances:startInstances instance startInstances
//
// StartInstances 批量开机
//     Responses:
//       200: startInstances
//       default: ErrorResponse
func (c *InstanceController) StartInstances() {

	defer c.CatchException()
	instancesAndOperation := c.Ctx.Input.Param(":param")
	if instancesAndOperation != "instances:startInstances" {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, fmt.Sprintf("url %s parse error", instancesAndOperation), errorCode.INVALID_ARGUMENT)
		return
	}

	req := requestTypes.StartInstancesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	req.Validate(c.logPoints)

	if err := instanceLogic.StartInstances(c.logPoints, req.InstanceIds); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route PUT /project/instances/batch/instances:stopInstances instance stopInstances
//
// StartInstances 批量关机
//     Responses:
//       200: stopInstances
//       default: ErrorResponse
func (c *InstanceController) StopInstances() {

	defer c.CatchException()
	instancesAndOperation := c.Ctx.Input.Param(":param")
	if instancesAndOperation != "instances:stopInstances" {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, fmt.Sprintf("url %s parse error", instancesAndOperation), errorCode.INVALID_ARGUMENT)
		return
	}
	req := requestTypes.StopInstancesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	req.Validate(c.logPoints)

	if err := instanceLogic.StopInstances(c.logPoints, req.InstanceIds); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route PUT /project/instances/batch/instances:restartInstances instance restartInstances
//
// StartInstances 批量重启
//     Responses:
//       200: restartInstances
//       default: ErrorResponse
func (c *InstanceController) RestartInstances() {

	defer c.CatchException()
	instancesAndOperation := c.Ctx.Input.Param(":param")
	if instancesAndOperation != "instances:restartInstances" {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, fmt.Sprintf("url %s parse error", instancesAndOperation), errorCode.INVALID_ARGUMENT)
		return
	}
	req := requestTypes.RestartInstancesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	req.Validate(c.logPoints)

	if err := instanceLogic.RestartInstances(c.logPoints, req.InstanceIds); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route PUT /project/instances/batch/instances:modifyInstances instance modifyInstances
//
// ModifyInstances 批量修改实例名称
//     Responses:
//       200: modifyInstances
//       default: ErrorResponse
func (c *InstanceController) ModifyInstances() {

	defer c.CatchException()
	instancesAndOperation := c.Ctx.Input.Param(":param")
	if instancesAndOperation != "instances:modifyInstances" {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, fmt.Sprintf("url %s parse error", instancesAndOperation), errorCode.INVALID_ARGUMENT)
		return
	}
	req := requestTypes.ModifyInstancesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	req.Validate(c.logPoints)

	if err := instanceLogic.ModifyInstances(c.logPoints, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// DeleteInstances ...
// swagger:route DELETE /project/instances/batch/instances:deleteInstances instance deleteInstances
//
// DeleteInstances 批量删除实例
//     Responses:
//       200: deleteInstances
//       default: ErrorResponse

func (c *InstanceController) DeleteInstances() {
	defer c.CatchException()

	defer c.CatchException()
	instancesAndOperation := c.Ctx.Input.Param(":param")
	if instancesAndOperation != "instances:deleteInstances" {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, fmt.Sprintf("url %s parse error", instancesAndOperation), errorCode.INVALID_ARGUMENT)
		return
	}
	req := requestTypes.DeleteInstancesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	req.Validate(c.logPoints)

	if err := instanceLogic.DeleteInstances(c.logPoints, req.InstanceIDs); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// ResetPasswd ...
// swagger:route PUT /project/instances/{instance_id}/resetpasswd instance resetProjectInstancePasswd
//
// ResetPasswd 重置密码
//     Responses:
//       200: resetProjectInstancePasswd
//       default: ErrorResponse

func (c *InstanceController) ResetPasswd() {
	defer c.CatchException()
	instanceId := c.Ctx.Input.Param(":instance_id")

	req := requestTypes.ResetInstancePasswdRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	req.Validate(c.logPoints)

	if err := instanceLogic.ResetInstancePasswd(c.logPoints, instanceId, req.Password); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// StartInstance ...
// swagger:route PUT /project/instances/batch/instances:resetPasswd instance resetProjectInstancesPasswd
//
// ResetInstancesPasswd 批量重置密码
//     Responses:
//       200: resetProjectInstancesPasswd
//       default: ErrorResponse

func (c *InstanceController) ResetInstancesPasswd() {
	defer c.CatchException()

	req := requestTypes.ResetInstancesPasswdRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	req.Validate(c.logPoints)

	if err := instanceLogic.ResetInstancesPasswd(c.logPoints, req.InstanceIDs, req.Password); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// ReinstallInstance ...
// swagger:route PUT /project/instances/{instance_id}/reinstallInstance instance reinstallProjectInstance
//
// ReinstallInstance 重装实例
//     Responses:
//       200: reinstallProjectInstance
//       default: ErrorResponse
func (c *InstanceController) ReinstallInstance() {
	defer c.CatchException()
	instance_id := c.Ctx.Input.Param(":instance_id")
	if match, err := regexp.MatchString(validation.REGEX_ID, instance_id); !match {
		if err != nil {
			c.logPoints.Warn("ReinstallInstance path error:", err.Error())
		}
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req := &requestTypes.ReinstallInstanceRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("parse post body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	if err := req.Validate(c.logPoints); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	if err := instanceLogic.ReinstallInstance(c.logPoints, instance_id, req); err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// GetInstanceList ...
// swagger:route GET /project/instances/share/describeInstances instance describeInstancesByProjectIdAndOwnerNameAndSharerName
//
// DescribeInstancesByProjectIdAndOwnerNameAndSharerName 根据projectid，ownername，username来获取实例列表。为了支持部分资源转移中的实例列表
//     Responses:
//       200: describeInstancesByProjectIdAndOwnerNameAndSharerName
//       default: ErrorResponse

func (c *InstanceController) DescribeInstancesByProjectIdAndOwnerNameAndSharerName() {
	defer c.CatchException()

	req := &requestTypes.DescribeInstancesByProjectIdAndOwnerNameAndSharerNameRequest{

		ProjectID: c.GetString("projectId"),
		// IsAll:      c.GetString("isAll"),
		OwnerName:  c.GetString("ownerName"),
		SharerName: c.GetString("sharerName"),

		InstanceName: c.GetString("instanceName"),
		InstanceID:   c.GetString("instanceId"),
	}
	req.Validate(c.logPoints)
	res, err := instanceLogic.GetInstanceListByProjectIdAndOwnerNameAndSharerName(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.InstanceForShareList{
		Instances: res,
	}

}
