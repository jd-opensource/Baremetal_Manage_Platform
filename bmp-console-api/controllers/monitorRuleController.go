package controllers

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-console-api/logic/monitorRuleLogic"
	"coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-console-api/util"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// MonitorRuleController 告警规则
type MonitorRuleController struct {
	BaseController
}

// swagger:route POST /monitorRule/addRule monitorRule addRule
//
// AddRule 添加规则
//
//     Responses:
//       200: addRule
//       default: ErrorResponse
func (c *MonitorRuleController) AddRule() {
	defer c.CatchException()
	c.logPoints.Infof("AddRule body:%s", string(c.Ctx.Input.RequestBody))
	req := &request.AddRuleRequest{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("AddRule parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	res, err := monitorRuleLogic.AddRule(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route GET /monitorRule/describeRule monitorRule describeRule
//
// DescribeRule 规则详情
//
//     Responses:
//       200: describeRule
//       default: ErrorResponse
func (c *MonitorRuleController) DescribeRule() {
	defer c.CatchException()
	req := request.DescribeRuleRequest{
		RuleID: c.GetString("ruleId"),
	}
	req.Validate(c.logPoints)
	res, err := monitorRuleLogic.DescribeRule(c.logPoints, req.RuleID)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}

// swagger:route GET /monitorRule/describeRules monitorRule describeRules
//
// DescribeRules 规则列表
//
//     Responses:
//       200: describeRules
//       default: ErrorResponse
func (c *MonitorRuleController) DescribeRules() {
	defer c.CatchException()

	p := util.Pageable{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	s, _ := c.GetInt("status")
	req := &request.DescribeRulesRequest{
		IsAll:     c.GetString("isAll"),
		UserName:  c.GetString("userName"),
		UserId:    c.GetString("userId"),
		RuleName:  c.GetString("ruleName"),
		RuleId:    c.GetString("ruleId"),
		Status:    s,
		ProjectID: c.GetString("projectId"),
	}
	req.Validate(c.logPoints)

	res, err := monitorRuleLogic.DescribeRules(c.logPoints, req, p)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	if c.GetString("exportType") != "" {
		fileName, downloadFileName := monitorRuleLogic.ExportRules(c.logPoints, res.Rules)
		//c.Ctx.Input.SetData("download", "yes")
		c.Ctx.Output.Download(fileName, downloadFileName)
		return
	} else {
		c.Res.Result = res
	}
}

// swagger:route PUT /monitorRule/editRule monitorRule editRule
//
// EditRule 编辑规则
//
//     Responses:
//       200: editRule
//       default: ErrorResponse
func (c *MonitorRuleController) EditRule() {
	defer c.CatchException()
	c.logPoints.Infof("EditRule body:%s", string(c.Ctx.Input.RequestBody))
	req := &request.EditRuleRequest{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("EditRule parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	res, err := monitorRuleLogic.EditRule(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route PUT /monitorRule/enableRule monitorRule enableRule
//
// EnableRule 启用规则
//
//     Responses:
//       200: enableRule
//       default: ErrorResponse
func (c *MonitorRuleController) EnableRule() {
	defer c.CatchException()
	c.logPoints.Infof("EnableRule body:%s", string(c.Ctx.Input.RequestBody))
	req := &request.EnableRuleRequest{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("EnableRule parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	res, err := monitorRuleLogic.EnableRule(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route PUT /monitorRule/disableRule monitorRule disableRule
//
// DisableRule 禁用规则
//
//     Responses:
//       200: disableRule
//       default: ErrorResponse
func (c *MonitorRuleController) DisableRule() {
	defer c.CatchException()
	c.logPoints.Infof("DisableRule body:%s", string(c.Ctx.Input.RequestBody))
	req := &request.DisableRuleRequest{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("DisableRule parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	res, err := monitorRuleLogic.DisableRule(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route DELETE /monitorRule/deleteRule monitorRule deleteRule
//
// DeleteRule 删除规则
//
//     Responses:
//       200: deleteRule
//       default: ErrorResponse
func (c *MonitorRuleController) DeleteRule() {
	defer c.CatchException()
	c.logPoints.Infof("EnableRule body:%s", string(c.Ctx.Input.RequestBody))
	req := &request.DeleteRuleRequest{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("DeleteRule parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	res, err := monitorRuleLogic.DeleteRule(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}
