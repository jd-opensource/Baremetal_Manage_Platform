// Code generated by go-swagger; DO NOT EDIT.

package os

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

// NewDeleteOSParams creates a new DeleteOSParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOSParams() *DeleteOSParams {
	return &DeleteOSParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOSParamsWithTimeout creates a new DeleteOSParams object
// with the ability to set a timeout on a request.
func NewDeleteOSParamsWithTimeout(timeout time.Duration) *DeleteOSParams {
	return &DeleteOSParams{
		timeout: timeout,
	}
}

// NewDeleteOSParamsWithContext creates a new DeleteOSParams object
// with the ability to set a context for a request.
func NewDeleteOSParamsWithContext(ctx context.Context) *DeleteOSParams {
	return &DeleteOSParams{
		Context: ctx,
	}
}

// NewDeleteOSParamsWithHTTPClient creates a new DeleteOSParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOSParamsWithHTTPClient(client *http.Client) *DeleteOSParams {
	return &DeleteOSParams{
		HTTPClient: client,
	}
}

/* DeleteOSParams contains all the parameters to send to the API endpoint
   for the delete o s operation.

   Typically these are written to a http.Request.
*/
type DeleteOSParams struct {

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

	// OsID.
	OsID string

	/* TraceID.

	   流量唯一id
	*/
	TraceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete o s params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOSParams) WithDefaults() *DeleteOSParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete o s params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOSParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete o s params
func (o *DeleteOSParams) WithTimeout(timeout time.Duration) *DeleteOSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete o s params
func (o *DeleteOSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete o s params
func (o *DeleteOSParams) WithContext(ctx context.Context) *DeleteOSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete o s params
func (o *DeleteOSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete o s params
func (o *DeleteOSParams) WithHTTPClient(client *http.Client) *DeleteOSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete o s params
func (o *DeleteOSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete o s params
func (o *DeleteOSParams) WithAuthorization(authorization string) *DeleteOSParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete o s params
func (o *DeleteOSParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the delete o s params
func (o *DeleteOSParams) WithBmpLanguage(bmpLanguage *string) *DeleteOSParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the delete o s params
func (o *DeleteOSParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the delete o s params
func (o *DeleteOSParams) WithBmpUserID(bmpUserID *string) *DeleteOSParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the delete o s params
func (o *DeleteOSParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithOsID adds the osID to the delete o s params
func (o *DeleteOSParams) WithOsID(osID string) *DeleteOSParams {
	o.SetOsID(osID)
	return o
}

// SetOsID adds the osId to the delete o s params
func (o *DeleteOSParams) SetOsID(osID string) {
	o.OsID = osID
}

// WithTraceID adds the traceID to the delete o s params
func (o *DeleteOSParams) WithTraceID(traceID string) *DeleteOSParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the delete o s params
func (o *DeleteOSParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param os_id
	if err := r.SetPathParam("os_id", o.OsID); err != nil {
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