// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new role API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for role API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRole(params *CreateRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRoleOK, error)

	CurrentRole(params *CurrentRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CurrentRoleOK, error)

	DeleteRole(params *DeleteRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRoleOK, error)

	DescribeRole(params *DescribeRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeRoleOK, error)

	DescribeRoles(params *DescribeRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeRolesOK, error)

	ModifyRole(params *ModifyRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyRoleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateRole CreateRole 创建角色(暂不启用)
*/
func (a *Client) CreateRole(params *CreateRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createRole",
		Method:             "POST",
		PathPattern:        "/roles",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRoleReader{formats: a.formats},
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
	success, ok := result.(*CreateRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateRoleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CurrentRole CurrentRole 获取当前登录用户的角色
*/
func (a *Client) CurrentRole(params *CurrentRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CurrentRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCurrentRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "currentRole",
		Method:             "GET",
		PathPattern:        "/roles/roleInfo/current",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CurrentRoleReader{formats: a.formats},
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
	success, ok := result.(*CurrentRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CurrentRoleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteRole DeleteRole 删除角色(暂不启用)
*/
func (a *Client) DeleteRole(params *DeleteRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteRole",
		Method:             "DELETE",
		PathPattern:        "/roles/{role_id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRoleReader{formats: a.formats},
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
	success, ok := result.(*DeleteRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteRoleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DescribeRole DescribeRole 获取角色详情
*/
func (a *Client) DescribeRole(params *DescribeRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "describeRole",
		Method:             "GET",
		PathPattern:        "/roles/{role_id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeRoleReader{formats: a.formats},
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
	success, ok := result.(*DescribeRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DescribeRoleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DescribeRoles DescribeRoles 获取角色列表
*/
func (a *Client) DescribeRoles(params *DescribeRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeRolesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "describeRoles",
		Method:             "GET",
		PathPattern:        "/roles",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeRolesReader{formats: a.formats},
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
	success, ok := result.(*DescribeRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DescribeRolesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ModifyRole ModifyRole 修改角色信息(暂不启用)
*/
func (a *Client) ModifyRole(params *ModifyRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "modifyRole",
		Method:             "PUT",
		PathPattern:        "/roles/{role_id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyRoleReader{formats: a.formats},
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
	success, ok := result.(*ModifyRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ModifyRoleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}