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

// DescribeLocalUserReader is a Reader for the DescribeLocalUser structure.
type DescribeLocalUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeLocalUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeLocalUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeLocalUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeLocalUserOK creates a DescribeLocalUserOK with default headers values
func NewDescribeLocalUserOK() *DescribeLocalUserOK {
	return &DescribeLocalUserOK{}
}

/* DescribeLocalUserOK describes a response with status code 200, with default header values.

A DescribeLocalUserResult is an response struct that is used to describe getuser.
*/
type DescribeLocalUserOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DescribeLocalUserOKBody
}

// IsSuccess returns true when this describe local user o k response has a 2xx status code
func (o *DescribeLocalUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe local user o k response has a 3xx status code
func (o *DescribeLocalUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe local user o k response has a 4xx status code
func (o *DescribeLocalUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe local user o k response has a 5xx status code
func (o *DescribeLocalUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe local user o k response a status code equal to that given
func (o *DescribeLocalUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *DescribeLocalUserOK) Error() string {
	return fmt.Sprintf("[GET /local/users][%d] describeLocalUserOK  %+v", 200, o.Payload)
}

func (o *DescribeLocalUserOK) String() string {
	return fmt.Sprintf("[GET /local/users][%d] describeLocalUserOK  %+v", 200, o.Payload)
}

func (o *DescribeLocalUserOK) GetPayload() *DescribeLocalUserOKBody {
	return o.Payload
}

func (o *DescribeLocalUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DescribeLocalUserOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeLocalUserDefault creates a DescribeLocalUserDefault with default headers values
func NewDescribeLocalUserDefault(code int) *DescribeLocalUserDefault {
	return &DescribeLocalUserDefault{
		_statusCode: code,
	}
}

/* DescribeLocalUserDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type DescribeLocalUserDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DescribeLocalUserDefaultBody
}

// Code gets the status code for the describe local user default response
func (o *DescribeLocalUserDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this describe local user default response has a 2xx status code
func (o *DescribeLocalUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe local user default response has a 3xx status code
func (o *DescribeLocalUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe local user default response has a 4xx status code
func (o *DescribeLocalUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe local user default response has a 5xx status code
func (o *DescribeLocalUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe local user default response a status code equal to that given
func (o *DescribeLocalUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DescribeLocalUserDefault) Error() string {
	return fmt.Sprintf("[GET /local/users][%d] describeLocalUser default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeLocalUserDefault) String() string {
	return fmt.Sprintf("[GET /local/users][%d] describeLocalUser default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeLocalUserDefault) GetPayload() *DescribeLocalUserDefaultBody {
	return o.Payload
}

func (o *DescribeLocalUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DescribeLocalUserDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DescribeLocalUserDefaultBody describe local user default body
swagger:model DescribeLocalUserDefaultBody
*/
type DescribeLocalUserDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this describe local user default body
func (o *DescribeLocalUserDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *DescribeLocalUserDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("describeLocalUser default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *DescribeLocalUserDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("describeLocalUser default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeLocalUser default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeLocalUser default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe local user default body based on the context it is used
func (o *DescribeLocalUserDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeLocalUserDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeLocalUser default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeLocalUser default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DescribeLocalUserDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DescribeLocalUserDefaultBody) UnmarshalBinary(b []byte) error {
	var res DescribeLocalUserDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DescribeLocalUserOKBody describe local user o k body
swagger:model DescribeLocalUserOKBody
*/
type DescribeLocalUserOKBody struct {

	// error
	Error *models.ErrorResponse `json:"error,omitempty"`

	// result
	Result *models.UserInfo `json:"result,omitempty"`
}

// Validate validates this describe local user o k body
func (o *DescribeLocalUserOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DescribeLocalUserOKBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeLocalUserOK" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeLocalUserOK" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *DescribeLocalUserOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeLocalUserOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeLocalUserOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe local user o k body based on the context it is used
func (o *DescribeLocalUserOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *DescribeLocalUserOKBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeLocalUserOK" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeLocalUserOK" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *DescribeLocalUserOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeLocalUserOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeLocalUserOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DescribeLocalUserOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DescribeLocalUserOKBody) UnmarshalBinary(b []byte) error {
	var res DescribeLocalUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}