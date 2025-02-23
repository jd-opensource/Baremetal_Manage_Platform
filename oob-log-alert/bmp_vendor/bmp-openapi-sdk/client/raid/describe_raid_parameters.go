// Code generated by go-swagger; DO NOT EDIT.

package raid

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

// NewDescribeRaidParams creates a new DescribeRaidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeRaidParams() *DescribeRaidParams {
	return &DescribeRaidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeRaidParamsWithTimeout creates a new DescribeRaidParams object
// with the ability to set a timeout on a request.
func NewDescribeRaidParamsWithTimeout(timeout time.Duration) *DescribeRaidParams {
	return &DescribeRaidParams{
		timeout: timeout,
	}
}

// NewDescribeRaidParamsWithContext creates a new DescribeRaidParams object
// with the ability to set a context for a request.
func NewDescribeRaidParamsWithContext(ctx context.Context) *DescribeRaidParams {
	return &DescribeRaidParams{
		Context: ctx,
	}
}

// NewDescribeRaidParamsWithHTTPClient creates a new DescribeRaidParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeRaidParamsWithHTTPClient(client *http.Client) *DescribeRaidParams {
	return &DescribeRaidParams{
		HTTPClient: client,
	}
}

/* DescribeRaidParams contains all the parameters to send to the API endpoint
   for the describe raid operation.

   Typically these are written to a http.Request.
*/
type DescribeRaidParams struct {

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

	// RaidID.
	RaidID string

	/* TraceID.

	   流量唯一id
	*/
	TraceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe raid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeRaidParams) WithDefaults() *DescribeRaidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe raid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeRaidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe raid params
func (o *DescribeRaidParams) WithTimeout(timeout time.Duration) *DescribeRaidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe raid params
func (o *DescribeRaidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe raid params
func (o *DescribeRaidParams) WithContext(ctx context.Context) *DescribeRaidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe raid params
func (o *DescribeRaidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe raid params
func (o *DescribeRaidParams) WithHTTPClient(client *http.Client) *DescribeRaidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe raid params
func (o *DescribeRaidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the describe raid params
func (o *DescribeRaidParams) WithAuthorization(authorization string) *DescribeRaidParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the describe raid params
func (o *DescribeRaidParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the describe raid params
func (o *DescribeRaidParams) WithBmpLanguage(bmpLanguage *string) *DescribeRaidParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the describe raid params
func (o *DescribeRaidParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the describe raid params
func (o *DescribeRaidParams) WithBmpUserID(bmpUserID *string) *DescribeRaidParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the describe raid params
func (o *DescribeRaidParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithRaidID adds the raidID to the describe raid params
func (o *DescribeRaidParams) WithRaidID(raidID string) *DescribeRaidParams {
	o.SetRaidID(raidID)
	return o
}

// SetRaidID adds the raidId to the describe raid params
func (o *DescribeRaidParams) SetRaidID(raidID string) {
	o.RaidID = raidID
}

// WithTraceID adds the traceID to the describe raid params
func (o *DescribeRaidParams) WithTraceID(traceID string) *DescribeRaidParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the describe raid params
func (o *DescribeRaidParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeRaidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param raid_id
	if err := r.SetPathParam("raid_id", o.RaidID); err != nil {
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
