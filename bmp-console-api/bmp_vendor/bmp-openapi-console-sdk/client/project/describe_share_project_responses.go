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

// DescribeShareProjectReader is a Reader for the DescribeShareProject structure.
type DescribeShareProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeShareProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeShareProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeShareProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeShareProjectOK creates a DescribeShareProjectOK with default headers values
func NewDescribeShareProjectOK() *DescribeShareProjectOK {
	return &DescribeShareProjectOK{}
}

/* DescribeShareProjectOK describes a response with status code 200, with default header values.

A DescribeShareProjectResult is an response struct that is used to describe getapikey.
*/
type DescribeShareProjectOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DescribeShareProjectOKBody
}

// IsSuccess returns true when this describe share project o k response has a 2xx status code
func (o *DescribeShareProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe share project o k response has a 3xx status code
func (o *DescribeShareProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe share project o k response has a 4xx status code
func (o *DescribeShareProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe share project o k response has a 5xx status code
func (o *DescribeShareProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe share project o k response a status code equal to that given
func (o *DescribeShareProjectOK) IsCode(code int) bool {
	return code == 200
}

func (o *DescribeShareProjectOK) Error() string {
	return fmt.Sprintf("[GET /user/projects/{project_id}/describeSharProject][%d] describeShareProjectOK  %+v", 200, o.Payload)
}

func (o *DescribeShareProjectOK) String() string {
	return fmt.Sprintf("[GET /user/projects/{project_id}/describeSharProject][%d] describeShareProjectOK  %+v", 200, o.Payload)
}

func (o *DescribeShareProjectOK) GetPayload() *DescribeShareProjectOKBody {
	return o.Payload
}

func (o *DescribeShareProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DescribeShareProjectOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeShareProjectDefault creates a DescribeShareProjectDefault with default headers values
func NewDescribeShareProjectDefault(code int) *DescribeShareProjectDefault {
	return &DescribeShareProjectDefault{
		_statusCode: code,
	}
}

/* DescribeShareProjectDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type DescribeShareProjectDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DescribeShareProjectDefaultBody
}

// Code gets the status code for the describe share project default response
func (o *DescribeShareProjectDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this describe share project default response has a 2xx status code
func (o *DescribeShareProjectDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe share project default response has a 3xx status code
func (o *DescribeShareProjectDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe share project default response has a 4xx status code
func (o *DescribeShareProjectDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe share project default response has a 5xx status code
func (o *DescribeShareProjectDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe share project default response a status code equal to that given
func (o *DescribeShareProjectDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DescribeShareProjectDefault) Error() string {
	return fmt.Sprintf("[GET /user/projects/{project_id}/describeSharProject][%d] describeShareProject default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeShareProjectDefault) String() string {
	return fmt.Sprintf("[GET /user/projects/{project_id}/describeSharProject][%d] describeShareProject default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeShareProjectDefault) GetPayload() *DescribeShareProjectDefaultBody {
	return o.Payload
}

func (o *DescribeShareProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DescribeShareProjectDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DescribeShareProjectDefaultBody describe share project default body
swagger:model DescribeShareProjectDefaultBody
*/
type DescribeShareProjectDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this describe share project default body
func (o *DescribeShareProjectDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *DescribeShareProjectDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("describeShareProject default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *DescribeShareProjectDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("describeShareProject default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeShareProject default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeShareProject default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe share project default body based on the context it is used
func (o *DescribeShareProjectDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeShareProjectDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeShareProject default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeShareProject default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DescribeShareProjectDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DescribeShareProjectDefaultBody) UnmarshalBinary(b []byte) error {
	var res DescribeShareProjectDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DescribeShareProjectOKBody describe share project o k body
swagger:model DescribeShareProjectOKBody
*/
type DescribeShareProjectOKBody struct {

	// result
	Result *models.ShareProjectInfo `json:"result,omitempty"`
}

// Validate validates this describe share project o k body
func (o *DescribeShareProjectOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeShareProjectOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeShareProjectOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeShareProjectOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe share project o k body based on the context it is used
func (o *DescribeShareProjectOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeShareProjectOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeShareProjectOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeShareProjectOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DescribeShareProjectOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DescribeShareProjectOKBody) UnmarshalBinary(b []byte) error {
	var res DescribeShareProjectOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
