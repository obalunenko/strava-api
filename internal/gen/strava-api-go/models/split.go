// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Split split
//
// swagger:model split
type Split struct {
	// The average speed of this split, in meters per second
	AverageSpeed float32 `json:"average_speed,omitempty"`

	// The distance of this split, in meters
	Distance float32 `json:"distance,omitempty"`

	// The elapsed time of this split, in seconds
	ElapsedTime int64 `json:"elapsed_time,omitempty"`

	// The elevation difference of this split, in meters
	ElevationDifference float32 `json:"elevation_difference,omitempty"`

	// The moving time of this split, in seconds
	MovingTime int64 `json:"moving_time,omitempty"`

	// The pacing zone of this split
	PaceZone int64 `json:"pace_zone,omitempty"`

	// N/A
	Split int64 `json:"split,omitempty"`
}

// Validate validates this split
func (m *Split) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this split based on context it is used
func (m *Split) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Split) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Split) UnmarshalBinary(b []byte) error {
	var res Split
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}