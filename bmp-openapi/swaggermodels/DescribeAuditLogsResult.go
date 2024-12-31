package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response describeAuditLogs
type DescribeAuditLogsResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.AuditLogList `json:"result"`
	}
}
