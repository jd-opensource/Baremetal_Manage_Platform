// Code generated by go-swagger; DO NOT EDIT.

package sshkey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
)

// CreateUserSshkeyReader is a Reader for the CreateUserSshkey structure.
type CreateUserSshkeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserSshkeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUserSshkeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateUserSshkeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateUserSshkeyOK creates a CreateUserSshkeyOK with default headers values
func NewCreateUserSshkeyOK() *CreateUserSshkeyOK {
	return &CreateUserSshkeyOK{}
}

/* CreateUserSshkeyOK describes a response with status code 200, with default header values.

A CreateUserSshkeyResult is an response struct that is used to describe create sshkey result.
*/
type CreateUserSshkeyOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *CreateUserSshkeyOKBody
}

// IsSuccess returns true when this create user sshkey o k response has a 2xx status code
func (o *CreateUserSshkeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user sshkey o k response has a 3xx status code
func (o *CreateUserSshkeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user sshkey o k response has a 4xx status code
func (o *CreateUserSshkeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user sshkey o k response has a 5xx status code
func (o *CreateUserSshkeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create user sshkey o k response a status code equal to that given
func (o *CreateUserSshkeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateUserSshkeyOK) Error() string {
	return fmt.Sprintf("[POST /user/sshkeys][%d] createUserSshkeyOK  %+v", 200, o.Payload)
}

func (o *CreateUserSshkeyOK) String() string {
	return fmt.Sprintf("[POST /user/sshkeys][%d] createUserSshkeyOK  %+v", 200, o.Payload)
}

func (o *CreateUserSshkeyOK) GetPayload() *CreateUserSshkeyOKBody {
	return o.Payload
}

func (o *CreateUserSshkeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(CreateUserSshkeyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserSshkeyDefault creates a CreateUserSshkeyDefault with default headers values
func NewCreateUserSshkeyDefault(code int) *CreateUserSshkeyDefault {
	return &CreateUserSshkeyDefault{
		_statusCode: code,
	}
}

/* CreateUserSshkeyDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type CreateUserSshkeyDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *CreateUserSshkeyDefaultBody
}

// Code gets the status code for the create user sshkey default response
func (o *CreateUserSshkeyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create user sshkey default response has a 2xx status code
func (o *CreateUserSshkeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create user sshkey default response has a 3xx status code
func (o *CreateUserSshkeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create user sshkey default response has a 4xx status code
func (o *CreateUserSshkeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create user sshkey default response has a 5xx status code
func (o *CreateUserSshkeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create user sshkey default response a status code equal to that given
func (o *CreateUserSshkeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateUserSshkeyDefault) Error() string {
	return fmt.Sprintf("[POST /user/sshkeys][%d] createUserSshkey default  %+v", o._statusCode, o.Payload)
}

func (o *CreateUserSshkeyDefault) String() string {
	return fmt.Sprintf("[POST /user/sshkeys][%d] createUserSshkey default  %+v", o._statusCode, o.Payload)
}

func (o *CreateUserSshkeyDefault) GetPayload() *CreateUserSshkeyDefaultBody {
	return o.Payload
}

func (o *CreateUserSshkeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(CreateUserSshkeyDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateUserSshkeyDefaultBody create user sshkey default body
swagger:model CreateUserSshkeyDefaultBody
*/
type CreateUserSshkeyDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this create user sshkey default body
func (o *CreateUserSshkeyDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRequestID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateUserSshkeyDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("createUserSshkey default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *CreateUserSshkeyDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("createUserSshkey default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createUserSshkey default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createUserSshkey default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create user sshkey default body based on the context it is used
func (o *CreateUserSshkeyDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateUserSshkeyDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createUserSshkey default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createUserSshkey default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateUserSshkeyDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateUserSshkeyDefaultBody) UnmarshalBinary(b []byte) error {
	var res CreateUserSshkeyDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateUserSshkeyOKBody create user sshkey o k body
swagger:model CreateUserSshkeyOKBody
*/
type CreateUserSshkeyOKBody struct {

	// result
	Result *models.SshkeyID `json:"result,omitempty"`
}

// Validate validates this create user sshkey o k body
func (o *CreateUserSshkeyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateUserSshkeyOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createUserSshkeyOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createUserSshkeyOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create user sshkey o k body based on the context it is used
func (o *CreateUserSshkeyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateUserSshkeyOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createUserSshkeyOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createUserSshkeyOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateUserSshkeyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateUserSshkeyOKBody) UnmarshalBinary(b []byte) error {
	var res CreateUserSshkeyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}