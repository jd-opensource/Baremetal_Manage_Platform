package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
)

// swagger:parameters describeAuditLogs
type DescribeAuditLogsRequest struct {
	ReadHeader

	// in: query
	requestTypes.DescribeAuditLogsRequest
}
