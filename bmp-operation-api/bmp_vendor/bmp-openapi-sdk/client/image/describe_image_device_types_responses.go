// Code generated by go-swagger; DO NOT EDIT.

package image

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

// DescribeImageDeviceTypesReader is a Reader for the DescribeImageDeviceTypes structure.
type DescribeImageDeviceTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeImageDeviceTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeImageDeviceTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeImageDeviceTypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeImageDeviceTypesOK creates a DescribeImageDeviceTypesOK with default headers values
func NewDescribeImageDeviceTypesOK() *DescribeImageDeviceTypesOK {
	return &DescribeImageDeviceTypesOK{}
}

/* DescribeImageDeviceTypesOK describes a response with status code 200, with default header values.

DescribeImageDeviceTypesOK describe image device types o k
*/
type DescribeImageDeviceTypesOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DescribeImageDeviceTypesOKBody
}

// IsSuccess returns true when this describe image device types o k response has a 2xx status code
func (o *DescribeImageDeviceTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe image device types o k response has a 3xx status code
func (o *DescribeImageDeviceTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe image device types o k response has a 4xx status code
func (o *DescribeImageDeviceTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe image device types o k response has a 5xx status code
func (o *DescribeImageDeviceTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe image device types o k response a status code equal to that given
func (o *DescribeImageDeviceTypesOK) IsCode(code int) bool {
	return code == 200
}

func (o *DescribeImageDeviceTypesOK) Error() string {
	return fmt.Sprintf("[GET /images/imageDeviceTypes][%d] describeImageDeviceTypesOK  %+v", 200, o.Payload)
}

func (o *DescribeImageDeviceTypesOK) String() string {
	return fmt.Sprintf("[GET /images/imageDeviceTypes][%d] describeImageDeviceTypesOK  %+v", 200, o.Payload)
}

func (o *DescribeImageDeviceTypesOK) GetPayload() *DescribeImageDeviceTypesOKBody {
	return o.Payload
}

func (o *DescribeImageDeviceTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DescribeImageDeviceTypesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeImageDeviceTypesDefault creates a DescribeImageDeviceTypesDefault with default headers values
func NewDescribeImageDeviceTypesDefault(code int) *DescribeImageDeviceTypesDefault {
	return &DescribeImageDeviceTypesDefault{
		_statusCode: code,
	}
}

/* DescribeImageDeviceTypesDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type DescribeImageDeviceTypesDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DescribeImageDeviceTypesDefaultBody
}

// Code gets the status code for the describe image device types default response
func (o *DescribeImageDeviceTypesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this describe image device types default response has a 2xx status code
func (o *DescribeImageDeviceTypesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe image device types default response has a 3xx status code
func (o *DescribeImageDeviceTypesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe image device types default response has a 4xx status code
func (o *DescribeImageDeviceTypesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe image device types default response has a 5xx status code
func (o *DescribeImageDeviceTypesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe image device types default response a status code equal to that given
func (o *DescribeImageDeviceTypesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DescribeImageDeviceTypesDefault) Error() string {
	return fmt.Sprintf("[GET /images/imageDeviceTypes][%d] describeImageDeviceTypes default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeImageDeviceTypesDefault) String() string {
	return fmt.Sprintf("[GET /images/imageDeviceTypes][%d] describeImageDeviceTypes default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeImageDeviceTypesDefault) GetPayload() *DescribeImageDeviceTypesDefaultBody {
	return o.Payload
}

func (o *DescribeImageDeviceTypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DescribeImageDeviceTypesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DescribeImageDeviceTypesDefaultBody describe image device types default body
swagger:model DescribeImageDeviceTypesDefaultBody
*/
type DescribeImageDeviceTypesDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this describe image device types default body
func (o *DescribeImageDeviceTypesDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *DescribeImageDeviceTypesDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("describeImageDeviceTypes default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *DescribeImageDeviceTypesDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("describeImageDeviceTypes default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeImageDeviceTypes default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeImageDeviceTypes default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe image device types default body based on the context it is used
func (o *DescribeImageDeviceTypesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeImageDeviceTypesDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeImageDeviceTypes default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeImageDeviceTypes default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DescribeImageDeviceTypesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DescribeImageDeviceTypesDefaultBody) UnmarshalBinary(b []byte) error {
	var res DescribeImageDeviceTypesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DescribeImageDeviceTypesOKBody describe image device types o k body
swagger:model DescribeImageDeviceTypesOKBody
*/
type DescribeImageDeviceTypesOKBody struct {

	// result
	Result *models.DeviceTypeList `json:"result,omitempty"`
}

// Validate validates this describe image device types o k body
func (o *DescribeImageDeviceTypesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeImageDeviceTypesOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeImageDeviceTypesOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeImageDeviceTypesOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe image device types o k body based on the context it is used
func (o *DescribeImageDeviceTypesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeImageDeviceTypesOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeImageDeviceTypesOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeImageDeviceTypesOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DescribeImageDeviceTypesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DescribeImageDeviceTypesOKBody) UnmarshalBinary(b []byte) error {
	var res DescribeImageDeviceTypesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}