// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixGetRoleResponse openpitrix get role response
// swagger:model openpitrixGetRoleResponse
type OpenpitrixGetRoleResponse struct {

	// role
	Role *OpenpitrixRole `json:"role,omitempty"`
}

// Validate validates this openpitrix get role response
func (m *OpenpitrixGetRoleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixGetRoleResponse) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {

		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixGetRoleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixGetRoleResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixGetRoleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}