// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1Firmwares v1 firmwares
//
// swagger:model v1.Firmwares
type V1Firmwares struct {

	// the firmware kind to which the contained firmwares belong
	// Required: true
	Kind *string `json:"kind"`

	// list of firmwares per vendor
	// Required: true
	VendorFirmwares []*V1VendorFirmwares `json:"vendor_firmwares"`
}

// Validate validates this v1 firmwares
func (m *V1Firmwares) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVendorFirmwares(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Firmwares) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *V1Firmwares) validateVendorFirmwares(formats strfmt.Registry) error {

	if err := validate.Required("vendor_firmwares", "body", m.VendorFirmwares); err != nil {
		return err
	}

	for i := 0; i < len(m.VendorFirmwares); i++ {
		if swag.IsZero(m.VendorFirmwares[i]) { // not required
			continue
		}

		if m.VendorFirmwares[i] != nil {
			if err := m.VendorFirmwares[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vendor_firmwares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Firmwares) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Firmwares) UnmarshalBinary(b []byte) error {
	var res V1Firmwares
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
