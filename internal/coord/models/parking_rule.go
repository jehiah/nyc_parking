// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ParkingRule How long you can park for, and how much you have to pay. See above for more details.
//
// swagger:model ParkingRule
type ParkingRule struct {

	// The method by which the total cost of a parking stay is calculated. For `live_quote`,
	// the full price schedule is unknown, but the `total_price` field in `ParkingLocation` will
	// still be set for calls to `/location/*` where the `return_total_price` param is set to
	// `true`.
	//
	// Enum: [rules live_quote]
	PriceMechanism *string `json:"price_mechanism,omitempty"`

	// The prices that apply for parking. If none are present, parking
	// price is unknown (free parking would have a `PricingRule` indicating
	// that parking is free) or is available only via a live quote. See `price_mechanism` field.
	//
	// Required: true
	Prices []*PricingRule `json:"prices"`

	// How many spaces exist at this location.
	SpacesTotal int64 `json:"spaces_total,omitempty"`
}

// Validate validates this parking rule
func (m *ParkingRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePriceMechanism(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var parkingRuleTypePriceMechanismPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rules","live_quote"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		parkingRuleTypePriceMechanismPropEnum = append(parkingRuleTypePriceMechanismPropEnum, v)
	}
}

const (

	// ParkingRulePriceMechanismRules captures enum value "rules"
	ParkingRulePriceMechanismRules string = "rules"

	// ParkingRulePriceMechanismLiveQuote captures enum value "live_quote"
	ParkingRulePriceMechanismLiveQuote string = "live_quote"
)

// prop value enum
func (m *ParkingRule) validatePriceMechanismEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, parkingRuleTypePriceMechanismPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ParkingRule) validatePriceMechanism(formats strfmt.Registry) error {

	if swag.IsZero(m.PriceMechanism) { // not required
		return nil
	}

	// value enum
	if err := m.validatePriceMechanismEnum("price_mechanism", "body", *m.PriceMechanism); err != nil {
		return err
	}

	return nil
}

func (m *ParkingRule) validatePrices(formats strfmt.Registry) error {

	if err := validate.Required("prices", "body", m.Prices); err != nil {
		return err
	}

	for i := 0; i < len(m.Prices); i++ {
		if swag.IsZero(m.Prices[i]) { // not required
			continue
		}

		if m.Prices[i] != nil {
			if err := m.Prices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ParkingRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParkingRule) UnmarshalBinary(b []byte) error {
	var res ParkingRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}