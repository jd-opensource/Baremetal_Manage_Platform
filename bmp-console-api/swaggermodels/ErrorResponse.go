package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-console-api/controllers"

// unit err response for non 200
// swagger:response ErrorResponse
type ErrorResponse struct {
	//In: body
	Body struct {
		Err       controllers.ErrorResponse `json:"error"`
		RequestId string                    `json:"requestId"`
	}
}