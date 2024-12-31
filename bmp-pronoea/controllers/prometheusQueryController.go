package controllers

import (
	"encoding/json"
	"fmt"
	"git.jd.com/bmp-pronoea/models/query"
	"git.jd.com/bmp-pronoea/types/request"
	"github.com/astaxie/beego/logs"
)

type PrometheusQueryController struct {
	BaseController
}

// PrometheusSearchSamples @Title PrometheusSearchSamples
// @Description write yaml rule
// @Param	body		body 	request.MetricRangeQueryRequest	true		"The query info"
// @Success 200 {string}    success
// @Failure 400 invalid request {message}
// @Failure 500 internal error
// @router / [post]
func (c *PrometheusQueryController) PrometheusSearchSamples() {
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

	logs.Info(requestId, fmt.Sprintf("PrometheusSearchSamples param is %v", string(queryByte)))

	req := &request.MetricRangeQueryRequest{}
	err := json.Unmarshal(queryByte, req)
	if err != nil {
		c.resp.Code = HTTP_STATUS_BAD_REQUEST
		c.resp.Message = fmt.Sprintf("json.Unmarshal param error:%v", err)
		return
	}

	samples, err := query.SearchPrometheusSamples(requestId, req)
	if err != nil {
		c.resp.Code = HTTP_STATUS_INTERNAL
		c.resp.Message = err.Error()
		return
	}

	c.resp.Result = samples
	return
}
