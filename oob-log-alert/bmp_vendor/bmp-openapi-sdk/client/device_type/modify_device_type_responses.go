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

// ModifyDeviceTypeReader is a Reader for the ModifyDeviceType structure.
type ModifyDeviceTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyDeviceTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModifyDeviceTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewModifyDeviceTypeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewModifyDeviceTypeOK creates a ModifyDeviceTypeOK with default headers values
func NewModifyDeviceTypeOK() *ModifyDeviceTypeOK {
	return &ModifyDeviceTypeOK{}
}

/* ModifyDeviceTypeOK describes a response with status code 200, with default header values.

A ModifyDeviceTypeResult is an response struct that is used to modify device type.
*/
type ModifyDeviceTypeOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *ModifyDeviceTypeOKBody
}

// IsSuccess returns true when this modify device type o k response has a 2xx status code
func (o *ModifyDeviceTypeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this modify device type o k response has a 3xx status code
func (o *ModifyDeviceTypeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify device type o k response has a 4xx status code
func (o *ModifyDeviceTypeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this modify device type o k response has a 5xx status code
func (o *ModifyDeviceTypeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this modify device type o k response a status code equal to that given
func (o *ModifyDeviceTypeOK) IsCode(code int) bool {
	return code == 200
}

func (o *ModifyDeviceTypeOK) Error() string {
	return fmt.Sprintf("[PUT /deviceTypes/{device_type_id}][%d] modifyDeviceTypeOK  %+v", 200, o.Payload)
}

func (o *ModifyDeviceTypeOK) String() string {
	return fmt.Sprintf("[PUT /deviceTypes/{device_type_id}][%d] modifyDeviceTypeOK  %+v", 200, o.Payload)
}

func (o *ModifyDeviceTypeOK) GetPayload() *ModifyDeviceTypeOKBody {
	return o.Payload
}

func (o *ModifyDeviceTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(ModifyDeviceTypeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyDeviceTypeDefault creates a ModifyDeviceTypeDefault with default headers values
func NewModifyDeviceTypeDefault(code int) *ModifyDeviceTypeDefault {
	return &ModifyDeviceTypeDefault{
		_statusCode: code,
	}
}

/* ModifyDeviceTypeDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type ModifyDeviceTypeDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *ModifyDeviceTypeDefaultBody
}

// Code gets the status code for the modify device type default response
func (o *ModifyDeviceTypeDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this modify device type default response has a 2xx status code
func (o *ModifyDeviceTypeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this modify device type default response has a 3xx status code
func (o *ModifyDeviceTypeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this modify device type default response has a 4xx status code
func (o *ModifyDeviceTypeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this modify device type default response has a 5xx status code
func (o *ModifyDeviceTypeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this modify device type default response a status code equal to that given
func (o *ModifyDeviceTypeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ModifyDeviceTypeDefault) Error() string {
	return fmt.Sprintf("[PUT /deviceTypes/{device_type_id}][%d] modifyDeviceType default  %+v", o._statusCode, o.Payload)
}

func (o *ModifyDeviceTypeDefault) String() string {
	return fmt.Sprintf("[PUT /deviceTypes/{device_type_id}][%d] modifyDeviceType default  %+v", o._statusCode, o.Payload)
}

func (o *ModifyDeviceTypeDefault) GetPayload() *ModifyDeviceTypeDefaultBody {
	return o.Payload
}

func (o *ModifyDeviceTypeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(ModifyDeviceTypeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ModifyDeviceTypeDefaultBody modify device type default body
swagger:model ModifyDeviceTypeDefaultBody
*/
type ModifyDeviceTypeDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this modify device type default body
func (o *ModifyDeviceTypeDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *ModifyDeviceTypeDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("modifyDeviceType default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *ModifyDeviceTypeDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("modifyDeviceType default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifyDeviceType default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifyDeviceType default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this modify device type default body based on the context it is used
func (o *ModifyDeviceTypeDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ModifyDeviceTypeDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifyDeviceType default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifyDeviceType default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ModifyDeviceTypeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ModifyDeviceTypeDefaultBody) UnmarshalBinary(b []byte) error {
	var res ModifyDeviceTypeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ModifyDeviceTypeOKBody modify device type o k body
swagger:model ModifyDeviceTypeOKBody
*/
type ModifyDeviceTypeOKBody struct {

	// result
	Result *models.CommonResponse `json:"result,omitempty"`
}

// Validate validates this modify device type o k body
func (o *ModifyDeviceTypeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ModifyDeviceTypeOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifyDeviceTypeOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifyDeviceTypeOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this modify device type o k body based on the context it is used
func (o *ModifyDeviceTypeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ModifyDeviceTypeOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifyDeviceTypeOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifyDeviceTypeOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ModifyDeviceTypeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ModifyDeviceTypeOKBody) UnmarshalBinary(b []byte) error {
	var res ModifyDeviceTypeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
