package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters uploadLicense
type UploadLicenseRequest struct {
	ReadRequestHeader

	// in: body
	Body requestTypes.CreateLicenseRequest
}
