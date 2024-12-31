package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters desrcibeAgentStatus
type DesrcibeAgentStatusRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.DesrcibeAgentStatusRequest
}
