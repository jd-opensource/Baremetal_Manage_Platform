// Code generated by go-swagger; DO NOT EDIT.

package instance

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

// NewStartProjectInstanceParams creates a new StartProjectInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartProjectInstanceParams() *StartProjectInstanceParams {
	return &StartProjectInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartProjectInstanceParamsWithTimeout creates a new StartProjectInstanceParams object
// with the ability to set a timeout on a request.
func NewStartProjectInstanceParamsWithTimeout(timeout time.Duration) *StartProjectInstanceParams {
	return &StartProjectInstanceParams{
		timeout: timeout,
	}
}

// NewStartProjectInstanceParamsWithContext creates a new StartProjectInstanceParams object
// with the ability to set a context for a request.
func NewStartProjectInstanceParamsWithContext(ctx context.Context) *StartProjectInstanceParams {
	return &StartProjectInstanceParams{
		Context: ctx,
	}
}

// NewStartProjectInstanceParamsWithHTTPClient creates a new StartProjectInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartProjectInstanceParamsWithHTTPClient(client *http.Client) *StartProjectInstanceParams {
	return &StartProjectInstanceParams{
		HTTPClient: client,
	}
}

/* StartProjectInstanceParams contains all the parameters to send to the API endpoint
   for the start project instance operation.

   Typically these are written to a http.Request.
*/
type StartProjectInstanceParams struct {

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

	// InstanceID.
	InstanceID string

	/* TraceID.

	   流量唯一id
	*/
	TraceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start project instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartProjectInstanceParams) WithDefaults() *StartProjectInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start project instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartProjectInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start project instance params
func (o *StartProjectInstanceParams) WithTimeout(timeout time.Duration) *StartProjectInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start project instance params
func (o *StartProjectInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start project instance params
func (o *StartProjectInstanceParams) WithContext(ctx context.Context) *StartProjectInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start project instance params
func (o *StartProjectInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start project instance params
func (o *StartProjectInstanceParams) WithHTTPClient(client *http.Client) *StartProjectInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start project instance params
func (o *StartProjectInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the start project instance params
func (o *StartProjectInstanceParams) WithAuthorization(authorization string) *StartProjectInstanceParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the start project instance params
func (o *StartProjectInstanceParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the start project instance params
func (o *StartProjectInstanceParams) WithBmpLanguage(bmpLanguage *string) *StartProjectInstanceParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the start project instance params
func (o *StartProjectInstanceParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the start project instance params
func (o *StartProjectInstanceParams) WithBmpUserID(bmpUserID *string) *StartProjectInstanceParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the start project instance params
func (o *StartProjectInstanceParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithInstanceID adds the instanceID to the start project instance params
func (o *StartProjectInstanceParams) WithInstanceID(instanceID string) *StartProjectInstanceParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the start project instance params
func (o *StartProjectInstanceParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithTraceID adds the traceID to the start project instance params
func (o *StartProjectInstanceParams) WithTraceID(traceID string) *StartProjectInstanceParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the start project instance params
func (o *StartProjectInstanceParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *StartProjectInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param instance_id
	if err := r.SetPathParam("instance_id", o.InstanceID); err != nil {
		return err
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