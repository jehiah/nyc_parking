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

// TimeRange Defines a time range.
// swagger:model TimeRange
type TimeRange struct {

	// Range upper-bound.
	// Required: true
	// Format: date-time
	MaxValue *strfmt.DateTime `json:"max_value"`

	// Range lower-bound.
	// Required: true
	// Format: date-time
	MinValue *strfmt.DateTime `json:"min_value"`
}

// Validate validates this time range
func (m *TimeRange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaxValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimeRange) validateMaxValue(formats strfmt.Registry) error {

	if err := validate.Required("max_value", "body", m.MaxValue); err != nil {
		return err
	}

	if err := validate.FormatOf("max_value", "body", "date-time", m.MaxValue.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TimeRange) validateMinValue(formats strfmt.Registry) error {

	if err := validate.Required("min_value", "body", m.MinValue); err != nil {
		return err
	}

	if err := validate.FormatOf("min_value", "body", "date-time", m.MinValue.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TimeRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeRange) UnmarshalBinary(b []byte) error {
	var res TimeRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
