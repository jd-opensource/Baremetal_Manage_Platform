package controllers

import (
	"encoding/json"
	"fmt"
	"git.jd.com/bmp-pronoea/common/util"
	"git.jd.com/bmp-pronoea/types/request"
	"github.com/astaxie/beego/logs"
)

type PrometheusRulesController struct {
	BaseController
}

// PrometheusRulesWrite @Title PrometheusRulesWrite
// @Description write yaml rule
// @Param	body		body 	request.RulesRequest	true		"The rule info"
// @Success 200 {string}    success
// @Failure 400 invalid request {message}
// @Failure 500 internal error
// @router /write [post]
func (c *PrometheusRulesController) PrometheusRulesWrite() {
	requestId := c.GetRequestId()

	c.resp.Code = HTTP_STATUS_SUCCESS
	c.resp.RequestId = requestId
	c.resp.Message = "success"

	queryByte := c.Ctx.Input.RequestBody
	if queryByte == nil || len(queryByte) <= 0 {
		c.resp.Code = HTTP_STATUS_BAD_REQUEST
		c.resp.Message = fmt.Sprintf("rule info empty! ")
		return
	}

	logs.Info(requestId, fmt.Sprintf("PrometheusRulesWrite param is %v", string(queryByte)))

	req := &request.RulesRequest{}
	err := json.Unmarshal(queryByte, req)
	if err != nil {
		c.resp.Code = HTTP_STATUS_BAD_REQUEST
		c.resp.Message = fmt.Sprintf("json.Unmarshal param error:%v, %v", err, string(queryByte))
		return
	}

	err = req.Validate(req)
	if err != nil {
		c.resp.Code = HTTP_STATUS_BAD_REQUEST
		c.resp.Message = err.Error()
		return
	}

	for _, tmp := range req.Rules {
		if err = tmp.Validate(tmp); err != nil {
			c.resp.Code = HTTP_STATUS_BAD_REQUEST
			c.resp.Message = err.Error()
			return
		}

		for _, tri := range tmp.TriggerOption {
			if err = tri.Validate(tri); err != nil {
				c.resp.Code = HTTP_STATUS_BAD_REQUEST
				c.resp.Message = err.Error()
				return
			}
		}

		err = tmp.NoticeOption.Validate(tmp.NoticeOption)
		if err != nil {
			c.resp.Code = HTTP_STATUS_BAD_REQUEST
			c.resp.Message = err.Error()
			return
		}
	}

	ruleModel, err := c.GetRuleModel(req.Source)
	if err != nil {
		c.resp.Code = HTTP_STATUS_BAD_REQUEST
		c.resp.Message = fmt.Sprintf("invalid source :%v", req.Source)
		return
	}

	err = ruleModel.WriteRule(requestId, string(queryByte))
	if err != nil {
		c.resp.Code = HTTP_STATUS_INTERNAL
		c.resp.Message = err.Error()
		return
	}

	return
}

// PrometheusRulesDelete @Title PrometheusRulesDelete
// @Description delete yaml rule
// @Param	body		body 	request.RulesDeleteRequest	true		"The rule id info"
// @Success 200 {string}    success
// @Failure 400 invalid request {message}
// @Failure 500 internal error
// @router /delete [post]
func (c *PrometheusRulesController) PrometheusRulesDelete() {
	requestId := c.GetRequestId()

	c.resp.Code = HTTP_STATUS_SUCCESS
	c.resp.RequestId = requestId
	c.resp.Message = "success"

	queryByte := c.Ctx.Input.RequestBody
	if queryByte == nil || len(queryByte) <= 0 {
		c.resp.Code = HTTP_STATUS_BAD_REQUEST
		c.resp.Message = fmt.Sprintf("param empty !")
		return
	}

	logs.Info(requestId, fmt.Sprintf("PrometheusRulesDelete param is %v", string(queryByte)))

	req := &request.RulesDeleteRequest{}
	err := json.Unmarshal(queryByte, req)
	if err != nil {
		c.resp.Code = HTTP_STATUS_BAD_REQUEST
		c.resp.Message = fmt.Sprintf("json.Unmarshal param error:%v", err)
		return
	}

	ruleModel, err := c.GetRuleModel(req.Source)
	if err != nil {
		c.resp.Code = HTTP_STATUS_BAD_REQUEST
		c.resp.Message = fmt.Sprintf("invalid source :%v", req.Source)
		return
	}
	err = ruleModel.DeleteRules(requestId, req.RuleIds, req.Source)
	if err != nil {
		c.resp.Code = HTTP_STATUS_INTERNAL
		c.resp.Message = fmt.Sprintf("delete [%v-%v]rule fail: %v", req.RuleIds, req.Source, err)
		return
	}

	return
}

// PrometheusRulesList @Title PrometheusRulesList
// @Description delete yaml rule
// @Param	source query		string true		"The rule source"
// @Success 200 {string}    success
// @Failure 400 invalid request {message}
// @Failure 500 internal error
// @router /list [get]
func (c *PrometheusRulesController) PrometheusRulesList() {
	requestId := c.GetRequestId()

	c.resp.Code = HTTP_STATUS_SUCCESS
	c.resp.RequestId = requestId
	c.resp.Message = "success"

	source := c.GetString("source", "")
	if util.IsEmpty(source) {
		c.resp.Code = HTTP_STATUS_BAD_REQUEST
		c.resp.Message = fmt.Sprintf("param source empty !")
		return
	}

	ruleModel, err := c.GetRuleModel(source)
	if err != nil {
		c.resp.Code = HTTP_STATUS_BAD_REQUEST
		c.resp.Message = fmt.Sprintf("invalid source :%v", source)
		return
	}
	ruleList, err := ruleModel.RulesList(requestId, source)
	if err != nil {
		c.resp.Code = HTTP_STATUS_INTERNAL
		c.resp.Message = fmt.Sprintf("get rule list error : %v, %v ", err, source)
		return
	}

	c.resp.Result = ruleList
	return
}