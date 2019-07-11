// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RedemptionInfo Info that the end user must present at the physical location to gain access.
//
// swagger:model RedemptionInfo
type RedemptionInfo struct {

	// barcode
	Barcode *RedemptionInfoBarcode `json:"barcode,omitempty"`
}

// Validate validates this redemption info
func (m *RedemptionInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBarcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RedemptionInfo) validateBarcode(formats strfmt.Registry) error {

	if swag.IsZero(m.Barcode) { // not required
		return nil
	}

	if m.Barcode != nil {
		if err := m.Barcode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("barcode")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RedemptionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RedemptionInfo) UnmarshalBinary(b []byte) error {
	var res RedemptionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RedemptionInfoBarcode If present, contains all info necessary to display a scannable barcode on the user's
// device, which will allow access at the location. Please note the following:
// * Barcodes are best read if screen brightness is set to max.
// * Assume that the barcode is necessary to display upon entrance *and upon exit*, so
// it should be accessible in your app until the session has been closed. You can determine
// if a session has been closed by polling [get_session](#operation/get_session).
// * Instruct your users to keep their phones charged, so that they can properly checkout of
// the location after their stay.
//
// swagger:model RedemptionInfoBarcode
type RedemptionInfoBarcode struct {

	// The raw string data that should be used to build the barcode.
	Data string `json:"data,omitempty"`

	// The kind of barcode that should be generated from the data.
	// Enum: [aztec]
	Format string `json:"format,omitempty"`
}

// Validate validates this redemption info barcode
func (m *RedemptionInfoBarcode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var redemptionInfoBarcodeTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aztec"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		redemptionInfoBarcodeTypeFormatPropEnum = append(redemptionInfoBarcodeTypeFormatPropEnum, v)
	}
}

const (

	// RedemptionInfoBarcodeFormatAztec captures enum value "aztec"
	RedemptionInfoBarcodeFormatAztec string = "aztec"
)

// prop value enum
func (m *RedemptionInfoBarcode) validateFormatEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, redemptionInfoBarcodeTypeFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RedemptionInfoBarcode) validateFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.Format) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormatEnum("barcode"+"."+"format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RedemptionInfoBarcode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RedemptionInfoBarcode) UnmarshalBinary(b []byte) error {
	var res RedemptionInfoBarcode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}