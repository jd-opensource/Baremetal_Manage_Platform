package swaggermodels

import (
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeVolumesByDeviceTypeResult is an response struct that is used to describe get sshkey list.
// swagger:response describeVolumesByDeviceType
type DescribeVolumesByDeviceTypeResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result []*responseTypes.VolumeIt `json:"result"`
	}
}
