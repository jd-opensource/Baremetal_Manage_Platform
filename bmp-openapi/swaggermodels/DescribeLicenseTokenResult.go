package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeLocalUserResult is an response struct that is used to describe getuser.
// swagger:response describeLicenseToken
type DescribeLicenseTokenResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.LicenseToken `json:"result"`
	}
}
