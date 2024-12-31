package request

import (
	"errors"
	"fmt"
	"git.jd.com/cps-golang/ironic-common/exception"

	"github.com/beego/beego/v2/core/validation"
)

type IRequest interface {
	Validate(request IRequest) error
}

type BaseRequest struct{}

func (r *BaseRequest) Validate(request IRequest) error {
	valid := validation.Validation{}
	ok, err := valid.Valid(request)
	fmt.Println("校验结果", ok, err)
	if err != nil {
		panic(exception.Exception{Msg: err.Error()})
		return err
	}
	if !ok {
		fmt.Println("返回的err是", errors.New(valid.Errors[0].Key+" "+valid.Errors[0].Error()))
		panic(exception.Exception{Msg: errors.New(valid.Errors[0].Key + " " + valid.Errors[0].Error()).Error()})
	}
	return nil
}
