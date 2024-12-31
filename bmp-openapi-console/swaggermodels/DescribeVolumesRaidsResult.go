package swaggermodels

import (
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response describeVolumesRaids
type DescribeDeviceTypeRaidsResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result []*responseTypes.VolumeRaids `json:"result"`
	}
}
