/*
DDI Keys API

The DDI Keys application is a BloxOne DDI service for managing TSIG keys and GSS-TSIG (Kerberos) keys which are used by other BloxOne DDI applications. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keys

import (
	"encoding/json"
)

// checks if the KeysReadKerberosKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeysReadKerberosKeyResponse{}

// KeysReadKerberosKeyResponse The response format to retrieve the __KerberosKey__ resource extracted from the uploaded keytab file.
type KeysReadKerberosKeyResponse struct {
	Result *KerberosKey `json:"result,omitempty"`
}

// NewKeysReadKerberosKeyResponse instantiates a new KeysReadKerberosKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeysReadKerberosKeyResponse() *KeysReadKerberosKeyResponse {
	this := KeysReadKerberosKeyResponse{}
	return &this
}

// NewKeysReadKerberosKeyResponseWithDefaults instantiates a new KeysReadKerberosKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeysReadKerberosKeyResponseWithDefaults() *KeysReadKerberosKeyResponse {
	this := KeysReadKerberosKeyResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *KeysReadKerberosKeyResponse) GetResult() KerberosKey {
	if o == nil || IsNil(o.Result) {
		var ret KerberosKey
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeysReadKerberosKeyResponse) GetResultOk() (*KerberosKey, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *KeysReadKerberosKeyResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given KerberosKey and assigns it to the Result field.
func (o *KeysReadKerberosKeyResponse) SetResult(v KerberosKey) {
	o.Result = &v
}

func (o KeysReadKerberosKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeysReadKerberosKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableKeysReadKerberosKeyResponse struct {
	value *KeysReadKerberosKeyResponse
	isSet bool
}

func (v NullableKeysReadKerberosKeyResponse) Get() *KeysReadKerberosKeyResponse {
	return v.value
}

func (v *NullableKeysReadKerberosKeyResponse) Set(val *KeysReadKerberosKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeysReadKerberosKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeysReadKerberosKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeysReadKerberosKeyResponse(val *KeysReadKerberosKeyResponse) *NullableKeysReadKerberosKeyResponse {
	return &NullableKeysReadKerberosKeyResponse{value: val, isSet: true}
}

func (v NullableKeysReadKerberosKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeysReadKerberosKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
