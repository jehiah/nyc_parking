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

// SVLocationType A type of shared vehicle location. A station with docks is a 'dock'. A station without docks where vehicles may be left but not docked, is a 'hub'. A vehicle that is not in a station is 'free_floating'.
// swagger:model SVLocationType
type SVLocationType string

const (

	// SVLocationTypeDock captures enum value "dock"
	SVLocationTypeDock SVLocationType = "dock"

	// SVLocationTypeHub captures enum value "hub"
	SVLocationTypeHub SVLocationType = "hub"

	// SVLocationTypeFreeFloating captures enum value "free_floating"
	SVLocationTypeFreeFloating SVLocationType = "free_floating"
)

// for schema
var sVLocationTypeEnum []interface{}

func init() {
	var res []SVLocationType
	if err := json.Unmarshal([]byte(`["dock","hub","free_floating"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sVLocationTypeEnum = append(sVLocationTypeEnum, v)
	}
}

func (m SVLocationType) validateSVLocationTypeEnum(path, location string, value SVLocationType) error {
	if err := validate.Enum(path, location, value, sVLocationTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this s v location type
func (m SVLocationType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSVLocationTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
