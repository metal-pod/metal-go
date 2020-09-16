// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V1IPFindRequest v1 IP find request
// swagger:model v1.IPFindRequest
type V1IPFindRequest struct {

	// the address (ipv4 or ipv6) of this ip
	Ipaddress string `json:"ipaddress,omitempty"`

	// the machine an ip address is associated to
	Machineid string `json:"machineid,omitempty"`

	// the network this ip allocate request address belongs to
	Networkid string `json:"networkid,omitempty"`

	// the prefix of the network this ip address belongs to
	Networkprefix string `json:"networkprefix,omitempty"`

	// the project this ip address belongs to, empty if not strong coupled
	Projectid string `json:"projectid,omitempty"`

	// the tags that are assigned to this ip address
	Tags []string `json:"tags"`

	// the type of the ip address, ephemeral or static
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 IP find request
func (m *V1IPFindRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1IPFindRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IPFindRequest) UnmarshalBinary(b []byte) error {
	var res V1IPFindRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
