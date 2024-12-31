package request

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	log "coding.jd.com/aidc-bmp/bmp_log"
	"github.com/go-playground/validator/v10"
)

type CancelCommandsRequest struct {
	Sn string `json:"sn"`
}
type RetryCommandRequest struct {
	OffsetCommandId *int64 `json:"offset_command_id" validate:"required"`
}

func (c *RetryCommandRequest) Validate(logger *log.Logger) {
	if err := validator.New().Struct(c); err != nil {
		logger.Warn("RetryCommandRequest.Validate error:", err.Error())
		panic(constant.INVALID_ARGUMENT)
	}
}

type QueryCommandsRequest struct {
	RequestId  string
	Sn         string
	InstanceId string
	PageNumber int
	PageSize   int
}

func (c *QueryCommandsRequest) Validate(logger *log.Logger) {
	if err := validator.New().Struct(c); err != nil {
		logger.Warn("QueryCommandsRequest.Validate error:", err.Error())
		panic(constant.INVALID_ARGUMENT)
	}
}
