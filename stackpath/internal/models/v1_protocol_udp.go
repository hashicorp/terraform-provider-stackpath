// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V1ProtocolUDP UDP protocol matching
// swagger:model v1ProtocolUdp
type V1ProtocolUDP struct {

	// List of destination ports to allow 1-65535
	DestinationPorts []string `json:"destinationPorts"`

	// List of source ports to allow 1-65535, defaults to 1000-65535
	SourcePorts []string `json:"sourcePorts"`
}

// Validate validates this v1 protocol Udp
func (m *V1ProtocolUDP) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ProtocolUDP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ProtocolUDP) UnmarshalBinary(b []byte) error {
	var res V1ProtocolUDP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
