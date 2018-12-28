// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// VehicleTypes vehicle types
// swagger:model VehicleTypes
type VehicleTypes string

const (

	// VehicleTypesAll captures enum value "all"
	VehicleTypesAll VehicleTypes = "all"

	// VehicleTypesTaxi captures enum value "taxi"
	VehicleTypesTaxi VehicleTypes = "taxi"

	// VehicleTypesCommercial captures enum value "commercial"
	VehicleTypesCommercial VehicleTypes = "commercial"

	// VehicleTypesTruck captures enum value "truck"
	VehicleTypesTruck VehicleTypes = "truck"

	// VehicleTypesMotorcycle captures enum value "motorcycle"
	VehicleTypesMotorcycle VehicleTypes = "motorcycle"

	// VehicleTypesDisabled captures enum value "disabled"
	VehicleTypesDisabled VehicleTypes = "disabled"
)

// for schema
var vehicleTypesEnum []interface{}

func init() {
	var res []VehicleTypes
	if err := json.Unmarshal([]byte(`["all","taxi","commercial","truck","motorcycle","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vehicleTypesEnum = append(vehicleTypesEnum, v)
	}
}

func (m VehicleTypes) validateVehicleTypesEnum(path, location string, value VehicleTypes) error {
	if err := validate.Enum(path, location, value, vehicleTypesEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this vehicle types
func (m VehicleTypes) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVehicleTypesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
