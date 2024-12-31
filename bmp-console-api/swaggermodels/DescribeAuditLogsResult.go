package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-console-api/types/response"
)

// swagger:response describeAuditLogs
type DescribeAuditLogsResult struct {

	// in: body
	Body struct {
		Result response.AuditLogsList `json:"result"`
	}
}
