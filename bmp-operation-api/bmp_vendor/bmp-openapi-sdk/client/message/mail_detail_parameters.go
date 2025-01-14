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

// NewMailDetailParams creates a new MailDetailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMailDetailParams() *MailDetailParams {
	return &MailDetailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMailDetailParamsWithTimeout creates a new MailDetailParams object
// with the ability to set a timeout on a request.
func NewMailDetailParamsWithTimeout(timeout time.Duration) *MailDetailParams {
	return &MailDetailParams{
		timeout: timeout,
	}
}

// NewMailDetailParamsWithContext creates a new MailDetailParams object
// with the ability to set a context for a request.
func NewMailDetailParamsWithContext(ctx context.Context) *MailDetailParams {
	return &MailDetailParams{
		Context: ctx,
	}
}

// NewMailDetailParamsWithHTTPClient creates a new MailDetailParams object
// with the ability to set a custom HTTPClient for a request.
func NewMailDetailParamsWithHTTPClient(client *http.Client) *MailDetailParams {
	return &MailDetailParams{
		HTTPClient: client,
	}
}

/* MailDetailParams contains all the parameters to send to the API endpoint
   for the mail detail operation.

   Typically these are written to a http.Request.
*/
type MailDetailParams struct {

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

	/* TraceID.

	   流量唯一id
	*/
	TraceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the mail detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MailDetailParams) WithDefaults() *MailDetailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the mail detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MailDetailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the mail detail params
func (o *MailDetailParams) WithTimeout(timeout time.Duration) *MailDetailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mail detail params
func (o *MailDetailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mail detail params
func (o *MailDetailParams) WithContext(ctx context.Context) *MailDetailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mail detail params
func (o *MailDetailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mail detail params
func (o *MailDetailParams) WithHTTPClient(client *http.Client) *MailDetailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mail detail params
func (o *MailDetailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the mail detail params
func (o *MailDetailParams) WithAuthorization(authorization string) *MailDetailParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the mail detail params
func (o *MailDetailParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the mail detail params
func (o *MailDetailParams) WithBmpLanguage(bmpLanguage *string) *MailDetailParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the mail detail params
func (o *MailDetailParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the mail detail params
func (o *MailDetailParams) WithBmpUserID(bmpUserID *string) *MailDetailParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the mail detail params
func (o *MailDetailParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithTraceID adds the traceID to the mail detail params
func (o *MailDetailParams) WithTraceID(traceID string) *MailDetailParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the mail detail params
func (o *MailDetailParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *MailDetailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// header param traceId
	if err := r.SetHeaderParam("traceId", o.TraceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
