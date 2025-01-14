// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MonitorDataQuery monitor data query
//
// swagger:model MonitorDataQuery
type MonitorDataQuery struct {

	// end time
	EndTime int64 `json:"endTime,omitempty"`

	// 跟赛迪约定，转换单位放这里实现
	Func map[string]int64 `json:"func,omitempty"`

	// metric都放这里了
	Labels map[string]string `json:"labels,omitempty"`

	// MetricName string            `json:"metricName"`
	// [rate, increase]
	SampleMethod string `json:"sampleMethod,omitempty"`

	// start time
	StartTime int64 `json:"startTime,omitempty"`

	// step
	Step int64 `json:"step,omitempty"`

	// [bmp_monitor_counter bmp_monitor_gauge]
	TableName string `json:"tableName,omitempty"`
}

// Validate validates this monitor data query
func (m *MonitorDataQuery) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this monitor data query based on context it is used
func (m *MonitorDataQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MonitorDataQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MonitorDataQuery) UnmarshalBinary(b []byte) error {
	var res MonitorDataQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
