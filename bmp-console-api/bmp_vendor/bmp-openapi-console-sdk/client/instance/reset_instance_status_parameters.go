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

// NewResetInstanceStatusParams creates a new ResetInstanceStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResetInstanceStatusParams() *ResetInstanceStatusParams {
	return &ResetInstanceStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResetInstanceStatusParamsWithTimeout creates a new ResetInstanceStatusParams object
// with the ability to set a timeout on a request.
func NewResetInstanceStatusParamsWithTimeout(timeout time.Duration) *ResetInstanceStatusParams {
	return &ResetInstanceStatusParams{
		timeout: timeout,
	}
}

// NewResetInstanceStatusParamsWithContext creates a new ResetInstanceStatusParams object
// with the ability to set a context for a request.
func NewResetInstanceStatusParamsWithContext(ctx context.Context) *ResetInstanceStatusParams {
	return &ResetInstanceStatusParams{
		Context: ctx,
	}
}

// NewResetInstanceStatusParamsWithHTTPClient creates a new ResetInstanceStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewResetInstanceStatusParamsWithHTTPClient(client *http.Client) *ResetInstanceStatusParams {
	return &ResetInstanceStatusParams{
		HTTPClient: client,
	}
}

/* ResetInstanceStatusParams contains all the parameters to send to the API endpoint
   for the reset instance status operation.

   Typically these are written to a http.Request.
*/
type ResetInstanceStatusParams struct {

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

// WithDefaults hydrates default values in the reset instance status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetInstanceStatusParams) WithDefaults() *ResetInstanceStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reset instance status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetInstanceStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reset instance status params
func (o *ResetInstanceStatusParams) WithTimeout(timeout time.Duration) *ResetInstanceStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset instance status params
func (o *ResetInstanceStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset instance status params
func (o *ResetInstanceStatusParams) WithContext(ctx context.Context) *ResetInstanceStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset instance status params
func (o *ResetInstanceStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset instance status params
func (o *ResetInstanceStatusParams) WithHTTPClient(client *http.Client) *ResetInstanceStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset instance status params
func (o *ResetInstanceStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the reset instance status params
func (o *ResetInstanceStatusParams) WithAuthorization(authorization string) *ResetInstanceStatusParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the reset instance status params
func (o *ResetInstanceStatusParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the reset instance status params
func (o *ResetInstanceStatusParams) WithBmpLanguage(bmpLanguage *string) *ResetInstanceStatusParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the reset instance status params
func (o *ResetInstanceStatusParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the reset instance status params
func (o *ResetInstanceStatusParams) WithBmpUserID(bmpUserID *string) *ResetInstanceStatusParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the reset instance status params
func (o *ResetInstanceStatusParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithInstanceID adds the instanceID to the reset instance status params
func (o *ResetInstanceStatusParams) WithInstanceID(instanceID string) *ResetInstanceStatusParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the reset instance status params
func (o *ResetInstanceStatusParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithTraceID adds the traceID to the reset instance status params
func (o *ResetInstanceStatusParams) WithTraceID(traceID string) *ResetInstanceStatusParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the reset instance status params
func (o *ResetInstanceStatusParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *ResetInstanceStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
