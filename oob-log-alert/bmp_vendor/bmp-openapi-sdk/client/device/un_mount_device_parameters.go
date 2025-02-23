// Code generated by go-swagger; DO NOT EDIT.

package device

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

// NewUnMountDeviceParams creates a new UnMountDeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnMountDeviceParams() *UnMountDeviceParams {
	return &UnMountDeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnMountDeviceParamsWithTimeout creates a new UnMountDeviceParams object
// with the ability to set a timeout on a request.
func NewUnMountDeviceParamsWithTimeout(timeout time.Duration) *UnMountDeviceParams {
	return &UnMountDeviceParams{
		timeout: timeout,
	}
}

// NewUnMountDeviceParamsWithContext creates a new UnMountDeviceParams object
// with the ability to set a context for a request.
func NewUnMountDeviceParamsWithContext(ctx context.Context) *UnMountDeviceParams {
	return &UnMountDeviceParams{
		Context: ctx,
	}
}

// NewUnMountDeviceParamsWithHTTPClient creates a new UnMountDeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnMountDeviceParamsWithHTTPClient(client *http.Client) *UnMountDeviceParams {
	return &UnMountDeviceParams{
		HTTPClient: client,
	}
}

/* UnMountDeviceParams contains all the parameters to send to the API endpoint
   for the un mount device operation.

   Typically these are written to a http.Request.
*/
type UnMountDeviceParams struct {

	// Body.
	Body *models.UnMountDevicesRequest

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

// WithDefaults hydrates default values in the un mount device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnMountDeviceParams) WithDefaults() *UnMountDeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the un mount device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnMountDeviceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the un mount device params
func (o *UnMountDeviceParams) WithTimeout(timeout time.Duration) *UnMountDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the un mount device params
func (o *UnMountDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the un mount device params
func (o *UnMountDeviceParams) WithContext(ctx context.Context) *UnMountDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the un mount device params
func (o *UnMountDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the un mount device params
func (o *UnMountDeviceParams) WithHTTPClient(client *http.Client) *UnMountDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the un mount device params
func (o *UnMountDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the un mount device params
func (o *UnMountDeviceParams) WithBody(body *models.UnMountDevicesRequest) *UnMountDeviceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the un mount device params
func (o *UnMountDeviceParams) SetBody(body *models.UnMountDevicesRequest) {
	o.Body = body
}

// WithAuthorization adds the authorization to the un mount device params
func (o *UnMountDeviceParams) WithAuthorization(authorization string) *UnMountDeviceParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the un mount device params
func (o *UnMountDeviceParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the un mount device params
func (o *UnMountDeviceParams) WithBmpLanguage(bmpLanguage *string) *UnMountDeviceParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the un mount device params
func (o *UnMountDeviceParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the un mount device params
func (o *UnMountDeviceParams) WithBmpUserID(bmpUserID *string) *UnMountDeviceParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the un mount device params
func (o *UnMountDeviceParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithTraceID adds the traceID to the un mount device params
func (o *UnMountDeviceParams) WithTraceID(traceID string) *UnMountDeviceParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the un mount device params
func (o *UnMountDeviceParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *UnMountDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
