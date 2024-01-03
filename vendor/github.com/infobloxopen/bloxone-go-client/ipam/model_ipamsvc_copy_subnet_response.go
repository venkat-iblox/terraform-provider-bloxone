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

// checks if the IpamsvcCopySubnetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcCopySubnetResponse{}

// IpamsvcCopySubnetResponse The response format to copy the __Subnet__ object.
type IpamsvcCopySubnetResponse struct {
	Result *IpamsvcCopyResponse `json:"result,omitempty"`
}

// NewIpamsvcCopySubnetResponse instantiates a new IpamsvcCopySubnetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcCopySubnetResponse() *IpamsvcCopySubnetResponse {
	this := IpamsvcCopySubnetResponse{}
	return &this
}

// NewIpamsvcCopySubnetResponseWithDefaults instantiates a new IpamsvcCopySubnetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcCopySubnetResponseWithDefaults() *IpamsvcCopySubnetResponse {
	this := IpamsvcCopySubnetResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *IpamsvcCopySubnetResponse) GetResult() IpamsvcCopyResponse {
	if o == nil || IsNil(o.Result) {
		var ret IpamsvcCopyResponse
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcCopySubnetResponse) GetResultOk() (*IpamsvcCopyResponse, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *IpamsvcCopySubnetResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given IpamsvcCopyResponse and assigns it to the Result field.
func (o *IpamsvcCopySubnetResponse) SetResult(v IpamsvcCopyResponse) {
	o.Result = &v
}

func (o IpamsvcCopySubnetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcCopySubnetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableIpamsvcCopySubnetResponse struct {
	value *IpamsvcCopySubnetResponse
	isSet bool
}

func (v NullableIpamsvcCopySubnetResponse) Get() *IpamsvcCopySubnetResponse {
	return v.value
}

func (v *NullableIpamsvcCopySubnetResponse) Set(val *IpamsvcCopySubnetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcCopySubnetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcCopySubnetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcCopySubnetResponse(val *IpamsvcCopySubnetResponse) *NullableIpamsvcCopySubnetResponse {
	return &NullableIpamsvcCopySubnetResponse{value: val, isSet: true}
}

func (v NullableIpamsvcCopySubnetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcCopySubnetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
