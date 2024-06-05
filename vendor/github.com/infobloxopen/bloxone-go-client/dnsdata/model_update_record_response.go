/*
DNS Data API

The DNS Data is a BloxOne DDI service providing primary authoritative zone support. DNS Data is authoritative for all DNS resource records and is acting as a primary DNS server. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsdata

import (
	"encoding/json"
)

// checks if the UpdateRecordResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRecordResponse{}

// UpdateRecordResponse The response format to update the __Record__ object.
type UpdateRecordResponse struct {
	Result               *Record `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateRecordResponse UpdateRecordResponse

// NewUpdateRecordResponse instantiates a new UpdateRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRecordResponse() *UpdateRecordResponse {
	this := UpdateRecordResponse{}
	return &this
}

// NewUpdateRecordResponseWithDefaults instantiates a new UpdateRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRecordResponseWithDefaults() *UpdateRecordResponse {
	this := UpdateRecordResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateRecordResponse) GetResult() Record {
	if o == nil || IsNil(o.Result) {
		var ret Record
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRecordResponse) GetResultOk() (*Record, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateRecordResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Record and assigns it to the Result field.
func (o *UpdateRecordResponse) SetResult(v Record) {
	o.Result = &v
}

func (o UpdateRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRecordResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateRecordResponse) UnmarshalJSON(data []byte) (err error) {
	varUpdateRecordResponse := _UpdateRecordResponse{}

	err = json.Unmarshal(data, &varUpdateRecordResponse)

	if err != nil {
		return err
	}

	*o = UpdateRecordResponse(varUpdateRecordResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateRecordResponse struct {
	value *UpdateRecordResponse
	isSet bool
}

func (v NullableUpdateRecordResponse) Get() *UpdateRecordResponse {
	return v.value
}

func (v *NullableUpdateRecordResponse) Set(val *UpdateRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRecordResponse(val *UpdateRecordResponse) *NullableUpdateRecordResponse {
	return &NullableUpdateRecordResponse{value: val, isSet: true}
}

func (v NullableUpdateRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
