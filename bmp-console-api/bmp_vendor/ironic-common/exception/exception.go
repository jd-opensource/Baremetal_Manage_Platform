package exception

import (
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

type Exception struct {
	Status    int `json:"code"`
	ErrorCode string `json:"status"`
	Msg       string `json:"message"`
}

var CommonParamValid = Exception{
	Status:    httpStatus.BAD_REQUEST,
	ErrorCode: errorCode.INVALID_ARGUMENT,
	Msg:       "param invalid",
}

var InternalError = Exception{
	Status:    httpStatus.INTERNAL_SERVER_ERROR,
	ErrorCode: errorCode.INTERNAL,
	Msg:       "内部错误",
}
