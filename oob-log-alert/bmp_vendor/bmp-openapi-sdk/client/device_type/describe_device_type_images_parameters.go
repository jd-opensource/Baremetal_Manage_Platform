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
	"github.com/go-openapi/swag"
)

// NewDescribeDeviceTypeImagesParams creates a new DescribeDeviceTypeImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeDeviceTypeImagesParams() *DescribeDeviceTypeImagesParams {
	return &DescribeDeviceTypeImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeDeviceTypeImagesParamsWithTimeout creates a new DescribeDeviceTypeImagesParams object
// with the ability to set a timeout on a request.
func NewDescribeDeviceTypeImagesParamsWithTimeout(timeout time.Duration) *DescribeDeviceTypeImagesParams {
	return &DescribeDeviceTypeImagesParams{
		timeout: timeout,
	}
}

// NewDescribeDeviceTypeImagesParamsWithContext creates a new DescribeDeviceTypeImagesParams object
// with the ability to set a context for a request.
func NewDescribeDeviceTypeImagesParamsWithContext(ctx context.Context) *DescribeDeviceTypeImagesParams {
	return &DescribeDeviceTypeImagesParams{
		Context: ctx,
	}
}

// NewDescribeDeviceTypeImagesParamsWithHTTPClient creates a new DescribeDeviceTypeImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeDeviceTypeImagesParamsWithHTTPClient(client *http.Client) *DescribeDeviceTypeImagesParams {
	return &DescribeDeviceTypeImagesParams{
		HTTPClient: client,
	}
}

/* DescribeDeviceTypeImagesParams contains all the parameters to send to the API endpoint
   for the describe device type images operation.

   Typically these are written to a http.Request.
*/
type DescribeDeviceTypeImagesParams struct {

	/* Architecture.

	   体系架构
	*/
	Architecture *string

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

	/* ImageID.

	   镜像id
	*/
	ImageID *string

	/* ImageName.

	   镜像名称
	*/
	ImageName *string

	/* IsAll.

	   是否显示全部
	*/
	IsAll *string

	/* OsID.

	   操作系统ID
	*/
	OsID *string

	/* OsType.

	   操作系统平台
	*/
	OsType *string

	/* PageNumber.

	   页码

	   Format: int64
	*/
	PageNumber *int64

	/* PageSize.

	   每页数量

	   Format: int64
	*/
	PageSize *int64

	/* Source.

	   镜像类型，预置，自定义
	*/
	Source *string

	/* TraceID.

	   流量唯一id
	*/
	TraceID string

	/* Version.

	   版本号
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe device type images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeDeviceTypeImagesParams) WithDefaults() *DescribeDeviceTypeImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe device type images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeDeviceTypeImagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithTimeout(timeout time.Duration) *DescribeDeviceTypeImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithContext(ctx context.Context) *DescribeDeviceTypeImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithHTTPClient(client *http.Client) *DescribeDeviceTypeImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArchitecture adds the architecture to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithArchitecture(architecture *string) *DescribeDeviceTypeImagesParams {
	o.SetArchitecture(architecture)
	return o
}

// SetArchitecture adds the architecture to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetArchitecture(architecture *string) {
	o.Architecture = architecture
}

// WithAuthorization adds the authorization to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithAuthorization(authorization string) *DescribeDeviceTypeImagesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithBmpLanguage(bmpLanguage *string) *DescribeDeviceTypeImagesParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithBmpUserID(bmpUserID *string) *DescribeDeviceTypeImagesParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithDeviceTypeID adds the deviceTypeID to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithDeviceTypeID(deviceTypeID *string) *DescribeDeviceTypeImagesParams {
	o.SetDeviceTypeID(deviceTypeID)
	return o
}

// SetDeviceTypeID adds the deviceTypeId to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetDeviceTypeID(deviceTypeID *string) {
	o.DeviceTypeID = deviceTypeID
}

// WithImageID adds the imageID to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithImageID(imageID *string) *DescribeDeviceTypeImagesParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetImageID(imageID *string) {
	o.ImageID = imageID
}

// WithImageName adds the imageName to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithImageName(imageName *string) *DescribeDeviceTypeImagesParams {
	o.SetImageName(imageName)
	return o
}

// SetImageName adds the imageName to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetImageName(imageName *string) {
	o.ImageName = imageName
}

// WithIsAll adds the isAll to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithIsAll(isAll *string) *DescribeDeviceTypeImagesParams {
	o.SetIsAll(isAll)
	return o
}

// SetIsAll adds the isAll to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetIsAll(isAll *string) {
	o.IsAll = isAll
}

// WithOsID adds the osID to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithOsID(osID *string) *DescribeDeviceTypeImagesParams {
	o.SetOsID(osID)
	return o
}

// SetOsID adds the osId to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetOsID(osID *string) {
	o.OsID = osID
}

// WithOsType adds the osType to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithOsType(osType *string) *DescribeDeviceTypeImagesParams {
	o.SetOsType(osType)
	return o
}

// SetOsType adds the osType to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetOsType(osType *string) {
	o.OsType = osType
}

// WithPageNumber adds the pageNumber to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithPageNumber(pageNumber *int64) *DescribeDeviceTypeImagesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetPageNumber(pageNumber *int64) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithPageSize(pageSize *int64) *DescribeDeviceTypeImagesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSource adds the source to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithSource(source *string) *DescribeDeviceTypeImagesParams {
	o.SetSource(source)
	return o
}

// SetSource adds the source to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetSource(source *string) {
	o.Source = source
}

// WithTraceID adds the traceID to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithTraceID(traceID string) *DescribeDeviceTypeImagesParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WithVersion adds the version to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) WithVersion(version *string) *DescribeDeviceTypeImagesParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the describe device type images params
func (o *DescribeDeviceTypeImagesParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeDeviceTypeImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Architecture != nil {

		// query param architecture
		var qrArchitecture string

		if o.Architecture != nil {
			qrArchitecture = *o.Architecture
		}
		qArchitecture := qrArchitecture
		if qArchitecture != "" {

			if err := r.SetQueryParam("architecture", qArchitecture); err != nil {
				return err
			}
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

	if o.ImageID != nil {

		// query param imageId
		var qrImageID string

		if o.ImageID != nil {
			qrImageID = *o.ImageID
		}
		qImageID := qrImageID
		if qImageID != "" {

			if err := r.SetQueryParam("imageId", qImageID); err != nil {
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

	if o.OsID != nil {

		// query param osId
		var qrOsID string

		if o.OsID != nil {
			qrOsID = *o.OsID
		}
		qOsID := qrOsID
		if qOsID != "" {

			if err := r.SetQueryParam("osId", qOsID); err != nil {
				return err
			}
		}
	}

	if o.OsType != nil {

		// query param osType
		var qrOsType string

		if o.OsType != nil {
			qrOsType = *o.OsType
		}
		qOsType := qrOsType
		if qOsType != "" {

			if err := r.SetQueryParam("osType", qOsType); err != nil {
				return err
			}
		}
	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int64

		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt64(qrPageNumber)
		if qPageNumber != "" {

			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.Source != nil {

		// query param source
		var qrSource string

		if o.Source != nil {
			qrSource = *o.Source
		}
		qSource := qrSource
		if qSource != "" {

			if err := r.SetQueryParam("source", qSource); err != nil {
				return err
			}
		}
	}

	// header param traceId
	if err := r.SetHeaderParam("traceId", o.TraceID); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}