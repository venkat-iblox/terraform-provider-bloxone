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

// checks if the KeysListTSIGKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeysListTSIGKeyResponse{}

// KeysListTSIGKeyResponse The response format to retrieve __TSIGKey__ objects.
type KeysListTSIGKeyResponse struct {
	// The list of TSIGKey objects.
	Results []KeysTSIGKey `json:"results,omitempty"`
}

// NewKeysListTSIGKeyResponse instantiates a new KeysListTSIGKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeysListTSIGKeyResponse() *KeysListTSIGKeyResponse {
	this := KeysListTSIGKeyResponse{}
	return &this
}

// NewKeysListTSIGKeyResponseWithDefaults instantiates a new KeysListTSIGKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeysListTSIGKeyResponseWithDefaults() *KeysListTSIGKeyResponse {
	this := KeysListTSIGKeyResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *KeysListTSIGKeyResponse) GetResults() []KeysTSIGKey {
	if o == nil || IsNil(o.Results) {
		var ret []KeysTSIGKey
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeysListTSIGKeyResponse) GetResultsOk() ([]KeysTSIGKey, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *KeysListTSIGKeyResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []KeysTSIGKey and assigns it to the Results field.
func (o *KeysListTSIGKeyResponse) SetResults(v []KeysTSIGKey) {
	o.Results = v
}

func (o KeysListTSIGKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeysListTSIGKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableKeysListTSIGKeyResponse struct {
	value *KeysListTSIGKeyResponse
	isSet bool
}

func (v NullableKeysListTSIGKeyResponse) Get() *KeysListTSIGKeyResponse {
	return v.value
}

func (v *NullableKeysListTSIGKeyResponse) Set(val *KeysListTSIGKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeysListTSIGKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeysListTSIGKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeysListTSIGKeyResponse(val *KeysListTSIGKeyResponse) *NullableKeysListTSIGKeyResponse {
	return &NullableKeysListTSIGKeyResponse{value: val, isSet: true}
}

func (v NullableKeysListTSIGKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeysListTSIGKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
