// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/api_key"
	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/collect"
	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/device"
	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/device_type"
	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/idc"
	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/image"
	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/instance"
	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/license"
	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/os"
	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/project"
	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/raid"
	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/resource"
	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/role"
	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/sshkey"
	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/user"
)

// Default c p s HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "bmp-openapi.bmp.local"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new c p s HTTP client.
func NewHTTPClient(formats strfmt.Registry) *CPS {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new c p s HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *CPS {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new c p s client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *CPS {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(CPS)
	cli.Transport = transport
	cli.APIKey = api_key.New(transport, formats)
	cli.Collect = collect.New(transport, formats)
	cli.Device = device.New(transport, formats)
	cli.DeviceType = device_type.New(transport, formats)
	cli.Idc = idc.New(transport, formats)
	cli.Image = image.New(transport, formats)
	cli.Instance = instance.New(transport, formats)
	cli.License = license.New(transport, formats)
	cli.Os = os.New(transport, formats)
	cli.Project = project.New(transport, formats)
	cli.Raid = raid.New(transport, formats)
	cli.Resource = resource.New(transport, formats)
	cli.Role = role.New(transport, formats)
	cli.Sshkey = sshkey.New(transport, formats)
	cli.User = user.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// CPS is a client for c p s
type CPS struct {
	APIKey api_key.ClientService

	Collect collect.ClientService

	Device device.ClientService

	DeviceType device_type.ClientService

	Idc idc.ClientService

	Image image.ClientService

	Instance instance.ClientService

	License license.ClientService

	Os os.ClientService

	Project project.ClientService

	Raid raid.ClientService

	Resource resource.ClientService

	Role role.ClientService

	Sshkey sshkey.ClientService

	User user.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *CPS) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.APIKey.SetTransport(transport)
	c.Collect.SetTransport(transport)
	c.Device.SetTransport(transport)
	c.DeviceType.SetTransport(transport)
	c.Idc.SetTransport(transport)
	c.Image.SetTransport(transport)
	c.Instance.SetTransport(transport)
	c.License.SetTransport(transport)
	c.Os.SetTransport(transport)
	c.Project.SetTransport(transport)
	c.Raid.SetTransport(transport)
	c.Resource.SetTransport(transport)
	c.Role.SetTransport(transport)
	c.Sshkey.SetTransport(transport)
	c.User.SetTransport(transport)
}