// Code generated by go-swagger; DO NOT EDIT.

package monitor_proxy

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

// DesrcibeAgentStatusReader is a Reader for the DesrcibeAgentStatus structure.
type DesrcibeAgentStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DesrcibeAgentStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDesrcibeAgentStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDesrcibeAgentStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDesrcibeAgentStatusOK creates a DesrcibeAgentStatusOK with default headers values
func NewDesrcibeAgentStatusOK() *DesrcibeAgentStatusOK {
	return &DesrcibeAgentStatusOK{}
}

/* DesrcibeAgentStatusOK describes a response with status code 200, with default header values.

DesrcibeAgentStatusOK desrcibe agent status o k
*/
type DesrcibeAgentStatusOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DesrcibeAgentStatusOKBody
}

// IsSuccess returns true when this desrcibe agent status o k response has a 2xx status code
func (o *DesrcibeAgentStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this desrcibe agent status o k response has a 3xx status code
func (o *DesrcibeAgentStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this desrcibe agent status o k response has a 4xx status code
func (o *DesrcibeAgentStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this desrcibe agent status o k response has a 5xx status code
func (o *DesrcibeAgentStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this desrcibe agent status o k response a status code equal to that given
func (o *DesrcibeAgentStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *DesrcibeAgentStatusOK) Error() string {
	return fmt.Sprintf("[GET /monitorProxy/desrcibeAgentStatus][%d] desrcibeAgentStatusOK  %+v", 200, o.Payload)
}

func (o *DesrcibeAgentStatusOK) String() string {
	return fmt.Sprintf("[GET /monitorProxy/desrcibeAgentStatus][%d] desrcibeAgentStatusOK  %+v", 200, o.Payload)
}

func (o *DesrcibeAgentStatusOK) GetPayload() *DesrcibeAgentStatusOKBody {
	return o.Payload
}

func (o *DesrcibeAgentStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DesrcibeAgentStatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesrcibeAgentStatusDefault creates a DesrcibeAgentStatusDefault with default headers values
func NewDesrcibeAgentStatusDefault(code int) *DesrcibeAgentStatusDefault {
	return &DesrcibeAgentStatusDefault{
		_statusCode: code,
	}
}

/* DesrcibeAgentStatusDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type DesrcibeAgentStatusDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DesrcibeAgentStatusDefaultBody
}

// Code gets the status code for the desrcibe agent status default response
func (o *DesrcibeAgentStatusDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this desrcibe agent status default response has a 2xx status code
func (o *DesrcibeAgentStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this desrcibe agent status default response has a 3xx status code
func (o *DesrcibeAgentStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this desrcibe agent status default response has a 4xx status code
func (o *DesrcibeAgentStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this desrcibe agent status default response has a 5xx status code
func (o *DesrcibeAgentStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this desrcibe agent status default response a status code equal to that given
func (o *DesrcibeAgentStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DesrcibeAgentStatusDefault) Error() string {
	return fmt.Sprintf("[GET /monitorProxy/desrcibeAgentStatus][%d] desrcibeAgentStatus default  %+v", o._statusCode, o.Payload)
}

func (o *DesrcibeAgentStatusDefault) String() string {
	return fmt.Sprintf("[GET /monitorProxy/desrcibeAgentStatus][%d] desrcibeAgentStatus default  %+v", o._statusCode, o.Payload)
}

func (o *DesrcibeAgentStatusDefault) GetPayload() *DesrcibeAgentStatusDefaultBody {
	return o.Payload
}

func (o *DesrcibeAgentStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DesrcibeAgentStatusDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DesrcibeAgentStatusDefaultBody desrcibe agent status default body
swagger:model DesrcibeAgentStatusDefaultBody
*/
type DesrcibeAgentStatusDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this desrcibe agent status default body
func (o *DesrcibeAgentStatusDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *DesrcibeAgentStatusDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("desrcibeAgentStatus default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *DesrcibeAgentStatusDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("desrcibeAgentStatus default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("desrcibeAgentStatus default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("desrcibeAgentStatus default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this desrcibe agent status default body based on the context it is used
func (o *DesrcibeAgentStatusDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DesrcibeAgentStatusDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("desrcibeAgentStatus default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("desrcibeAgentStatus default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DesrcibeAgentStatusDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DesrcibeAgentStatusDefaultBody) UnmarshalBinary(b []byte) error {
	var res DesrcibeAgentStatusDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DesrcibeAgentStatusOKBody desrcibe agent status o k body
swagger:model DesrcibeAgentStatusOKBody
*/
type DesrcibeAgentStatusOKBody struct {

	// result
	Result *models.AgentStatusResponse `json:"result,omitempty"`
}

// Validate validates this desrcibe agent status o k body
func (o *DesrcibeAgentStatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DesrcibeAgentStatusOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("desrcibeAgentStatusOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("desrcibeAgentStatusOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this desrcibe agent status o k body based on the context it is used
func (o *DesrcibeAgentStatusOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DesrcibeAgentStatusOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("desrcibeAgentStatusOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("desrcibeAgentStatusOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DesrcibeAgentStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DesrcibeAgentStatusOKBody) UnmarshalBinary(b []byte) error {
	var res DesrcibeAgentStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
