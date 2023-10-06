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

// IpamsvcDHCPInfo DHCPInfo
//
// The __DHCPInfo__ object represents the DHCP information associated with an address object.
//
// swagger:model ipamsvcDHCPInfo
type IpamsvcDHCPInfo struct {

	// The DHCP host name associated with this client.
	// Read Only: true
	ClientHostname string `json:"client_hostname,omitempty"`

	// The hardware address associated with this client.
	// Read Only: true
	ClientHwaddr string `json:"client_hwaddr,omitempty"`

	// The ID associated with this client.
	// Read Only: true
	ClientID string `json:"client_id,omitempty"`

	// The timestamp at which the _state_, when set to _leased_, will be changed to _free_.
	// Read Only: true
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// The DHCP fingerprint for the associated lease.
	// Read Only: true
	Fingerprint string `json:"fingerprint,omitempty"`

	// The remaining time, in seconds, until which the _state_, when set to _leased_, will remain in that state.
	// Read Only: true
	Remain int64 `json:"remain,omitempty"`

	// The timestamp at which _state_ was first set to _leased_.
	// Read Only: true
	// Format: date-time
	Start strfmt.DateTime `json:"start,omitempty"`

	// Indicates the status of this IP address from a DHCP protocol standpoint as:
	//   * _none_: Address is not under DHCP control.
	//   * _free_: Address is under DHCP control but has no lease currently assigned.
	//   * _leased_: Address is under DHCP control and has a lease currently assigned. The lease details are contained in the matching _dhcp/lease_ resource.
	// Read Only: true
	State string `json:"state,omitempty"`

	// The timestamp at which the _state_ was last reported.
	// Read Only: true
	// Format: date-time
	StateTs strfmt.DateTime `json:"state_ts,omitempty"`
}

// Validate validates this ipamsvc d h c p info
func (m *IpamsvcDHCPInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateTs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcDHCPInfo) validateEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcDHCPInfo) validateStart(formats strfmt.Registry) error {
	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcDHCPInfo) validateStateTs(formats strfmt.Registry) error {
	if swag.IsZero(m.StateTs) { // not required
		return nil
	}

	if err := validate.FormatOf("state_ts", "body", "date-time", m.StateTs.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ipamsvc d h c p info based on the context it is used
func (m *IpamsvcDHCPInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClientHostname(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClientHwaddr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClientID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFingerprint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStateTs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcDHCPInfo) contextValidateClientHostname(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "client_hostname", "body", string(m.ClientHostname)); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcDHCPInfo) contextValidateClientHwaddr(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "client_hwaddr", "body", string(m.ClientHwaddr)); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcDHCPInfo) contextValidateClientID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "client_id", "body", string(m.ClientID)); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcDHCPInfo) contextValidateEnd(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "end", "body", strfmt.DateTime(m.End)); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcDHCPInfo) contextValidateFingerprint(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "fingerprint", "body", string(m.Fingerprint)); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcDHCPInfo) contextValidateRemain(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "remain", "body", int64(m.Remain)); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcDHCPInfo) contextValidateStart(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "start", "body", strfmt.DateTime(m.Start)); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcDHCPInfo) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcDHCPInfo) contextValidateStateTs(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state_ts", "body", strfmt.DateTime(m.StateTs)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcDHCPInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcDHCPInfo) UnmarshalBinary(b []byte) error {
	var res IpamsvcDHCPInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
