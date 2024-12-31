package swaggermodels

import (
	"coding.jd.com/aidc-bmp/oob-log-alert/controllers"
)

type ReadRequestHeader struct {
	// 流量唯一id
	// required: true
	// in: header
	TraceID string `json:"traceId"`
	// // demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token
	// // required: true
	// // in: header
	// Authorization string `json:"authorization"`
	// // 用户uuid, deprecated
	// // in: header
	// BmpUserID string `json:"bmpUserId"`
	// // 用户语言 [zh_CN, en_US]
	// // in: header
	// BmpLanguage string `json:"bmpLanguage"`
}

type WriteRequestHeader struct {
	// 流量唯一id
	// required: true
	// in: header
	TraceID string `json:"traceId"`
	// // demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token
	// // required: true
	// // in: header
	// Authorization string `json:"authorization"`
	// // 用户uuid
	// // in: header
	// BmpUserID string `json:"bmpUserId"`
	// // 用户语言 [zh_CN, en_US]
	// // in: header
	// BmpLanguage string `json:"bmpLanguage"`
}

// unit err response for non 200
// swagger:response ErrorResponse
type ErrorResponse struct {
	ResponseHeader
	//In: body
	Body struct {
		// 错误信息
		// required: true
		Err controllers.ErrorResponse `json:"error"`
		// 流量ID
		// required: true
		RequestId string `json:"requestId"`
	}
}

// 返回body里面的公共字段
type ResponseHeader struct {
	// 流量ID
	// in: header
	TraceID string `json:"traceId"`
}
