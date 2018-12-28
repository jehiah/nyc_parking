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

// SegmentTimeUse A rule for which uses can be performed for how long on a particular segment of curb.
//
// swagger:model SegmentTimeUse
type SegmentTimeUse struct {

	// The ISO-8601 formatted string representing the time until which a vehicle may remain
	// at this curb while engaged in a this use if it arrives at the specified time.
	// If empty, the use is permitted for at least the next 24 hours.
	//
	// Note that this is the time when a vehicle must leave **if it arrived at the specified
	// time**. This is different than the point at which this use stops being permitted. For
	// instance, if a curb had 2-hour parking from 8am-8pm, a vehicle that arrived at 8am would
	// only be allowed to park until 10am, which is what we would represent here.
	//
	// Format: date-time
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// use
	Use Uses `json:"use,omitempty"`

	// vehicle type
	VehicleType VehicleTypes `json:"vehicle_type,omitempty"`
}

// Validate validates this segment time use
func (m *SegmentTimeUse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVehicleType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SegmentTimeUse) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SegmentTimeUse) validateUse(formats strfmt.Registry) error {

	if swag.IsZero(m.Use) { // not required
		return nil
	}

	if err := m.Use.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("use")
		}
		return err
	}

	return nil
}

func (m *SegmentTimeUse) validateVehicleType(formats strfmt.Registry) error {

	if swag.IsZero(m.VehicleType) { // not required
		return nil
	}

	if err := m.VehicleType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vehicle_type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SegmentTimeUse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SegmentTimeUse) UnmarshalBinary(b []byte) error {
	var res SegmentTimeUse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}