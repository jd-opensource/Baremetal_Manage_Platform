package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/auditLogLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/AuditLogsType"
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
		Sn:        c.GetString("sn"),
		Operation: c.GetString("operation"),
		UserName:  c.GetString("username"),
		Result:    c.GetString("result"),
		IsAll:     c.GetString("isAll"),
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

func (c *AuditLogsController) MockMultiAuditLogs() {

	c.logPoints.Point("userId", "user-ta5c5tsos2wkm8d2qtmvx3vufr2h")
	//device_types, err := deviceTypeLogic.QueryByRequest(c.logPoints, req)
	for i := 0; i < 100000; i++ {
		auditLogLogic.SaveAuditLogs(c.logPoints, "d-minping-mock-deviceid", "minping-mock-instanceId", AuditLogsType.AuditLogsStopInstances)
	}

	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}
