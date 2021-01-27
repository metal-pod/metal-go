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

// V1MachineIpmiReport v1 machine ipmi report
//
// swagger:model v1.MachineIpmiReport
type V1MachineIpmiReport struct {

	// b i o s version
	// Required: true
	BIOSVersion *string `json:"BIOSVersion"`

	// b m c Ip
	// Required: true
	BMCIP *string `json:"BMCIp"`

	// b m c version
	// Required: true
	BMCVersion *string `json:"BMCVersion"`

	// f r u
	// Required: true
	FRU *V1MachineFru `json:"FRU"`
}

// Validate validates this v1 machine ipmi report
func (m *V1MachineIpmiReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBIOSVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBMCIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBMCVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFRU(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineIpmiReport) validateBIOSVersion(formats strfmt.Registry) error {

	if err := validate.Required("BIOSVersion", "body", m.BIOSVersion); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineIpmiReport) validateBMCIP(formats strfmt.Registry) error {

	if err := validate.Required("BMCIp", "body", m.BMCIP); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineIpmiReport) validateBMCVersion(formats strfmt.Registry) error {

	if err := validate.Required("BMCVersion", "body", m.BMCVersion); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineIpmiReport) validateFRU(formats strfmt.Registry) error {

	if err := validate.Required("FRU", "body", m.FRU); err != nil {
		return err
	}

	if m.FRU != nil {
		if err := m.FRU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FRU")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 machine ipmi report based on the context it is used
func (m *V1MachineIpmiReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFRU(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineIpmiReport) contextValidateFRU(ctx context.Context, formats strfmt.Registry) error {

	if m.FRU != nil {
		if err := m.FRU.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FRU")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineIpmiReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineIpmiReport) UnmarshalBinary(b []byte) error {
	var res V1MachineIpmiReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}