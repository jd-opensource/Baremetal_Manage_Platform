// Code generated by go-swagger; DO NOT EDIT.

package monitor_alert

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

	"coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
)

// NewAddAlertParams creates a new AddAlertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddAlertParams() *AddAlertParams {
	return &AddAlertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddAlertParamsWithTimeout creates a new AddAlertParams object
// with the ability to set a timeout on a request.
func NewAddAlertParamsWithTimeout(timeout time.Duration) *AddAlertParams {
	return &AddAlertParams{
		timeout: timeout,
	}
}

// NewAddAlertParamsWithContext creates a new AddAlertParams object
// with the ability to set a context for a request.
func NewAddAlertParamsWithContext(ctx context.Context) *AddAlertParams {
	return &AddAlertParams{
		Context: ctx,
	}
}

// NewAddAlertParamsWithHTTPClient creates a new AddAlertParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddAlertParamsWithHTTPClient(client *http.Client) *AddAlertParams {
	return &AddAlertParams{
		HTTPClient: client,
	}
}

/* AddAlertParams contains all the parameters to send to the API endpoint
   for the add alert operation.

   Typically these are written to a http.Request.
*/
type AddAlertParams struct {

	// Body.
	Body *models.AddAlertRequest

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

// WithDefaults hydrates default values in the add alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAlertParams) WithDefaults() *AddAlertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAlertParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add alert params
func (o *AddAlertParams) WithTimeout(timeout time.Duration) *AddAlertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add alert params
func (o *AddAlertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add alert params
func (o *AddAlertParams) WithContext(ctx context.Context) *AddAlertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add alert params
func (o *AddAlertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add alert params
func (o *AddAlertParams) WithHTTPClient(client *http.Client) *AddAlertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add alert params
func (o *AddAlertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add alert params
func (o *AddAlertParams) WithBody(body *models.AddAlertRequest) *AddAlertParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add alert params
func (o *AddAlertParams) SetBody(body *models.AddAlertRequest) {
	o.Body = body
}

// WithAuthorization adds the authorization to the add alert params
func (o *AddAlertParams) WithAuthorization(authorization string) *AddAlertParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the add alert params
func (o *AddAlertParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the add alert params
func (o *AddAlertParams) WithBmpLanguage(bmpLanguage *string) *AddAlertParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the add alert params
func (o *AddAlertParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the add alert params
func (o *AddAlertParams) WithBmpUserID(bmpUserID *string) *AddAlertParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the add alert params
func (o *AddAlertParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithTraceID adds the traceID to the add alert params
func (o *AddAlertParams) WithTraceID(traceID string) *AddAlertParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the add alert params
func (o *AddAlertParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *AddAlertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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