// Code generated by go-swagger; DO NOT EDIT.

package role

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

// NewCurrentRoleParams creates a new CurrentRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCurrentRoleParams() *CurrentRoleParams {
	return &CurrentRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCurrentRoleParamsWithTimeout creates a new CurrentRoleParams object
// with the ability to set a timeout on a request.
func NewCurrentRoleParamsWithTimeout(timeout time.Duration) *CurrentRoleParams {
	return &CurrentRoleParams{
		timeout: timeout,
	}
}

// NewCurrentRoleParamsWithContext creates a new CurrentRoleParams object
// with the ability to set a context for a request.
func NewCurrentRoleParamsWithContext(ctx context.Context) *CurrentRoleParams {
	return &CurrentRoleParams{
		Context: ctx,
	}
}

// NewCurrentRoleParamsWithHTTPClient creates a new CurrentRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCurrentRoleParamsWithHTTPClient(client *http.Client) *CurrentRoleParams {
	return &CurrentRoleParams{
		HTTPClient: client,
	}
}

/* CurrentRoleParams contains all the parameters to send to the API endpoint
   for the current role operation.

   Typically these are written to a http.Request.
*/
type CurrentRoleParams struct {

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

// WithDefaults hydrates default values in the current role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CurrentRoleParams) WithDefaults() *CurrentRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the current role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CurrentRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the current role params
func (o *CurrentRoleParams) WithTimeout(timeout time.Duration) *CurrentRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the current role params
func (o *CurrentRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the current role params
func (o *CurrentRoleParams) WithContext(ctx context.Context) *CurrentRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the current role params
func (o *CurrentRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the current role params
func (o *CurrentRoleParams) WithHTTPClient(client *http.Client) *CurrentRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the current role params
func (o *CurrentRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the current role params
func (o *CurrentRoleParams) WithAuthorization(authorization string) *CurrentRoleParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the current role params
func (o *CurrentRoleParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the current role params
func (o *CurrentRoleParams) WithBmpLanguage(bmpLanguage *string) *CurrentRoleParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the current role params
func (o *CurrentRoleParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the current role params
func (o *CurrentRoleParams) WithBmpUserID(bmpUserID *string) *CurrentRoleParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the current role params
func (o *CurrentRoleParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithTraceID adds the traceID to the current role params
func (o *CurrentRoleParams) WithTraceID(traceID string) *CurrentRoleParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the current role params
func (o *CurrentRoleParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *CurrentRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
