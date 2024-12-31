package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeAlertResult is an response struct that is used to describe get Alert.
// swagger:response describeAlert
type DescribeAlertResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.Alert `json:"result"`
	}
}
