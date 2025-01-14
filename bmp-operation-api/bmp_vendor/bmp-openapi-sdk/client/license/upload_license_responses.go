// Code generated by go-swagger; DO NOT EDIT.

package license

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

// UploadLicenseReader is a Reader for the UploadLicense structure.
type UploadLicenseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadLicenseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadLicenseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUploadLicenseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUploadLicenseOK creates a UploadLicenseOK with default headers values
func NewUploadLicenseOK() *UploadLicenseOK {
	return &UploadLicenseOK{}
}

/* UploadLicenseOK describes a response with status code 200, with default header values.

A UploadLicenseResult is an response struct that is used to describe modify instance result.
*/
type UploadLicenseOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *UploadLicenseOKBody
}

// IsSuccess returns true when this upload license o k response has a 2xx status code
func (o *UploadLicenseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload license o k response has a 3xx status code
func (o *UploadLicenseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload license o k response has a 4xx status code
func (o *UploadLicenseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload license o k response has a 5xx status code
func (o *UploadLicenseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this upload license o k response a status code equal to that given
func (o *UploadLicenseOK) IsCode(code int) bool {
	return code == 200
}

func (o *UploadLicenseOK) Error() string {
	return fmt.Sprintf("[POST /license/upload][%d] uploadLicenseOK  %+v", 200, o.Payload)
}

func (o *UploadLicenseOK) String() string {
	return fmt.Sprintf("[POST /license/upload][%d] uploadLicenseOK  %+v", 200, o.Payload)
}

func (o *UploadLicenseOK) GetPayload() *UploadLicenseOKBody {
	return o.Payload
}

func (o *UploadLicenseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(UploadLicenseOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadLicenseDefault creates a UploadLicenseDefault with default headers values
func NewUploadLicenseDefault(code int) *UploadLicenseDefault {
	return &UploadLicenseDefault{
		_statusCode: code,
	}
}

/* UploadLicenseDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type UploadLicenseDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *UploadLicenseDefaultBody
}

// Code gets the status code for the upload license default response
func (o *UploadLicenseDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this upload license default response has a 2xx status code
func (o *UploadLicenseDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this upload license default response has a 3xx status code
func (o *UploadLicenseDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this upload license default response has a 4xx status code
func (o *UploadLicenseDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this upload license default response has a 5xx status code
func (o *UploadLicenseDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this upload license default response a status code equal to that given
func (o *UploadLicenseDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UploadLicenseDefault) Error() string {
	return fmt.Sprintf("[POST /license/upload][%d] uploadLicense default  %+v", o._statusCode, o.Payload)
}

func (o *UploadLicenseDefault) String() string {
	return fmt.Sprintf("[POST /license/upload][%d] uploadLicense default  %+v", o._statusCode, o.Payload)
}

func (o *UploadLicenseDefault) GetPayload() *UploadLicenseDefaultBody {
	return o.Payload
}

func (o *UploadLicenseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(UploadLicenseDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UploadLicenseDefaultBody upload license default body
swagger:model UploadLicenseDefaultBody
*/
type UploadLicenseDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this upload license default body
func (o *UploadLicenseDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *UploadLicenseDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("uploadLicense default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *UploadLicenseDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("uploadLicense default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadLicense default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadLicense default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this upload license default body based on the context it is used
func (o *UploadLicenseDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UploadLicenseDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadLicense default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadLicense default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UploadLicenseDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UploadLicenseDefaultBody) UnmarshalBinary(b []byte) error {
	var res UploadLicenseDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UploadLicenseOKBody upload license o k body
swagger:model UploadLicenseOKBody
*/
type UploadLicenseOKBody struct {

	// result
	Result *models.CommonResponse `json:"result,omitempty"`
}

// Validate validates this upload license o k body
func (o *UploadLicenseOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UploadLicenseOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadLicenseOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadLicenseOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this upload license o k body based on the context it is used
func (o *UploadLicenseOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UploadLicenseOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadLicenseOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadLicenseOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UploadLicenseOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UploadLicenseOKBody) UnmarshalBinary(b []byte) error {
	var res UploadLicenseOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
