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

// ModifyDeviceReader is a Reader for the ModifyDevice structure.
type ModifyDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModifyDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewModifyDeviceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewModifyDeviceOK creates a ModifyDeviceOK with default headers values
func NewModifyDeviceOK() *ModifyDeviceOK {
	return &ModifyDeviceOK{}
}

/* ModifyDeviceOK describes a response with status code 200, with default header values.

A ModifyDeviceResult is an response struct that is used to describe modify device result.
*/
type ModifyDeviceOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *ModifyDeviceOKBody
}

// IsSuccess returns true when this modify device o k response has a 2xx status code
func (o *ModifyDeviceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this modify device o k response has a 3xx status code
func (o *ModifyDeviceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify device o k response has a 4xx status code
func (o *ModifyDeviceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this modify device o k response has a 5xx status code
func (o *ModifyDeviceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this modify device o k response a status code equal to that given
func (o *ModifyDeviceOK) IsCode(code int) bool {
	return code == 200
}

func (o *ModifyDeviceOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{device_id}][%d] modifyDeviceOK  %+v", 200, o.Payload)
}

func (o *ModifyDeviceOK) String() string {
	return fmt.Sprintf("[PUT /devices/{device_id}][%d] modifyDeviceOK  %+v", 200, o.Payload)
}

func (o *ModifyDeviceOK) GetPayload() *ModifyDeviceOKBody {
	return o.Payload
}

func (o *ModifyDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(ModifyDeviceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyDeviceDefault creates a ModifyDeviceDefault with default headers values
func NewModifyDeviceDefault(code int) *ModifyDeviceDefault {
	return &ModifyDeviceDefault{
		_statusCode: code,
	}
}

/* ModifyDeviceDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type ModifyDeviceDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *ModifyDeviceDefaultBody
}

// Code gets the status code for the modify device default response
func (o *ModifyDeviceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this modify device default response has a 2xx status code
func (o *ModifyDeviceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this modify device default response has a 3xx status code
func (o *ModifyDeviceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this modify device default response has a 4xx status code
func (o *ModifyDeviceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this modify device default response has a 5xx status code
func (o *ModifyDeviceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this modify device default response a status code equal to that given
func (o *ModifyDeviceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ModifyDeviceDefault) Error() string {
	return fmt.Sprintf("[PUT /devices/{device_id}][%d] modifyDevice default  %+v", o._statusCode, o.Payload)
}

func (o *ModifyDeviceDefault) String() string {
	return fmt.Sprintf("[PUT /devices/{device_id}][%d] modifyDevice default  %+v", o._statusCode, o.Payload)
}

func (o *ModifyDeviceDefault) GetPayload() *ModifyDeviceDefaultBody {
	return o.Payload
}

func (o *ModifyDeviceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(ModifyDeviceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ModifyDeviceDefaultBody modify device default body
swagger:model ModifyDeviceDefaultBody
*/
type ModifyDeviceDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this modify device default body
func (o *ModifyDeviceDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *ModifyDeviceDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("modifyDevice default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *ModifyDeviceDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("modifyDevice default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifyDevice default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifyDevice default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this modify device default body based on the context it is used
func (o *ModifyDeviceDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ModifyDeviceDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifyDevice default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifyDevice default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ModifyDeviceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ModifyDeviceDefaultBody) UnmarshalBinary(b []byte) error {
	var res ModifyDeviceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ModifyDeviceOKBody modify device o k body
swagger:model ModifyDeviceOKBody
*/
type ModifyDeviceOKBody struct {

	// result
	Result *models.CommonResponse `json:"result,omitempty"`
}

// Validate validates this modify device o k body
func (o *ModifyDeviceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ModifyDeviceOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifyDeviceOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifyDeviceOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this modify device o k body based on the context it is used
func (o *ModifyDeviceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ModifyDeviceOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifyDeviceOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifyDeviceOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ModifyDeviceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ModifyDeviceOKBody) UnmarshalBinary(b []byte) error {
	var res ModifyDeviceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}