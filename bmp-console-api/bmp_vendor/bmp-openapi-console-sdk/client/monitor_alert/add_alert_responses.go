// Code generated by go-swagger; DO NOT EDIT.

package monitor_alert

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

// AddAlertReader is a Reader for the AddAlert structure.
type AddAlertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddAlertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddAlertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddAlertDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddAlertOK creates a AddAlertOK with default headers values
func NewAddAlertOK() *AddAlertOK {
	return &AddAlertOK{}
}

/* AddAlertOK describes a response with status code 200, with default header values.

A AddAlertResult is an response struct that is used to describe create Alert result.
*/
type AddAlertOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *AddAlertOKBody
}

// IsSuccess returns true when this add alert o k response has a 2xx status code
func (o *AddAlertOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add alert o k response has a 3xx status code
func (o *AddAlertOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add alert o k response has a 4xx status code
func (o *AddAlertOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add alert o k response has a 5xx status code
func (o *AddAlertOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add alert o k response a status code equal to that given
func (o *AddAlertOK) IsCode(code int) bool {
	return code == 200
}

func (o *AddAlertOK) Error() string {
	return fmt.Sprintf("[POST /monitorAlert/addAlert][%d] addAlertOK  %+v", 200, o.Payload)
}

func (o *AddAlertOK) String() string {
	return fmt.Sprintf("[POST /monitorAlert/addAlert][%d] addAlertOK  %+v", 200, o.Payload)
}

func (o *AddAlertOK) GetPayload() *AddAlertOKBody {
	return o.Payload
}

func (o *AddAlertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(AddAlertOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAlertDefault creates a AddAlertDefault with default headers values
func NewAddAlertDefault(code int) *AddAlertDefault {
	return &AddAlertDefault{
		_statusCode: code,
	}
}

/* AddAlertDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type AddAlertDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *AddAlertDefaultBody
}

// Code gets the status code for the add alert default response
func (o *AddAlertDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this add alert default response has a 2xx status code
func (o *AddAlertDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add alert default response has a 3xx status code
func (o *AddAlertDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add alert default response has a 4xx status code
func (o *AddAlertDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add alert default response has a 5xx status code
func (o *AddAlertDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add alert default response a status code equal to that given
func (o *AddAlertDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AddAlertDefault) Error() string {
	return fmt.Sprintf("[POST /monitorAlert/addAlert][%d] addAlert default  %+v", o._statusCode, o.Payload)
}

func (o *AddAlertDefault) String() string {
	return fmt.Sprintf("[POST /monitorAlert/addAlert][%d] addAlert default  %+v", o._statusCode, o.Payload)
}

func (o *AddAlertDefault) GetPayload() *AddAlertDefaultBody {
	return o.Payload
}

func (o *AddAlertDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(AddAlertDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddAlertDefaultBody add alert default body
swagger:model AddAlertDefaultBody
*/
type AddAlertDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this add alert default body
func (o *AddAlertDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *AddAlertDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("addAlert default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *AddAlertDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("addAlert default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addAlert default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addAlert default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add alert default body based on the context it is used
func (o *AddAlertDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddAlertDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addAlert default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addAlert default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddAlertDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddAlertDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddAlertDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddAlertOKBody add alert o k body
swagger:model AddAlertOKBody
*/
type AddAlertOKBody struct {

	// result
	Result *models.CommonResponse `json:"result,omitempty"`
}

// Validate validates this add alert o k body
func (o *AddAlertOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddAlertOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addAlertOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addAlertOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add alert o k body based on the context it is used
func (o *AddAlertOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddAlertOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addAlertOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addAlertOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddAlertOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddAlertOKBody) UnmarshalBinary(b []byte) error {
	var res AddAlertOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
