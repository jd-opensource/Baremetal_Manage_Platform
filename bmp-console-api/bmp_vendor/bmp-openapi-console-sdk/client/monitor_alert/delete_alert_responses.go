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

// DeleteAlertReader is a Reader for the DeleteAlert structure.
type DeleteAlertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAlertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAlertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteAlertDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAlertOK creates a DeleteAlertOK with default headers values
func NewDeleteAlertOK() *DeleteAlertOK {
	return &DeleteAlertOK{}
}

/* DeleteAlertOK describes a response with status code 200, with default header values.

A DeleteAlertResult is an response struct that is used to describe create alert result.
*/
type DeleteAlertOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DeleteAlertOKBody
}

// IsSuccess returns true when this delete alert o k response has a 2xx status code
func (o *DeleteAlertOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete alert o k response has a 3xx status code
func (o *DeleteAlertOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alert o k response has a 4xx status code
func (o *DeleteAlertOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete alert o k response has a 5xx status code
func (o *DeleteAlertOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alert o k response a status code equal to that given
func (o *DeleteAlertOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteAlertOK) Error() string {
	return fmt.Sprintf("[DELETE /monitorAlert/deleteAlert][%d] deleteAlertOK  %+v", 200, o.Payload)
}

func (o *DeleteAlertOK) String() string {
	return fmt.Sprintf("[DELETE /monitorAlert/deleteAlert][%d] deleteAlertOK  %+v", 200, o.Payload)
}

func (o *DeleteAlertOK) GetPayload() *DeleteAlertOKBody {
	return o.Payload
}

func (o *DeleteAlertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DeleteAlertOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertDefault creates a DeleteAlertDefault with default headers values
func NewDeleteAlertDefault(code int) *DeleteAlertDefault {
	return &DeleteAlertDefault{
		_statusCode: code,
	}
}

/* DeleteAlertDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type DeleteAlertDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DeleteAlertDefaultBody
}

// Code gets the status code for the delete alert default response
func (o *DeleteAlertDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete alert default response has a 2xx status code
func (o *DeleteAlertDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete alert default response has a 3xx status code
func (o *DeleteAlertDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete alert default response has a 4xx status code
func (o *DeleteAlertDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete alert default response has a 5xx status code
func (o *DeleteAlertDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete alert default response a status code equal to that given
func (o *DeleteAlertDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteAlertDefault) Error() string {
	return fmt.Sprintf("[DELETE /monitorAlert/deleteAlert][%d] deleteAlert default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAlertDefault) String() string {
	return fmt.Sprintf("[DELETE /monitorAlert/deleteAlert][%d] deleteAlert default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAlertDefault) GetPayload() *DeleteAlertDefaultBody {
	return o.Payload
}

func (o *DeleteAlertDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DeleteAlertDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteAlertDefaultBody delete alert default body
swagger:model DeleteAlertDefaultBody
*/
type DeleteAlertDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this delete alert default body
func (o *DeleteAlertDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *DeleteAlertDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("deleteAlert default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *DeleteAlertDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("deleteAlert default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteAlert default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteAlert default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this delete alert default body based on the context it is used
func (o *DeleteAlertDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteAlertDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteAlert default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteAlert default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteAlertDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteAlertDefaultBody) UnmarshalBinary(b []byte) error {
	var res DeleteAlertDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteAlertOKBody delete alert o k body
swagger:model DeleteAlertOKBody
*/
type DeleteAlertOKBody struct {

	// result
	Result *models.CommonResponse `json:"result,omitempty"`
}

// Validate validates this delete alert o k body
func (o *DeleteAlertOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteAlertOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteAlertOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteAlertOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this delete alert o k body based on the context it is used
func (o *DeleteAlertOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteAlertOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteAlertOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteAlertOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteAlertOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteAlertOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteAlertOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}