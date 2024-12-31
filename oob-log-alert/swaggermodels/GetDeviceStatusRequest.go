package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/oob-log-alert/dao"
)

// swagger:parameters getDeviceStatus
type GetDeviceStatusRequest struct {
	ReadRequestHeader
	// in: query
	requestTypes.GetDeviceStatusRequest
}
