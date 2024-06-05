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

// checks if the KerberosKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KerberosKey{}

// KerberosKey A __Key__ object represents a Kerberos key.
type KerberosKey struct {
	// Encryption algorithm of the key in accordance with RFC 3961.
	Algorithm *string `json:"algorithm,omitempty"`
	// The description for Kerberos key. May contain 0 to 1024 characters. Can include UTF-8.
	Comment *string `json:"comment,omitempty"`
	// Kerberos realm of the principal.
	Domain *string `json:"domain,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// Kerberos principal associated with key.
	Principal *string `json:"principal,omitempty"`
	// The tags for the Kerberos key in JSON format.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Upload time for the key.
	UploadedAt *string `json:"uploaded_at,omitempty"`
	// The version number (KVNO) of the key.
	Version              *int64 `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KerberosKey KerberosKey

// NewKerberosKey instantiates a new KerberosKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKerberosKey() *KerberosKey {
	this := KerberosKey{}
	return &this
}

// NewKerberosKeyWithDefaults instantiates a new KerberosKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKerberosKeyWithDefaults() *KerberosKey {
	this := KerberosKey{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *KerberosKey) GetAlgorithm() string {
	if o == nil || IsNil(o.Algorithm) {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosKey) GetAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *KerberosKey) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *KerberosKey) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *KerberosKey) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosKey) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *KerberosKey) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *KerberosKey) SetComment(v string) {
	o.Comment = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *KerberosKey) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosKey) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *KerberosKey) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *KerberosKey) SetDomain(v string) {
	o.Domain = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KerberosKey) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosKey) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KerberosKey) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KerberosKey) SetId(v string) {
	o.Id = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *KerberosKey) GetPrincipal() string {
	if o == nil || IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosKey) GetPrincipalOk() (*string, bool) {
	if o == nil || IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *KerberosKey) HasPrincipal() bool {
	if o != nil && !IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *KerberosKey) SetPrincipal(v string) {
	o.Principal = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *KerberosKey) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosKey) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *KerberosKey) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *KerberosKey) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetUploadedAt returns the UploadedAt field value if set, zero value otherwise.
func (o *KerberosKey) GetUploadedAt() string {
	if o == nil || IsNil(o.UploadedAt) {
		var ret string
		return ret
	}
	return *o.UploadedAt
}

// GetUploadedAtOk returns a tuple with the UploadedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosKey) GetUploadedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UploadedAt) {
		return nil, false
	}
	return o.UploadedAt, true
}

// HasUploadedAt returns a boolean if a field has been set.
func (o *KerberosKey) HasUploadedAt() bool {
	if o != nil && !IsNil(o.UploadedAt) {
		return true
	}

	return false
}

// SetUploadedAt gets a reference to the given string and assigns it to the UploadedAt field.
func (o *KerberosKey) SetUploadedAt(v string) {
	o.UploadedAt = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KerberosKey) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosKey) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KerberosKey) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *KerberosKey) SetVersion(v int64) {
	o.Version = &v
}

func (o KerberosKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KerberosKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.UploadedAt) {
		toSerialize["uploaded_at"] = o.UploadedAt
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KerberosKey) UnmarshalJSON(data []byte) (err error) {
	varKerberosKey := _KerberosKey{}

	err = json.Unmarshal(data, &varKerberosKey)

	if err != nil {
		return err
	}

	*o = KerberosKey(varKerberosKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "algorithm")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "id")
		delete(additionalProperties, "principal")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "uploaded_at")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKerberosKey struct {
	value *KerberosKey
	isSet bool
}

func (v NullableKerberosKey) Get() *KerberosKey {
	return v.value
}

func (v *NullableKerberosKey) Set(val *KerberosKey) {
	v.value = val
	v.isSet = true
}

func (v NullableKerberosKey) IsSet() bool {
	return v.isSet
}

func (v *NullableKerberosKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKerberosKey(val *KerberosKey) *NullableKerberosKey {
	return &NullableKerberosKey{value: val, isSet: true}
}

func (v NullableKerberosKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKerberosKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
