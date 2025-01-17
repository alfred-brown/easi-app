// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SystemDetail system detail
//
// swagger:model SystemDetail
type SystemDetail struct {

	// acronym
	Acronym string `json:"acronym,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Example: 326-1-0
	// Required: true
	ID *string `json:"id"`

	// name
	// Example: CMS System
	Name string `json:"name,omitempty"`

	// status
	// Example: Active
	Status string `json:"status,omitempty"`
}

// Validate validates this system detail
func (m *SystemDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemDetail) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this system detail based on context it is used
func (m *SystemDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemDetail) UnmarshalBinary(b []byte) error {
	var res SystemDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
