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
	"github.com/go-openapi/swag"
)

// NewDescribeUsersParams creates a new DescribeUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeUsersParams() *DescribeUsersParams {
	return &DescribeUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeUsersParamsWithTimeout creates a new DescribeUsersParams object
// with the ability to set a timeout on a request.
func NewDescribeUsersParamsWithTimeout(timeout time.Duration) *DescribeUsersParams {
	return &DescribeUsersParams{
		timeout: timeout,
	}
}

// NewDescribeUsersParamsWithContext creates a new DescribeUsersParams object
// with the ability to set a context for a request.
func NewDescribeUsersParamsWithContext(ctx context.Context) *DescribeUsersParams {
	return &DescribeUsersParams{
		Context: ctx,
	}
}

// NewDescribeUsersParamsWithHTTPClient creates a new DescribeUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeUsersParamsWithHTTPClient(client *http.Client) *DescribeUsersParams {
	return &DescribeUsersParams{
		HTTPClient: client,
	}
}

/* DescribeUsersParams contains all the parameters to send to the API endpoint
   for the describe users operation.

   Typically these are written to a http.Request.
*/
type DescribeUsersParams struct {

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

	/* DefaultProjectID.

	   项目uuid
	*/
	DefaultProjectID *string

	/* IsAll.

	   是否显示全部, isAll=1表示全部
	*/
	IsAll *string

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

	/* RoleID.

	   角色uuid
	*/
	RoleID *string

	/* TraceID.

	   流量唯一id
	*/
	TraceID string

	/* UserName.

	   用户名
	*/
	UserName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeUsersParams) WithDefaults() *DescribeUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe users params
func (o *DescribeUsersParams) WithTimeout(timeout time.Duration) *DescribeUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe users params
func (o *DescribeUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe users params
func (o *DescribeUsersParams) WithContext(ctx context.Context) *DescribeUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe users params
func (o *DescribeUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe users params
func (o *DescribeUsersParams) WithHTTPClient(client *http.Client) *DescribeUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe users params
func (o *DescribeUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the describe users params
func (o *DescribeUsersParams) WithAuthorization(authorization string) *DescribeUsersParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the describe users params
func (o *DescribeUsersParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the describe users params
func (o *DescribeUsersParams) WithBmpLanguage(bmpLanguage *string) *DescribeUsersParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the describe users params
func (o *DescribeUsersParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the describe users params
func (o *DescribeUsersParams) WithBmpUserID(bmpUserID *string) *DescribeUsersParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the describe users params
func (o *DescribeUsersParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithDefaultProjectID adds the defaultProjectID to the describe users params
func (o *DescribeUsersParams) WithDefaultProjectID(defaultProjectID *string) *DescribeUsersParams {
	o.SetDefaultProjectID(defaultProjectID)
	return o
}

// SetDefaultProjectID adds the defaultProjectId to the describe users params
func (o *DescribeUsersParams) SetDefaultProjectID(defaultProjectID *string) {
	o.DefaultProjectID = defaultProjectID
}

// WithIsAll adds the isAll to the describe users params
func (o *DescribeUsersParams) WithIsAll(isAll *string) *DescribeUsersParams {
	o.SetIsAll(isAll)
	return o
}

// SetIsAll adds the isAll to the describe users params
func (o *DescribeUsersParams) SetIsAll(isAll *string) {
	o.IsAll = isAll
}

// WithPageNumber adds the pageNumber to the describe users params
func (o *DescribeUsersParams) WithPageNumber(pageNumber *int64) *DescribeUsersParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the describe users params
func (o *DescribeUsersParams) SetPageNumber(pageNumber *int64) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the describe users params
func (o *DescribeUsersParams) WithPageSize(pageSize *int64) *DescribeUsersParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the describe users params
func (o *DescribeUsersParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithRoleID adds the roleID to the describe users params
func (o *DescribeUsersParams) WithRoleID(roleID *string) *DescribeUsersParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the describe users params
func (o *DescribeUsersParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithTraceID adds the traceID to the describe users params
func (o *DescribeUsersParams) WithTraceID(traceID string) *DescribeUsersParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the describe users params
func (o *DescribeUsersParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WithUserName adds the userName to the describe users params
func (o *DescribeUsersParams) WithUserName(userName *string) *DescribeUsersParams {
	o.SetUserName(userName)
	return o
}

// SetUserName adds the userName to the describe users params
func (o *DescribeUsersParams) SetUserName(userName *string) {
	o.UserName = userName
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DefaultProjectID != nil {

		// query param defaultProjectId
		var qrDefaultProjectID string

		if o.DefaultProjectID != nil {
			qrDefaultProjectID = *o.DefaultProjectID
		}
		qDefaultProjectID := qrDefaultProjectID
		if qDefaultProjectID != "" {

			if err := r.SetQueryParam("defaultProjectId", qDefaultProjectID); err != nil {
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

	if o.RoleID != nil {

		// query param roleId
		var qrRoleID string

		if o.RoleID != nil {
			qrRoleID = *o.RoleID
		}
		qRoleID := qrRoleID
		if qRoleID != "" {

			if err := r.SetQueryParam("roleId", qRoleID); err != nil {
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
