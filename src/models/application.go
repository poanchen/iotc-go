// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Application application
//
// swagger:model Application
type Application struct {

	// Display name of the application.
	DisplayName string `json:"displayName,omitempty"`

	// The URL host of the application.
	// Required: true
	Host *string `json:"host"`

	// Unique ID of the application.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The URL subdomain of the application.
	// Required: true
	Subdomain *string `json:"subdomain"`
}

// Validate validates this application
func (m *Application) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubdomain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Application) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateSubdomain(formats strfmt.Registry) error {

	if err := validate.Required("subdomain", "body", m.Subdomain); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Application) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Application) UnmarshalBinary(b []byte) error {
	var res Application
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}