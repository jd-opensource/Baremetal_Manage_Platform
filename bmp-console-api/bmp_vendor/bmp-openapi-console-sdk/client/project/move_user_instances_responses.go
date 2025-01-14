// Code generated by go-swagger; DO NOT EDIT.

package project

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

// MoveUserInstancesReader is a Reader for the MoveUserInstances structure.
type MoveUserInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MoveUserInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMoveUserInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMoveUserInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMoveUserInstancesOK creates a MoveUserInstancesOK with default headers values
func NewMoveUserInstancesOK() *MoveUserInstancesOK {
	return &MoveUserInstancesOK{}
}

/* MoveUserInstancesOK describes a response with status code 200, with default header values.

A MoveInstanceResult is an response struct
*/
type MoveUserInstancesOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *MoveUserInstancesOKBody
}

// IsSuccess returns true when this move user instances o k response has a 2xx status code
func (o *MoveUserInstancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this move user instances o k response has a 3xx status code
func (o *MoveUserInstancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move user instances o k response has a 4xx status code
func (o *MoveUserInstancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this move user instances o k response has a 5xx status code
func (o *MoveUserInstancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this move user instances o k response a status code equal to that given
func (o *MoveUserInstancesOK) IsCode(code int) bool {
	return code == 200
}

func (o *MoveUserInstancesOK) Error() string {
	return fmt.Sprintf("[PUT /user/projects/move/instances][%d] moveUserInstancesOK  %+v", 200, o.Payload)
}

func (o *MoveUserInstancesOK) String() string {
	return fmt.Sprintf("[PUT /user/projects/move/instances][%d] moveUserInstancesOK  %+v", 200, o.Payload)
}

func (o *MoveUserInstancesOK) GetPayload() *MoveUserInstancesOKBody {
	return o.Payload
}

func (o *MoveUserInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(MoveUserInstancesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveUserInstancesDefault creates a MoveUserInstancesDefault with default headers values
func NewMoveUserInstancesDefault(code int) *MoveUserInstancesDefault {
	return &MoveUserInstancesDefault{
		_statusCode: code,
	}
}

/* MoveUserInstancesDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type MoveUserInstancesDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *MoveUserInstancesDefaultBody
}

// Code gets the status code for the move user instances default response
func (o *MoveUserInstancesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this move user instances default response has a 2xx status code
func (o *MoveUserInstancesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this move user instances default response has a 3xx status code
func (o *MoveUserInstancesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this move user instances default response has a 4xx status code
func (o *MoveUserInstancesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this move user instances default response has a 5xx status code
func (o *MoveUserInstancesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this move user instances default response a status code equal to that given
func (o *MoveUserInstancesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *MoveUserInstancesDefault) Error() string {
	return fmt.Sprintf("[PUT /user/projects/move/instances][%d] moveUserInstances default  %+v", o._statusCode, o.Payload)
}

func (o *MoveUserInstancesDefault) String() string {
	return fmt.Sprintf("[PUT /user/projects/move/instances][%d] moveUserInstances default  %+v", o._statusCode, o.Payload)
}

func (o *MoveUserInstancesDefault) GetPayload() *MoveUserInstancesDefaultBody {
	return o.Payload
}

func (o *MoveUserInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(MoveUserInstancesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*MoveUserInstancesDefaultBody move user instances default body
swagger:model MoveUserInstancesDefaultBody
*/
type MoveUserInstancesDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this move user instances default body
func (o *MoveUserInstancesDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *MoveUserInstancesDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("moveUserInstances default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *MoveUserInstancesDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("moveUserInstances default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moveUserInstances default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("moveUserInstances default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this move user instances default body based on the context it is used
func (o *MoveUserInstancesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *MoveUserInstancesDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moveUserInstances default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("moveUserInstances default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *MoveUserInstancesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MoveUserInstancesDefaultBody) UnmarshalBinary(b []byte) error {
	var res MoveUserInstancesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MoveUserInstancesOKBody move user instances o k body
swagger:model MoveUserInstancesOKBody
*/
type MoveUserInstancesOKBody struct {

	// result
	Result *models.CommonResponse `json:"result,omitempty"`
}

// Validate validates this move user instances o k body
func (o *MoveUserInstancesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *MoveUserInstancesOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moveUserInstancesOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("moveUserInstancesOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this move user instances o k body based on the context it is used
func (o *MoveUserInstancesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *MoveUserInstancesOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moveUserInstancesOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("moveUserInstancesOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *MoveUserInstancesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MoveUserInstancesOKBody) UnmarshalBinary(b []byte) error {
	var res MoveUserInstancesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
