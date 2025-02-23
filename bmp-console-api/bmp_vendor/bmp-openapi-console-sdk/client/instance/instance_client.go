// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new instance API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for instance API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateProjectInstance(params *CreateProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateProjectInstanceOK, error)

	DeleteInstances(params *DeleteInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInstancesOK, error)

	DeleteProjectInstance(params *DeleteProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteProjectInstanceOK, error)

	DescribeInstancesByProjectIDAndOwnerNameAndSharerName(params *DescribeInstancesByProjectIDAndOwnerNameAndSharerNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeInstancesByProjectIDAndOwnerNameAndSharerNameOK, error)

	DescribeProjectInstance(params *DescribeProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeProjectInstanceOK, error)

	DescribeProjectInstances(params *DescribeProjectInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeProjectInstancesOK, error)

	LockProjectInstance(params *LockProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LockProjectInstanceOK, error)

	ModifyInstances(params *ModifyInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyInstancesOK, error)

	ModifyProjectInstance(params *ModifyProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyProjectInstanceOK, error)

	ReinstallProjectInstance(params *ReinstallProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReinstallProjectInstanceOK, error)

	ResetInstanceStatus(params *ResetInstanceStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetInstanceStatusOK, error)

	ResetProjectInstancePasswd(params *ResetProjectInstancePasswdParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetProjectInstancePasswdOK, error)

	ResetProjectInstancesPasswd(params *ResetProjectInstancesPasswdParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetProjectInstancesPasswdOK, error)

	RestartInstances(params *RestartInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestartInstancesOK, error)

	RestartProjectInstance(params *RestartProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestartProjectInstanceOK, error)

	StartInstances(params *StartInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StartInstancesOK, error)

	StartProjectInstance(params *StartProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StartProjectInstanceOK, error)

	StopInstances(params *StopInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopInstancesOK, error)

	StopProjectInstance(params *StopProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopProjectInstanceOK, error)

	UnLockProjectInstance(params *UnLockProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnLockProjectInstanceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateProjectInstance CreateProjectInstance 创建实例
*/
func (a *Client) CreateProjectInstance(params *CreateProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateProjectInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProjectInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createProjectInstance",
		Method:             "POST",
		PathPattern:        "/project/instances",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateProjectInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateProjectInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateProjectInstanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteInstances DeleteInstances 批量删除实例
*/
func (a *Client) DeleteInstances(params *DeleteInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInstancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteInstances",
		Method:             "DELETE",
		PathPattern:        "/project/instances/batch/instances:deleteInstances",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteInstancesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteInstancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteInstancesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteProjectInstance DeleteProjectInstance 删除实例
*/
func (a *Client) DeleteProjectInstance(params *DeleteProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteProjectInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProjectInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteProjectInstance",
		Method:             "DELETE",
		PathPattern:        "/project/instances/{instance_id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteProjectInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteProjectInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteProjectInstanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DescribeInstancesByProjectIDAndOwnerNameAndSharerName DescribeInstancesByProjectIdAndOwnerNameAndSharerName 根据projectid，ownername，username来获取实例列表。为了支持部分资源转移中的实例列表
*/
func (a *Client) DescribeInstancesByProjectIDAndOwnerNameAndSharerName(params *DescribeInstancesByProjectIDAndOwnerNameAndSharerNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeInstancesByProjectIDAndOwnerNameAndSharerNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeInstancesByProjectIDAndOwnerNameAndSharerNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "describeInstancesByProjectIdAndOwnerNameAndSharerName",
		Method:             "GET",
		PathPattern:        "/project/instances/share/describeInstances",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeInstancesByProjectIDAndOwnerNameAndSharerNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DescribeInstancesByProjectIDAndOwnerNameAndSharerNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DescribeInstancesByProjectIDAndOwnerNameAndSharerNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DescribeProjectInstance DescribeProjectInstance 获取实例详情
*/
func (a *Client) DescribeProjectInstance(params *DescribeProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeProjectInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeProjectInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "describeProjectInstance",
		Method:             "GET",
		PathPattern:        "/project/instances/{instance_id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeProjectInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DescribeProjectInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DescribeProjectInstanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DescribeProjectInstances DescribeProjectInstances 获取实例列表
*/
func (a *Client) DescribeProjectInstances(params *DescribeProjectInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeProjectInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeProjectInstancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "describeProjectInstances",
		Method:             "GET",
		PathPattern:        "/project/instances",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeProjectInstancesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DescribeProjectInstancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DescribeProjectInstancesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  LockProjectInstance LockProjectInstance 锁定实例
*/
func (a *Client) LockProjectInstance(params *LockProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LockProjectInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLockProjectInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "lockProjectInstance",
		Method:             "PUT",
		PathPattern:        "/project/instances/{instance_id}/lock",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LockProjectInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LockProjectInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LockProjectInstanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ModifyInstances ModifyInstances 批量修改实例名称
*/
func (a *Client) ModifyInstances(params *ModifyInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyInstancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "modifyInstances",
		Method:             "PUT",
		PathPattern:        "/project/instances/batch/instances:modifyInstances",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyInstancesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ModifyInstancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ModifyInstancesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ModifyProjectInstance ModifyProjectInstance 修改实例信息
*/
func (a *Client) ModifyProjectInstance(params *ModifyProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyProjectInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyProjectInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "modifyProjectInstance",
		Method:             "PUT",
		PathPattern:        "/project/instances/{instance_id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyProjectInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ModifyProjectInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ModifyProjectInstanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReinstallProjectInstance ReinstallInstance 重装实例
*/
func (a *Client) ReinstallProjectInstance(params *ReinstallProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReinstallProjectInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReinstallProjectInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "reinstallProjectInstance",
		Method:             "PUT",
		PathPattern:        "/project/instances/{instance_id}/reinstallInstance",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReinstallProjectInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReinstallProjectInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReinstallProjectInstanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ResetInstanceStatus ResetInstanceStatus 重置实例状态
*/
func (a *Client) ResetInstanceStatus(params *ResetInstanceStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetInstanceStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetInstanceStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "resetInstanceStatus",
		Method:             "PUT",
		PathPattern:        "/project/instances/{instance_id}/resetStatus",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ResetInstanceStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResetInstanceStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ResetInstanceStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ResetProjectInstancePasswd ResetPasswd 重置密码
*/
func (a *Client) ResetProjectInstancePasswd(params *ResetProjectInstancePasswdParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetProjectInstancePasswdOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetProjectInstancePasswdParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "resetProjectInstancePasswd",
		Method:             "PUT",
		PathPattern:        "/project/instances/{instance_id}/resetpasswd",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ResetProjectInstancePasswdReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResetProjectInstancePasswdOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ResetProjectInstancePasswdDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ResetProjectInstancesPasswd ResetInstancesPasswd 批量重置密码
*/
func (a *Client) ResetProjectInstancesPasswd(params *ResetProjectInstancesPasswdParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetProjectInstancesPasswdOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetProjectInstancesPasswdParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "resetProjectInstancesPasswd",
		Method:             "PUT",
		PathPattern:        "/project/instances/batch/instances:resetPasswd",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ResetProjectInstancesPasswdReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResetProjectInstancesPasswdOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ResetProjectInstancesPasswdDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RestartInstances StartInstances 批量重启
*/
func (a *Client) RestartInstances(params *RestartInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestartInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestartInstancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "restartInstances",
		Method:             "PUT",
		PathPattern:        "/project/instances/batch/instances:restartInstances",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RestartInstancesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestartInstancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RestartInstancesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RestartProjectInstance RestartProjectInstance 实例重启
*/
func (a *Client) RestartProjectInstance(params *RestartProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestartProjectInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestartProjectInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "restartProjectInstance",
		Method:             "PUT",
		PathPattern:        "/project/instances/{instance_id}/restart",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RestartProjectInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestartProjectInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RestartProjectInstanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StartInstances StartInstances 批量开机
*/
func (a *Client) StartInstances(params *StartInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StartInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartInstancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "startInstances",
		Method:             "PUT",
		PathPattern:        "/project/instances/batch/instances:startInstances",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StartInstancesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartInstancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartInstancesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StartProjectInstance StartProjectInstance 实例开机
*/
func (a *Client) StartProjectInstance(params *StartProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StartProjectInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartProjectInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "startProjectInstance",
		Method:             "PUT",
		PathPattern:        "/project/instances/{instance_id}/start",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StartProjectInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartProjectInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartProjectInstanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StopInstances StartInstances 批量关机
*/
func (a *Client) StopInstances(params *StopInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopInstancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stopInstances",
		Method:             "PUT",
		PathPattern:        "/project/instances/batch/instances:stopInstances",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StopInstancesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopInstancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StopInstancesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StopProjectInstance StopProjectInstance 实例关机
*/
func (a *Client) StopProjectInstance(params *StopProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopProjectInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopProjectInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stopProjectInstance",
		Method:             "PUT",
		PathPattern:        "/project/instances/{instance_id}/stop",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StopProjectInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopProjectInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StopProjectInstanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UnLockProjectInstance UnLockProjectInstance 解锁实例
*/
func (a *Client) UnLockProjectInstance(params *UnLockProjectInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnLockProjectInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnLockProjectInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "unLockProjectInstance",
		Method:             "PUT",
		PathPattern:        "/project/instances/{instance_id}/unlock",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UnLockProjectInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UnLockProjectInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UnLockProjectInstanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
