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

// NewModifyRoleParams creates a new ModifyRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModifyRoleParams() *ModifyRoleParams {
	return &ModifyRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModifyRoleParamsWithTimeout creates a new ModifyRoleParams object
// with the ability to set a timeout on a request.
func NewModifyRoleParamsWithTimeout(timeout time.Duration) *ModifyRoleParams {
	return &ModifyRoleParams{
		timeout: timeout,
	}
}

// NewModifyRoleParamsWithContext creates a new ModifyRoleParams object
// with the ability to set a context for a request.
func NewModifyRoleParamsWithContext(ctx context.Context) *ModifyRoleParams {
	return &ModifyRoleParams{
		Context: ctx,
	}
}

// NewModifyRoleParamsWithHTTPClient creates a new ModifyRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewModifyRoleParamsWithHTTPClient(client *http.Client) *ModifyRoleParams {
	return &ModifyRoleParams{
		HTTPClient: client,
	}
}

/* ModifyRoleParams contains all the parameters to send to the API endpoint
   for the modify role operation.

   Typically these are written to a http.Request.
*/
type ModifyRoleParams struct {

	// Body.
	Body *models.ModifyRoleRequest

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

	// RoleID.
	RoleID string

	/* TraceID.

	   流量唯一id
	*/
	TraceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the modify role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyRoleParams) WithDefaults() *ModifyRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the modify role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the modify role params
func (o *ModifyRoleParams) WithTimeout(timeout time.Duration) *ModifyRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify role params
func (o *ModifyRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify role params
func (o *ModifyRoleParams) WithContext(ctx context.Context) *ModifyRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify role params
func (o *ModifyRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify role params
func (o *ModifyRoleParams) WithHTTPClient(client *http.Client) *ModifyRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify role params
func (o *ModifyRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the modify role params
func (o *ModifyRoleParams) WithBody(body *models.ModifyRoleRequest) *ModifyRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify role params
func (o *ModifyRoleParams) SetBody(body *models.ModifyRoleRequest) {
	o.Body = body
}

// WithAuthorization adds the authorization to the modify role params
func (o *ModifyRoleParams) WithAuthorization(authorization string) *ModifyRoleParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the modify role params
func (o *ModifyRoleParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the modify role params
func (o *ModifyRoleParams) WithBmpLanguage(bmpLanguage *string) *ModifyRoleParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the modify role params
func (o *ModifyRoleParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the modify role params
func (o *ModifyRoleParams) WithBmpUserID(bmpUserID *string) *ModifyRoleParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the modify role params
func (o *ModifyRoleParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithRoleID adds the roleID to the modify role params
func (o *ModifyRoleParams) WithRoleID(roleID string) *ModifyRoleParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the modify role params
func (o *ModifyRoleParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WithTraceID adds the traceID to the modify role params
func (o *ModifyRoleParams) WithTraceID(traceID string) *ModifyRoleParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the modify role params
func (o *ModifyRoleParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param role_id
	if err := r.SetPathParam("role_id", o.RoleID); err != nil {
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