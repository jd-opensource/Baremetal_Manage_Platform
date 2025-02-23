// Code generated by go-swagger; DO NOT EDIT.

package message

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

// NewGetMessageListParams creates a new GetMessageListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMessageListParams() *GetMessageListParams {
	return &GetMessageListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMessageListParamsWithTimeout creates a new GetMessageListParams object
// with the ability to set a timeout on a request.
func NewGetMessageListParamsWithTimeout(timeout time.Duration) *GetMessageListParams {
	return &GetMessageListParams{
		timeout: timeout,
	}
}

// NewGetMessageListParamsWithContext creates a new GetMessageListParams object
// with the ability to set a context for a request.
func NewGetMessageListParamsWithContext(ctx context.Context) *GetMessageListParams {
	return &GetMessageListParams{
		Context: ctx,
	}
}

// NewGetMessageListParamsWithHTTPClient creates a new GetMessageListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMessageListParamsWithHTTPClient(client *http.Client) *GetMessageListParams {
	return &GetMessageListParams{
		HTTPClient: client,
	}
}

/* GetMessageListParams contains all the parameters to send to the API endpoint
   for the get message list operation.

   Typically these are written to a http.Request.
*/
type GetMessageListParams struct {

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

	/* Detail.

	   内容模糊搜索
	*/
	Detail *string

	/* HasRead.

	   0：未读；1：已读；""不传表示全部
	*/
	HasRead *string

	/* IsAll.

	   是否显示全部，取值为1时表示全部
	*/
	IsAll *string

	/* MessageSubType.

	   消息子类型
	*/
	MessageSubType *string

	/* MessageType.

	   消息类型
	*/
	MessageType *string

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

	/* TraceID.

	   流量唯一id
	*/
	TraceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get message list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMessageListParams) WithDefaults() *GetMessageListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get message list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMessageListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get message list params
func (o *GetMessageListParams) WithTimeout(timeout time.Duration) *GetMessageListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get message list params
func (o *GetMessageListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get message list params
func (o *GetMessageListParams) WithContext(ctx context.Context) *GetMessageListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get message list params
func (o *GetMessageListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get message list params
func (o *GetMessageListParams) WithHTTPClient(client *http.Client) *GetMessageListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get message list params
func (o *GetMessageListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get message list params
func (o *GetMessageListParams) WithAuthorization(authorization string) *GetMessageListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get message list params
func (o *GetMessageListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the get message list params
func (o *GetMessageListParams) WithBmpLanguage(bmpLanguage *string) *GetMessageListParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the get message list params
func (o *GetMessageListParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the get message list params
func (o *GetMessageListParams) WithBmpUserID(bmpUserID *string) *GetMessageListParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the get message list params
func (o *GetMessageListParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithDetail adds the detail to the get message list params
func (o *GetMessageListParams) WithDetail(detail *string) *GetMessageListParams {
	o.SetDetail(detail)
	return o
}

// SetDetail adds the detail to the get message list params
func (o *GetMessageListParams) SetDetail(detail *string) {
	o.Detail = detail
}

// WithHasRead adds the hasRead to the get message list params
func (o *GetMessageListParams) WithHasRead(hasRead *string) *GetMessageListParams {
	o.SetHasRead(hasRead)
	return o
}

// SetHasRead adds the hasRead to the get message list params
func (o *GetMessageListParams) SetHasRead(hasRead *string) {
	o.HasRead = hasRead
}

// WithIsAll adds the isAll to the get message list params
func (o *GetMessageListParams) WithIsAll(isAll *string) *GetMessageListParams {
	o.SetIsAll(isAll)
	return o
}

// SetIsAll adds the isAll to the get message list params
func (o *GetMessageListParams) SetIsAll(isAll *string) {
	o.IsAll = isAll
}

// WithMessageSubType adds the messageSubType to the get message list params
func (o *GetMessageListParams) WithMessageSubType(messageSubType *string) *GetMessageListParams {
	o.SetMessageSubType(messageSubType)
	return o
}

// SetMessageSubType adds the messageSubType to the get message list params
func (o *GetMessageListParams) SetMessageSubType(messageSubType *string) {
	o.MessageSubType = messageSubType
}

// WithMessageType adds the messageType to the get message list params
func (o *GetMessageListParams) WithMessageType(messageType *string) *GetMessageListParams {
	o.SetMessageType(messageType)
	return o
}

// SetMessageType adds the messageType to the get message list params
func (o *GetMessageListParams) SetMessageType(messageType *string) {
	o.MessageType = messageType
}

// WithPageNumber adds the pageNumber to the get message list params
func (o *GetMessageListParams) WithPageNumber(pageNumber *int64) *GetMessageListParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get message list params
func (o *GetMessageListParams) SetPageNumber(pageNumber *int64) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get message list params
func (o *GetMessageListParams) WithPageSize(pageSize *int64) *GetMessageListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get message list params
func (o *GetMessageListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithTraceID adds the traceID to the get message list params
func (o *GetMessageListParams) WithTraceID(traceID string) *GetMessageListParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the get message list params
func (o *GetMessageListParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMessageListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Detail != nil {

		// query param detail
		var qrDetail string

		if o.Detail != nil {
			qrDetail = *o.Detail
		}
		qDetail := qrDetail
		if qDetail != "" {

			if err := r.SetQueryParam("detail", qDetail); err != nil {
				return err
			}
		}
	}

	if o.HasRead != nil {

		// query param hasRead
		var qrHasRead string

		if o.HasRead != nil {
			qrHasRead = *o.HasRead
		}
		qHasRead := qrHasRead
		if qHasRead != "" {

			if err := r.SetQueryParam("hasRead", qHasRead); err != nil {
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

	if o.MessageSubType != nil {

		// query param messageSubType
		var qrMessageSubType string

		if o.MessageSubType != nil {
			qrMessageSubType = *o.MessageSubType
		}
		qMessageSubType := qrMessageSubType
		if qMessageSubType != "" {

			if err := r.SetQueryParam("messageSubType", qMessageSubType); err != nil {
				return err
			}
		}
	}

	if o.MessageType != nil {

		// query param messageType
		var qrMessageType string

		if o.MessageType != nil {
			qrMessageType = *o.MessageType
		}
		qMessageType := qrMessageType
		if qMessageType != "" {

			if err := r.SetQueryParam("messageType", qMessageType); err != nil {
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

	// header param traceId
	if err := r.SetHeaderParam("traceId", o.TraceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
