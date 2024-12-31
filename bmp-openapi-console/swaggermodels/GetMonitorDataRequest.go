package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters getMonitorData
type GetMonitorDataRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.GetMonitorDataRequest
}
