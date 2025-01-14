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

	"coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
)

// NewDeleteInstancesParams creates a new DeleteInstancesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteInstancesParams() *DeleteInstancesParams {
	return &DeleteInstancesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstancesParamsWithTimeout creates a new DeleteInstancesParams object
// with the ability to set a timeout on a request.
func NewDeleteInstancesParamsWithTimeout(timeout time.Duration) *DeleteInstancesParams {
	return &DeleteInstancesParams{
		timeout: timeout,
	}
}

// NewDeleteInstancesParamsWithContext creates a new DeleteInstancesParams object
// with the ability to set a context for a request.
func NewDeleteInstancesParamsWithContext(ctx context.Context) *DeleteInstancesParams {
	return &DeleteInstancesParams{
		Context: ctx,
	}
}

// NewDeleteInstancesParamsWithHTTPClient creates a new DeleteInstancesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteInstancesParamsWithHTTPClient(client *http.Client) *DeleteInstancesParams {
	return &DeleteInstancesParams{
		HTTPClient: client,
	}
}

/* DeleteInstancesParams contains all the parameters to send to the API endpoint
   for the delete instances operation.

   Typically these are written to a http.Request.
*/
type DeleteInstancesParams struct {

	// Body.
	Body *models.DeleteInstancesRequest

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

// WithDefaults hydrates default values in the delete instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteInstancesParams) WithDefaults() *DeleteInstancesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteInstancesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete instances params
func (o *DeleteInstancesParams) WithTimeout(timeout time.Duration) *DeleteInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instances params
func (o *DeleteInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instances params
func (o *DeleteInstancesParams) WithContext(ctx context.Context) *DeleteInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instances params
func (o *DeleteInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instances params
func (o *DeleteInstancesParams) WithHTTPClient(client *http.Client) *DeleteInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instances params
func (o *DeleteInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete instances params
func (o *DeleteInstancesParams) WithBody(body *models.DeleteInstancesRequest) *DeleteInstancesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete instances params
func (o *DeleteInstancesParams) SetBody(body *models.DeleteInstancesRequest) {
	o.Body = body
}

// WithAuthorization adds the authorization to the delete instances params
func (o *DeleteInstancesParams) WithAuthorization(authorization string) *DeleteInstancesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete instances params
func (o *DeleteInstancesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the delete instances params
func (o *DeleteInstancesParams) WithBmpLanguage(bmpLanguage *string) *DeleteInstancesParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the delete instances params
func (o *DeleteInstancesParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the delete instances params
func (o *DeleteInstancesParams) WithBmpUserID(bmpUserID *string) *DeleteInstancesParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the delete instances params
func (o *DeleteInstancesParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithTraceID adds the traceID to the delete instances params
func (o *DeleteInstancesParams) WithTraceID(traceID string) *DeleteInstancesParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the delete instances params
func (o *DeleteInstancesParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
