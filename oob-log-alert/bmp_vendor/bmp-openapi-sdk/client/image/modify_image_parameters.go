// Code generated by go-swagger; DO NOT EDIT.

package image

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

// NewModifyImageParams creates a new ModifyImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModifyImageParams() *ModifyImageParams {
	return &ModifyImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModifyImageParamsWithTimeout creates a new ModifyImageParams object
// with the ability to set a timeout on a request.
func NewModifyImageParamsWithTimeout(timeout time.Duration) *ModifyImageParams {
	return &ModifyImageParams{
		timeout: timeout,
	}
}

// NewModifyImageParamsWithContext creates a new ModifyImageParams object
// with the ability to set a context for a request.
func NewModifyImageParamsWithContext(ctx context.Context) *ModifyImageParams {
	return &ModifyImageParams{
		Context: ctx,
	}
}

// NewModifyImageParamsWithHTTPClient creates a new ModifyImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewModifyImageParamsWithHTTPClient(client *http.Client) *ModifyImageParams {
	return &ModifyImageParams{
		HTTPClient: client,
	}
}

/* ModifyImageParams contains all the parameters to send to the API endpoint
   for the modify image operation.

   Typically these are written to a http.Request.
*/
type ModifyImageParams struct {

	// Body.
	Body *models.ModifyImageRequest

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

	// ImageID.
	ImageID string

	/* TraceID.

	   流量唯一id
	*/
	TraceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the modify image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyImageParams) WithDefaults() *ModifyImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the modify image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the modify image params
func (o *ModifyImageParams) WithTimeout(timeout time.Duration) *ModifyImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify image params
func (o *ModifyImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify image params
func (o *ModifyImageParams) WithContext(ctx context.Context) *ModifyImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify image params
func (o *ModifyImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify image params
func (o *ModifyImageParams) WithHTTPClient(client *http.Client) *ModifyImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify image params
func (o *ModifyImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the modify image params
func (o *ModifyImageParams) WithBody(body *models.ModifyImageRequest) *ModifyImageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify image params
func (o *ModifyImageParams) SetBody(body *models.ModifyImageRequest) {
	o.Body = body
}

// WithAuthorization adds the authorization to the modify image params
func (o *ModifyImageParams) WithAuthorization(authorization string) *ModifyImageParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the modify image params
func (o *ModifyImageParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the modify image params
func (o *ModifyImageParams) WithBmpLanguage(bmpLanguage *string) *ModifyImageParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the modify image params
func (o *ModifyImageParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the modify image params
func (o *ModifyImageParams) WithBmpUserID(bmpUserID *string) *ModifyImageParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the modify image params
func (o *ModifyImageParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithImageID adds the imageID to the modify image params
func (o *ModifyImageParams) WithImageID(imageID string) *ModifyImageParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the modify image params
func (o *ModifyImageParams) SetImageID(imageID string) {
	o.ImageID = imageID
}

// WithTraceID adds the traceID to the modify image params
func (o *ModifyImageParams) WithTraceID(traceID string) *ModifyImageParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the modify image params
func (o *ModifyImageParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param image_id
	if err := r.SetPathParam("image_id", o.ImageID); err != nil {
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