package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeAlerts
type DescribeAlertsRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.DescribeAlertsRequest
}
