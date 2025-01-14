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

	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
)

// NewCreateRoleParams creates a new CreateRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRoleParams() *CreateRoleParams {
	return &CreateRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRoleParamsWithTimeout creates a new CreateRoleParams object
// with the ability to set a timeout on a request.
func NewCreateRoleParamsWithTimeout(timeout time.Duration) *CreateRoleParams {
	return &CreateRoleParams{
		timeout: timeout,
	}
}

// NewCreateRoleParamsWithContext creates a new CreateRoleParams object
// with the ability to set a context for a request.
func NewCreateRoleParamsWithContext(ctx context.Context) *CreateRoleParams {
	return &CreateRoleParams{
		Context: ctx,
	}
}

// NewCreateRoleParamsWithHTTPClient creates a new CreateRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRoleParamsWithHTTPClient(client *http.Client) *CreateRoleParams {
	return &CreateRoleParams{
		HTTPClient: client,
	}
}

/* CreateRoleParams contains all the parameters to send to the API endpoint
   for the create role operation.

   Typically these are written to a http.Request.
*/
type CreateRoleParams struct {

	// Body.
	Body *models.CreateRoleRequest

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

// WithDefaults hydrates default values in the create role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRoleParams) WithDefaults() *CreateRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create role params
func (o *CreateRoleParams) WithTimeout(timeout time.Duration) *CreateRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create role params
func (o *CreateRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create role params
func (o *CreateRoleParams) WithContext(ctx context.Context) *CreateRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create role params
func (o *CreateRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create role params
func (o *CreateRoleParams) WithHTTPClient(client *http.Client) *CreateRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create role params
func (o *CreateRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create role params
func (o *CreateRoleParams) WithBody(body *models.CreateRoleRequest) *CreateRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create role params
func (o *CreateRoleParams) SetBody(body *models.CreateRoleRequest) {
	o.Body = body
}

// WithAuthorization adds the authorization to the create role params
func (o *CreateRoleParams) WithAuthorization(authorization string) *CreateRoleParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the create role params
func (o *CreateRoleParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the create role params
func (o *CreateRoleParams) WithBmpLanguage(bmpLanguage *string) *CreateRoleParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the create role params
func (o *CreateRoleParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the create role params
func (o *CreateRoleParams) WithBmpUserID(bmpUserID *string) *CreateRoleParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the create role params
func (o *CreateRoleParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithTraceID adds the traceID to the create role params
func (o *CreateRoleParams) WithTraceID(traceID string) *CreateRoleParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the create role params
func (o *CreateRoleParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
