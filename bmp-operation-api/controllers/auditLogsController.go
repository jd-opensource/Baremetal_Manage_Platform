package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/AuditLogLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
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

	p := requestTypes.PagingRequest{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	req := &requestTypes.DescribeAuditLogsRequest{
		PagingRequest: p,
		Sn:            c.GetString("sn"),
		Operation:     c.GetString("operation"),
		Result:        c.GetString("result"),
		UserName:      c.GetString("userName"),
		IsAll:         c.GetString("isAll"),
	}
	st, _ := c.GetInt64("startTime")
	et, _ := c.GetInt64("endTime")
	req.StartTime = st
	req.EndTime = et

	//device_types, err := deviceTypeLogic.QueryByRequest(c.logPoints, req)
	res, err := AuditLogLogic.QueryAuditLogs(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	if c.GetString("exportType") != "" {
		fileName, downloadFileName := AuditLogLogic.ExportAuditLogs(c.logPoints, res.AuditLogs)
		//c.Ctx.Input.SetData("download", "yes")
		c.Ctx.Output.Download(fileName, downloadFileName)
		return
	} else {
		c.Res.Result = res
	}

}

func (c *AuditLogsController) DescribeAuditLogTypes() {

	//device_types, err := deviceTypeLogic.QueryByRequest(c.logPoints, req)
	res, err := AuditLogLogic.DescribeAuditLogsTypes(c.logPoints)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}
