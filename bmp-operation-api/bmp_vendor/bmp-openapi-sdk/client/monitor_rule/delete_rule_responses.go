// Code generated by go-swagger; DO NOT EDIT.

package monitor_rule

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

// DeleteRuleReader is a Reader for the DeleteRule structure.
type DeleteRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRuleOK creates a DeleteRuleOK with default headers values
func NewDeleteRuleOK() *DeleteRuleOK {
	return &DeleteRuleOK{}
}

/* DeleteRuleOK describes a response with status code 200, with default header values.

A DeleteRuleResult is an response struct that is used to describe create rule result.
*/
type DeleteRuleOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DeleteRuleOKBody
}

// IsSuccess returns true when this delete rule o k response has a 2xx status code
func (o *DeleteRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete rule o k response has a 3xx status code
func (o *DeleteRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rule o k response has a 4xx status code
func (o *DeleteRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete rule o k response has a 5xx status code
func (o *DeleteRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rule o k response a status code equal to that given
func (o *DeleteRuleOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteRuleOK) Error() string {
	return fmt.Sprintf("[DELETE /monitorRule/deleteRule][%d] deleteRuleOK  %+v", 200, o.Payload)
}

func (o *DeleteRuleOK) String() string {
	return fmt.Sprintf("[DELETE /monitorRule/deleteRule][%d] deleteRuleOK  %+v", 200, o.Payload)
}

func (o *DeleteRuleOK) GetPayload() *DeleteRuleOKBody {
	return o.Payload
}

func (o *DeleteRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DeleteRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRuleDefault creates a DeleteRuleDefault with default headers values
func NewDeleteRuleDefault(code int) *DeleteRuleDefault {
	return &DeleteRuleDefault{
		_statusCode: code,
	}
}

/* DeleteRuleDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type DeleteRuleDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *DeleteRuleDefaultBody
}

// Code gets the status code for the delete rule default response
func (o *DeleteRuleDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete rule default response has a 2xx status code
func (o *DeleteRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete rule default response has a 3xx status code
func (o *DeleteRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete rule default response has a 4xx status code
func (o *DeleteRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete rule default response has a 5xx status code
func (o *DeleteRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete rule default response a status code equal to that given
func (o *DeleteRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteRuleDefault) Error() string {
	return fmt.Sprintf("[DELETE /monitorRule/deleteRule][%d] deleteRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRuleDefault) String() string {
	return fmt.Sprintf("[DELETE /monitorRule/deleteRule][%d] deleteRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRuleDefault) GetPayload() *DeleteRuleDefaultBody {
	return o.Payload
}

func (o *DeleteRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(DeleteRuleDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteRuleDefaultBody delete rule default body
swagger:model DeleteRuleDefaultBody
*/
type DeleteRuleDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this delete rule default body
func (o *DeleteRuleDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *DeleteRuleDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("deleteRule default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *DeleteRuleDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("deleteRule default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteRule default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteRule default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this delete rule default body based on the context it is used
func (o *DeleteRuleDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteRuleDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteRule default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteRule default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteRuleDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteRuleDefaultBody) UnmarshalBinary(b []byte) error {
	var res DeleteRuleDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteRuleOKBody delete rule o k body
swagger:model DeleteRuleOKBody
*/
type DeleteRuleOKBody struct {

	// result
	Result *models.CommonResponse `json:"result,omitempty"`
}

// Validate validates this delete rule o k body
func (o *DeleteRuleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteRuleOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteRuleOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteRuleOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this delete rule o k body based on the context it is used
func (o *DeleteRuleOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteRuleOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteRuleOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteRuleOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteRuleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteRuleOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteRuleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}