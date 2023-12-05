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

// checks if the KeysGenerateTSIGResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeysGenerateTSIGResponse{}

// KeysGenerateTSIGResponse The response format to generate the TSIG key.
type KeysGenerateTSIGResponse struct {
	Result *KeysGenerateTSIGResult `json:"result,omitempty"`
}

// NewKeysGenerateTSIGResponse instantiates a new KeysGenerateTSIGResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeysGenerateTSIGResponse() *KeysGenerateTSIGResponse {
	this := KeysGenerateTSIGResponse{}
	return &this
}

// NewKeysGenerateTSIGResponseWithDefaults instantiates a new KeysGenerateTSIGResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeysGenerateTSIGResponseWithDefaults() *KeysGenerateTSIGResponse {
	this := KeysGenerateTSIGResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *KeysGenerateTSIGResponse) GetResult() KeysGenerateTSIGResult {
	if o == nil || IsNil(o.Result) {
		var ret KeysGenerateTSIGResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeysGenerateTSIGResponse) GetResultOk() (*KeysGenerateTSIGResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *KeysGenerateTSIGResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given KeysGenerateTSIGResult and assigns it to the Result field.
func (o *KeysGenerateTSIGResponse) SetResult(v KeysGenerateTSIGResult) {
	o.Result = &v
}

func (o KeysGenerateTSIGResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeysGenerateTSIGResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableKeysGenerateTSIGResponse struct {
	value *KeysGenerateTSIGResponse
	isSet bool
}

func (v NullableKeysGenerateTSIGResponse) Get() *KeysGenerateTSIGResponse {
	return v.value
}

func (v *NullableKeysGenerateTSIGResponse) Set(val *KeysGenerateTSIGResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeysGenerateTSIGResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeysGenerateTSIGResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeysGenerateTSIGResponse(val *KeysGenerateTSIGResponse) *NullableKeysGenerateTSIGResponse {
	return &NullableKeysGenerateTSIGResponse{value: val, isSet: true}
}

func (v NullableKeysGenerateTSIGResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeysGenerateTSIGResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
