/*
BloxOne Redirect API

You can configure BloxOne Threat Defense Cloud to redirect traffic to the Infoblox server that displays the default or customized redirect page. You can redirect traffic to a custom destination using custom redirects.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redirect

import (
	"encoding/json"
)

// checks if the RedirectPageUpdateRedirectPage400ResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedirectPageUpdateRedirectPage400ResponseError{}

// RedirectPageUpdateRedirectPage400ResponseError struct for RedirectPageUpdateRedirectPage400ResponseError
type RedirectPageUpdateRedirectPage400ResponseError struct {
	Code                 *string `json:"code,omitempty"`
	Message              *string `json:"message,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RedirectPageUpdateRedirectPage400ResponseError RedirectPageUpdateRedirectPage400ResponseError

// NewRedirectPageUpdateRedirectPage400ResponseError instantiates a new RedirectPageUpdateRedirectPage400ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectPageUpdateRedirectPage400ResponseError() *RedirectPageUpdateRedirectPage400ResponseError {
	this := RedirectPageUpdateRedirectPage400ResponseError{}
	return &this
}

// NewRedirectPageUpdateRedirectPage400ResponseErrorWithDefaults instantiates a new RedirectPageUpdateRedirectPage400ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectPageUpdateRedirectPage400ResponseErrorWithDefaults() *RedirectPageUpdateRedirectPage400ResponseError {
	this := RedirectPageUpdateRedirectPage400ResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *RedirectPageUpdateRedirectPage400ResponseError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectPageUpdateRedirectPage400ResponseError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *RedirectPageUpdateRedirectPage400ResponseError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *RedirectPageUpdateRedirectPage400ResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RedirectPageUpdateRedirectPage400ResponseError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectPageUpdateRedirectPage400ResponseError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RedirectPageUpdateRedirectPage400ResponseError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RedirectPageUpdateRedirectPage400ResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RedirectPageUpdateRedirectPage400ResponseError) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectPageUpdateRedirectPage400ResponseError) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RedirectPageUpdateRedirectPage400ResponseError) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RedirectPageUpdateRedirectPage400ResponseError) SetStatus(v string) {
	o.Status = &v
}

func (o RedirectPageUpdateRedirectPage400ResponseError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedirectPageUpdateRedirectPage400ResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RedirectPageUpdateRedirectPage400ResponseError) UnmarshalJSON(data []byte) (err error) {
	varRedirectPageUpdateRedirectPage400ResponseError := _RedirectPageUpdateRedirectPage400ResponseError{}

	err = json.Unmarshal(data, &varRedirectPageUpdateRedirectPage400ResponseError)

	if err != nil {
		return err
	}

	*o = RedirectPageUpdateRedirectPage400ResponseError(varRedirectPageUpdateRedirectPage400ResponseError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRedirectPageUpdateRedirectPage400ResponseError struct {
	value *RedirectPageUpdateRedirectPage400ResponseError
	isSet bool
}

func (v NullableRedirectPageUpdateRedirectPage400ResponseError) Get() *RedirectPageUpdateRedirectPage400ResponseError {
	return v.value
}

func (v *NullableRedirectPageUpdateRedirectPage400ResponseError) Set(val *RedirectPageUpdateRedirectPage400ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectPageUpdateRedirectPage400ResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectPageUpdateRedirectPage400ResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectPageUpdateRedirectPage400ResponseError(val *RedirectPageUpdateRedirectPage400ResponseError) *NullableRedirectPageUpdateRedirectPage400ResponseError {
	return &NullableRedirectPageUpdateRedirectPage400ResponseError{value: val, isSet: true}
}

func (v NullableRedirectPageUpdateRedirectPage400ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectPageUpdateRedirectPage400ResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
