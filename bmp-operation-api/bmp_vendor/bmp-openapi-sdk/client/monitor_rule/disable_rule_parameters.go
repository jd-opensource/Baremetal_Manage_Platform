// Code generated by go-swagger; DO NOT EDIT.

package monitor_rule

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

// NewDisableRuleParams creates a new DisableRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDisableRuleParams() *DisableRuleParams {
	return &DisableRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDisableRuleParamsWithTimeout creates a new DisableRuleParams object
// with the ability to set a timeout on a request.
func NewDisableRuleParamsWithTimeout(timeout time.Duration) *DisableRuleParams {
	return &DisableRuleParams{
		timeout: timeout,
	}
}

// NewDisableRuleParamsWithContext creates a new DisableRuleParams object
// with the ability to set a context for a request.
func NewDisableRuleParamsWithContext(ctx context.Context) *DisableRuleParams {
	return &DisableRuleParams{
		Context: ctx,
	}
}

// NewDisableRuleParamsWithHTTPClient creates a new DisableRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewDisableRuleParamsWithHTTPClient(client *http.Client) *DisableRuleParams {
	return &DisableRuleParams{
		HTTPClient: client,
	}
}

/* DisableRuleParams contains all the parameters to send to the API endpoint
   for the disable rule operation.

   Typically these are written to a http.Request.
*/
type DisableRuleParams struct {

	// Body.
	Body *models.DisableRuleRequest

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

// WithDefaults hydrates default values in the disable rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisableRuleParams) WithDefaults() *DisableRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the disable rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisableRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the disable rule params
func (o *DisableRuleParams) WithTimeout(timeout time.Duration) *DisableRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disable rule params
func (o *DisableRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disable rule params
func (o *DisableRuleParams) WithContext(ctx context.Context) *DisableRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disable rule params
func (o *DisableRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disable rule params
func (o *DisableRuleParams) WithHTTPClient(client *http.Client) *DisableRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disable rule params
func (o *DisableRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the disable rule params
func (o *DisableRuleParams) WithBody(body *models.DisableRuleRequest) *DisableRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the disable rule params
func (o *DisableRuleParams) SetBody(body *models.DisableRuleRequest) {
	o.Body = body
}

// WithAuthorization adds the authorization to the disable rule params
func (o *DisableRuleParams) WithAuthorization(authorization string) *DisableRuleParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the disable rule params
func (o *DisableRuleParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the disable rule params
func (o *DisableRuleParams) WithBmpLanguage(bmpLanguage *string) *DisableRuleParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the disable rule params
func (o *DisableRuleParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the disable rule params
func (o *DisableRuleParams) WithBmpUserID(bmpUserID *string) *DisableRuleParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the disable rule params
func (o *DisableRuleParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithTraceID adds the traceID to the disable rule params
func (o *DisableRuleParams) WithTraceID(traceID string) *DisableRuleParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the disable rule params
func (o *DisableRuleParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *DisableRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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