// Code generated by go-swagger; DO NOT EDIT.

package message

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

// DialMailReader is a Reader for the DialMail structure.
type DialMailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DialMailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDialMailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDialMailDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDialMailOK creates a DialMailOK with default headers values
func NewDialMailOK() *DialMailOK {
	return &DialMailOK{}
}

/* DialMailOK describes a response with status code 200, with default header values.

A DialMailResult is an response struct
*/
type DialMailOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DialMailOKBody
}

// IsSuccess returns true when this dial mail o k response has a 2xx status code
func (o *DialMailOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dial mail o k response has a 3xx status code
func (o *DialMailOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dial mail o k response has a 4xx status code
func (o *DialMailOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dial mail o k response has a 5xx status code
func (o *DialMailOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dial mail o k response a status code equal to that given
func (o *DialMailOK) IsCode(code int) bool {
	return code == 200
}

func (o *DialMailOK) Error() string {
	return fmt.Sprintf("[POST /messages/dialMail][%d] dialMailOK  %+v", 200, o.Payload)
}

func (o *DialMailOK) String() string {
	return fmt.Sprintf("[POST /messages/dialMail][%d] dialMailOK  %+v", 200, o.Payload)
}

func (o *DialMailOK) GetPayload() *DialMailOKBody {
	return o.Payload
}

func (o *DialMailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DialMailOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDialMailDefault creates a DialMailDefault with default headers values
func NewDialMailDefault(code int) *DialMailDefault {
	return &DialMailDefault{
		_statusCode: code,
	}
}

/* DialMailDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type DialMailDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DialMailDefaultBody
}

// Code gets the status code for the dial mail default response
func (o *DialMailDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dial mail default response has a 2xx status code
func (o *DialMailDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dial mail default response has a 3xx status code
func (o *DialMailDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dial mail default response has a 4xx status code
func (o *DialMailDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dial mail default response has a 5xx status code
func (o *DialMailDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dial mail default response a status code equal to that given
func (o *DialMailDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DialMailDefault) Error() string {
	return fmt.Sprintf("[POST /messages/dialMail][%d] dialMail default  %+v", o._statusCode, o.Payload)
}

func (o *DialMailDefault) String() string {
	return fmt.Sprintf("[POST /messages/dialMail][%d] dialMail default  %+v", o._statusCode, o.Payload)
}

func (o *DialMailDefault) GetPayload() *DialMailDefaultBody {
	return o.Payload
}

func (o *DialMailDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DialMailDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DialMailDefaultBody dial mail default body
swagger:model DialMailDefaultBody
*/
type DialMailDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this dial mail default body
func (o *DialMailDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *DialMailDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("dialMail default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *DialMailDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("dialMail default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dialMail default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dialMail default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this dial mail default body based on the context it is used
func (o *DialMailDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DialMailDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dialMail default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dialMail default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DialMailDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DialMailDefaultBody) UnmarshalBinary(b []byte) error {
	var res DialMailDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DialMailOKBody dial mail o k body
swagger:model DialMailOKBody
*/
type DialMailOKBody struct {

	// result
	Result *models.CommonResponse `json:"result,omitempty"`
}

// Validate validates this dial mail o k body
func (o *DialMailOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DialMailOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dialMailOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dialMailOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this dial mail o k body based on the context it is used
func (o *DialMailOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DialMailOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dialMailOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dialMailOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DialMailOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DialMailOKBody) UnmarshalBinary(b []byte) error {
	var res DialMailOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}