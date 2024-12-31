package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeAlertsResult is an response struct that is used to describe get Alert.
// swagger:response describeAlerts
type DescribeAlertsResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.AlertList `json:"result"`
	}
}
