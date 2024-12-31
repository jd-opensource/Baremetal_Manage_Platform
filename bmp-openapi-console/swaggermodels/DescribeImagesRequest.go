package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeImages
type DescribeImagesRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.QueryImagesRequest
}
