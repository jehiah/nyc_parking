// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	common_models "github.com/coordcity/go/common/models"
)

// SegmentTemporaryRule A temporary rule that applies along a particular segment of curb. Rather than recurring
// on a weekly basis, temporary rules are scheduled once. These **always** take precedence
// over regular rules.
//
// swagger:model SegmentTemporaryRule
type SegmentTemporaryRule struct {

	// When this rule stops applying, in ISO-8601 format.
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// The longest a vehicle may remain at this curb while engaged in a permitted use, in
	// hours.
	//
	// If a new rule starts applying before `max_duration_h` has elapsed, the new rule's
	// max_duration_h takes effect, but counting from when the vehicle first arrived.
	// For instance, If a curb had 2-hour parking until 5pm but 3-hour parking thereafter, and
	// a vehicle arrived at 4pm, they could continue parking until 7pm.
	//
	MaxDurationH float64 `json:"max_duration_h,omitempty"`

	// The uses that are permitted for vehicles not of this segment's primary vehicle type.
	//
	OtherVehiclesPermitted []Uses `json:"other_vehicles_permitted"`

	// All the uses that are permitted, including the primary use.
	Permitted []Uses `json:"permitted"`

	// The price a vehicle must pay while on this segment. In general, this price applies regardless
	// of use or vehicle type.
	//
	// If a new rule starts applying, that rule's prices take effect, but counting from when the
	// vehicle first arrived. For instance, if a curb had:
	//   * Parking at $1 an hour until 8am;
	//   * Parking at $4 for the first hour and $5 for the second hour thereafter,
	// A vehicle arriving at 7am would pay $1 for the first hour and $5 for the second.
	//
	Price []*common_models.PricingRule `json:"price"`

	// primary
	Primary Uses `json:"primary,omitempty"`

	// The reasons that this rule applies at this time.
	Reasons []Reason `json:"reasons"`

	// When this rule starts to apply, in ISO-8601 format.
	// Format: date-time
	Start strfmt.DateTime `json:"start,omitempty"`

	// vehicle type
	VehicleType VehicleTypes `json:"vehicle_type,omitempty"`
}

// Validate validates this segment temporary rule
func (m *SegmentTemporaryRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOtherVehiclesPermitted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermitted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReasons(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
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

func (m *SegmentTemporaryRule) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SegmentTemporaryRule) validateOtherVehiclesPermitted(formats strfmt.Registry) error {

	if swag.IsZero(m.OtherVehiclesPermitted) { // not required
		return nil
	}

	for i := 0; i < len(m.OtherVehiclesPermitted); i++ {

		if err := m.OtherVehiclesPermitted[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("other_vehicles_permitted" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SegmentTemporaryRule) validatePermitted(formats strfmt.Registry) error {

	if swag.IsZero(m.Permitted) { // not required
		return nil
	}

	for i := 0; i < len(m.Permitted); i++ {

		if err := m.Permitted[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permitted" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SegmentTemporaryRule) validatePrice(formats strfmt.Registry) error {

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

func (m *SegmentTemporaryRule) validatePrimary(formats strfmt.Registry) error {

	if swag.IsZero(m.Primary) { // not required
		return nil
	}

	if err := m.Primary.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("primary")
		}
		return err
	}

	return nil
}

func (m *SegmentTemporaryRule) validateReasons(formats strfmt.Registry) error {

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

func (m *SegmentTemporaryRule) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SegmentTemporaryRule) validateVehicleType(formats strfmt.Registry) error {

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
func (m *SegmentTemporaryRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SegmentTemporaryRule) UnmarshalBinary(b []byte) error {
	var res SegmentTemporaryRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}