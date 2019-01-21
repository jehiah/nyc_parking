// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TimestampedLatLng A point in time and space represented by timestamp, latitude, and longitude.
// swagger:model TimestampedLatLng
type TimestampedLatLng struct {

	// Latitude
	// Required: true
	Lat *float64 `json:"lat"`

	// Longitude
	// Required: true
	Lng *float64 `json:"lng"`

	// The ISO-8601 formatted string representing the timestamp of a GPS reading.
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`
}

// Validate validates this timestamped lat lng
func (m *TimestampedLatLng) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLng(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimestampedLatLng) validateLat(formats strfmt.Registry) error {

	if err := validate.Required("lat", "body", m.Lat); err != nil {
		return err
	}

	return nil
}

func (m *TimestampedLatLng) validateLng(formats strfmt.Registry) error {

	if err := validate.Required("lng", "body", m.Lng); err != nil {
		return err
	}

	return nil
}

func (m *TimestampedLatLng) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TimestampedLatLng) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimestampedLatLng) UnmarshalBinary(b []byte) error {
	var res TimestampedLatLng
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
