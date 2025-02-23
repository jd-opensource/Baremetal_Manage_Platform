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

// DeleteOSReader is a Reader for the DeleteOS structure.
type DeleteOSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteOSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOSOK creates a DeleteOSOK with default headers values
func NewDeleteOSOK() *DeleteOSOK {
	return &DeleteOSOK{}
}

/* DeleteOSOK describes a response with status code 200, with default header values.

A DeleteOsResult is an response struct that is used to delete stop os result.
*/
type DeleteOSOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DeleteOSOKBody
}

// IsSuccess returns true when this delete o s o k response has a 2xx status code
func (o *DeleteOSOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete o s o k response has a 3xx status code
func (o *DeleteOSOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete o s o k response has a 4xx status code
func (o *DeleteOSOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete o s o k response has a 5xx status code
func (o *DeleteOSOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete o s o k response a status code equal to that given
func (o *DeleteOSOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteOSOK) Error() string {
	return fmt.Sprintf("[DELETE /oss/{os_id}][%d] deleteOSOK  %+v", 200, o.Payload)
}

func (o *DeleteOSOK) String() string {
	return fmt.Sprintf("[DELETE /oss/{os_id}][%d] deleteOSOK  %+v", 200, o.Payload)
}

func (o *DeleteOSOK) GetPayload() *DeleteOSOKBody {
	return o.Payload
}

func (o *DeleteOSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DeleteOSOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOSDefault creates a DeleteOSDefault with default headers values
func NewDeleteOSDefault(code int) *DeleteOSDefault {
	return &DeleteOSDefault{
		_statusCode: code,
	}
}

/* DeleteOSDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type DeleteOSDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DeleteOSDefaultBody
}

// Code gets the status code for the delete o s default response
func (o *DeleteOSDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete o s default response has a 2xx status code
func (o *DeleteOSDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete o s default response has a 3xx status code
func (o *DeleteOSDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete o s default response has a 4xx status code
func (o *DeleteOSDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete o s default response has a 5xx status code
func (o *DeleteOSDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete o s default response a status code equal to that given
func (o *DeleteOSDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteOSDefault) Error() string {
	return fmt.Sprintf("[DELETE /oss/{os_id}][%d] deleteOS default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOSDefault) String() string {
	return fmt.Sprintf("[DELETE /oss/{os_id}][%d] deleteOS default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOSDefault) GetPayload() *DeleteOSDefaultBody {
	return o.Payload
}

func (o *DeleteOSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DeleteOSDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteOSDefaultBody delete o s default body
swagger:model DeleteOSDefaultBody
*/
type DeleteOSDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this delete o s default body
func (o *DeleteOSDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *DeleteOSDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("deleteOS default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *DeleteOSDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("deleteOS default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteOS default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteOS default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this delete o s default body based on the context it is used
func (o *DeleteOSDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteOSDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteOS default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteOS default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteOSDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteOSDefaultBody) UnmarshalBinary(b []byte) error {
	var res DeleteOSDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteOSOKBody delete o s o k body
swagger:model DeleteOSOKBody
*/
type DeleteOSOKBody struct {

	// result
	Result *models.CommonResponse `json:"result,omitempty"`
}

// Validate validates this delete o s o k body
func (o *DeleteOSOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteOSOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteOSOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteOSOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this delete o s o k body based on the context it is used
func (o *DeleteOSOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteOSOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteOSOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteOSOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteOSOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteOSOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteOSOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
