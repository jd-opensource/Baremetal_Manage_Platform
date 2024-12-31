package request

import (
	"errors"

	log "coding.jd.com/aidc-bmp/bmp_log"
	validator "github.com/go-playground/validator/v10"
)

var va = validator.New()

func validate(req interface{}, logger *log.Logger) error {
	if err := va.Struct(req); err != nil {
		logger.Warn("Validate error:", err.Error())
		var e string
		v, ok := err.(validator.ValidationErrors)
		if !ok {
			e = err.Error()
		} else {
			e = v[0].Field()
		}
		return errors.New(e + " param invalid")
	}
	return nil
}
