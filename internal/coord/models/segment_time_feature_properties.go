// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SegmentTimeFeatureProperties segment time feature properties
// swagger:model SegmentTimeFeatureProperties
type SegmentTimeFeatureProperties struct {

	// How far in meters this curb segment is from the search point, at its minimum distance.
	// Only set in responses to `bylocation` requests.
	//
	DistanceFromCenterMeters float64 `json:"distance_from_center_meters,omitempty"`

	// metadata
	Metadata *SegmentMetadata `json:"metadata,omitempty"`

	// The price a vehicle must pay while on this segment starting at this time.
	//
	Price []*PricingRule `json:"price"`

	// The reasons for the rules that apply at this time.
	Reasons []Reason `json:"reasons"`

	// uses
	Uses *SegmentTimeUses `json:"uses,omitempty"`
}

// Validate validates this segment time feature properties
func (m *SegmentTimeFeatureProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReasons(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SegmentTimeFeatureProperties) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *SegmentTimeFeatureProperties) validatePrice(formats strfmt.Registry) error {

	if swag.IsZero(m.Price) { // not required
		return nil
	}

	for i := 0; i < len(m.Price); i++ {
		if swag.IsZero(m.Price[i]) { // not required
			continue
		}

		if m.Price[i] != nil {
			if err := m.Price[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("price" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SegmentTimeFeatureProperties) validateReasons(formats strfmt.Registry) error {

	if swag.IsZero(m.Reasons) { // not required
		return nil
	}

	for i := 0; i < len(m.Reasons); i++ {

		if err := m.Reasons[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reasons" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SegmentTimeFeatureProperties) validateUses(formats strfmt.Registry) error {

	if swag.IsZero(m.Uses) { // not required
		return nil
	}

	if m.Uses != nil {
		if err := m.Uses.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uses")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SegmentTimeFeatureProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SegmentTimeFeatureProperties) UnmarshalBinary(b []byte) error {
	var res SegmentTimeFeatureProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
