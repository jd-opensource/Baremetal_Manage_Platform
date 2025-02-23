// Code generated by go-swagger; DO NOT EDIT.

package os

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

	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
)

// CreateOSReader is a Reader for the CreateOS structure.
type CreateOSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateOSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOSOK creates a CreateOSOK with default headers values
func NewCreateOSOK() *CreateOSOK {
	return &CreateOSOK{}
}

/* CreateOSOK describes a response with status code 200, with default header values.

A CreateOsResult is an response struct that is used to describe create os result.
*/
type CreateOSOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *CreateOSOKBody
}

// IsSuccess returns true when this create o s o k response has a 2xx status code
func (o *CreateOSOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create o s o k response has a 3xx status code
func (o *CreateOSOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create o s o k response has a 4xx status code
func (o *CreateOSOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create o s o k response has a 5xx status code
func (o *CreateOSOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create o s o k response a status code equal to that given
func (o *CreateOSOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateOSOK) Error() string {
	return fmt.Sprintf("[POST /oss][%d] createOSOK  %+v", 200, o.Payload)
}

func (o *CreateOSOK) String() string {
	return fmt.Sprintf("[POST /oss][%d] createOSOK  %+v", 200, o.Payload)
}

func (o *CreateOSOK) GetPayload() *CreateOSOKBody {
	return o.Payload
}

func (o *CreateOSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(CreateOSOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOSDefault creates a CreateOSDefault with default headers values
func NewCreateOSDefault(code int) *CreateOSDefault {
	return &CreateOSDefault{
		_statusCode: code,
	}
}

/* CreateOSDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type CreateOSDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *CreateOSDefaultBody
}

// Code gets the status code for the create o s default response
func (o *CreateOSDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create o s default response has a 2xx status code
func (o *CreateOSDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create o s default response has a 3xx status code
func (o *CreateOSDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create o s default response has a 4xx status code
func (o *CreateOSDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create o s default response has a 5xx status code
func (o *CreateOSDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create o s default response a status code equal to that given
func (o *CreateOSDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateOSDefault) Error() string {
	return fmt.Sprintf("[POST /oss][%d] createOS default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOSDefault) String() string {
	return fmt.Sprintf("[POST /oss][%d] createOS default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOSDefault) GetPayload() *CreateOSDefaultBody {
	return o.Payload
}

func (o *CreateOSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(CreateOSDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateOSDefaultBody create o s default body
swagger:model CreateOSDefaultBody
*/
type CreateOSDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this create o s default body
func (o *CreateOSDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *CreateOSDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("createOS default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *CreateOSDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("createOS default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOS default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createOS default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create o s default body based on the context it is used
func (o *CreateOSDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOSDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOS default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createOS default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateOSDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOSDefaultBody) UnmarshalBinary(b []byte) error {
	var res CreateOSDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateOSOKBody create o s o k body
swagger:model CreateOSOKBody
*/
type CreateOSOKBody struct {

	// result
	Result *models.OsID `json:"result,omitempty"`
}

// Validate validates this create o s o k body
func (o *CreateOSOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOSOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOSOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createOSOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create o s o k body based on the context it is used
func (o *CreateOSOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOSOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOSOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createOSOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateOSOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOSOKBody) UnmarshalBinary(b []byte) error {
	var res CreateOSOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
