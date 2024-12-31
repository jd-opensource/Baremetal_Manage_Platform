package swaggermodels

import (
	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
)

// swagger:response getDeviceStatus
type GetDeviceStatusResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result dao.GetDeviceStatusResponse `json:"result"`
	}
}
