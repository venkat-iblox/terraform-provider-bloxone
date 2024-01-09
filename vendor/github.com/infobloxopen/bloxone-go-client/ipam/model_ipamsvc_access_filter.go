/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the IpamsvcAccessFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcAccessFilter{}

// IpamsvcAccessFilter The __AccessFilter__ object represents an allow/deny filter for a DHCP range.
type IpamsvcAccessFilter struct {
	// The access type of DHCP filter (_allow_ or _deny_).  Defaults to _allow_.
	Access string `json:"access"`
	// The resource identifier.
	HardwareFilterId *string `json:"hardware_filter_id,omitempty"`
	// The resource identifier.
	OptionFilterId *string `json:"option_filter_id,omitempty"`
}

// NewIpamsvcAccessFilter instantiates a new IpamsvcAccessFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcAccessFilter(access string) *IpamsvcAccessFilter {
	this := IpamsvcAccessFilter{}
	this.Access = access
	return &this
}

// NewIpamsvcAccessFilterWithDefaults instantiates a new IpamsvcAccessFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcAccessFilterWithDefaults() *IpamsvcAccessFilter {
	this := IpamsvcAccessFilter{}
	return &this
}

// GetAccess returns the Access field value
func (o *IpamsvcAccessFilter) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *IpamsvcAccessFilter) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *IpamsvcAccessFilter) SetAccess(v string) {
	o.Access = v
}

// GetHardwareFilterId returns the HardwareFilterId field value if set, zero value otherwise.
func (o *IpamsvcAccessFilter) GetHardwareFilterId() string {
	if o == nil || IsNil(o.HardwareFilterId) {
		var ret string
		return ret
	}
	return *o.HardwareFilterId
}

// GetHardwareFilterIdOk returns a tuple with the HardwareFilterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAccessFilter) GetHardwareFilterIdOk() (*string, bool) {
	if o == nil || IsNil(o.HardwareFilterId) {
		return nil, false
	}
	return o.HardwareFilterId, true
}

// HasHardwareFilterId returns a boolean if a field has been set.
func (o *IpamsvcAccessFilter) HasHardwareFilterId() bool {
	if o != nil && !IsNil(o.HardwareFilterId) {
		return true
	}

	return false
}

// SetHardwareFilterId gets a reference to the given string and assigns it to the HardwareFilterId field.
func (o *IpamsvcAccessFilter) SetHardwareFilterId(v string) {
	o.HardwareFilterId = &v
}

// GetOptionFilterId returns the OptionFilterId field value if set, zero value otherwise.
func (o *IpamsvcAccessFilter) GetOptionFilterId() string {
	if o == nil || IsNil(o.OptionFilterId) {
		var ret string
		return ret
	}
	return *o.OptionFilterId
}

// GetOptionFilterIdOk returns a tuple with the OptionFilterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAccessFilter) GetOptionFilterIdOk() (*string, bool) {
	if o == nil || IsNil(o.OptionFilterId) {
		return nil, false
	}
	return o.OptionFilterId, true
}

// HasOptionFilterId returns a boolean if a field has been set.
func (o *IpamsvcAccessFilter) HasOptionFilterId() bool {
	if o != nil && !IsNil(o.OptionFilterId) {
		return true
	}

	return false
}

// SetOptionFilterId gets a reference to the given string and assigns it to the OptionFilterId field.
func (o *IpamsvcAccessFilter) SetOptionFilterId(v string) {
	o.OptionFilterId = &v
}

func (o IpamsvcAccessFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcAccessFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access"] = o.Access
	if !IsNil(o.HardwareFilterId) {
		toSerialize["hardware_filter_id"] = o.HardwareFilterId
	}
	if !IsNil(o.OptionFilterId) {
		toSerialize["option_filter_id"] = o.OptionFilterId
	}
	return toSerialize, nil
}

type NullableIpamsvcAccessFilter struct {
	value *IpamsvcAccessFilter
	isSet bool
}

func (v NullableIpamsvcAccessFilter) Get() *IpamsvcAccessFilter {
	return v.value
}

func (v *NullableIpamsvcAccessFilter) Set(val *IpamsvcAccessFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcAccessFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcAccessFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcAccessFilter(val *IpamsvcAccessFilter) *NullableIpamsvcAccessFilter {
	return &NullableIpamsvcAccessFilter{value: val, isSet: true}
}

func (v NullableIpamsvcAccessFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcAccessFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
