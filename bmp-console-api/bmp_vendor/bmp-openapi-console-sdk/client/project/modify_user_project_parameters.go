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

// NewModifyUserProjectParams creates a new ModifyUserProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModifyUserProjectParams() *ModifyUserProjectParams {
	return &ModifyUserProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModifyUserProjectParamsWithTimeout creates a new ModifyUserProjectParams object
// with the ability to set a timeout on a request.
func NewModifyUserProjectParamsWithTimeout(timeout time.Duration) *ModifyUserProjectParams {
	return &ModifyUserProjectParams{
		timeout: timeout,
	}
}

// NewModifyUserProjectParamsWithContext creates a new ModifyUserProjectParams object
// with the ability to set a context for a request.
func NewModifyUserProjectParamsWithContext(ctx context.Context) *ModifyUserProjectParams {
	return &ModifyUserProjectParams{
		Context: ctx,
	}
}

// NewModifyUserProjectParamsWithHTTPClient creates a new ModifyUserProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewModifyUserProjectParamsWithHTTPClient(client *http.Client) *ModifyUserProjectParams {
	return &ModifyUserProjectParams{
		HTTPClient: client,
	}
}

/* ModifyUserProjectParams contains all the parameters to send to the API endpoint
   for the modify user project operation.

   Typically these are written to a http.Request.
*/
type ModifyUserProjectParams struct {

	// Body.
	Body *models.ModifyProjectRequest

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

// WithDefaults hydrates default values in the modify user project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyUserProjectParams) WithDefaults() *ModifyUserProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the modify user project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyUserProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the modify user project params
func (o *ModifyUserProjectParams) WithTimeout(timeout time.Duration) *ModifyUserProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify user project params
func (o *ModifyUserProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify user project params
func (o *ModifyUserProjectParams) WithContext(ctx context.Context) *ModifyUserProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify user project params
func (o *ModifyUserProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify user project params
func (o *ModifyUserProjectParams) WithHTTPClient(client *http.Client) *ModifyUserProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify user project params
func (o *ModifyUserProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the modify user project params
func (o *ModifyUserProjectParams) WithBody(body *models.ModifyProjectRequest) *ModifyUserProjectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify user project params
func (o *ModifyUserProjectParams) SetBody(body *models.ModifyProjectRequest) {
	o.Body = body
}

// WithAuthorization adds the authorization to the modify user project params
func (o *ModifyUserProjectParams) WithAuthorization(authorization string) *ModifyUserProjectParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the modify user project params
func (o *ModifyUserProjectParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the modify user project params
func (o *ModifyUserProjectParams) WithBmpLanguage(bmpLanguage *string) *ModifyUserProjectParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the modify user project params
func (o *ModifyUserProjectParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the modify user project params
func (o *ModifyUserProjectParams) WithBmpUserID(bmpUserID *string) *ModifyUserProjectParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the modify user project params
func (o *ModifyUserProjectParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithProjectID adds the projectID to the modify user project params
func (o *ModifyUserProjectParams) WithProjectID(projectID string) *ModifyUserProjectParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the modify user project params
func (o *ModifyUserProjectParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithTraceID adds the traceID to the modify user project params
func (o *ModifyUserProjectParams) WithTraceID(traceID string) *ModifyUserProjectParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the modify user project params
func (o *ModifyUserProjectParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyUserProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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