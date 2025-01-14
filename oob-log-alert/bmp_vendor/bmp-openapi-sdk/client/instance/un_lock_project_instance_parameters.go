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

// NewUnLockProjectInstanceParams creates a new UnLockProjectInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnLockProjectInstanceParams() *UnLockProjectInstanceParams {
	return &UnLockProjectInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnLockProjectInstanceParamsWithTimeout creates a new UnLockProjectInstanceParams object
// with the ability to set a timeout on a request.
func NewUnLockProjectInstanceParamsWithTimeout(timeout time.Duration) *UnLockProjectInstanceParams {
	return &UnLockProjectInstanceParams{
		timeout: timeout,
	}
}

// NewUnLockProjectInstanceParamsWithContext creates a new UnLockProjectInstanceParams object
// with the ability to set a context for a request.
func NewUnLockProjectInstanceParamsWithContext(ctx context.Context) *UnLockProjectInstanceParams {
	return &UnLockProjectInstanceParams{
		Context: ctx,
	}
}

// NewUnLockProjectInstanceParamsWithHTTPClient creates a new UnLockProjectInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnLockProjectInstanceParamsWithHTTPClient(client *http.Client) *UnLockProjectInstanceParams {
	return &UnLockProjectInstanceParams{
		HTTPClient: client,
	}
}

/* UnLockProjectInstanceParams contains all the parameters to send to the API endpoint
   for the un lock project instance operation.

   Typically these are written to a http.Request.
*/
type UnLockProjectInstanceParams struct {

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

// WithDefaults hydrates default values in the un lock project instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnLockProjectInstanceParams) WithDefaults() *UnLockProjectInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the un lock project instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnLockProjectInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the un lock project instance params
func (o *UnLockProjectInstanceParams) WithTimeout(timeout time.Duration) *UnLockProjectInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the un lock project instance params
func (o *UnLockProjectInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the un lock project instance params
func (o *UnLockProjectInstanceParams) WithContext(ctx context.Context) *UnLockProjectInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the un lock project instance params
func (o *UnLockProjectInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the un lock project instance params
func (o *UnLockProjectInstanceParams) WithHTTPClient(client *http.Client) *UnLockProjectInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the un lock project instance params
func (o *UnLockProjectInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the un lock project instance params
func (o *UnLockProjectInstanceParams) WithAuthorization(authorization string) *UnLockProjectInstanceParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the un lock project instance params
func (o *UnLockProjectInstanceParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the un lock project instance params
func (o *UnLockProjectInstanceParams) WithBmpLanguage(bmpLanguage *string) *UnLockProjectInstanceParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the un lock project instance params
func (o *UnLockProjectInstanceParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the un lock project instance params
func (o *UnLockProjectInstanceParams) WithBmpUserID(bmpUserID *string) *UnLockProjectInstanceParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the un lock project instance params
func (o *UnLockProjectInstanceParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithInstanceID adds the instanceID to the un lock project instance params
func (o *UnLockProjectInstanceParams) WithInstanceID(instanceID string) *UnLockProjectInstanceParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the un lock project instance params
func (o *UnLockProjectInstanceParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithTraceID adds the traceID to the un lock project instance params
func (o *UnLockProjectInstanceParams) WithTraceID(traceID string) *UnLockProjectInstanceParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the un lock project instance params
func (o *UnLockProjectInstanceParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *UnLockProjectInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
