package response

import sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"

type AuditLogsList struct {
	AuditLogs  []*sdkModels.AuditLog `json:"messages"`
	PageNumber int64                 `json:"pageNumber"`
	PageSize   int64                 `json:"pageSize"`
	TotalCount int64                 `json:"totalCount"`
}
