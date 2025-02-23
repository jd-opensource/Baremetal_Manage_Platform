// Code generated by go-swagger; DO NOT EDIT.

package message

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

// GetMessageStatisticReader is a Reader for the GetMessageStatistic structure.
type GetMessageStatisticReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMessageStatisticReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMessageStatisticOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMessageStatisticDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMessageStatisticOK creates a GetMessageStatisticOK with default headers values
func NewGetMessageStatisticOK() *GetMessageStatisticOK {
	return &GetMessageStatisticOK{}
}

/* GetMessageStatisticOK describes a response with status code 200, with default header values.

GetMessageStatisticOK get message statistic o k
*/
type GetMessageStatisticOK struct {

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *GetMessageStatisticOKBody
}

// IsSuccess returns true when this get message statistic o k response has a 2xx status code
func (o *GetMessageStatisticOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get message statistic o k response has a 3xx status code
func (o *GetMessageStatisticOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get message statistic o k response has a 4xx status code
func (o *GetMessageStatisticOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get message statistic o k response has a 5xx status code
func (o *GetMessageStatisticOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get message statistic o k response a status code equal to that given
func (o *GetMessageStatisticOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetMessageStatisticOK) Error() string {
	return fmt.Sprintf("[GET /messages/statistic][%d] getMessageStatisticOK  %+v", 200, o.Payload)
}

func (o *GetMessageStatisticOK) String() string {
	return fmt.Sprintf("[GET /messages/statistic][%d] getMessageStatisticOK  %+v", 200, o.Payload)
}

func (o *GetMessageStatisticOK) GetPayload() *GetMessageStatisticOKBody {
	return o.Payload
}

func (o *GetMessageStatisticOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(GetMessageStatisticOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessageStatisticDefault creates a GetMessageStatisticDefault with default headers values
func NewGetMessageStatisticDefault(code int) *GetMessageStatisticDefault {
	return &GetMessageStatisticDefault{
		_statusCode: code,
	}
}

/* GetMessageStatisticDefault describes a response with status code -1, with default header values.

unit err response for non 200
*/
type GetMessageStatisticDefault struct {
	_statusCode int

	/* 流量ID
	in: header
	*/
	TraceID string

	Payload *GetMessageStatisticDefaultBody
}

// Code gets the status code for the get message statistic default response
func (o *GetMessageStatisticDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get message statistic default response has a 2xx status code
func (o *GetMessageStatisticDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get message statistic default response has a 3xx status code
func (o *GetMessageStatisticDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get message statistic default response has a 4xx status code
func (o *GetMessageStatisticDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get message statistic default response has a 5xx status code
func (o *GetMessageStatisticDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get message statistic default response a status code equal to that given
func (o *GetMessageStatisticDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetMessageStatisticDefault) Error() string {
	return fmt.Sprintf("[GET /messages/statistic][%d] getMessageStatistic default  %+v", o._statusCode, o.Payload)
}

func (o *GetMessageStatisticDefault) String() string {
	return fmt.Sprintf("[GET /messages/statistic][%d] getMessageStatistic default  %+v", o._statusCode, o.Payload)
}

func (o *GetMessageStatisticDefault) GetPayload() *GetMessageStatisticDefaultBody {
	return o.Payload
}

func (o *GetMessageStatisticDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header traceId
	hdrTraceID := response.GetHeader("traceId")

	if hdrTraceID != "" {
		o.TraceID = hdrTraceID
	}

	o.Payload = new(GetMessageStatisticDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetMessageStatisticDefaultBody get message statistic default body
swagger:model GetMessageStatisticDefaultBody
*/
type GetMessageStatisticDefaultBody struct {

	// 流量ID
	// Required: true
	RequestID *string `json:"requestId"`

	// error
	// Required: true
	Error *models.ErrorResponse `json:"error"`
}

// Validate validates this get message statistic default body
func (o *GetMessageStatisticDefaultBody) Validate(formats strfmt.Registry) error {
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

func (o *GetMessageStatisticDefaultBody) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("getMessageStatistic default"+"."+"requestId", "body", o.RequestID); err != nil {
		return err
	}

	return nil
}

func (o *GetMessageStatisticDefaultBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getMessageStatistic default"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getMessageStatistic default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getMessageStatistic default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get message statistic default body based on the context it is used
func (o *GetMessageStatisticDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMessageStatisticDefaultBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getMessageStatistic default" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getMessageStatistic default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetMessageStatisticDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMessageStatisticDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetMessageStatisticDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetMessageStatisticOKBody get message statistic o k body
swagger:model GetMessageStatisticOKBody
*/
type GetMessageStatisticOKBody struct {

	// result
	Result *models.MessageStatistic `json:"result,omitempty"`
}

// Validate validates this get message statistic o k body
func (o *GetMessageStatisticOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMessageStatisticOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getMessageStatisticOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getMessageStatisticOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get message statistic o k body based on the context it is used
func (o *GetMessageStatisticOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMessageStatisticOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getMessageStatisticOK" + "." + "result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getMessageStatisticOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetMessageStatisticOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMessageStatisticOKBody) UnmarshalBinary(b []byte) error {
	var res GetMessageStatisticOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
