// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewShareUserProjectParams creates a new ShareUserProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShareUserProjectParams() *ShareUserProjectParams {
	return &ShareUserProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShareUserProjectParamsWithTimeout creates a new ShareUserProjectParams object
// with the ability to set a timeout on a request.
func NewShareUserProjectParamsWithTimeout(timeout time.Duration) *ShareUserProjectParams {
	return &ShareUserProjectParams{
		timeout: timeout,
	}
}

// NewShareUserProjectParamsWithContext creates a new ShareUserProjectParams object
// with the ability to set a context for a request.
func NewShareUserProjectParamsWithContext(ctx context.Context) *ShareUserProjectParams {
	return &ShareUserProjectParams{
		Context: ctx,
	}
}

// NewShareUserProjectParamsWithHTTPClient creates a new ShareUserProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewShareUserProjectParamsWithHTTPClient(client *http.Client) *ShareUserProjectParams {
	return &ShareUserProjectParams{
		HTTPClient: client,
	}
}

/* ShareUserProjectParams contains all the parameters to send to the API endpoint
   for the share user project operation.

   Typically these are written to a http.Request.
*/
type ShareUserProjectParams struct {

	// Body.
	Body *models.ShareProjectRequest

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

	// ProjectID.
	ProjectID string

	/* TraceID.

	   流量唯一id
	*/
	TraceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the share user project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShareUserProjectParams) WithDefaults() *ShareUserProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the share user project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShareUserProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the share user project params
func (o *ShareUserProjectParams) WithTimeout(timeout time.Duration) *ShareUserProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the share user project params
func (o *ShareUserProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the share user project params
func (o *ShareUserProjectParams) WithContext(ctx context.Context) *ShareUserProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the share user project params
func (o *ShareUserProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the share user project params
func (o *ShareUserProjectParams) WithHTTPClient(client *http.Client) *ShareUserProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the share user project params
func (o *ShareUserProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the share user project params
func (o *ShareUserProjectParams) WithBody(body *models.ShareProjectRequest) *ShareUserProjectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the share user project params
func (o *ShareUserProjectParams) SetBody(body *models.ShareProjectRequest) {
	o.Body = body
}

// WithAuthorization adds the authorization to the share user project params
func (o *ShareUserProjectParams) WithAuthorization(authorization string) *ShareUserProjectParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the share user project params
func (o *ShareUserProjectParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the share user project params
func (o *ShareUserProjectParams) WithBmpLanguage(bmpLanguage *string) *ShareUserProjectParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the share user project params
func (o *ShareUserProjectParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the share user project params
func (o *ShareUserProjectParams) WithBmpUserID(bmpUserID *string) *ShareUserProjectParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the share user project params
func (o *ShareUserProjectParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithProjectID adds the projectID to the share user project params
func (o *ShareUserProjectParams) WithProjectID(projectID string) *ShareUserProjectParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the share user project params
func (o *ShareUserProjectParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithTraceID adds the traceID to the share user project params
func (o *ShareUserProjectParams) WithTraceID(traceID string) *ShareUserProjectParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the share user project params
func (o *ShareUserProjectParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *ShareUserProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
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