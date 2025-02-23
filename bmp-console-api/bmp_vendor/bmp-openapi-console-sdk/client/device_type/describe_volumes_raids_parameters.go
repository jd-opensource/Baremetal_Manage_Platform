// Code generated by go-swagger; DO NOT EDIT.

package device_type

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

// NewDescribeVolumesRaidsParams creates a new DescribeVolumesRaidsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeVolumesRaidsParams() *DescribeVolumesRaidsParams {
	return &DescribeVolumesRaidsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeVolumesRaidsParamsWithTimeout creates a new DescribeVolumesRaidsParams object
// with the ability to set a timeout on a request.
func NewDescribeVolumesRaidsParamsWithTimeout(timeout time.Duration) *DescribeVolumesRaidsParams {
	return &DescribeVolumesRaidsParams{
		timeout: timeout,
	}
}

// NewDescribeVolumesRaidsParamsWithContext creates a new DescribeVolumesRaidsParams object
// with the ability to set a context for a request.
func NewDescribeVolumesRaidsParamsWithContext(ctx context.Context) *DescribeVolumesRaidsParams {
	return &DescribeVolumesRaidsParams{
		Context: ctx,
	}
}

// NewDescribeVolumesRaidsParamsWithHTTPClient creates a new DescribeVolumesRaidsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeVolumesRaidsParamsWithHTTPClient(client *http.Client) *DescribeVolumesRaidsParams {
	return &DescribeVolumesRaidsParams{
		HTTPClient: client,
	}
}

/* DescribeVolumesRaidsParams contains all the parameters to send to the API endpoint
   for the describe volumes raids operation.

   Typically these are written to a http.Request.
*/
type DescribeVolumesRaidsParams struct {

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

	/* DeviceTypeID.

	   机型id
	*/
	DeviceTypeID *string

	/* IsAll.

	   是否显示所有
	*/
	IsAll *string

	/* TraceID.

	   流量唯一id
	*/
	TraceID string

	/* VolumeType.

	   系统盘还是数据盘
	*/
	VolumeType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe volumes raids params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeVolumesRaidsParams) WithDefaults() *DescribeVolumesRaidsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe volumes raids params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeVolumesRaidsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) WithTimeout(timeout time.Duration) *DescribeVolumesRaidsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) WithContext(ctx context.Context) *DescribeVolumesRaidsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) WithHTTPClient(client *http.Client) *DescribeVolumesRaidsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) WithAuthorization(authorization string) *DescribeVolumesRaidsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) WithBmpLanguage(bmpLanguage *string) *DescribeVolumesRaidsParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) WithBmpUserID(bmpUserID *string) *DescribeVolumesRaidsParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithDeviceTypeID adds the deviceTypeID to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) WithDeviceTypeID(deviceTypeID *string) *DescribeVolumesRaidsParams {
	o.SetDeviceTypeID(deviceTypeID)
	return o
}

// SetDeviceTypeID adds the deviceTypeId to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) SetDeviceTypeID(deviceTypeID *string) {
	o.DeviceTypeID = deviceTypeID
}

// WithIsAll adds the isAll to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) WithIsAll(isAll *string) *DescribeVolumesRaidsParams {
	o.SetIsAll(isAll)
	return o
}

// SetIsAll adds the isAll to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) SetIsAll(isAll *string) {
	o.IsAll = isAll
}

// WithTraceID adds the traceID to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) WithTraceID(traceID string) *DescribeVolumesRaidsParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WithVolumeType adds the volumeType to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) WithVolumeType(volumeType *string) *DescribeVolumesRaidsParams {
	o.SetVolumeType(volumeType)
	return o
}

// SetVolumeType adds the volumeType to the describe volumes raids params
func (o *DescribeVolumesRaidsParams) SetVolumeType(volumeType *string) {
	o.VolumeType = volumeType
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeVolumesRaidsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IsAll != nil {

		// query param isAll
		var qrIsAll string

		if o.IsAll != nil {
			qrIsAll = *o.IsAll
		}
		qIsAll := qrIsAll
		if qIsAll != "" {

			if err := r.SetQueryParam("isAll", qIsAll); err != nil {
				return err
			}
		}
	}

	// header param traceId
	if err := r.SetHeaderParam("traceId", o.TraceID); err != nil {
		return err
	}

	if o.VolumeType != nil {

		// query param volumeType
		var qrVolumeType string

		if o.VolumeType != nil {
			qrVolumeType = *o.VolumeType
		}
		qVolumeType := qrVolumeType
		if qVolumeType != "" {

			if err := r.SetQueryParam("volumeType", qVolumeType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
