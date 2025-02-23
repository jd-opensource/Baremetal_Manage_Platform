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

	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
)

// AddRuleReader is a Reader for the AddRule structure.
type AddRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddRuleOK creates a AddRuleOK with default headers values
func NewAddRuleOK() *AddRuleOK {
	return &AddRuleOK{}
}

/* AddRuleOK describes a response with status code 200, with default header values.

A AddRuleResult is an response struct that is used to describe create rule result.
*/
type AddRuleOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *AddRuleOKBody
}

// IsSuccess returns true when this add rule o k response has a 2xx status code
func (o *AddRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add rule o k response has a 3xx status code
func (o *AddRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add rule o k response has a 4xx status code
func (o *AddRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add rule o k response has a 5xx status code
func (o *AddRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add rule o k response a status code equal to that given
func (o *AddRuleOK) IsCode(code int) bool {
	return code == 200
}

func (o *AddRuleOK) Error() string {
	return fmt.Sprintf("[POST /monitorRule/addRule][%d] addRuleOK  %+v", 200, o.Payload)
}

func (o *AddRuleOK) String() string {
	return fmt.Sprintf("[POST /monitorRule/addRule][%d] addRuleOK  %+v", 200, o.Payload)
}

func (o *AddRuleOK) GetPayload() *AddRuleOKBody {
	return o.Payload
}

func (o *AddRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(AddRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddRuleDefault creates a AddRuleDefault with default headers values
func NewAddRuleDefault(code int) *AddRuleDefault {
	return &AddRuleDefault{
		_statusCode: code,
	}
}

/* AddRuleDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type AddRuleDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *AddRuleDefaultBody
}

// Code gets the status code for the add rule default response
func (o *AddRuleDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this add rule default response has a 2xx status code
func (o *AddRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add rule default response has a 3xx status code
func (o *AddRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add rule default response has a 4xx status code
func (o *AddRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add rule default response has a 5xx status code
func (o *AddRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add rule default response a status code equal to that given
func (o *AddRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AddRuleDefault) Error() string {
	return fmt.Sprintf("[POST /monitorRule/addRule][%d] addRule default  %+v", o._statusCode, o.Payload)
}

func (o *AddRuleDefault) String() string {
	return fmt.Sprintf("[POST /monitorRule/addRule][%d] addRule default  %+v", o._statusCode, o.Payload)
}

func (o *AddRuleDefault) GetPayload() *AddRuleDefaultBody {
	return o.Payload
}

func (o *AddRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(AddRuleDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddRuleDefaultBody add rule default body
swagger:model AddRuleDefaultBody
*/
type AddRuleDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this add rule default body
func (o *AddRuleDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *AddRuleDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("addRule default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *AddRuleDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("addRule default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addRule default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addRule default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add rule default body based on the context it is used
func (o *AddRuleDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddRuleDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addRule default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addRule default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddRuleDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRuleDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddRuleDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRuleOKBody add rule o k body
swagger:model AddRuleOKBody
*/
type AddRuleOKBody struct {

	// result
	Result *models.CommonResponse `json:"result,omitempty"`
}

// Validate validates this add rule o k body
func (o *AddRuleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddRuleOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addRuleOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addRuleOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add rule o k body based on the context it is used
func (o *AddRuleOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddRuleOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addRuleOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addRuleOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddRuleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRuleOKBody) UnmarshalBinary(b []byte) error {
	var res AddRuleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
