package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response describeAuditLogTypes
type DescribeAuditLogTypesResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result []response.AuditLogsType `json:"result"`
	}
}
