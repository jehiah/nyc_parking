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

// SegmentRulesFeatureProperties segment rules feature properties
// swagger:model SegmentRulesFeatureProperties
type SegmentRulesFeatureProperties struct {

	// How far in meters this curb segment is from the search point, at its minimum distance.
	// Only set in responses to `bylocation` requests.
	//
	DistanceFromCenterMeters float64 `json:"distance_from_center_meters,omitempty"`

	// metadata
	Metadata *SegmentMetadata `json:"metadata,omitempty"`

	// rules
	Rules []*SegmentRule `json:"rules"`

	// Rules that only apply temporarily (usually due to construction or events). Temporary rules
	// **always** take priority over regular rules. Note that temporary rules in the future may
	// be subject to change.
	//
	TemporaryRules []*SegmentTemporaryRule `json:"temporary_rules"`
}

// Validate validates this segment rules feature properties
func (m *SegmentRulesFeatureProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemporaryRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SegmentRulesFeatureProperties) validateMetadata(formats strfmt.Registry) error {

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

func (m *SegmentRulesFeatureProperties) validateRules(formats strfmt.Registry) error {

	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SegmentRulesFeatureProperties) validateTemporaryRules(formats strfmt.Registry) error {

	if swag.IsZero(m.TemporaryRules) { // not required
		return nil
	}

	for i := 0; i < len(m.TemporaryRules); i++ {
		if swag.IsZero(m.TemporaryRules[i]) { // not required
			continue
		}

		if m.TemporaryRules[i] != nil {
			if err := m.TemporaryRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("temporary_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SegmentRulesFeatureProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SegmentRulesFeatureProperties) UnmarshalBinary(b []byte) error {
	var res SegmentRulesFeatureProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}