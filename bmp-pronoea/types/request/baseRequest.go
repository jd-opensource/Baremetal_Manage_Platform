package request

import (
	"errors"
	"github.com/astaxie/beego/validation"
)

type IRequest interface {
	Validate(request IRequest) error
}

type BaseRequest struct{}

func (r *BaseRequest) Validate(request IRequest) error {
	valid := validation.Validation{}
	ok, err := valid.Valid(request)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New(valid.Errors[0].Key + " " + valid.Errors[0].Error())
	}
	return nil

}

