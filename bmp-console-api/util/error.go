package util

import (
	"fmt"
	"reflect"

	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
	apiRuntime "github.com/go-openapi/runtime"
)

/*
sdk的原因，写了此GetOpenapiError用来解析openapi的返回err。err如果已知，为具体类型(比如GetUserListDefault)或者APIError类型(多了一层)，如果未知解析失败，第二个参数为false

type GetUserListDefault struct {
	_statusCode int
	 TraceID string
	 Payload *GetUserListDefaultBody
 }

 type GetUserListDefaultBody struct {

	// error
	Error *models.ErrorResponse `json:"error,omitempty"`
}

APIError{
	OperationName: opName,
	Response:      interface{}(GetUserListDefault),
	Code:          code,
}

*/

func GetOpenapiError(err error) (*sdkModels.ErrorResponse, bool) {

	var v interface{}
	e, ok := err.(*apiRuntime.APIError)
	if ok {
		v = e.Response
	} else {
		v = err
	}
	if reflect.ValueOf(v).MethodByName("GetPayload").IsValid() {
		val := reflect.ValueOf(v).MethodByName("GetPayload").Call([]reflect.Value{})
		if reflect.Indirect(val[0]).FieldByName("Error").IsValid() {
			err1, ok := reflect.Indirect(val[0]).FieldByName("Error").Interface().(*sdkModels.ErrorResponse)
			if ok {
				fmt.Println(err1.Code, err1.Message, err1.Status)
				return err1, true
			}
		}
	}
	return nil, false

}
