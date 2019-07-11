// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CardMetadata Credit card metadata.
// swagger:model CardMetadata
type CardMetadata struct {

	// Card brand (Visa, Mastercard, etc.)
	CardBrand string `json:"card_brand,omitempty"`

	// Card's expiration month.
	ExpMonth int32 `json:"exp_month,omitempty"`

	// Card's expiration year.
	ExpYear int32 `json:"exp_year,omitempty"`

	// Last four digits of card.
	LastFourDigits string `json:"last_four_digits,omitempty"`
}

// Validate validates this card metadata
func (m *CardMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CardMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CardMetadata) UnmarshalBinary(b []byte) error {
	var res CardMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}