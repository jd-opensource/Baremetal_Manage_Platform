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

	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
)

// NewDialMailParams creates a new DialMailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDialMailParams() *DialMailParams {
	return &DialMailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDialMailParamsWithTimeout creates a new DialMailParams object
// with the ability to set a timeout on a request.
func NewDialMailParamsWithTimeout(timeout time.Duration) *DialMailParams {
	return &DialMailParams{
		timeout: timeout,
	}
}

// NewDialMailParamsWithContext creates a new DialMailParams object
// with the ability to set a context for a request.
func NewDialMailParamsWithContext(ctx context.Context) *DialMailParams {
	return &DialMailParams{
		Context: ctx,
	}
}

// NewDialMailParamsWithHTTPClient creates a new DialMailParams object
// with the ability to set a custom HTTPClient for a request.
func NewDialMailParamsWithHTTPClient(client *http.Client) *DialMailParams {
	return &DialMailParams{
		HTTPClient: client,
	}
}

/* DialMailParams contains all the parameters to send to the API endpoint
   for the dial mail operation.

   Typically these are written to a http.Request.
*/
type DialMailParams struct {

	// Body.
	Body *models.MailDialRequest

	/* Authorization.

	   demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token
	*/
	Authorization string

	/* BmpLanguage.

	   用户语言 [zh_CN, en_US]
	*/
	BmpLanguage *string

	/* BmpUserID.

	   用户uuid
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

// WithDefaults hydrates default values in the dial mail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DialMailParams) WithDefaults() *DialMailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dial mail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DialMailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dial mail params
func (o *DialMailParams) WithTimeout(timeout time.Duration) *DialMailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dial mail params
func (o *DialMailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dial mail params
func (o *DialMailParams) WithContext(ctx context.Context) *DialMailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dial mail params
func (o *DialMailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dial mail params
func (o *DialMailParams) WithHTTPClient(client *http.Client) *DialMailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dial mail params
func (o *DialMailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the dial mail params
func (o *DialMailParams) WithBody(body *models.MailDialRequest) *DialMailParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the dial mail params
func (o *DialMailParams) SetBody(body *models.MailDialRequest) {
	o.Body = body
}

// WithAuthorization adds the authorization to the dial mail params
func (o *DialMailParams) WithAuthorization(authorization string) *DialMailParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the dial mail params
func (o *DialMailParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the dial mail params
func (o *DialMailParams) WithBmpLanguage(bmpLanguage *string) *DialMailParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the dial mail params
func (o *DialMailParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the dial mail params
func (o *DialMailParams) WithBmpUserID(bmpUserID *string) *DialMailParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the dial mail params
func (o *DialMailParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithTraceID adds the traceID to the dial mail params
func (o *DialMailParams) WithTraceID(traceID string) *DialMailParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the dial mail params
func (o *DialMailParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *DialMailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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