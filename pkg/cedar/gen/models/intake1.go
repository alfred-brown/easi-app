// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Intake1 intake 1
// swagger:model Intake_1
type Intake1 struct {

	// governance
	// Required: true
	Governance *GovernanceIntake `json:"Governance"`
}

// Validate validates this intake 1
func (m *Intake1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGovernance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Intake1) validateGovernance(formats strfmt.Registry) error {

	if err := validate.Required("Governance", "body", m.Governance); err != nil {
		return err
	}

	if m.Governance != nil {
		if err := m.Governance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Governance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Intake1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Intake1) UnmarshalBinary(b []byte) error {
	var res Intake1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
