package controllers

import (
	"encoding/json"
	"time"

	"coding.jd.com/bmp/agent-proxy-jdstack/api/models"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"

	"coding.jd.com/bmp/agent-proxy-jdstack/global"

	beego "github.com/beego/beego/v2/server/web"
)

type StatusController struct {
	BaseController
}

type DescribeAgentStatusRequest struct {
	InstanceIds []string `json:"instanceIds"`
}

func (c *StatusController) DescribeAgentStatus() {

	req := DescribeAgentStatusRequest{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.logPoints.Warn("DescribeAgentStatus parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	if len(req.InstanceIds) == 0 {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, "instanceId empty, invalid", errorCode.INVALID_ARGUMENT)

		return
	}

	var tmp *models.AgentStatus
	var o []*models.AgentStatus
	var ok bool
	var ids = req.InstanceIds
	for _, id := range ids {
		if v, found := global.Get(id); !found {
			tmp = &models.AgentStatus{InstanceID: id, Status: models.AGENT_UNKNOWN}
		} else {
			tmp, ok = v.(*models.AgentStatus)
			if tmp == nil || !ok {
				tmp = &models.AgentStatus{InstanceID: id, Status: models.AGENT_UNKNOWN}
			} else {
				tmp.InstanceID = id
				if time.Now().Unix()-tmp.Timestamp > beego.AppConfig.DefaultInt64("app.agent.heartbeat.timeout", 300) && tmp.Status == models.AGENT_OK {
					tmp.Status = models.AGENT_UNKNOWN
				}
			}
		}

		o = append(o, tmp)
	}

	c.Res.Result = o
}
