package request

import (
	"fmt"
	"reflect"
	"strings"

	"coding.jd.com/bmp/agent-proxy-jdstack/constant"
	log "git.jd.com/cps-golang/log"
	validator "github.com/go-playground/validator/v10"
)

type Pageable struct {
	// 页码
	PageNumber int64 `json:"pageNumber"`
	// 每页数量
	PageSize int64 `json:"pageSize"`
}

var va = validator.New()

func InitValidator() {
	va.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		// if name == "-" {
		// 	return ""
		// }

		return name
	})

	if err := va.RegisterValidation("required_in_ptr", ValidateRequiredInPtr); err != nil {
		fmt.Println("required_in_ptr register error!!!")
	} else {
		fmt.Println("required_in_ptr register success!!!")
	}
}

func ValidateRequiredInPtr(fl validator.FieldLevel) bool {

	return true
}

func validate(req interface{}, logger *log.Logger) {
	if err := va.Struct(req); err != nil {
		fmt.Println(err, err.Error())
		logger.Warn("Validate error:", err.Error())
		var e string
		v, ok := err.(validator.ValidationErrors)
		if !ok {
			e = err.Error()
		} else {
			e = v[0].Field()
		}
		panic(constant.BuildInvalidArgumentWithArgs(e+" 非法", e+" invalid"))
	}
}
