package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/auditLogLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// AuditLogsController operations for auditlog
type AuditLogsController struct {
	BaseController
}

// swagger:route GET /auditLogs auditLog describeAuditLogs
// DescribeAuditLogs 获取审计日志列表
//     Responses:
//       200: describeAuditLogs
//       default: ErrorResponse

func (c *AuditLogsController) DescribeAuditLogs() {
	defer c.CatchException()
	req := &requestTypes.DescribeAuditLogsRequest{
		InstanceID: c.GetString("instance_id"),
		Operation:  c.GetString("operation"),
		UserName:   c.GetString("username"),
		IsAll:      c.GetString("isAll"),
		Result:     c.GetString("result"),
	}
	st, _ := c.GetInt("startTime")
	et, _ := c.GetInt("endTime")
	req.StartTime = st
	req.EndTime = et

	req.Validate(c.logPoints)

	p := util.Pageable{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	//device_types, err := deviceTypeLogic.QueryByRequest(c.logPoints, req)
	res, count, err := auditLogLogic.QueryAuditLogs(c.logPoints, req, p)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.AuditLogList{
		AuditLogs:  res,
		PageNumber: c.pageable.PageNumber,
		PageSize:   c.pageable.PageSize,
		TotalCount: count,
	}
}

// swagger:route GET /auditLogs/types auditLog describeAuditLogTypes
// DescribeAuditLogTypes 获取审计日志类型
//     Responses:
//       200: describeAuditLogTypes
//       default: ErrorResponse

func (c *AuditLogsController) DescribeAuditLogTypes() {

	//device_types, err := deviceTypeLogic.QueryByRequest(c.logPoints, req)
	res := auditLogLogic.DescribeAuditLogsTypes(c.logPoints)

	c.Res.Result = res
}
