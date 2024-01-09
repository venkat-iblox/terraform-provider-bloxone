/*
Host Activation Service

Host activation service provides a RESTful interface to manage cert and join token object. Join tokens are essentially a password that allows on-prem hosts to auto-associate themselves to a customer's account and receive a signed cert.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infra_provision

import (
	"encoding/json"
)

// checks if the HostactivationRevokeCertRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostactivationRevokeCertRequest{}

// HostactivationRevokeCertRequest struct for HostactivationRevokeCertRequest
type HostactivationRevokeCertRequest struct {
	CertSerial *string `json:"cert_serial,omitempty"`
	// On-prem host ID which can be obtained either from on-prem or BloxOne UI portal(Manage > Infrastructure > Hosts > Select the onprem > click on 3 dots on top right side > General Information > Ophid) .
	Ophid        *string `json:"ophid,omitempty"`
	RevokeReason *string `json:"revoke_reason,omitempty"`
}

// NewHostactivationRevokeCertRequest instantiates a new HostactivationRevokeCertRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostactivationRevokeCertRequest() *HostactivationRevokeCertRequest {
	this := HostactivationRevokeCertRequest{}
	return &this
}

// NewHostactivationRevokeCertRequestWithDefaults instantiates a new HostactivationRevokeCertRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostactivationRevokeCertRequestWithDefaults() *HostactivationRevokeCertRequest {
	this := HostactivationRevokeCertRequest{}
	return &this
}

// GetCertSerial returns the CertSerial field value if set, zero value otherwise.
func (o *HostactivationRevokeCertRequest) GetCertSerial() string {
	if o == nil || IsNil(o.CertSerial) {
		var ret string
		return ret
	}
	return *o.CertSerial
}

// GetCertSerialOk returns a tuple with the CertSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationRevokeCertRequest) GetCertSerialOk() (*string, bool) {
	if o == nil || IsNil(o.CertSerial) {
		return nil, false
	}
	return o.CertSerial, true
}

// HasCertSerial returns a boolean if a field has been set.
func (o *HostactivationRevokeCertRequest) HasCertSerial() bool {
	if o != nil && !IsNil(o.CertSerial) {
		return true
	}

	return false
}

// SetCertSerial gets a reference to the given string and assigns it to the CertSerial field.
func (o *HostactivationRevokeCertRequest) SetCertSerial(v string) {
	o.CertSerial = &v
}

// GetOphid returns the Ophid field value if set, zero value otherwise.
func (o *HostactivationRevokeCertRequest) GetOphid() string {
	if o == nil || IsNil(o.Ophid) {
		var ret string
		return ret
	}
	return *o.Ophid
}

// GetOphidOk returns a tuple with the Ophid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationRevokeCertRequest) GetOphidOk() (*string, bool) {
	if o == nil || IsNil(o.Ophid) {
		return nil, false
	}
	return o.Ophid, true
}

// HasOphid returns a boolean if a field has been set.
func (o *HostactivationRevokeCertRequest) HasOphid() bool {
	if o != nil && !IsNil(o.Ophid) {
		return true
	}

	return false
}

// SetOphid gets a reference to the given string and assigns it to the Ophid field.
func (o *HostactivationRevokeCertRequest) SetOphid(v string) {
	o.Ophid = &v
}

// GetRevokeReason returns the RevokeReason field value if set, zero value otherwise.
func (o *HostactivationRevokeCertRequest) GetRevokeReason() string {
	if o == nil || IsNil(o.RevokeReason) {
		var ret string
		return ret
	}
	return *o.RevokeReason
}

// GetRevokeReasonOk returns a tuple with the RevokeReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationRevokeCertRequest) GetRevokeReasonOk() (*string, bool) {
	if o == nil || IsNil(o.RevokeReason) {
		return nil, false
	}
	return o.RevokeReason, true
}

// HasRevokeReason returns a boolean if a field has been set.
func (o *HostactivationRevokeCertRequest) HasRevokeReason() bool {
	if o != nil && !IsNil(o.RevokeReason) {
		return true
	}

	return false
}

// SetRevokeReason gets a reference to the given string and assigns it to the RevokeReason field.
func (o *HostactivationRevokeCertRequest) SetRevokeReason(v string) {
	o.RevokeReason = &v
}

func (o HostactivationRevokeCertRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostactivationRevokeCertRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertSerial) {
		toSerialize["cert_serial"] = o.CertSerial
	}
	if !IsNil(o.Ophid) {
		toSerialize["ophid"] = o.Ophid
	}
	if !IsNil(o.RevokeReason) {
		toSerialize["revoke_reason"] = o.RevokeReason
	}
	return toSerialize, nil
}

type NullableHostactivationRevokeCertRequest struct {
	value *HostactivationRevokeCertRequest
	isSet bool
}

func (v NullableHostactivationRevokeCertRequest) Get() *HostactivationRevokeCertRequest {
	return v.value
}

func (v *NullableHostactivationRevokeCertRequest) Set(val *HostactivationRevokeCertRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHostactivationRevokeCertRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHostactivationRevokeCertRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostactivationRevokeCertRequest(val *HostactivationRevokeCertRequest) *NullableHostactivationRevokeCertRequest {
	return &NullableHostactivationRevokeCertRequest{value: val, isSet: true}
}

func (v NullableHostactivationRevokeCertRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostactivationRevokeCertRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
