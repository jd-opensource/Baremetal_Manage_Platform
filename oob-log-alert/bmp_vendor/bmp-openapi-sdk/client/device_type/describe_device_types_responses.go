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

// DescribeDeviceTypesReader is a Reader for the DescribeDeviceTypes structure.
type DescribeDeviceTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeDeviceTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeDeviceTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeDeviceTypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeDeviceTypesOK creates a DescribeDeviceTypesOK with default headers values
func NewDescribeDeviceTypesOK() *DescribeDeviceTypesOK {
	return &DescribeDeviceTypesOK{}
}

/* DescribeDeviceTypesOK describes a response with status code 200, with default header values.

DescribeDeviceTypesOK describe device types o k
*/
type DescribeDeviceTypesOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DescribeDeviceTypesOKBody
}

// IsSuccess returns true when this describe device types o k response has a 2xx status code
func (o *DescribeDeviceTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe device types o k response has a 3xx status code
func (o *DescribeDeviceTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe device types o k response has a 4xx status code
func (o *DescribeDeviceTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe device types o k response has a 5xx status code
func (o *DescribeDeviceTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe device types o k response a status code equal to that given
func (o *DescribeDeviceTypesOK) IsCode(code int) bool {
	return code == 200
}

func (o *DescribeDeviceTypesOK) Error() string {
	return fmt.Sprintf("[GET /deviceTypes][%d] describeDeviceTypesOK  %+v", 200, o.Payload)
}

func (o *DescribeDeviceTypesOK) String() string {
	return fmt.Sprintf("[GET /deviceTypes][%d] describeDeviceTypesOK  %+v", 200, o.Payload)
}

func (o *DescribeDeviceTypesOK) GetPayload() *DescribeDeviceTypesOKBody {
	return o.Payload
}

func (o *DescribeDeviceTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DescribeDeviceTypesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeDeviceTypesDefault creates a DescribeDeviceTypesDefault with default headers values
func NewDescribeDeviceTypesDefault(code int) *DescribeDeviceTypesDefault {
	return &DescribeDeviceTypesDefault{
		_statusCode: code,
	}
}

/* DescribeDeviceTypesDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type DescribeDeviceTypesDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DescribeDeviceTypesDefaultBody
}

// Code gets the status code for the describe device types default response
func (o *DescribeDeviceTypesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this describe device types default response has a 2xx status code
func (o *DescribeDeviceTypesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe device types default response has a 3xx status code
func (o *DescribeDeviceTypesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe device types default response has a 4xx status code
func (o *DescribeDeviceTypesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe device types default response has a 5xx status code
func (o *DescribeDeviceTypesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe device types default response a status code equal to that given
func (o *DescribeDeviceTypesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DescribeDeviceTypesDefault) Error() string {
	return fmt.Sprintf("[GET /deviceTypes][%d] describeDeviceTypes default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeDeviceTypesDefault) String() string {
	return fmt.Sprintf("[GET /deviceTypes][%d] describeDeviceTypes default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeDeviceTypesDefault) GetPayload() *DescribeDeviceTypesDefaultBody {
	return o.Payload
}

func (o *DescribeDeviceTypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DescribeDeviceTypesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DescribeDeviceTypesDefaultBody describe device types default body
swagger:model DescribeDeviceTypesDefaultBody
*/
type DescribeDeviceTypesDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this describe device types default body
func (o *DescribeDeviceTypesDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *DescribeDeviceTypesDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("describeDeviceTypes default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *DescribeDeviceTypesDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("describeDeviceTypes default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeDeviceTypes default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeDeviceTypes default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe device types default body based on the context it is used
func (o *DescribeDeviceTypesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeDeviceTypesDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeDeviceTypes default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeDeviceTypes default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DescribeDeviceTypesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DescribeDeviceTypesDefaultBody) UnmarshalBinary(b []byte) error {
	var res DescribeDeviceTypesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DescribeDeviceTypesOKBody describe device types o k body
swagger:model DescribeDeviceTypesOKBody
*/
type DescribeDeviceTypesOKBody struct {

	// result
	Result *models.DeviceTypeList `json:"result,omitempty"`
}

// Validate validates this describe device types o k body
func (o *DescribeDeviceTypesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeDeviceTypesOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeDeviceTypesOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeDeviceTypesOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe device types o k body based on the context it is used
func (o *DescribeDeviceTypesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeDeviceTypesOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeDeviceTypesOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeDeviceTypesOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DescribeDeviceTypesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DescribeDeviceTypesOKBody) UnmarshalBinary(b []byte) error {
	var res DescribeDeviceTypesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
