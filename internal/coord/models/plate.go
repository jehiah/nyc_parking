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

// Plate Information identifying a particular car.
// swagger:model Plate
type Plate struct {

	// The two letter, upper-cased state code.
	// Required: true
	State string `json:"state"`

	// The actual text of the plate.
	// Required: true
	Text string `json:"text"`
}

// Validate validates this plate
func (m *Plate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plate) validateState(formats strfmt.Registry) error {

	if err := validate.RequiredString("state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *Plate) validateText(formats strfmt.Registry) error {

	if err := validate.RequiredString("text", "body", string(m.Text)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Plate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Plate) UnmarshalBinary(b []byte) error {
	var res Plate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}