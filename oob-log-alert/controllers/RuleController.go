package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
	"coding.jd.com/aidc-bmp/oob-log-alert/service"

	"github.com/beego/beego/v2/core/logs"
)

type RuleController struct {
	BaseController
}

// swagger:route GET /v1/oob-alert/rules/alert-spec-list alert-rule alertPartList
//
// AlertPartList 故障报警规则-故障配件列表
//
//     Responses:
//       200: alertPartList
//       default: ErrorResponse
func (c *RuleController) AlertPartList() {
	data, err := service.GetAlertPartList(c.logPoints)
	if err != nil {
		c.SetErrorResponse(InternalErrorResp, err.Error())
		return
	}
	c.Res.Result = data
}

// swagger:route GET /v1/oob-alert/rules/alert-level-list alert-rule alertLevelList
//
// AlertLevelList 故障报警规则-故障等级列表
//
//     Responses:
//       200: alertLevelList
//       default: ErrorResponse
func (c *RuleController) AlertLevelList() {
	data, err := service.GetAlertLevelList(c.logPoints)
	if err != nil {
		c.SetErrorResponse(InternalErrorResp, err.Error())
		return
	}
	c.Res.Result = data
}

// swagger:route GET /v1/oob-alert/rules alert-rule ruleList
//
// RuleList 故障报警规则列表
//
//     Responses:
//       200: ruleList
//       default: ErrorResponse
func (c *RuleController) RuleList() {
	logs.Debug(">>>>>>>>RuleController.RuleList()")
	defer logs.Debug("<<<<<<<<RuleController.RuleList()")

	req := dao.RuleListRequest{
		AlertName:   c.GetString("alert_name"),
		AlertSpec:   c.GetString("alert_spec"),
		AlertLevel:  c.GetString("alert_level"),
		PageRequest: c.GetPageReqInfo(),
	}

	logs.Debug("GetLogs.req = %+v", req)

	data, err := service.GetRuleList(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(InternalErrorResp, err.Error())
		return
	}

	if c.GetString("exportType") != "" {
		fileName, downloadFileName, err := service.ExportRulesExcel(c.logPoints, data.Detail)
		if err != nil {
			c.SetErrorResponse(InternalErrorResp, "导出故障规则列表excel错误 "+err.Error())
			return
		}
		//c.Ctx.Input.SetData("download", "yes")
		c.Ctx.Output.Download(fileName, downloadFileName)
		// if err := os.Remove(fileName); err != nil {
		// 	c.SetErrorResponse(InternalErrorResp, "删除上传文件失败"+err.Error())
		// 	return
		// }
		return
	}

	c.Res.Result = data
}

// swagger:route POST /v1/oob-alert/rules/change-push alert-rule changePush
//
// ChangePush 故障报警规则列表-故障推送开关
//
//     Responses:
//       200: changePush
//       default: ErrorResponse
func (c *RuleController) ChangePush() {
	logs.Debug(">>>>>>>>RuleController.ChangePush()")
	defer logs.Debug("<<<<<<<<RuleController.ChangePush()")

	req := dao.ChangePushRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.SetErrorResponse(ParamInvalid, err.Error())
		return
	}
	logs.Debug("ChangePush.req = %+v", req)
	data, err := service.ChangeRulePush(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(InternalErrorResp, err.Error())
		return
	}
	c.Res.Result = dao.CommonResponse{
		Success: data,
	}
}

// swagger:route POST /v1/oob-alert/rules/change-use alert-rule changeUse
//
// ChangeUse 故障报警规则列表-故障启用开关
//
//     Responses:
//       200: changeUse
//       default: ErrorResponse
func (c *RuleController) ChangeUse() {
	logs.Debug(">>>>>>>>RuleController.ChangeUse()")
	defer logs.Debug("<<<<<<<<RuleController.ChangeUse()")

	req := dao.ChangeUseRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.SetErrorResponse(ParamInvalid, err.Error())
		return
	}
	logs.Debug("ChangeUse.req = %+v", req)

	data, err := service.ChangeRuleUse(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(InternalErrorResp, err.Error())
		return
	}
	c.Res.Result = dao.CommonResponse{
		Success: data,
	}
}

// swagger:route POST /v1/oob-alert/rules/reset alert-rule ruleResetPushAndUse
//
// RuleResetPushAndUse 故障报警规则列表-恢复默认规则重置
//
//     Responses:
//       200: ruleResetPushAndUse
//       default: ErrorResponse
func (c *RuleController) RuleResetPushAndUse() {
	logs.Debug(">>>>>>>>RuleController.RuleResetPushAndUse()")
	defer logs.Debug("<<<<<<<<RuleController.RuleResetPushAndUse()")

	req := dao.RuleResetPushAndUseRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.SetErrorResponse(ParamInvalid, err.Error())
		return
	}
	logs.Debug("RuleResetPushAndUse.req = %+v", req)

	ids := []int{}
	tmp := strings.Split(req.RuleIDs, ",")
	for _, v := range tmp {
		id, err := strconv.Atoi(v)
		if err != nil {
			c.SetErrorResponse(ParamInvalid, err.Error())
			return
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		c.SetErrorResponse(ParamInvalid, "rule_ids invalid")
		return
	}
	data, err := service.RuleResetPushAndUse(c.logPoints, ids)
	if err != nil {
		c.SetErrorResponse(InternalErrorResp, err.Error())
		return
	}
	c.Res.Result = dao.CommonResponse{
		Success: data,
	}
}
