// Code generated by go-swagger; DO NOT EDIT.

package monitor_rule

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

// DisableRuleReader is a Reader for the DisableRule structure.
type DisableRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDisableRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDisableRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDisableRuleOK creates a DisableRuleOK with default headers values
func NewDisableRuleOK() *DisableRuleOK {
	return &DisableRuleOK{}
}

/* DisableRuleOK describes a response with status code 200, with default header values.

A DisableRuleResult is an response struct that is used to describe create rule result.
*/
type DisableRuleOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DisableRuleOKBody
}

// IsSuccess returns true when this disable rule o k response has a 2xx status code
func (o *DisableRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this disable rule o k response has a 3xx status code
func (o *DisableRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable rule o k response has a 4xx status code
func (o *DisableRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this disable rule o k response has a 5xx status code
func (o *DisableRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this disable rule o k response a status code equal to that given
func (o *DisableRuleOK) IsCode(code int) bool {
	return code == 200
}

func (o *DisableRuleOK) Error() string {
	return fmt.Sprintf("[PUT /monitorRule/disableRule][%d] disableRuleOK  %+v", 200, o.Payload)
}

func (o *DisableRuleOK) String() string {
	return fmt.Sprintf("[PUT /monitorRule/disableRule][%d] disableRuleOK  %+v", 200, o.Payload)
}

func (o *DisableRuleOK) GetPayload() *DisableRuleOKBody {
	return o.Payload
}

func (o *DisableRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DisableRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisableRuleDefault creates a DisableRuleDefault with default headers values
func NewDisableRuleDefault(code int) *DisableRuleDefault {
	return &DisableRuleDefault{
		_statusCode: code,
	}
}

/* DisableRuleDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type DisableRuleDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DisableRuleDefaultBody
}

// Code gets the status code for the disable rule default response
func (o *DisableRuleDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this disable rule default response has a 2xx status code
func (o *DisableRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this disable rule default response has a 3xx status code
func (o *DisableRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this disable rule default response has a 4xx status code
func (o *DisableRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this disable rule default response has a 5xx status code
func (o *DisableRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this disable rule default response a status code equal to that given
func (o *DisableRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DisableRuleDefault) Error() string {
	return fmt.Sprintf("[PUT /monitorRule/disableRule][%d] disableRule default  %+v", o._statusCode, o.Payload)
}

func (o *DisableRuleDefault) String() string {
	return fmt.Sprintf("[PUT /monitorRule/disableRule][%d] disableRule default  %+v", o._statusCode, o.Payload)
}

func (o *DisableRuleDefault) GetPayload() *DisableRuleDefaultBody {
	return o.Payload
}

func (o *DisableRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DisableRuleDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DisableRuleDefaultBody disable rule default body
swagger:model DisableRuleDefaultBody
*/
type DisableRuleDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this disable rule default body
func (o *DisableRuleDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *DisableRuleDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("disableRule default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *DisableRuleDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("disableRule default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disableRule default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disableRule default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this disable rule default body based on the context it is used
func (o *DisableRuleDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DisableRuleDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disableRule default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disableRule default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DisableRuleDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DisableRuleDefaultBody) UnmarshalBinary(b []byte) error {
	var res DisableRuleDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DisableRuleOKBody disable rule o k body
swagger:model DisableRuleOKBody
*/
type DisableRuleOKBody struct {

	// result
	Result *models.CommonResponse `json:"result,omitempty"`
}

// Validate validates this disable rule o k body
func (o *DisableRuleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DisableRuleOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disableRuleOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disableRuleOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this disable rule o k body based on the context it is used
func (o *DisableRuleOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DisableRuleOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disableRuleOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disableRuleOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DisableRuleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DisableRuleOKBody) UnmarshalBinary(b []byte) error {
	var res DisableRuleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}