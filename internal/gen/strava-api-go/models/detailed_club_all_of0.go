// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DetailedClubAllOf0 detailed club all of0
//
// swagger:model detailedClubAllOf0
type DetailedClubAllOf0 struct {
	MetaClub

	// The activity types that count for a club. This takes precedence over sport_type.
	ActivityTypes []ActivityType `json:"activity_types"`

	// The club's city.
	City string `json:"city,omitempty"`

	// The club's country.
	Country string `json:"country,omitempty"`

	// URL to a ~1185x580 pixel cover photo.
	CoverPhoto string `json:"cover_photo,omitempty"`

	// URL to a ~360x176  pixel cover photo.
	CoverPhotoSmall string `json:"cover_photo_small,omitempty"`

	// Whether the club is featured or not.
	Featured bool `json:"featured,omitempty"`

	// The club's member count.
	MemberCount int64 `json:"member_count,omitempty"`

	// Whether the club is private.
	Private bool `json:"private,omitempty"`

	// URL to a 60x60 pixel profile picture.
	ProfileMedium string `json:"profile_medium,omitempty"`

	// Deprecated. Prefer to use activity_types.
	// Enum: ["cycling","running","triathlon","other"]
	SportType string `json:"sport_type,omitempty"`

	// The club's state or geographical region.
	State string `json:"state,omitempty"`

	// The club's vanity URL.
	URL string `json:"url,omitempty"`

	// Whether the club is verified or not.
	Verified bool `json:"verified,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DetailedClubAllOf0) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MetaClub
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MetaClub = aO0

	// AO1
	var dataAO1 struct {
		ActivityTypes []ActivityType `json:"activity_types"`

		City string `json:"city,omitempty"`

		Country string `json:"country,omitempty"`

		CoverPhoto string `json:"cover_photo,omitempty"`

		CoverPhotoSmall string `json:"cover_photo_small,omitempty"`

		Featured bool `json:"featured,omitempty"`

		MemberCount int64 `json:"member_count,omitempty"`

		Private bool `json:"private,omitempty"`

		ProfileMedium string `json:"profile_medium,omitempty"`

		SportType string `json:"sport_type,omitempty"`

		State string `json:"state,omitempty"`

		URL string `json:"url,omitempty"`

		Verified bool `json:"verified,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ActivityTypes = dataAO1.ActivityTypes

	m.City = dataAO1.City

	m.Country = dataAO1.Country

	m.CoverPhoto = dataAO1.CoverPhoto

	m.CoverPhotoSmall = dataAO1.CoverPhotoSmall

	m.Featured = dataAO1.Featured

	m.MemberCount = dataAO1.MemberCount

	m.Private = dataAO1.Private

	m.ProfileMedium = dataAO1.ProfileMedium

	m.SportType = dataAO1.SportType

	m.State = dataAO1.State

	m.URL = dataAO1.URL

	m.Verified = dataAO1.Verified

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DetailedClubAllOf0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MetaClub)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ActivityTypes []ActivityType `json:"activity_types"`

		City string `json:"city,omitempty"`

		Country string `json:"country,omitempty"`

		CoverPhoto string `json:"cover_photo,omitempty"`

		CoverPhotoSmall string `json:"cover_photo_small,omitempty"`

		Featured bool `json:"featured,omitempty"`

		MemberCount int64 `json:"member_count,omitempty"`

		Private bool `json:"private,omitempty"`

		ProfileMedium string `json:"profile_medium,omitempty"`

		SportType string `json:"sport_type,omitempty"`

		State string `json:"state,omitempty"`

		URL string `json:"url,omitempty"`

		Verified bool `json:"verified,omitempty"`
	}

	dataAO1.ActivityTypes = m.ActivityTypes

	dataAO1.City = m.City

	dataAO1.Country = m.Country

	dataAO1.CoverPhoto = m.CoverPhoto

	dataAO1.CoverPhotoSmall = m.CoverPhotoSmall

	dataAO1.Featured = m.Featured

	dataAO1.MemberCount = m.MemberCount

	dataAO1.Private = m.Private

	dataAO1.ProfileMedium = m.ProfileMedium

	dataAO1.SportType = m.SportType

	dataAO1.State = m.State

	dataAO1.URL = m.URL

	dataAO1.Verified = m.Verified

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this detailed club all of0
func (m *DetailedClubAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MetaClub
	if err := m.MetaClub.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActivityTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSportType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedClubAllOf0) validateActivityTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivityTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.ActivityTypes); i++ {

		if err := m.ActivityTypes[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activity_types" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("activity_types" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

var detailedClubAllOf0TypeSportTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cycling","running","triathlon","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		detailedClubAllOf0TypeSportTypePropEnum = append(detailedClubAllOf0TypeSportTypePropEnum, v)
	}
}

// property enum
func (m *DetailedClubAllOf0) validateSportTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, detailedClubAllOf0TypeSportTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DetailedClubAllOf0) validateSportType(formats strfmt.Registry) error {

	if swag.IsZero(m.SportType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSportTypeEnum("sport_type", "body", m.SportType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this detailed club all of0 based on the context it is used
func (m *DetailedClubAllOf0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MetaClub
	if err := m.MetaClub.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateActivityTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedClubAllOf0) contextValidateActivityTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ActivityTypes); i++ {

		if swag.IsZero(m.ActivityTypes[i]) { // not required
			return nil
		}

		if err := m.ActivityTypes[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activity_types" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("activity_types" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DetailedClubAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetailedClubAllOf0) UnmarshalBinary(b []byte) error {
	var res DetailedClubAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
