package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeDeviceTypeImages
type DescribeDeviceTypeImagesRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.QueryDeviceTypeImageRequest
}