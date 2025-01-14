// Code generated by go-swagger; DO NOT EDIT.

package device_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new device type API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for device type API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AssociatedImage(params *AssociatedImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssociatedImageOK, error)

	CreateDeviceType(params *CreateDeviceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDeviceTypeOK, error)

	DeleteDeviceType(params *DeleteDeviceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDeviceTypeOK, error)

	DescribeDeviceType(params *DescribeDeviceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeDeviceTypeOK, error)

	DescribeDeviceTypeImagePartitions(params *DescribeDeviceTypeImagePartitionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeDeviceTypeImagePartitionsOK, error)

	DescribeDeviceTypeImages(params *DescribeDeviceTypeImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeDeviceTypeImagesOK, error)

	DescribeDeviceTypeRaids(params *DescribeDeviceTypeRaidsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeDeviceTypeRaidsOK, error)

	DescribeDeviceTypes(params *DescribeDeviceTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeDeviceTypesOK, error)

	DissociatedImage(params *DissociatedImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DissociatedImageOK, error)

	ModifyDeviceType(params *ModifyDeviceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyDeviceTypeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AssociatedImage AssociatedImage 机型绑定镜像
*/
func (a *Client) AssociatedImage(params *AssociatedImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssociatedImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssociatedImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "associatedImage",
		Method:             "POST",
		PathPattern:        "/deviceTypes/associatedImage",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AssociatedImageReader{formats: a.formats},
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
	success, ok := result.(*AssociatedImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AssociatedImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreateDeviceType CreateDeviceType 添加机型
*/
func (a *Client) CreateDeviceType(params *CreateDeviceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDeviceTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeviceTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createDeviceType",
		Method:             "POST",
		PathPattern:        "/deviceTypes",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDeviceTypeReader{formats: a.formats},
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
	success, ok := result.(*CreateDeviceTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateDeviceTypeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteDeviceType DeleteDeviceType 删除机型
*/
func (a *Client) DeleteDeviceType(params *DeleteDeviceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDeviceTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeviceTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteDeviceType",
		Method:             "DELETE",
		PathPattern:        "/deviceTypes/{device_type_id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDeviceTypeReader{formats: a.formats},
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
	success, ok := result.(*DeleteDeviceTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteDeviceTypeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DescribeDeviceType DescribeDeviceType 获取机型详情
*/
func (a *Client) DescribeDeviceType(params *DescribeDeviceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeDeviceTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeDeviceTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "describeDeviceType",
		Method:             "GET",
		PathPattern:        "/deviceTypes/{device_type_id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeDeviceTypeReader{formats: a.formats},
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
	success, ok := result.(*DescribeDeviceTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DescribeDeviceTypeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DescribeDeviceTypeImagePartitions QueryDeviceTypeImagePartition 根据机型，镜像，获取分区
*/
func (a *Client) DescribeDeviceTypeImagePartitions(params *DescribeDeviceTypeImagePartitionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeDeviceTypeImagePartitionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeDeviceTypeImagePartitionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "describeDeviceTypeImagePartitions",
		Method:             "GET",
		PathPattern:        "/deviceTypes/deviceTypeImagePartition",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeDeviceTypeImagePartitionsReader{formats: a.formats},
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
	success, ok := result.(*DescribeDeviceTypeImagePartitionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DescribeDeviceTypeImagePartitionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DescribeDeviceTypeImages QueryDeviceTypeImage 根据机型获取镜像
*/
func (a *Client) DescribeDeviceTypeImages(params *DescribeDeviceTypeImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeDeviceTypeImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeDeviceTypeImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "describeDeviceTypeImages",
		Method:             "GET",
		PathPattern:        "/deviceTypes/deviceTypeImage",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeDeviceTypeImagesReader{formats: a.formats},
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
	success, ok := result.(*DescribeDeviceTypeImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DescribeDeviceTypeImagesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DescribeDeviceTypeRaids DescribeDeviceTypeRaids 根据机型获取raid
*/
func (a *Client) DescribeDeviceTypeRaids(params *DescribeDeviceTypeRaidsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeDeviceTypeRaidsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeDeviceTypeRaidsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "describeDeviceTypeRaids",
		Method:             "GET",
		PathPattern:        "/deviceTypes/deviceTypeRaid",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeDeviceTypeRaidsReader{formats: a.formats},
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
	success, ok := result.(*DescribeDeviceTypeRaidsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DescribeDeviceTypeRaidsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DescribeDeviceTypes DescribeDeviceTypes 获取机型列表
*/
func (a *Client) DescribeDeviceTypes(params *DescribeDeviceTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DescribeDeviceTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeDeviceTypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "describeDeviceTypes",
		Method:             "GET",
		PathPattern:        "/deviceTypes",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeDeviceTypesReader{formats: a.formats},
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
	success, ok := result.(*DescribeDeviceTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DescribeDeviceTypesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DissociatedImage DissociatedImage 机型解绑镜像
*/
func (a *Client) DissociatedImage(params *DissociatedImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DissociatedImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDissociatedImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dissociatedImage",
		Method:             "DELETE",
		PathPattern:        "/deviceTypes/dissociatedImage",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DissociatedImageReader{formats: a.formats},
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
	success, ok := result.(*DissociatedImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DissociatedImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ModifyDeviceType ModifyDeviceType 修改机型信息
*/
func (a *Client) ModifyDeviceType(params *ModifyDeviceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyDeviceTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyDeviceTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "modifyDeviceType",
		Method:             "PUT",
		PathPattern:        "/deviceTypes/{device_type_id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyDeviceTypeReader{formats: a.formats},
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
	success, ok := result.(*ModifyDeviceTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ModifyDeviceTypeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
