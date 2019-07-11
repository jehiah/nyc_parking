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

// Partner An external company that has the ability to use our APIs.
// swagger:model Partner
type Partner struct {

	// A flag denoting if the partner has accepted Surveyor terms.
	// Format: datetime
	AcceptedSurveyorTerms *strfmt.DateTime `json:"accepted_surveyor_terms,omitempty"`

	// A flag denoting if the partner has their API access enabled.
	APIAccessEnabled bool `json:"api_access_enabled,omitempty"`

	// A list of API keys assigned to the partner. These keys are safe to use in urls called
	// directly from end-user client code (JS, iOS, Android).
	//
	APIKeys []string `json:"api_keys"`

	// card metadata
	CardMetadata *CardMetadata `json:"card_metadata,omitempty"`

	// The partner's number of employees. On creation, this field is required.
	// Represents the low end of the range from the UI.
	//
	CompanySize int64 `json:"company_size,omitempty"`

	// The time this partner was created, either internally, or via an external login.
	// Format: datetime
	CreationTime strfmt.DateTime `json:"creation_time,omitempty"`

	// The first time a member in this partner logged in.
	// Format: datetime
	FirstLoginTime *strfmt.DateTime `json:"first_login_time,omitempty"`

	// The unique ID for the partner. On creation, this field is set by the
	// server.
	//
	ID string `json:"id,omitempty"`

	// The industry of the partner. On creation, this field is required.
	//
	Industry string `json:"industry,omitempty"`

	// A flag denoting if the partner is a supply provider on the platform.
	IsProvider bool `json:"is_provider,omitempty"`

	// The name of the partner. On creation, this field is required.
	//
	Name string `json:"name,omitempty"`

	// A flag denoting if the partner has transactions enabled.
	ProdEnabled bool `json:"prod_enabled,omitempty"`

	// A list of secret keys assigned to the partner. These keys should only be used in calls
	// originating from partner backends as they afford extra privileges such as generating
	// production JWTs.
	//
	SecretKeys []string `json:"secret_keys"`

	// The partner's total amount of surveyor credit (in meters).
	//
	SurveyorCreditM *float64 `json:"surveyor_credit_m,omitempty"`

	// An array of whitelisted domain for the partner, e.g. `google.com`.
	//
	WhitelistedDomains []string `json:"whitelisted_domains"`
}

// Validate validates this partner
func (m *Partner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptedSurveyorTerms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstLoginTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Partner) validateAcceptedSurveyorTerms(formats strfmt.Registry) error {

	if swag.IsZero(m.AcceptedSurveyorTerms) { // not required
		return nil
	}

	if err := validate.FormatOf("accepted_surveyor_terms", "body", "datetime", m.AcceptedSurveyorTerms.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Partner) validateCardMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.CardMetadata) { // not required
		return nil
	}

	if m.CardMetadata != nil {
		if err := m.CardMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("card_metadata")
			}
			return err
		}
	}

	return nil
}

func (m *Partner) validateCreationTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_time", "body", "datetime", m.CreationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Partner) validateFirstLoginTime(formats strfmt.Registry) error {

	if swag.IsZero(m.FirstLoginTime) { // not required
		return nil
	}

	if err := validate.FormatOf("first_login_time", "body", "datetime", m.FirstLoginTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Partner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Partner) UnmarshalBinary(b []byte) error {
	var res Partner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}