// Code generated by go-swagger; DO NOT EDIT.

package instance

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

// CreateProjectInstanceReader is a Reader for the CreateProjectInstance structure.
type CreateProjectInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProjectInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateProjectInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateProjectInstanceOK creates a CreateProjectInstanceOK with default headers values
func NewCreateProjectInstanceOK() *CreateProjectInstanceOK {
	return &CreateProjectInstanceOK{}
}

/* CreateProjectInstanceOK describes a response with status code 200, with default header values.

A CreateProjectInstanceResult is an response struct that is used to describe create instance result.
*/
type CreateProjectInstanceOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *CreateProjectInstanceOKBody
}

// IsSuccess returns true when this create project instance o k response has a 2xx status code
func (o *CreateProjectInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create project instance o k response has a 3xx status code
func (o *CreateProjectInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create project instance o k response has a 4xx status code
func (o *CreateProjectInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create project instance o k response has a 5xx status code
func (o *CreateProjectInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create project instance o k response a status code equal to that given
func (o *CreateProjectInstanceOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateProjectInstanceOK) Error() string {
	return fmt.Sprintf("[POST /project/instances][%d] createProjectInstanceOK  %+v", 200, o.Payload)
}

func (o *CreateProjectInstanceOK) String() string {
	return fmt.Sprintf("[POST /project/instances][%d] createProjectInstanceOK  %+v", 200, o.Payload)
}

func (o *CreateProjectInstanceOK) GetPayload() *CreateProjectInstanceOKBody {
	return o.Payload
}

func (o *CreateProjectInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(CreateProjectInstanceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectInstanceDefault creates a CreateProjectInstanceDefault with default headers values
func NewCreateProjectInstanceDefault(code int) *CreateProjectInstanceDefault {
	return &CreateProjectInstanceDefault{
		_statusCode: code,
	}
}

/* CreateProjectInstanceDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type CreateProjectInstanceDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *CreateProjectInstanceDefaultBody
}

// Code gets the status code for the create project instance default response
func (o *CreateProjectInstanceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create project instance default response has a 2xx status code
func (o *CreateProjectInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create project instance default response has a 3xx status code
func (o *CreateProjectInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create project instance default response has a 4xx status code
func (o *CreateProjectInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create project instance default response has a 5xx status code
func (o *CreateProjectInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create project instance default response a status code equal to that given
func (o *CreateProjectInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateProjectInstanceDefault) Error() string {
	return fmt.Sprintf("[POST /project/instances][%d] createProjectInstance default  %+v", o._statusCode, o.Payload)
}

func (o *CreateProjectInstanceDefault) String() string {
	return fmt.Sprintf("[POST /project/instances][%d] createProjectInstance default  %+v", o._statusCode, o.Payload)
}

func (o *CreateProjectInstanceDefault) GetPayload() *CreateProjectInstanceDefaultBody {
	return o.Payload
}

func (o *CreateProjectInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(CreateProjectInstanceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateProjectInstanceDefaultBody create project instance default body
swagger:model CreateProjectInstanceDefaultBody
*/
type CreateProjectInstanceDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this create project instance default body
func (o *CreateProjectInstanceDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *CreateProjectInstanceDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("createProjectInstance default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *CreateProjectInstanceDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("createProjectInstance default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createProjectInstance default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createProjectInstance default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create project instance default body based on the context it is used
func (o *CreateProjectInstanceDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateProjectInstanceDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createProjectInstance default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createProjectInstance default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateProjectInstanceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateProjectInstanceDefaultBody) UnmarshalBinary(b []byte) error {
	var res CreateProjectInstanceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateProjectInstanceOKBody create project instance o k body
swagger:model CreateProjectInstanceOKBody
*/
type CreateProjectInstanceOKBody struct {

	// result
	Result *models.InstanceIds `json:"result,omitempty"`
}

// Validate validates this create project instance o k body
func (o *CreateProjectInstanceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateProjectInstanceOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createProjectInstanceOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createProjectInstanceOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create project instance o k body based on the context it is used
func (o *CreateProjectInstanceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateProjectInstanceOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createProjectInstanceOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createProjectInstanceOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateProjectInstanceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateProjectInstanceOKBody) UnmarshalBinary(b []byte) error {
	var res CreateProjectInstanceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
