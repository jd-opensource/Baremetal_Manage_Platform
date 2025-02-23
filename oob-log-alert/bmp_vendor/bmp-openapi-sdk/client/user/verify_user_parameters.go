// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewVerifyUserParams creates a new VerifyUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVerifyUserParams() *VerifyUserParams {
	return &VerifyUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVerifyUserParamsWithTimeout creates a new VerifyUserParams object
// with the ability to set a timeout on a request.
func NewVerifyUserParamsWithTimeout(timeout time.Duration) *VerifyUserParams {
	return &VerifyUserParams{
		timeout: timeout,
	}
}

// NewVerifyUserParamsWithContext creates a new VerifyUserParams object
// with the ability to set a context for a request.
func NewVerifyUserParamsWithContext(ctx context.Context) *VerifyUserParams {
	return &VerifyUserParams{
		Context: ctx,
	}
}

// NewVerifyUserParamsWithHTTPClient creates a new VerifyUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewVerifyUserParamsWithHTTPClient(client *http.Client) *VerifyUserParams {
	return &VerifyUserParams{
		HTTPClient: client,
	}
}

/* VerifyUserParams contains all the parameters to send to the API endpoint
   for the verify user operation.

   Typically these are written to a http.Request.
*/
type VerifyUserParams struct {

	// Body.
	Body *models.VerifyUserRequest

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

// WithDefaults hydrates default values in the verify user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifyUserParams) WithDefaults() *VerifyUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the verify user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifyUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the verify user params
func (o *VerifyUserParams) WithTimeout(timeout time.Duration) *VerifyUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the verify user params
func (o *VerifyUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the verify user params
func (o *VerifyUserParams) WithContext(ctx context.Context) *VerifyUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the verify user params
func (o *VerifyUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the verify user params
func (o *VerifyUserParams) WithHTTPClient(client *http.Client) *VerifyUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the verify user params
func (o *VerifyUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the verify user params
func (o *VerifyUserParams) WithBody(body *models.VerifyUserRequest) *VerifyUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the verify user params
func (o *VerifyUserParams) SetBody(body *models.VerifyUserRequest) {
	o.Body = body
}

// WithAuthorization adds the authorization to the verify user params
func (o *VerifyUserParams) WithAuthorization(authorization string) *VerifyUserParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the verify user params
func (o *VerifyUserParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the verify user params
func (o *VerifyUserParams) WithBmpLanguage(bmpLanguage *string) *VerifyUserParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the verify user params
func (o *VerifyUserParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the verify user params
func (o *VerifyUserParams) WithBmpUserID(bmpUserID *string) *VerifyUserParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the verify user params
func (o *VerifyUserParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithTraceID adds the traceID to the verify user params
func (o *VerifyUserParams) WithTraceID(traceID string) *VerifyUserParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the verify user params
func (o *VerifyUserParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *VerifyUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
