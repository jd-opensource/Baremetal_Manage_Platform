// Code generated by go-swagger; DO NOT EDIT.

package message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetMessageByIDParams creates a new GetMessageByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMessageByIDParams() *GetMessageByIDParams {
	return &GetMessageByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMessageByIDParamsWithTimeout creates a new GetMessageByIDParams object
// with the ability to set a timeout on a request.
func NewGetMessageByIDParamsWithTimeout(timeout time.Duration) *GetMessageByIDParams {
	return &GetMessageByIDParams{
		timeout: timeout,
	}
}

// NewGetMessageByIDParamsWithContext creates a new GetMessageByIDParams object
// with the ability to set a context for a request.
func NewGetMessageByIDParamsWithContext(ctx context.Context) *GetMessageByIDParams {
	return &GetMessageByIDParams{
		Context: ctx,
	}
}

// NewGetMessageByIDParamsWithHTTPClient creates a new GetMessageByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMessageByIDParamsWithHTTPClient(client *http.Client) *GetMessageByIDParams {
	return &GetMessageByIDParams{
		HTTPClient: client,
	}
}

/* GetMessageByIDParams contains all the parameters to send to the API endpoint
   for the get message by Id operation.

   Typically these are written to a http.Request.
*/
type GetMessageByIDParams struct {

	/* Authorization.

	   demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token
	*/
	Authorization string

	/* BmpLanguage.

	   用户语言 [zh_CN, en_US]
	*/
	BmpLanguage *string

	/* BmpUserID.

	   用户uuid, deprecated
	*/
	BmpUserID *string

	// MessageID.
	MessageID string

	/* TraceID.

	   流量唯一id
	*/
	TraceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get message by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMessageByIDParams) WithDefaults() *GetMessageByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get message by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMessageByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get message by Id params
func (o *GetMessageByIDParams) WithTimeout(timeout time.Duration) *GetMessageByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get message by Id params
func (o *GetMessageByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get message by Id params
func (o *GetMessageByIDParams) WithContext(ctx context.Context) *GetMessageByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get message by Id params
func (o *GetMessageByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get message by Id params
func (o *GetMessageByIDParams) WithHTTPClient(client *http.Client) *GetMessageByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get message by Id params
func (o *GetMessageByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get message by Id params
func (o *GetMessageByIDParams) WithAuthorization(authorization string) *GetMessageByIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get message by Id params
func (o *GetMessageByIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the get message by Id params
func (o *GetMessageByIDParams) WithBmpLanguage(bmpLanguage *string) *GetMessageByIDParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the get message by Id params
func (o *GetMessageByIDParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the get message by Id params
func (o *GetMessageByIDParams) WithBmpUserID(bmpUserID *string) *GetMessageByIDParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the get message by Id params
func (o *GetMessageByIDParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithMessageID adds the messageID to the get message by Id params
func (o *GetMessageByIDParams) WithMessageID(messageID string) *GetMessageByIDParams {
	o.SetMessageID(messageID)
	return o
}

// SetMessageID adds the messageId to the get message by Id params
func (o *GetMessageByIDParams) SetMessageID(messageID string) {
	o.MessageID = messageID
}

// WithTraceID adds the traceID to the get message by Id params
func (o *GetMessageByIDParams) WithTraceID(traceID string) *GetMessageByIDParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the get message by Id params
func (o *GetMessageByIDParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMessageByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param authorization
	if err := r.SetHeaderParam("authorization", o.Authorization); err != nil {
		return err
	}

	if o.BmpLanguage != nil {

		// header param bmpLanguage
		if err := r.SetHeaderParam("bmpLanguage", *o.BmpLanguage); err != nil {
			return err
		}
	}

	if o.BmpUserID != nil {

		// header param bmpUserId
		if err := r.SetHeaderParam("bmpUserId", *o.BmpUserID); err != nil {
			return err
		}
	}

	// query param messageId
	qrMessageID := o.MessageID
	qMessageID := qrMessageID
	if qMessageID != "" {

		if err := r.SetQueryParam("messageId", qMessageID); err != nil {
			return err
		}
	}

	// header param traceId
	if err := r.SetHeaderParam("traceId", o.TraceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}