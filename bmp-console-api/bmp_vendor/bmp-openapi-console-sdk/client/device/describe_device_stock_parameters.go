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
)

// NewDescribeDeviceStockParams creates a new DescribeDeviceStockParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeDeviceStockParams() *DescribeDeviceStockParams {
	return &DescribeDeviceStockParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeDeviceStockParamsWithTimeout creates a new DescribeDeviceStockParams object
// with the ability to set a timeout on a request.
func NewDescribeDeviceStockParamsWithTimeout(timeout time.Duration) *DescribeDeviceStockParams {
	return &DescribeDeviceStockParams{
		timeout: timeout,
	}
}

// NewDescribeDeviceStockParamsWithContext creates a new DescribeDeviceStockParams object
// with the ability to set a context for a request.
func NewDescribeDeviceStockParamsWithContext(ctx context.Context) *DescribeDeviceStockParams {
	return &DescribeDeviceStockParams{
		Context: ctx,
	}
}

// NewDescribeDeviceStockParamsWithHTTPClient creates a new DescribeDeviceStockParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeDeviceStockParamsWithHTTPClient(client *http.Client) *DescribeDeviceStockParams {
	return &DescribeDeviceStockParams{
		HTTPClient: client,
	}
}

/* DescribeDeviceStockParams contains all the parameters to send to the API endpoint
   for the describe device stock operation.

   Typically these are written to a http.Request.
*/
type DescribeDeviceStockParams struct {

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

	// DeviceTypeID.
	DeviceTypeID *string

	/* TraceID.

	   流量唯一id
	*/
	TraceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe device stock params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeDeviceStockParams) WithDefaults() *DescribeDeviceStockParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe device stock params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeDeviceStockParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe device stock params
func (o *DescribeDeviceStockParams) WithTimeout(timeout time.Duration) *DescribeDeviceStockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe device stock params
func (o *DescribeDeviceStockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe device stock params
func (o *DescribeDeviceStockParams) WithContext(ctx context.Context) *DescribeDeviceStockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe device stock params
func (o *DescribeDeviceStockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe device stock params
func (o *DescribeDeviceStockParams) WithHTTPClient(client *http.Client) *DescribeDeviceStockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe device stock params
func (o *DescribeDeviceStockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the describe device stock params
func (o *DescribeDeviceStockParams) WithAuthorization(authorization string) *DescribeDeviceStockParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the describe device stock params
func (o *DescribeDeviceStockParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the describe device stock params
func (o *DescribeDeviceStockParams) WithBmpLanguage(bmpLanguage *string) *DescribeDeviceStockParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the describe device stock params
func (o *DescribeDeviceStockParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the describe device stock params
func (o *DescribeDeviceStockParams) WithBmpUserID(bmpUserID *string) *DescribeDeviceStockParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the describe device stock params
func (o *DescribeDeviceStockParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithDeviceTypeID adds the deviceTypeID to the describe device stock params
func (o *DescribeDeviceStockParams) WithDeviceTypeID(deviceTypeID *string) *DescribeDeviceStockParams {
	o.SetDeviceTypeID(deviceTypeID)
	return o
}

// SetDeviceTypeID adds the deviceTypeId to the describe device stock params
func (o *DescribeDeviceStockParams) SetDeviceTypeID(deviceTypeID *string) {
	o.DeviceTypeID = deviceTypeID
}

// WithTraceID adds the traceID to the describe device stock params
func (o *DescribeDeviceStockParams) WithTraceID(traceID string) *DescribeDeviceStockParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the describe device stock params
func (o *DescribeDeviceStockParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeDeviceStockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DeviceTypeID != nil {

		// query param deviceTypeId
		var qrDeviceTypeID string

		if o.DeviceTypeID != nil {
			qrDeviceTypeID = *o.DeviceTypeID
		}
		qDeviceTypeID := qrDeviceTypeID
		if qDeviceTypeID != "" {

			if err := r.SetQueryParam("deviceTypeId", qDeviceTypeID); err != nil {
				return err
			}
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