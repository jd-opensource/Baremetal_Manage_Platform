package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeLicenseContentResult is an response struct that is used to describe getuser.
// swagger:response describeLicenseContent
type DescribeLicenseContentResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.LicenseContent `json:"result"`
	}
}
