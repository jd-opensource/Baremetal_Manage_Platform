// Code generated by go-swagger; DO NOT EDIT.

package monitor_rule

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

// NewDescribeRulesParams creates a new DescribeRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeRulesParams() *DescribeRulesParams {
	return &DescribeRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeRulesParamsWithTimeout creates a new DescribeRulesParams object
// with the ability to set a timeout on a request.
func NewDescribeRulesParamsWithTimeout(timeout time.Duration) *DescribeRulesParams {
	return &DescribeRulesParams{
		timeout: timeout,
	}
}

// NewDescribeRulesParamsWithContext creates a new DescribeRulesParams object
// with the ability to set a context for a request.
func NewDescribeRulesParamsWithContext(ctx context.Context) *DescribeRulesParams {
	return &DescribeRulesParams{
		Context: ctx,
	}
}

// NewDescribeRulesParamsWithHTTPClient creates a new DescribeRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeRulesParamsWithHTTPClient(client *http.Client) *DescribeRulesParams {
	return &DescribeRulesParams{
		HTTPClient: client,
	}
}

/* DescribeRulesParams contains all the parameters to send to the API endpoint
   for the describe rules operation.

   Typically these are written to a http.Request.
*/
type DescribeRulesParams struct {

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

	/* IsAll.

	   是否显示全部，取值为1时表示全部
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

	/* RuleID.

	   规则uuid
	*/
	RuleID *string

	/* RuleName.

	   规则名称
	*/
	RuleName *string

	/* Status.

	   1 正常,2 禁用,3 报警

	   Format: int64
	*/
	Status *int64

	/* TraceID.

	   流量唯一id
	*/
	TraceID string

	/* UserID.

	   user uuid
	*/
	UserID *string

	/* UserName.

	   username
	*/
	UserName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeRulesParams) WithDefaults() *DescribeRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe rules params
func (o *DescribeRulesParams) WithTimeout(timeout time.Duration) *DescribeRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe rules params
func (o *DescribeRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe rules params
func (o *DescribeRulesParams) WithContext(ctx context.Context) *DescribeRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe rules params
func (o *DescribeRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe rules params
func (o *DescribeRulesParams) WithHTTPClient(client *http.Client) *DescribeRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe rules params
func (o *DescribeRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the describe rules params
func (o *DescribeRulesParams) WithAuthorization(authorization string) *DescribeRulesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the describe rules params
func (o *DescribeRulesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBmpLanguage adds the bmpLanguage to the describe rules params
func (o *DescribeRulesParams) WithBmpLanguage(bmpLanguage *string) *DescribeRulesParams {
	o.SetBmpLanguage(bmpLanguage)
	return o
}

// SetBmpLanguage adds the bmpLanguage to the describe rules params
func (o *DescribeRulesParams) SetBmpLanguage(bmpLanguage *string) {
	o.BmpLanguage = bmpLanguage
}

// WithBmpUserID adds the bmpUserID to the describe rules params
func (o *DescribeRulesParams) WithBmpUserID(bmpUserID *string) *DescribeRulesParams {
	o.SetBmpUserID(bmpUserID)
	return o
}

// SetBmpUserID adds the bmpUserId to the describe rules params
func (o *DescribeRulesParams) SetBmpUserID(bmpUserID *string) {
	o.BmpUserID = bmpUserID
}

// WithIsAll adds the isAll to the describe rules params
func (o *DescribeRulesParams) WithIsAll(isAll *string) *DescribeRulesParams {
	o.SetIsAll(isAll)
	return o
}

// SetIsAll adds the isAll to the describe rules params
func (o *DescribeRulesParams) SetIsAll(isAll *string) {
	o.IsAll = isAll
}

// WithPageNumber adds the pageNumber to the describe rules params
func (o *DescribeRulesParams) WithPageNumber(pageNumber *int64) *DescribeRulesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the describe rules params
func (o *DescribeRulesParams) SetPageNumber(pageNumber *int64) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the describe rules params
func (o *DescribeRulesParams) WithPageSize(pageSize *int64) *DescribeRulesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the describe rules params
func (o *DescribeRulesParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithRuleID adds the ruleID to the describe rules params
func (o *DescribeRulesParams) WithRuleID(ruleID *string) *DescribeRulesParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the describe rules params
func (o *DescribeRulesParams) SetRuleID(ruleID *string) {
	o.RuleID = ruleID
}

// WithRuleName adds the ruleName to the describe rules params
func (o *DescribeRulesParams) WithRuleName(ruleName *string) *DescribeRulesParams {
	o.SetRuleName(ruleName)
	return o
}

// SetRuleName adds the ruleName to the describe rules params
func (o *DescribeRulesParams) SetRuleName(ruleName *string) {
	o.RuleName = ruleName
}

// WithStatus adds the status to the describe rules params
func (o *DescribeRulesParams) WithStatus(status *int64) *DescribeRulesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe rules params
func (o *DescribeRulesParams) SetStatus(status *int64) {
	o.Status = status
}

// WithTraceID adds the traceID to the describe rules params
func (o *DescribeRulesParams) WithTraceID(traceID string) *DescribeRulesParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the describe rules params
func (o *DescribeRulesParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WithUserID adds the userID to the describe rules params
func (o *DescribeRulesParams) WithUserID(userID *string) *DescribeRulesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the describe rules params
func (o *DescribeRulesParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WithUserName adds the userName to the describe rules params
func (o *DescribeRulesParams) WithUserName(userName *string) *DescribeRulesParams {
	o.SetUserName(userName)
	return o
}

// SetUserName adds the userName to the describe rules params
func (o *DescribeRulesParams) SetUserName(userName *string) {
	o.UserName = userName
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.RuleID != nil {

		// query param ruleId
		var qrRuleID string

		if o.RuleID != nil {
			qrRuleID = *o.RuleID
		}
		qRuleID := qrRuleID
		if qRuleID != "" {

			if err := r.SetQueryParam("ruleId", qRuleID); err != nil {
				return err
			}
		}
	}

	if o.RuleName != nil {

		// query param ruleName
		var qrRuleName string

		if o.RuleName != nil {
			qrRuleName = *o.RuleName
		}
		qRuleName := qrRuleName
		if qRuleName != "" {

			if err := r.SetQueryParam("ruleName", qRuleName); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus int64

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := swag.FormatInt64(qrStatus)
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

	if o.UserID != nil {

		// query param userId
		var qrUserID string

		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {

			if err := r.SetQueryParam("userId", qUserID); err != nil {
				return err
			}
		}
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
