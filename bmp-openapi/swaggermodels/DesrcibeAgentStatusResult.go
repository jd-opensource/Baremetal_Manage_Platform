package swaggermodels

import (
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response desrcibeAgentStatus
type DesrcibeAgentStatusResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result responseTypes.AgentStatusResponse `json:"result"`
	}
}
