// Code generated by go-swagger; DO NOT EDIT.

package device_type

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

// AssociatedImageReader is a Reader for the AssociatedImage structure.
type AssociatedImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssociatedImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssociatedImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAssociatedImageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssociatedImageOK creates a AssociatedImageOK with default headers values
func NewAssociatedImageOK() *AssociatedImageOK {
	return &AssociatedImageOK{}
}

/* AssociatedImageOK describes a response with status code 200, with default header values.

A AssociatedImageResult is an response struct that is used to associated image result.
*/
type AssociatedImageOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *AssociatedImageOKBody
}

// IsSuccess returns true when this associated image o k response has a 2xx status code
func (o *AssociatedImageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this associated image o k response has a 3xx status code
func (o *AssociatedImageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this associated image o k response has a 4xx status code
func (o *AssociatedImageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this associated image o k response has a 5xx status code
func (o *AssociatedImageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this associated image o k response a status code equal to that given
func (o *AssociatedImageOK) IsCode(code int) bool {
	return code == 200
}

func (o *AssociatedImageOK) Error() string {
	return fmt.Sprintf("[POST /deviceTypes/associatedImage][%d] associatedImageOK  %+v", 200, o.Payload)
}

func (o *AssociatedImageOK) String() string {
	return fmt.Sprintf("[POST /deviceTypes/associatedImage][%d] associatedImageOK  %+v", 200, o.Payload)
}

func (o *AssociatedImageOK) GetPayload() *AssociatedImageOKBody {
	return o.Payload
}

func (o *AssociatedImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(AssociatedImageOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssociatedImageDefault creates a AssociatedImageDefault with default headers values
func NewAssociatedImageDefault(code int) *AssociatedImageDefault {
	return &AssociatedImageDefault{
		_statusCode: code,
	}
}

/* AssociatedImageDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type AssociatedImageDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *AssociatedImageDefaultBody
}

// Code gets the status code for the associated image default response
func (o *AssociatedImageDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this associated image default response has a 2xx status code
func (o *AssociatedImageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this associated image default response has a 3xx status code
func (o *AssociatedImageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this associated image default response has a 4xx status code
func (o *AssociatedImageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this associated image default response has a 5xx status code
func (o *AssociatedImageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this associated image default response a status code equal to that given
func (o *AssociatedImageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AssociatedImageDefault) Error() string {
	return fmt.Sprintf("[POST /deviceTypes/associatedImage][%d] associatedImage default  %+v", o._statusCode, o.Payload)
}

func (o *AssociatedImageDefault) String() string {
	return fmt.Sprintf("[POST /deviceTypes/associatedImage][%d] associatedImage default  %+v", o._statusCode, o.Payload)
}

func (o *AssociatedImageDefault) GetPayload() *AssociatedImageDefaultBody {
	return o.Payload
}

func (o *AssociatedImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(AssociatedImageDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AssociatedImageDefaultBody associated image default body
swagger:model AssociatedImageDefaultBody
*/
type AssociatedImageDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this associated image default body
func (o *AssociatedImageDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *AssociatedImageDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("associatedImage default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *AssociatedImageDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("associatedImage default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("associatedImage default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("associatedImage default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this associated image default body based on the context it is used
func (o *AssociatedImageDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AssociatedImageDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("associatedImage default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("associatedImage default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AssociatedImageDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AssociatedImageDefaultBody) UnmarshalBinary(b []byte) error {
	var res AssociatedImageDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AssociatedImageOKBody associated image o k body
swagger:model AssociatedImageOKBody
*/
type AssociatedImageOKBody struct {

	// result
	Result *models.CommonResponse `json:"result,omitempty"`
}

// Validate validates this associated image o k body
func (o *AssociatedImageOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AssociatedImageOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("associatedImageOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("associatedImageOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this associated image o k body based on the context it is used
func (o *AssociatedImageOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AssociatedImageOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("associatedImageOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("associatedImageOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AssociatedImageOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AssociatedImageOKBody) UnmarshalBinary(b []byte) error {
	var res AssociatedImageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
