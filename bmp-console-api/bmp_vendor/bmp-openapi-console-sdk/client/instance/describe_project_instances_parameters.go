// Code generated by go-swagger; DO NOT EDIT.

package instance

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

// NewDescribeProjectInstancesParams creates a new DescribeProjectInstancesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeProjectInstancesParams() *DescribeProjectInstancesParams {
	return &DescribeProjectInstancesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeProjectInstancesParamsWithTimeout creates a new DescribeProjectInstancesParams object
// with the ability to set a timeout on a request.
func NewDescribeProjectInstancesParamsWithTimeout(timeout time.Duration) *DescribeProjectInstancesParams {
	return &DescribeProjectInstancesParams{
		timeout: timeout,
	}
}

// NewDescribeProjectInstancesParamsWithContext creates a new DescribeProjectInstancesParams object
// with the ability to set a context for a request.
func NewDescribeProjectInstancesParamsWithContext(ctx context.Context) *DescribeProjectInstancesParams {
	return &DescribeProjectInstancesParams{
		Context: ctx,
	}
}

// NewDescribeProjectInstancesParamsWithHTTPClient creates a new DescribeProjectInstancesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeProjectInstancesParamsWithHTTPClient(client *http.Client) *DescribeProjectInstancesParams {
	return &DescribeProjectInstancesParams{
		HTTPClient: client,
	}
}

/* DescribeProjectInstancesParams contains all the parameters to send to the API endpoint
   for the describe project instances operation.

   Typically these are written to a http.Request.
*/
type DescribeProjectInstancesParams struct {

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

	/* DeviceID.

	   设备ID
	*/
	DeviceID *string

	/* DeviceTypeID.

	   设备类型ID
	*/
	DeviceTypeID *string

	/* IdcID.

	   机房ID
	*/
	IdcID *string

	/* IloIP.

	   带外ip，精确搜索
	*/
	IloIP *string

	/* InstanceID.

	   实例ID
	*/
	InstanceID *string

	/* InstanceName.

	   实例名称,模糊搜索
	*/
	InstanceName *string

	/* IPV4.

	   ipv4，精确搜索
	*/
	IPV4 *string

	/* IPV6.

	   ipv6，精确搜索
	*/
	IPV6 *string

	/* IsAll.

	   是否显示全部，取值为1时表示全部
	*/
	IsAll *string

	/* IsInstallAgent.

	   是否安装了agent，"0"表示未安装，"1"表示已安装
	*/
	IsInstallAgent *string

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

	/* ProjectID.

	   项目ID
	*/
	ProjectID *string

	/* Sn.

	   SN
	*/
	Sn *string

	/* Status.

	   运行状态
	*/
	Status *string

	/* TraceID.

	   流量唯一id
	*/
	TraceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe project instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeProjectInstancesParams) WithDefaults() *DescribeProjectInstancesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe project instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeProjectInstancesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe project instances params
