package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-operation-api/controllers"

// unit err response for non 200
// swagger:response ErrorResponse
type ErrorResponse struct {
	ResponseHeader
	//In: body
	Body struct {
		Err       controllers.ErrorResponse `json:"error"`
		RequestId string                    `json:"requestId"`
	}
}

type CommonRespBody struct {
	RequestId string `json:"requestId"`
}
