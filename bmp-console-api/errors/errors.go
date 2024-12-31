package errors

import "fmt"

type Error interface {
	error
	GetErrorCode() int
	GetMessage() string
	GetOriginError() error
}

type CommonError struct {
	errorCode   int
	message     string
	originError error
}

func NewCommonError(errorCode int, message string, originErr error) Error {
	return &CommonError{
		errorCode:   errorCode,
		message:     message,
		originError: originErr,
	}
}

func (e *CommonError) Error() string {
	var msg string
	if e.GetMessage() == "" {
		msg = fmt.Sprintf("%s", GetCodeMsg(e.GetErrorCode()))
	} else {
		msg = fmt.Sprintf("%s:%s", GetCodeMsg(e.GetErrorCode()), e.GetMessage())
	}
	if e.GetOriginError() != nil {
		return msg + " caused by: " + e.GetOriginError().Error()
	}
	return msg
}

func (e *CommonError) GetErrorCode() int {
	return e.errorCode
}

func (e *CommonError) GetOriginError() error {
	return e.originError
}

func (e *CommonError) GetMessage() string {
	return e.message
}
