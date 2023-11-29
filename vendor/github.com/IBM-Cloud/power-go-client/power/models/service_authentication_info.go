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

// ServiceAuthenticationInfo service authentication info
//
// swagger:model ServiceAuthenticationInfo
type ServiceAuthenticationInfo struct {

	// Map of internal capabilities the user has access to and their level of access
	Capabilities ResourceAccess `json:"capabilities,omitempty"`

	// Service ID
	// Required: true
	ID *string `json:"id"`

	// Service Name
	// Required: true
	Name *string `json:"name"`

	// OAuth2 Token
	// Required: true
	Token *Token `json:"token"`
}

// Validate validates this service authentication info
func (m *ServiceAuthenticationInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceAuthenticationInfo) validateCapabilities(formats strfmt.Registry) error {
	if swag.IsZero(m.Capabilities) { // not required
		return nil
	}

	if m.Capabilities != nil {
		if err := m.Capabilities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capabilities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capabilities")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceAuthenticationInfo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceAuthenticationInfo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ServiceAuthenticationInfo) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	if m.Token != nil {
		if err := m.Token.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service authentication info based on the context it is used
func (m *ServiceAuthenticationInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceAuthenticationInfo) contextValidateCapabilities(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Capabilities) { // not required
		return nil
	}

	if err := m.Capabilities.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("capabilities")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("capabilities")
		}
		return err
	}

	return nil
}

func (m *ServiceAuthenticationInfo) contextValidateToken(ctx context.Context, formats strfmt.Registry) error {

	if m.Token != nil {

		if err := m.Token.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceAuthenticationInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceAuthenticationInfo) UnmarshalBinary(b []byte) error {
	var res ServiceAuthenticationInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
