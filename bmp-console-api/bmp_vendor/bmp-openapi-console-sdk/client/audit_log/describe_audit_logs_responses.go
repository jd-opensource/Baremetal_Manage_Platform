// Code generated by go-swagger; DO NOT EDIT.

package audit_log

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

// DescribeAuditLogsReader is a Reader for the DescribeAuditLogs structure.
type DescribeAuditLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeAuditLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeAuditLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeAuditLogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeAuditLogsOK creates a DescribeAuditLogsOK with default headers values
func NewDescribeAuditLogsOK() *DescribeAuditLogsOK {
	return &DescribeAuditLogsOK{}
}

/* DescribeAuditLogsOK describes a response with status code 200, with default header values.

DescribeAuditLogsOK describe audit logs o k
*/
type DescribeAuditLogsOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DescribeAuditLogsOKBody
}

// IsSuccess returns true when this describe audit logs o k response has a 2xx status code
func (o *DescribeAuditLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe audit logs o k response has a 3xx status code
func (o *DescribeAuditLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe audit logs o k response has a 4xx status code
func (o *DescribeAuditLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe audit logs o k response has a 5xx status code
func (o *DescribeAuditLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe audit logs o k response a status code equal to that given
func (o *DescribeAuditLogsOK) IsCode(code int) bool {
	return code == 200
}

func (o *DescribeAuditLogsOK) Error() string {
	return fmt.Sprintf("[GET /auditLogs][%d] describeAuditLogsOK  %+v", 200, o.Payload)
}

func (o *DescribeAuditLogsOK) String() string {
	return fmt.Sprintf("[GET /auditLogs][%d] describeAuditLogsOK  %+v", 200, o.Payload)
}

func (o *DescribeAuditLogsOK) GetPayload() *DescribeAuditLogsOKBody {
	return o.Payload
}

func (o *DescribeAuditLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DescribeAuditLogsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeAuditLogsDefault creates a DescribeAuditLogsDefault with default headers values
func NewDescribeAuditLogsDefault(code int) *DescribeAuditLogsDefault {
	return &DescribeAuditLogsDefault{
		_statusCode: code,
	}
}

/* DescribeAuditLogsDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type DescribeAuditLogsDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DescribeAuditLogsDefaultBody
}

// Code gets the status code for the describe audit logs default response
func (o *DescribeAuditLogsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this describe audit logs default response has a 2xx status code
func (o *DescribeAuditLogsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe audit logs default response has a 3xx status code
func (o *DescribeAuditLogsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe audit logs default response has a 4xx status code
func (o *DescribeAuditLogsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe audit logs default response has a 5xx status code
func (o *DescribeAuditLogsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe audit logs default response a status code equal to that given
func (o *DescribeAuditLogsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DescribeAuditLogsDefault) Error() string {
	return fmt.Sprintf("[GET /auditLogs][%d] describeAuditLogs default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeAuditLogsDefault) String() string {
	return fmt.Sprintf("[GET /auditLogs][%d] describeAuditLogs default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeAuditLogsDefault) GetPayload() *DescribeAuditLogsDefaultBody {
	return o.Payload
}

func (o *DescribeAuditLogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DescribeAuditLogsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DescribeAuditLogsDefaultBody describe audit logs default body
swagger:model DescribeAuditLogsDefaultBody
*/
type DescribeAuditLogsDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this describe audit logs default body
func (o *DescribeAuditLogsDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *DescribeAuditLogsDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("describeAuditLogs default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *DescribeAuditLogsDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("describeAuditLogs default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeAuditLogs default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeAuditLogs default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe audit logs default body based on the context it is used
func (o *DescribeAuditLogsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeAuditLogsDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeAuditLogs default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeAuditLogs default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DescribeAuditLogsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DescribeAuditLogsDefaultBody) UnmarshalBinary(b []byte) error {
	var res DescribeAuditLogsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DescribeAuditLogsOKBody describe audit logs o k body
swagger:model DescribeAuditLogsOKBody
*/
type DescribeAuditLogsOKBody struct {

	// result
	Result *models.AuditLogList `json:"result,omitempty"`
}

// Validate validates this describe audit logs o k body
func (o *DescribeAuditLogsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeAuditLogsOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeAuditLogsOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeAuditLogsOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe audit logs o k body based on the context it is used
func (o *DescribeAuditLogsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DescribeAuditLogsOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describeAuditLogsOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describeAuditLogsOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DescribeAuditLogsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DescribeAuditLogsOKBody) UnmarshalBinary(b []byte) error {
	var res DescribeAuditLogsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}