func (o *DescribeProjectInstancesParams) WithTimeout(timeout time.Duration) *DescribeProjectInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe project instances params
func (o *DescribeProjectInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe project instances params
func (o *DescribeProjectInstancesParams) WithContext(ctx context.Context) *DescribeProjectInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe project instances params
func (o *DescribeProjectInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe project instances params
func (o *DescribeProjectInstancesParams) WithHTTPClient(client *http.Client) *DescribeProjectInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe project instances params
func (o *DescribeProjectInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the describe project instances params
func (o *DescribeProjectInstancesParams) WithAuthorization(authorization string) *DescribeProjectInstancesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the describe project instances params
func (o *DescribeProjectInstancesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the describe project instances params
func (o *DescribeProjectInstancesParams) WithBmpLanguage(bmpLanguage *string) *DescribeProjectInstancesParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the describe project instances params
func (o *DescribeProjectInstancesParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the describe project instances params
func (o *DescribeProjectInstancesParams) WithBmpUserID(bmpUserID *string) *DescribeProjectInstancesParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the describe project instances params
func (o *DescribeProjectInstancesParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithDeviceID adds the deviceID to the describe project instances params
func (o *DescribeProjectInstancesParams) WithDeviceID(deviceID *string) *DescribeProjectInstancesParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the describe project instances params
func (o *DescribeProjectInstancesParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithDeviceTypeID adds the deviceTypeID to the describe project instances params
func (o *DescribeProjectInstancesParams) WithDeviceTypeID(deviceTypeID *string) *DescribeProjectInstancesParams {
	o.SetDeviceTypeID(deviceTypeID)
	return o
}

// SetDeviceTypeID adds the deviceTypeId to the describe project instances params
func (o *DescribeProjectInstancesParams) SetDeviceTypeID(deviceTypeID *string) {
	o.DeviceTypeID = deviceTypeID
}

// WithIdcID adds the idcID to the describe project instances params
func (o *DescribeProjectInstancesParams) WithIdcID(idcID *string) *DescribeProjectInstancesParams {
	o.SetIdcID(idcID)
	return o
}

// SetIdcID adds the idcId to the describe project instances params
func (o *DescribeProjectInstancesParams) SetIdcID(idcID *string) {
	o.IdcID = idcID
}

// WithIloIP adds the iloIP to the describe project instances params
func (o *DescribeProjectInstancesParams) WithIloIP(iloIP *string) *DescribeProjectInstancesParams {
	o.SetIloIP(iloIP)
	return o
}

// SetIloIP adds the iloIp to the describe project instances params
func (o *DescribeProjectInstancesParams) SetIloIP(iloIP *string) {
	o.IloIP = iloIP
}

// WithInstanceID adds the instanceID to the describe project instances params
func (o *DescribeProjectInstancesParams) WithInstanceID(instanceID *string) *DescribeProjectInstancesParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the describe project instances params
func (o *DescribeProjectInstancesParams) SetInstanceID(instanceID *string) {
	o.InstanceID = instanceID
}

// WithInstanceName adds the instanceName to the describe project instances params
func (o *DescribeProjectInstancesParams) WithInstanceName(instanceName *string) *DescribeProjectInstancesParams {
	o.SetInstanceName(instanceName)
	return o
}

// SetInstanceName adds the instanceName to the describe project instances params
func (o *DescribeProjectInstancesParams) SetInstanceName(instanceName *string) {
	o.InstanceName = instanceName
}

// WithIPV4 adds the iPV4 to the describe project instances params
func (o *DescribeProjectInstancesParams) WithIPV4(iPV4 *string) *DescribeProjectInstancesParams {
	o.SetIPV4(iPV4)
	return o
}

// SetIPV4 adds the ipv4 to the describe project instances params
func (o *DescribeProjectInstancesParams) SetIPV4(iPV4 *string) {
	o.IPV4 = iPV4
}

// WithIPV6 adds the iPV6 to the describe project instances params
func (o *DescribeProjectInstancesParams) WithIPV6(iPV6 *string) *DescribeProjectInstancesParams {
	o.SetIPV6(iPV6)
	return o
}

// SetIPV6 adds the ipv6 to the describe project instances params
func (o *DescribeProjectInstancesParams) SetIPV6(iPV6 *string) {
	o.IPV6 = iPV6
}

// WithIsAll adds the isAll to the describe project instances params
func (o *DescribeProjectInstancesParams) WithIsAll(isAll *string) *DescribeProjectInstancesParams {
	o.SetIsAll(isAll)
	return o
}

// SetIsAll adds the isAll to the describe project instances params
func (o *DescribeProjectInstancesParams) SetIsAll(isAll *string) {
	o.IsAll = isAll
}

// WithIsInstallAgent adds the isInstallAgent to the describe project instances params
func (o *DescribeProjectInstancesParams) WithIsInstallAgent(isInstallAgent *string) *DescribeProjectInstancesParams {
	o.SetIsInstallAgent(isInstallAgent)
	return o
}

// SetIsInstallAgent adds the isInstallAgent to the describe project instances params
func (o *DescribeProjectInstancesParams) SetIsInstallAgent(isInstallAgent *string) {
	o.IsInstallAgent = isInstallAgent
}

// WithPageNumber adds the pageNumber to the describe project instances params
func (o *DescribeProjectInstancesParams) WithPageNumber(pageNumber *int64) *DescribeProjectInstancesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the describe project instances params
func (o *DescribeProjectInstancesParams) SetPageNumber(pageNumber *int64) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the describe project instances params
func (o *DescribeProjectInstancesParams) WithPageSize(pageSize *int64) *DescribeProjectInstancesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the describe project instances params
func (o *DescribeProjectInstancesParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithProjectID adds the projectID to the describe project instances params
func (o *DescribeProjectInstancesParams) WithProjectID(projectID *string) *DescribeProjectInstancesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the describe project instances params
func (o *DescribeProjectInstancesParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithSn adds the sn to the describe project instances params
func (o *DescribeProjectInstancesParams) WithSn(sn *string) *DescribeProjectInstancesParams {
	o.SetSn(sn)
	return o
}

// SetSn adds the sn to the describe project instances params
func (o *DescribeProjectInstancesParams) SetSn(sn *string) {
	o.Sn = sn
}

// WithStatus adds the status to the describe project instances params
func (o *DescribeProjectInstancesParams) WithStatus(status *string) *DescribeProjectInstancesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe project instances params
func (o *DescribeProjectInstancesParams) SetStatus(status *string) {
	o.Status = status
}

// WithTraceID adds the traceID to the describe project instances params
func (o *DescribeProjectInstancesParams) WithTraceID(traceID string) *DescribeProjectInstancesParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the describe project instances params
func (o *DescribeProjectInstancesParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeProjectInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DeviceID != nil {

		// query param deviceId
		var qrDeviceID string

		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {

			if err := r.SetQueryParam("deviceId", qDeviceID); err != nil {
				return err
			}
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

	if o.IdcID != nil {

		// query param idcId
		var qrIdcID string

		if o.IdcID != nil {
			qrIdcID = *o.IdcID
		}
		qIdcID := qrIdcID
		if qIdcID != "" {

			if err := r.SetQueryParam("idcId", qIdcID); err != nil {
				return err
			}
		}
	}

	if o.IloIP != nil {

		// query param ilo_ip
		var qrIloIP string

		if o.IloIP != nil {
			qrIloIP = *o.IloIP
		}
		qIloIP := qrIloIP
		if qIloIP != "" {

			if err := r.SetQueryParam("ilo_ip", qIloIP); err != nil {
				return err
			}
		}
	}

	if o.InstanceID != nil {

		// query param instanceId
		var qrInstanceID string

		if o.InstanceID != nil {
			qrInstanceID = *o.InstanceID
		}
		qInstanceID := qrInstanceID
		if qInstanceID != "" {

			if err := r.SetQueryParam("instanceId", qInstanceID); err != nil {
				return err
			}
		}
	}

	if o.InstanceName != nil {

		// query param instanceName
		var qrInstanceName string

		if o.InstanceName != nil {
			qrInstanceName = *o.InstanceName
		}
		qInstanceName := qrInstanceName
		if qInstanceName != "" {

			if err := r.SetQueryParam("instanceName", qInstanceName); err != nil {
				return err
			}
		}
	}

	if o.IPV4 != nil {

		// query param ipv4
		var qrIPV4 string

		if o.IPV4 != nil {
			qrIPV4 = *o.IPV4
		}
		qIPV4 := qrIPV4
		if qIPV4 != "" {

			if err := r.SetQueryParam("ipv4", qIPV4); err != nil {
				return err
			}
		}
	}

	if o.IPV6 != nil {

		// query param ipv6
		var qrIPV6 string

		if o.IPV6 != nil {
			qrIPV6 = *o.IPV6
		}
		qIPV6 := qrIPV6
		if qIPV6 != "" {

			if err := r.SetQueryParam("ipv6", qIPV6); err != nil {
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

	if o.IsInstallAgent != nil {

		// query param isInstallAgent
		var qrIsInstallAgent string

		if o.IsInstallAgent != nil {
			qrIsInstallAgent = *o.IsInstallAgent
		}
		qIsInstallAgent := qrIsInstallAgent
		if qIsInstallAgent != "" {

			if err := r.SetQueryParam("isInstallAgent", qIsInstallAgent); err != nil {
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

	if o.ProjectID != nil {

		// query param projectId
		var qrProjectID string

		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := qrProjectID
		if qProjectID != "" {

			if err := r.SetQueryParam("projectId", qProjectID); err != nil {
				return err
			}
		}
	}

	if o.Sn != nil {

		// query param sn
		var qrSn string

		if o.Sn != nil {
			qrSn = *o.Sn
		}
		qSn := qrSn
		if qSn != "" {

			if err := r.SetQueryParam("sn", qSn); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
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