// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Comment comment
//
// swagger:model comment
type Comment struct {

	// The identifier of the activity this comment is related to
	ActivityID int64 `json:"activity_id,omitempty"`

	// athlete
	Athlete *CommentAthlete `json:"athlete,omitempty"`

	// The time at which this comment was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The unique identifier of this comment
	ID int64 `json:"id,omitempty"`

	// The content of the comment
	Text string `json:"text,omitempty"`
}

// Validate validates this comment
func (m *Comment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAthlete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Comment) validateAthlete(formats strfmt.Registry) error {
	if swag.IsZero(m.Athlete) { // not required
		return nil
	}

	if m.Athlete != nil {
		if err := m.Athlete.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("athlete")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("athlete")
			}
			return err
		}
	}

	return nil
}

func (m *Comment) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this comment based on the context it is used
func (m *Comment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAthlete(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Comment) contextValidateAthlete(ctx context.Context, formats strfmt.Registry) error {

	if m.Athlete != nil {

		if swag.IsZero(m.Athlete) { // not required
			return nil
		}

		if err := m.Athlete.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("athlete")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("athlete")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Comment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Comment) UnmarshalBinary(b []byte) error {
	var res Comment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
