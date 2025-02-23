// Code generated by go-swagger; DO NOT EDIT.

package instance

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

// DescribeProjectInstanceReader is a Reader for the DescribeProjectInstance structure.
type DescribeProjectInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeProjectInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeProjectInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeProjectInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeProjectInstanceOK creates a DescribeProjectInstanceOK with default headers values
func NewDescribeProjectInstanceOK() *DescribeProjectInstanceOK {
	return &DescribeProjectInstanceOK{}
}

/* DescribeProjectInstanceOK describes a response with status code 200, with default header values.

A GetInstanceResult is an response struct that is used to describe instance.
*/
type DescribeProjectInstanceOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DescribeProjectInstanceOKBody
}

// IsSuccess returns true when this describe project instance o k response has a 2xx status code
func (o *DescribeProjectInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe project instance o k response has a 3xx status code
func (o *DescribeProjectInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe project instance o k response has a 4xx status code
func (o *DescribeProjectInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe project instance o k response has a 5xx status code
func (o *DescribeProjectInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe project instance o k response a status code equal to that given
func (o *DescribeProjectInstanceOK) IsCode(code int) bool {
	return code == 200
}

func (o *DescribeProjectInstanceOK) Error() string {
	return fmt.Sprintf("[GET /project/instances/{instance_id}][%d] describeProjectInstanceOK  %+v", 200, o.Payload)
}

func (o *DescribeProjectInstanceOK) String() string {
	return fmt.Sprintf("[GET /project/instances/{instance_id}][%d] describeProjectInstanceOK  %+v", 200, o.Payload)
}

func (o *DescribeProjectInstanceOK) GetPayload() *DescribeProjectInstanceOKBody {
	return o.Payload
}

func (o *DescribeProjectInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DescribeProjectInstanceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeProjectInstanceDefault creates a DescribeProjectInstanceDefault with default headers values
func NewDescribeProjectInstanceDefault(code int) *DescribeProjectInstanceDefault {
	return &DescribeProjectInstanceDefault{
		_statusCode: code,
	}
}

/* DescribeProjectInstanceDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type DescribeProjectInstanceDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DescribeProjectInstanceDefaultBody
}

// Code gets the status code for the describe project instance default response
func (o *DescribeProjectInstanceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this describe project instance default response has a 2xx status code
func (o *DescribeProjectInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe project instance default response has a 3xx status code
func (o *DescribeProjectInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe project instance default response has a 4xx status code
func (o *DescribeProjectInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe project instance default response has a 5xx status code
func (o *DescribeProjectInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe project instance default response a status code equal to that given
func (o *DescribeProjectInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DescribeProjectInstanceDefault) Error() string {
	return fmt.Sprintf("[GET /project/instances/{instance_id}][%d] describeProjectInstance default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeProjectInstanceDefault) String() string {
	return fmt.Sprintf("[GET /project/instances/{instance_id}][%d] describeProjectInstance default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeProjectInstanceDefault) GetPayload() *DescribeProjectInstanceDefaultBody {
	return o.Payload
}

func (o *DescribeProjectInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DescribeProjectInstanceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DescribeProjectInstanceDefaultBody describe project instance default body
swagger:model DescribeProjectInstanceDefaultBody
*/
type DescribeProjectInstanceDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this describe project instance default body
func (o *DescribeProjectInstanceDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *DescribeProjectInstanceDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("describeProjectInstance default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *DescribeProjectInstanceDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("describeProjectInstance default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeProjectInstance default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeProjectInstance default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe project instance default body based on the context it is used
func (o *DescribeProjectInstanceDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeProjectInstanceDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeProjectInstance default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeProjectInstance default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DescribeProjectInstanceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DescribeProjectInstanceDefaultBody) UnmarshalBinary(b []byte) error {
	var res DescribeProjectInstanceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DescribeProjectInstanceOKBody describe project instance o k body
swagger:model DescribeProjectInstanceOKBody
*/
type DescribeProjectInstanceOKBody struct {

	// result
	Result *models.InstanceInfo `json:"result,omitempty"`
}

// Validate validates this describe project instance o k body
func (o *DescribeProjectInstanceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeProjectInstanceOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeProjectInstanceOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeProjectInstanceOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe project instance o k body based on the context it is used
func (o *DescribeProjectInstanceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeProjectInstanceOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeProjectInstanceOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeProjectInstanceOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DescribeProjectInstanceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DescribeProjectInstanceOKBody) UnmarshalBinary(b []byte) error {
	var res DescribeProjectInstanceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
