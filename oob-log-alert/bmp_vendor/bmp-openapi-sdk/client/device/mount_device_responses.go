// Code generated by go-swagger; DO NOT EDIT.

package device

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

// MountDeviceReader is a Reader for the MountDevice structure.
type MountDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MountDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMountDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMountDeviceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMountDeviceOK creates a MountDeviceOK with default headers values
func NewMountDeviceOK() *MountDeviceOK {
	return &MountDeviceOK{}
}

/* MountDeviceOK describes a response with status code 200, with default header values.

A MountDeviceResult is an response struct that is used to describe mount device result.
*/
type MountDeviceOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *MountDeviceOKBody
}

// IsSuccess returns true when this mount device o k response has a 2xx status code
func (o *MountDeviceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this mount device o k response has a 3xx status code
func (o *MountDeviceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this mount device o k response has a 4xx status code
func (o *MountDeviceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this mount device o k response has a 5xx status code
func (o *MountDeviceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this mount device o k response a status code equal to that given
func (o *MountDeviceOK) IsCode(code int) bool {
	return code == 200
}

func (o *MountDeviceOK) Error() string {
	return fmt.Sprintf("[PUT /devices/mount][%d] mountDeviceOK  %+v", 200, o.Payload)
}

func (o *MountDeviceOK) String() string {
	return fmt.Sprintf("[PUT /devices/mount][%d] mountDeviceOK  %+v", 200, o.Payload)
}

func (o *MountDeviceOK) GetPayload() *MountDeviceOKBody {
	return o.Payload
}

func (o *MountDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(MountDeviceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMountDeviceDefault creates a MountDeviceDefault with default headers values
func NewMountDeviceDefault(code int) *MountDeviceDefault {
	return &MountDeviceDefault{
		_statusCode: code,
	}
}

/* MountDeviceDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type MountDeviceDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *MountDeviceDefaultBody
}

// Code gets the status code for the mount device default response
func (o *MountDeviceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this mount device default response has a 2xx status code
func (o *MountDeviceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this mount device default response has a 3xx status code
func (o *MountDeviceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this mount device default response has a 4xx status code
func (o *MountDeviceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this mount device default response has a 5xx status code
func (o *MountDeviceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this mount device default response a status code equal to that given
func (o *MountDeviceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *MountDeviceDefault) Error() string {
	return fmt.Sprintf("[PUT /devices/mount][%d] mountDevice default  %+v", o._statusCode, o.Payload)
}

func (o *MountDeviceDefault) String() string {
	return fmt.Sprintf("[PUT /devices/mount][%d] mountDevice default  %+v", o._statusCode, o.Payload)
}

func (o *MountDeviceDefault) GetPayload() *MountDeviceDefaultBody {
	return o.Payload
}

func (o *MountDeviceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(MountDeviceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*MountDeviceDefaultBody mount device default body
swagger:model MountDeviceDefaultBody
*/
type MountDeviceDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this mount device default body
func (o *MountDeviceDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *MountDeviceDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("mountDevice default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *MountDeviceDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("mountDevice default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mountDevice default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mountDevice default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mount device default body based on the context it is used
func (o *MountDeviceDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *MountDeviceDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mountDevice default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mountDevice default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *MountDeviceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MountDeviceDefaultBody) UnmarshalBinary(b []byte) error {
	var res MountDeviceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MountDeviceOKBody mount device o k body
swagger:model MountDeviceOKBody
*/
type MountDeviceOKBody struct {

	// result
	Result *models.CommonResponse `json:"result,omitempty"`
}

// Validate validates this mount device o k body
func (o *MountDeviceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *MountDeviceOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mountDeviceOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mountDeviceOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mount device o k body based on the context it is used
func (o *MountDeviceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *MountDeviceOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mountDeviceOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mountDeviceOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *MountDeviceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MountDeviceOKBody) UnmarshalBinary(b []byte) error {
	var res MountDeviceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}