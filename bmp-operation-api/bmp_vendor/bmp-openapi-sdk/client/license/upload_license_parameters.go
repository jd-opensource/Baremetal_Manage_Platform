// Code generated by go-swagger; DO NOT EDIT.

package license

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

// NewUploadLicenseParams creates a new UploadLicenseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadLicenseParams() *UploadLicenseParams {
	return &UploadLicenseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadLicenseParamsWithTimeout creates a new UploadLicenseParams object
// with the ability to set a timeout on a request.
func NewUploadLicenseParamsWithTimeout(timeout time.Duration) *UploadLicenseParams {
	return &UploadLicenseParams{
		timeout: timeout,
	}
}

// NewUploadLicenseParamsWithContext creates a new UploadLicenseParams object
// with the ability to set a context for a request.
func NewUploadLicenseParamsWithContext(ctx context.Context) *UploadLicenseParams {
	return &UploadLicenseParams{
		Context: ctx,
	}
}

// NewUploadLicenseParamsWithHTTPClient creates a new UploadLicenseParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadLicenseParamsWithHTTPClient(client *http.Client) *UploadLicenseParams {
	return &UploadLicenseParams{
		HTTPClient: client,
	}
}

/* UploadLicenseParams contains all the parameters to send to the API endpoint
   for the upload license operation.

   Typically these are written to a http.Request.
*/
type UploadLicenseParams struct {

	// Body.
	Body *models.CreateLicenseRequest

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

// WithDefaults hydrates default values in the upload license params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadLicenseParams) WithDefaults() *UploadLicenseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload license params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadLicenseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload license params
func (o *UploadLicenseParams) WithTimeout(timeout time.Duration) *UploadLicenseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload license params
func (o *UploadLicenseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload license params
func (o *UploadLicenseParams) WithContext(ctx context.Context) *UploadLicenseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload license params
func (o *UploadLicenseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload license params
func (o *UploadLicenseParams) WithHTTPClient(client *http.Client) *UploadLicenseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload license params
func (o *UploadLicenseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upload license params
func (o *UploadLicenseParams) WithBody(body *models.CreateLicenseRequest) *UploadLicenseParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upload license params
func (o *UploadLicenseParams) SetBody(body *models.CreateLicenseRequest) {
	o.Body = body
}

// WithAuthorization adds the authorization to the upload license params
func (o *UploadLicenseParams) WithAuthorization(authorization string) *UploadLicenseParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the upload license params
func (o *UploadLicenseParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the upload license params
func (o *UploadLicenseParams) WithBmpLanguage(bmpLanguage *string) *UploadLicenseParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the upload license params
func (o *UploadLicenseParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the upload license params
func (o *UploadLicenseParams) WithBmpUserID(bmpUserID *string) *UploadLicenseParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the upload license params
func (o *UploadLicenseParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithTraceID adds the traceID to the upload license params
func (o *UploadLicenseParams) WithTraceID(traceID string) *UploadLicenseParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the upload license params
func (o *UploadLicenseParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *UploadLicenseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
