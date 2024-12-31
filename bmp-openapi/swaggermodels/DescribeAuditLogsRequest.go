package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeAuditLogs
type DescribeAuditLogsRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.DescribeAuditLogsRequest
}
