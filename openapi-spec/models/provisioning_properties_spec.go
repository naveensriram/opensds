// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProvisioningPropertiesSpec ProvisioningPropertiesSpec represents some suggested properties for performing provisioning policies.
// swagger:model ProvisioningPropertiesSpec
type ProvisioningPropertiesSpec struct {

	// data storage
	DataStorage *DataStorageLoS `json:"dataStorage,omitempty"`

	// io connectivity
	IoConnectivity *IOConnectivityLoS `json:"ioConnectivity,omitempty"`
}

// Validate validates this provisioning properties spec
func (m *ProvisioningPropertiesSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoConnectivity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvisioningPropertiesSpec) validateDataStorage(formats strfmt.Registry) error {

	if swag.IsZero(m.DataStorage) { // not required
		return nil
	}

	if m.DataStorage != nil {
		if err := m.DataStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataStorage")
			}
			return err
		}
	}

	return nil
}

func (m *ProvisioningPropertiesSpec) validateIoConnectivity(formats strfmt.Registry) error {

	if swag.IsZero(m.IoConnectivity) { // not required
		return nil
	}

	if m.IoConnectivity != nil {
		if err := m.IoConnectivity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ioConnectivity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProvisioningPropertiesSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProvisioningPropertiesSpec) UnmarshalBinary(b []byte) error {
	var res ProvisioningPropertiesSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
