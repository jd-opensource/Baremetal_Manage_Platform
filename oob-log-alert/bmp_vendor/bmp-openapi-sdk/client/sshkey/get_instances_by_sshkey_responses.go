// Code generated by go-swagger; DO NOT EDIT.

package sshkey

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

// GetInstancesBySshkeyReader is a Reader for the GetInstancesBySshkey structure.
type GetInstancesBySshkeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesBySshkeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstancesBySshkeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetInstancesBySshkeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetInstancesBySshkeyOK creates a GetInstancesBySshkeyOK with default headers values
func NewGetInstancesBySshkeyOK() *GetInstancesBySshkeyOK {
	return &GetInstancesBySshkeyOK{}
}

/* GetInstancesBySshkeyOK describes a response with status code 200, with default header values.

A GetInstancesBySshkeyResult is an response struct that is used to get instance list by sshkey.
*/
type GetInstancesBySshkeyOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *GetInstancesBySshkeyOKBody
}

// IsSuccess returns true when this get instances by sshkey o k response has a 2xx status code
func (o *GetInstancesBySshkeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get instances by sshkey o k response has a 3xx status code
func (o *GetInstancesBySshkeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get instances by sshkey o k response has a 4xx status code
func (o *GetInstancesBySshkeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get instances by sshkey o k response has a 5xx status code
func (o *GetInstancesBySshkeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get instances by sshkey o k response a status code equal to that given
func (o *GetInstancesBySshkeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInstancesBySshkeyOK) Error() string {
	return fmt.Sprintf("[GET /user/sshkeys/instances/{sshkey_id}][%d] getInstancesBySshkeyOK  %+v", 200, o.Payload)
}

func (o *GetInstancesBySshkeyOK) String() string {
	return fmt.Sprintf("[GET /user/sshkeys/instances/{sshkey_id}][%d] getInstancesBySshkeyOK  %+v", 200, o.Payload)
}

func (o *GetInstancesBySshkeyOK) GetPayload() *GetInstancesBySshkeyOKBody {
	return o.Payload
}

func (o *GetInstancesBySshkeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(GetInstancesBySshkeyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesBySshkeyDefault creates a GetInstancesBySshkeyDefault with default headers values
func NewGetInstancesBySshkeyDefault(code int) *GetInstancesBySshkeyDefault {
	return &GetInstancesBySshkeyDefault{
		_statusCode: code,
	}
}

/* GetInstancesBySshkeyDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type GetInstancesBySshkeyDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *GetInstancesBySshkeyDefaultBody
}

// Code gets the status code for the get instances by sshkey default response
func (o *GetInstancesBySshkeyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get instances by sshkey default response has a 2xx status code
func (o *GetInstancesBySshkeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get instances by sshkey default response has a 3xx status code
func (o *GetInstancesBySshkeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get instances by sshkey default response has a 4xx status code
func (o *GetInstancesBySshkeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get instances by sshkey default response has a 5xx status code
func (o *GetInstancesBySshkeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get instances by sshkey default response a status code equal to that given
func (o *GetInstancesBySshkeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetInstancesBySshkeyDefault) Error() string {
	return fmt.Sprintf("[GET /user/sshkeys/instances/{sshkey_id}][%d] getInstancesBySshkey default  %+v", o._statusCode, o.Payload)
}

func (o *GetInstancesBySshkeyDefault) String() string {
	return fmt.Sprintf("[GET /user/sshkeys/instances/{sshkey_id}][%d] getInstancesBySshkey default  %+v", o._statusCode, o.Payload)
}

func (o *GetInstancesBySshkeyDefault) GetPayload() *GetInstancesBySshkeyDefaultBody {
	return o.Payload
}

func (o *GetInstancesBySshkeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(GetInstancesBySshkeyDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetInstancesBySshkeyDefaultBody get instances by sshkey default body
swagger:model GetInstancesBySshkeyDefaultBody
*/
type GetInstancesBySshkeyDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this get instances by sshkey default body
func (o *GetInstancesBySshkeyDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *GetInstancesBySshkeyDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("getInstancesBySshkey default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *GetInstancesBySshkeyDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getInstancesBySshkey default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getInstancesBySshkey default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getInstancesBySshkey default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get instances by sshkey default body based on the context it is used
func (o *GetInstancesBySshkeyDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInstancesBySshkeyDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getInstancesBySshkey default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getInstancesBySshkey default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetInstancesBySshkeyDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstancesBySshkeyDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetInstancesBySshkeyDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetInstancesBySshkeyOKBody get instances by sshkey o k body
swagger:model GetInstancesBySshkeyOKBody
*/
type GetInstancesBySshkeyOKBody struct {

	// result
	Result *models.InstancesSshkeyInfoResponse `json:"result,omitempty"`
}

// Validate validates this get instances by sshkey o k body
func (o *GetInstancesBySshkeyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInstancesBySshkeyOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getInstancesBySshkeyOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getInstancesBySshkeyOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get instances by sshkey o k body based on the context it is used
func (o *GetInstancesBySshkeyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInstancesBySshkeyOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getInstancesBySshkeyOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getInstancesBySshkeyOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetInstancesBySshkeyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstancesBySshkeyOKBody) UnmarshalBinary(b []byte) error {
	var res GetInstancesBySshkeyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
