// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LatLngStream lat lng stream
//
// swagger:model latLngStream
type LatLngStream struct {
	BaseStream

	// The sequence of lat/long values for this stream
	Data []LatLng `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *LatLngStream) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseStream
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseStream = aO0

	// AO1
	var dataAO1 struct {
		Data []LatLng `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Data = dataAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m LatLngStream) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseStream)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Data []LatLng `json:"data"`
	}

	dataAO1.Data = m.Data

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this lat lng stream
func (m *LatLngStream) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseStream
	if err := m.BaseStream.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LatLngStream) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if err := m.Data[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + strconv.Itoa(i))
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this lat lng stream based on the context it is used
func (m *LatLngStream) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseStream
	if err := m.BaseStream.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LatLngStream) contextValidateData(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(m.Data); i++ {
		if err := m.Data[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + strconv.Itoa(i))
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LatLngStream) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LatLngStream) UnmarshalBinary(b []byte) error {
	var res LatLngStream
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}