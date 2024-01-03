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

// checks if the IpamsvcKerberosKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcKerberosKey{}

// IpamsvcKerberosKey A __KerberosKey__ object (_keys/kerberos_) represents a Kerberos key.
type IpamsvcKerberosKey struct {
	// Encryption algorithm of the key in accordance with RFC 3961.
	Algorithm *string `json:"algorithm,omitempty"`
	// Kerberos realm of the principal.
	Domain *string `json:"domain,omitempty"`
	// The resource identifier.
	Key string `json:"key"`
	// Kerberos principal associated with key.
	Principal *string `json:"principal,omitempty"`
	// Upload time for the key.
	UploadedAt *string `json:"uploaded_at,omitempty"`
	// The version number (KVNO) of the key.
	Version *int64 `json:"version,omitempty"`
}

// NewIpamsvcKerberosKey instantiates a new IpamsvcKerberosKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcKerberosKey(key string) *IpamsvcKerberosKey {
	this := IpamsvcKerberosKey{}
	this.Key = key
	return &this
}

// NewIpamsvcKerberosKeyWithDefaults instantiates a new IpamsvcKerberosKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcKerberosKeyWithDefaults() *IpamsvcKerberosKey {
	this := IpamsvcKerberosKey{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *IpamsvcKerberosKey) GetAlgorithm() string {
	if o == nil || IsNil(o.Algorithm) {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcKerberosKey) GetAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *IpamsvcKerberosKey) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *IpamsvcKerberosKey) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *IpamsvcKerberosKey) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcKerberosKey) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *IpamsvcKerberosKey) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *IpamsvcKerberosKey) SetDomain(v string) {
	o.Domain = &v
}

// GetKey returns the Key field value
func (o *IpamsvcKerberosKey) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *IpamsvcKerberosKey) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *IpamsvcKerberosKey) SetKey(v string) {
	o.Key = v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *IpamsvcKerberosKey) GetPrincipal() string {
	if o == nil || IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcKerberosKey) GetPrincipalOk() (*string, bool) {
	if o == nil || IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *IpamsvcKerberosKey) HasPrincipal() bool {
	if o != nil && !IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *IpamsvcKerberosKey) SetPrincipal(v string) {
	o.Principal = &v
}

// GetUploadedAt returns the UploadedAt field value if set, zero value otherwise.
func (o *IpamsvcKerberosKey) GetUploadedAt() string {
	if o == nil || IsNil(o.UploadedAt) {
		var ret string
		return ret
	}
	return *o.UploadedAt
}

// GetUploadedAtOk returns a tuple with the UploadedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcKerberosKey) GetUploadedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UploadedAt) {
		return nil, false
	}
	return o.UploadedAt, true
}

// HasUploadedAt returns a boolean if a field has been set.
func (o *IpamsvcKerberosKey) HasUploadedAt() bool {
	if o != nil && !IsNil(o.UploadedAt) {
		return true
	}

	return false
}

// SetUploadedAt gets a reference to the given string and assigns it to the UploadedAt field.
func (o *IpamsvcKerberosKey) SetUploadedAt(v string) {
	o.UploadedAt = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *IpamsvcKerberosKey) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcKerberosKey) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *IpamsvcKerberosKey) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *IpamsvcKerberosKey) SetVersion(v int64) {
	o.Version = &v
}

func (o IpamsvcKerberosKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcKerberosKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	toSerialize["key"] = o.Key
	if !IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}
	if !IsNil(o.UploadedAt) {
		toSerialize["uploaded_at"] = o.UploadedAt
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableIpamsvcKerberosKey struct {
	value *IpamsvcKerberosKey
	isSet bool
}

func (v NullableIpamsvcKerberosKey) Get() *IpamsvcKerberosKey {
	return v.value
}

func (v *NullableIpamsvcKerberosKey) Set(val *IpamsvcKerberosKey) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcKerberosKey) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcKerberosKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcKerberosKey(val *IpamsvcKerberosKey) *NullableIpamsvcKerberosKey {
	return &NullableIpamsvcKerberosKey{value: val, isSet: true}
}

func (v NullableIpamsvcKerberosKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcKerberosKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
