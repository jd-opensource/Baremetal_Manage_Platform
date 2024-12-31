package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeAlert
type DescribeAlertRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.DescribeAlertRequest
}
