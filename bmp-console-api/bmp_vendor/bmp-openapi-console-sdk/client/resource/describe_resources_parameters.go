// Code generated by go-swagger; DO NOT EDIT.

package resource

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

// NewDescribeResourcesParams creates a new DescribeResourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeResourcesParams() *DescribeResourcesParams {
	return &DescribeResourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeResourcesParamsWithTimeout creates a new DescribeResourcesParams object
// with the ability to set a timeout on a request.
func NewDescribeResourcesParamsWithTimeout(timeout time.Duration) *DescribeResourcesParams {
	return &DescribeResourcesParams{
		timeout: timeout,
	}
}

// NewDescribeResourcesParamsWithContext creates a new DescribeResourcesParams object
// with the ability to set a context for a request.
func NewDescribeResourcesParamsWithContext(ctx context.Context) *DescribeResourcesParams {
	return &DescribeResourcesParams{
		Context: ctx,
	}
}

// NewDescribeResourcesParamsWithHTTPClient creates a new DescribeResourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeResourcesParamsWithHTTPClient(client *http.Client) *DescribeResourcesParams {
	return &DescribeResourcesParams{
		HTTPClient: client,
	}
}

/* DescribeResourcesParams contains all the parameters to send to the API endpoint
   for the describe resources operation.

   Typically these are written to a http.Request.
*/
type DescribeResourcesParams struct {

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

	/* DeviceType.

	   机型规格
	*/
	DeviceType *string

	/* ImageName.

	   镜像名称
	*/
	ImageName *string

	/* Name.

	     机房名称
	IdcName string `json:"idcName"`
	机房英文名称
	IdcNameEn string `json:"idcNameEn"`[*]/
	机型名称
	*/
	Name *string

	/* TraceID.

	   流量唯一id
	*/
	TraceID string

	/* UserName.

	   用户名称
	*/
	UserName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeResourcesParams) WithDefaults() *DescribeResourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeResourcesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe resources params
func (o *DescribeResourcesParams) WithTimeout(timeout time.Duration) *DescribeResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe resources params
func (o *DescribeResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe resources params
func (o *DescribeResourcesParams) WithContext(ctx context.Context) *DescribeResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe resources params
func (o *DescribeResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe resources params
func (o *DescribeResourcesParams) WithHTTPClient(client *http.Client) *DescribeResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe resources params
func (o *DescribeResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the describe resources params
func (o *DescribeResourcesParams) WithAuthorization(authorization string) *DescribeResourcesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the describe resources params
func (o *DescribeResourcesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the describe resources params
func (o *DescribeResourcesParams) WithBmpLanguage(bmpLanguage *string) *DescribeResourcesParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the describe resources params
func (o *DescribeResourcesParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the describe resources params
func (o *DescribeResourcesParams) WithBmpUserID(bmpUserID *string) *DescribeResourcesParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the describe resources params
func (o *DescribeResourcesParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithDeviceType adds the deviceType to the describe resources params
func (o *DescribeResourcesParams) WithDeviceType(deviceType *string) *DescribeResourcesParams {
	o.SetDeviceType(deviceType)
	return o
}

// SetDeviceType adds the deviceType to the describe resources params
func (o *DescribeResourcesParams) SetDeviceType(deviceType *string) {
	o.DeviceType = deviceType
}

// WithImageName adds the imageName to the describe resources params
func (o *DescribeResourcesParams) WithImageName(imageName *string) *DescribeResourcesParams {
	o.SetImageName(imageName)
	return o
}

// SetImageName adds the imageName to the describe resources params
func (o *DescribeResourcesParams) SetImageName(imageName *string) {
	o.ImageName = imageName
}

// WithName adds the name to the describe resources params
func (o *DescribeResourcesParams) WithName(name *string) *DescribeResourcesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the describe resources params
func (o *DescribeResourcesParams) SetName(name *string) {
	o.Name = name
}

// WithTraceID adds the traceID to the describe resources params
func (o *DescribeResourcesParams) WithTraceID(traceID string) *DescribeResourcesParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the describe resources params
func (o *DescribeResourcesParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WithUserName adds the userName to the describe resources params
func (o *DescribeResourcesParams) WithUserName(userName *string) *DescribeResourcesParams {
	o.SetUserName(userName)
	return o
}

// SetUserName adds the userName to the describe resources params
func (o *DescribeResourcesParams) SetUserName(userName *string) {
	o.UserName = userName
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DeviceType != nil {

		// query param deviceType
		var qrDeviceType string

		if o.DeviceType != nil {
			qrDeviceType = *o.DeviceType
		}
		qDeviceType := qrDeviceType
		if qDeviceType != "" {

			if err := r.SetQueryParam("deviceType", qDeviceType); err != nil {
				return err
			}
		}
	}

	if o.ImageName != nil {

		// query param imageName
		var qrImageName string

		if o.ImageName != nil {
			qrImageName = *o.ImageName
		}
		qImageName := qrImageName
		if qImageName != "" {

			if err := r.SetQueryParam("imageName", qImageName); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	// header param traceId
	if err := r.SetHeaderParam("traceId", o.TraceID); err != nil {
		return err
	}

	if o.UserName != nil {

		// query param userName
		var qrUserName string

		if o.UserName != nil {
			qrUserName = *o.UserName
		}
		qUserName := qrUserName
		if qUserName != "" {

			if err := r.SetQueryParam("userName", qUserName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
