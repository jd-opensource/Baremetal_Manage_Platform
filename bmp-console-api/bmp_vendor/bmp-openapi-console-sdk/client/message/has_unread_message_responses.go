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

	"coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
)

// HasUnreadMessageReader is a Reader for the HasUnreadMessage structure.
type HasUnreadMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HasUnreadMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHasUnreadMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHasUnreadMessageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHasUnreadMessageOK creates a HasUnreadMessageOK with default headers values
func NewHasUnreadMessageOK() *HasUnreadMessageOK {
	return &HasUnreadMessageOK{}
}

/* HasUnreadMessageOK describes a response with status code 200, with default header values.

HasUnreadMessageOK has unread message o k
*/
type HasUnreadMessageOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *HasUnreadMessageOKBody
}

// IsSuccess returns true when this has unread message o k response has a 2xx status code
func (o *HasUnreadMessageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this has unread message o k response has a 3xx status code
func (o *HasUnreadMessageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this has unread message o k response has a 4xx status code
func (o *HasUnreadMessageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this has unread message o k response has a 5xx status code
func (o *HasUnreadMessageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this has unread message o k response a status code equal to that given
func (o *HasUnreadMessageOK) IsCode(code int) bool {
	return code == 200
}

func (o *HasUnreadMessageOK) Error() string {
	return fmt.Sprintf("[GET /messages/hasUnreadMessage][%d] hasUnreadMessageOK  %+v", 200, o.Payload)
}

func (o *HasUnreadMessageOK) String() string {
	return fmt.Sprintf("[GET /messages/hasUnreadMessage][%d] hasUnreadMessageOK  %+v", 200, o.Payload)
}

func (o *HasUnreadMessageOK) GetPayload() *HasUnreadMessageOKBody {
	return o.Payload
}

func (o *HasUnreadMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(HasUnreadMessageOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHasUnreadMessageDefault creates a HasUnreadMessageDefault with default headers values
func NewHasUnreadMessageDefault(code int) *HasUnreadMessageDefault {
	return &HasUnreadMessageDefault{
		_statusCode: code,
	}
}

/* HasUnreadMessageDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type HasUnreadMessageDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *HasUnreadMessageDefaultBody
}

// Code gets the status code for the has unread message default response
func (o *HasUnreadMessageDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this has unread message default response has a 2xx status code
func (o *HasUnreadMessageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this has unread message default response has a 3xx status code
func (o *HasUnreadMessageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this has unread message default response has a 4xx status code
func (o *HasUnreadMessageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this has unread message default response has a 5xx status code
func (o *HasUnreadMessageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this has unread message default response a status code equal to that given
func (o *HasUnreadMessageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HasUnreadMessageDefault) Error() string {
	return fmt.Sprintf("[GET /messages/hasUnreadMessage][%d] hasUnreadMessage default  %+v", o._statusCode, o.Payload)
}

func (o *HasUnreadMessageDefault) String() string {
	return fmt.Sprintf("[GET /messages/hasUnreadMessage][%d] hasUnreadMessage default  %+v", o._statusCode, o.Payload)
}

func (o *HasUnreadMessageDefault) GetPayload() *HasUnreadMessageDefaultBody {
	return o.Payload
}

func (o *HasUnreadMessageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(HasUnreadMessageDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*HasUnreadMessageDefaultBody has unread message default body
swagger:model HasUnreadMessageDefaultBody
*/
type HasUnreadMessageDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this has unread message default body
func (o *HasUnreadMessageDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *HasUnreadMessageDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("hasUnreadMessage default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *HasUnreadMessageDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("hasUnreadMessage default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hasUnreadMessage default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hasUnreadMessage default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this has unread message default body based on the context it is used
func (o *HasUnreadMessageDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *HasUnreadMessageDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hasUnreadMessage default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hasUnreadMessage default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *HasUnreadMessageDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *HasUnreadMessageDefaultBody) UnmarshalBinary(b []byte) error {
	var res HasUnreadMessageDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*HasUnreadMessageOKBody has unread message o k body
swagger:model HasUnreadMessageOKBody
*/
type HasUnreadMessageOKBody struct {

	// result
	Result *models.HasUnreadMessage `json:"result,omitempty"`
}

// Validate validates this has unread message o k body
func (o *HasUnreadMessageOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *HasUnreadMessageOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hasUnreadMessageOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hasUnreadMessageOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this has unread message o k body based on the context it is used
func (o *HasUnreadMessageOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *HasUnreadMessageOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hasUnreadMessageOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hasUnreadMessageOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *HasUnreadMessageOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *HasUnreadMessageOKBody) UnmarshalBinary(b []byte) error {
	var res HasUnreadMessageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
