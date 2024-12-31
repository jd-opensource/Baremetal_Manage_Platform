package errors

import "fmt"

const (
	DEFAULT_BMP_ERROR_CODE             = "ERROR"
	DEFAULT_BMP_ERROR_MESSAGE          = "Error"
	USERNAME_OR_PASSWORD_ERROR_CODE    = "USERNAME_OR_PASSWORD_ERROR"
	USERNAME_OR_PASSWORD_ERROR_MESSAGE = "User name or password error:%s"
	USER_PRIVILEGE_ERROR_CODE          = "USER_PRIVILEGE_ERROR"
	USER_PRIVILEGE_ERROR_MESSAGE       = "Current user:%s privilege is %s not %s"
	PING_ILOIP_ERROR_CODE              = "PING_ILOIP_ERROR"
	PING_ILOIP_ERROR_MESSAGE           = "Ping iloip error."
	PING_SUBNET_GATEWAY_ERROR_CODE     = "PING_SUBNET_GATEWAY_ERROR"
	PING_SUBNET_GATEWAY_ERROR_MESSAGE  = "Ping subnet gateway error."
	NETWORK_STATUS_ERROR_CODE          = "NETWORK_STATUS_ERROR_CODE"
	NETWORK_STATUS_ERROR_MESSAGE       = "Network status error:%s."
	MAC1_ERROR_CODE                    = "MAC1_ERROR"
	MAC1_ERROR_MESSAGE                 = "Mac1 connectivity error"
	MAC2_ERROR_CODE                    = "MAC2_ERROR"
	MAC2_ERROR_MESSAGE                 = "Mac2 connectivity error"
)

type BmpError struct {
	errorCode   string
	message     string
	originError error
}

func NewBmpError(errorCode, message string, originError error) Error {
	return &BmpError{
		errorCode:   errorCode,
		message:     message,
		originError: originError,
	}
}

func (err *BmpError) ErrorCode() string {
	if err.errorCode == "" {
		return DEFAULT_BMP_ERROR_CODE
	} else {
		return err.errorCode
	}
}

func (err *BmpError) Error() string {
	bmpErrMsg := fmt.Sprintf("[%s] %s", err.ErrorCode(), err.message)
	if err.originError != nil {
		return bmpErrMsg + "\ncaused by:\n" + err.originError.Error()
	}
	return bmpErrMsg
}

func (err *BmpError) OriginError() error {
	return err.originError
}

func (err *BmpError) Message() string {
	return err.message
}

func (err *BmpError) String() string {
	return err.Error()
}
