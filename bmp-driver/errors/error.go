package errors

type Error interface {
	error
	ErrorCode() string
	Message() string
	OriginError() error
}
