// Code generated by go-swagger; DO NOT EDIT.

package user

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

// DescribeUserReader is a Reader for the DescribeUser structure.
type DescribeUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeUserOK creates a DescribeUserOK with default headers values
func NewDescribeUserOK() *DescribeUserOK {
	return &DescribeUserOK{}
}

/* DescribeUserOK describes a response with status code 200, with default header values.

DescribeUserOK describe user o k
*/
type DescribeUserOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DescribeUserOKBody
}

// IsSuccess returns true when this describe user o k response has a 2xx status code
func (o *DescribeUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe user o k response has a 3xx status code
func (o *DescribeUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe user o k response has a 4xx status code
func (o *DescribeUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe user o k response has a 5xx status code
func (o *DescribeUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe user o k response a status code equal to that given
func (o *DescribeUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *DescribeUserOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] describeUserOK  %+v", 200, o.Payload)
}

func (o *DescribeUserOK) String() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] describeUserOK  %+v", 200, o.Payload)
}

func (o *DescribeUserOK) GetPayload() *DescribeUserOKBody {
	return o.Payload
}

func (o *DescribeUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DescribeUserOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeUserDefault creates a DescribeUserDefault with default headers values
func NewDescribeUserDefault(code int) *DescribeUserDefault {
	return &DescribeUserDefault{
		_statusCode: code,
	}
}

/* DescribeUserDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type DescribeUserDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DescribeUserDefaultBody
}

// Code gets the status code for the describe user default response
func (o *DescribeUserDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this describe user default response has a 2xx status code
func (o *DescribeUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe user default response has a 3xx status code
func (o *DescribeUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe user default response has a 4xx status code
func (o *DescribeUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe user default response has a 5xx status code
func (o *DescribeUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe user default response a status code equal to that given
func (o *DescribeUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DescribeUserDefault) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] describeUser default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeUserDefault) String() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] describeUser default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeUserDefault) GetPayload() *DescribeUserDefaultBody {
	return o.Payload
}

func (o *DescribeUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DescribeUserDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DescribeUserDefaultBody describe user default body
swagger:model DescribeUserDefaultBody
*/
type DescribeUserDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this describe user default body
func (o *DescribeUserDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *DescribeUserDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("describeUser default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *DescribeUserDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("describeUser default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeUser default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeUser default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe user default body based on the context it is used
func (o *DescribeUserDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeUserDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeUser default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeUser default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DescribeUserDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DescribeUserDefaultBody) UnmarshalBinary(b []byte) error {
	var res DescribeUserDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DescribeUserOKBody describe user o k body
swagger:model DescribeUserOKBody
*/
type DescribeUserOKBody struct {

	// error
	Error *models.ErrorResponse `json:"error,omitempty"`

	// result
	Result *models.UserInfo `json:"result,omitempty"`
}

// Validate validates this describe user o k body
func (o *DescribeUserOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeUserOKBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeUserOK" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeUserOK" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *DescribeUserOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeUserOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeUserOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe user o k body based on the context it is used
func (o *DescribeUserOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeUserOKBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeUserOK" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeUserOK" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *DescribeUserOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeUserOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeUserOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DescribeUserOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DescribeUserOKBody) UnmarshalBinary(b []byte) error {
	var res DescribeUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}