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

// IpamsvcTSIGKey TSIGKey
//
// A __TSIGKey__ object represents a TSIG key.
//
// swagger:model ipamsvcTSIGKey
type IpamsvcTSIGKey struct {

	// TSIG key algorithm.
	//
	// Valid values are:
	//  * _hmac_sha256_
	//  * _hmac_sha1_
	//  * _hmac_sha224_
	//  * _hmac_sha384_
	//  * _hmac_sha512_
	Algorithm string `json:"algorithm,omitempty"`

	// The description for the TSIG key. May contain 0 to 1024 characters. Can include UTF-8.
	Comment string `json:"comment,omitempty"`

	// The resource identifier.
	// Required: true
	Key *string `json:"key,omitempty"`

	// The TSIG key name, FQDN.
	Name string `json:"name,omitempty"`

	// The TSIG key name in punycode.
	// Read Only: true
	ProtocolName string `json:"protocol_name,omitempty"`

	// The TSIG key secret, base64 string.
	Secret string `json:"secret,omitempty"`
}

// Validate validates this ipamsvc t s i g key
func (m *IpamsvcTSIGKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcTSIGKey) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ipamsvc t s i g key based on the context it is used
func (m *IpamsvcTSIGKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProtocolName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcTSIGKey) contextValidateProtocolName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "protocol_name", "body", string(m.ProtocolName)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcTSIGKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcTSIGKey) UnmarshalBinary(b []byte) error {
	var res IpamsvcTSIGKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
