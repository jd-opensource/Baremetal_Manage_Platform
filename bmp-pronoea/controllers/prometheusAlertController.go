package controllers

import (
	"encoding/json"
	"fmt"
	"git.jd.com/bmp-pronoea/types/request"
	"github.com/astaxie/beego/logs"
)

type PrometheusAlertController struct {
	BaseController
}

// PrometheusAlertsReceive @Title PrometheusAlertsReceive
// @Description receive alert info
// @Param	body		body 	request.AlertReceiverRequest	true		"The alert info"
// @Success 200 {string}    success
// @Failure 400 invalid request {message}
// @Failure 500 internal error
// @router /receiver [post]
func (c *PrometheusAlertController) PrometheusAlertsReceive() {
	requestId := c.GetRequestId()

	c.resp.Code = HTTP_STATUS_SUCCESS
	c.resp.RequestId = requestId
	c.resp.Message = "success"

	queryByte := c.Ctx.Input.RequestBody
	if queryByte == nil || len(queryByte) <= 0 {
		c.resp.Code = HTTP_STATUS_BAD_REQUEST
		c.resp.Message = fmt.Sprintf("alert info empty! ")
		return
	}

	logs.Info(requestId, "pronoea receive alertmanager info is ", string(queryByte))

	alertMsg := &request.AlertReceiverRequest{}
	err := json.Unmarshal(queryByte, alertMsg)
	if err != nil {
		c.resp.Code = HTTP_STATUS_BAD_REQUEST
		c.resp.Message = fmt.Sprintf("request.AlertReceiverRequest json.Unmarshal fail:%v", err)
		return
	}

	alertModel, err := c.GetAlertModel(alertMsg.Receiver)
	if err != nil {
		c.resp.Code = HTTP_STATUS_BAD_REQUEST
		c.resp.Message = fmt.Sprintf("invalid receiver :%v", alertMsg.Receiver)
		return
	}

	err = alertModel.ReportAlerts(requestId, alertMsg)
	if err != nil {
		c.resp.Code = HTTP_STATUS_INTERNAL
		c.resp.Message = err.Error()
		return
	}

	return
}


