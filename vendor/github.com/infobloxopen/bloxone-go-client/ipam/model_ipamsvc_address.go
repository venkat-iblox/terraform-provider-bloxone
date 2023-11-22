/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"time"
)

// checks if the IpamsvcAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcAddress{}

// IpamsvcAddress An __Address__ object (_ipam/address_) represents any single IP address within a given IP space.
type IpamsvcAddress struct {
	// The address in form \"a.b.c.d\".
	Address string `json:"address"`
	// The description for the address object. May contain 0 to 1024 characters. Can include UTF-8.
	Comment *string `json:"comment,omitempty"`
	// Time when the object has been created.
	CreatedAt *time.Time       `json:"created_at,omitempty"`
	DhcpInfo  *IpamsvcDHCPInfo `json:"dhcp_info,omitempty"`
	// Read only. Represent the value of the same field in the associated _dhcp/fixed_address_ object.
	DisableDhcp *bool `json:"disable_dhcp,omitempty"`
	// The discovery attributes for this address in JSON format.
	DiscoveryAttrs map[string]interface{} `json:"discovery_attrs,omitempty"`
	// The discovery metadata for this address in JSON format.
	DiscoveryMetadata map[string]interface{} `json:"discovery_metadata,omitempty"`
	// The resource identifier.
	Host *string `json:"host,omitempty"`
	// The hardware address associated with this IP address.
	Hwaddr *string `json:"hwaddr,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The name of the network interface card (NIC) associated with the address, if any.
	Interface *string `json:"interface,omitempty"`
	// The list of all names associated with this address.
	Names []IpamsvcName `json:"names,omitempty"`
	// The resource identifier.
	Parent *string `json:"parent,omitempty"`
	// The type of protocol (_ip4_ or _ip6_).
	Protocol *string `json:"protocol,omitempty"`
	// The resource identifier.
	Range *string `json:"range,omitempty"`
	// The resource identifier.
	Space *string `json:"space,omitempty"`
	// The state of the address (_used_ or _free_).
	State *string `json:"state,omitempty"`
	// The tags for this address in JSON format.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The usage is a combination of indicators, each tracking a specific associated use. Listed below are usage indicators with their meaning:  usage indicator        | description  ---------------------- | --------------------------------  _IPAM_                 |  Address was created by the IPAM component.  _IPAM_, _RESERVED_     |  Address was created by the API call _ipam/address_ or _ipam/host_.  _IPAM_, _NETWORK_      |  Address was automatically created by the IPAM component and is the network address of the parent subnet.  _IPAM_, _BROADCAST_    |  Address was automatically created by the IPAM component and is the broadcast address of the parent subnet.  _DHCP_                 |  Address was created by the DHCP component.  _DHCP_, _FIXEDADDRESS_ |  Address was created by the API call _dhcp/fixed_address_.  _DHCP_, _LEASED_       |  An active lease for that address was issued by a DHCP server.  _DHCP_, _DISABLED_     |  Address is disabled.  _DNS_                  |  Address is used by one or more DNS records.  _DISCOVERED_           |  Address is discovered by some network discovery probe like Network Insight or NetMRI in NIOS.
	Usage []string `json:"usage,omitempty"`
}

// NewIpamsvcAddress instantiates a new IpamsvcAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcAddress(address string) *IpamsvcAddress {
	this := IpamsvcAddress{}
	this.Address = address
	return &this
}

// NewIpamsvcAddressWithDefaults instantiates a new IpamsvcAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcAddressWithDefaults() *IpamsvcAddress {
	this := IpamsvcAddress{}
	return &this
}

// GetAddress returns the Address field value
func (o *IpamsvcAddress) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *IpamsvcAddress) SetAddress(v string) {
	o.Address = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *IpamsvcAddress) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *IpamsvcAddress) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDhcpInfo returns the DhcpInfo field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetDhcpInfo() IpamsvcDHCPInfo {
	if o == nil || IsNil(o.DhcpInfo) {
		var ret IpamsvcDHCPInfo
		return ret
	}
	return *o.DhcpInfo
}

// GetDhcpInfoOk returns a tuple with the DhcpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetDhcpInfoOk() (*IpamsvcDHCPInfo, bool) {
	if o == nil || IsNil(o.DhcpInfo) {
		return nil, false
	}
	return o.DhcpInfo, true
}

// HasDhcpInfo returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasDhcpInfo() bool {
	if o != nil && !IsNil(o.DhcpInfo) {
		return true
	}

	return false
}

// SetDhcpInfo gets a reference to the given IpamsvcDHCPInfo and assigns it to the DhcpInfo field.
func (o *IpamsvcAddress) SetDhcpInfo(v IpamsvcDHCPInfo) {
	o.DhcpInfo = &v
}

// GetDisableDhcp returns the DisableDhcp field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetDisableDhcp() bool {
	if o == nil || IsNil(o.DisableDhcp) {
		var ret bool
		return ret
	}
	return *o.DisableDhcp
}

// GetDisableDhcpOk returns a tuple with the DisableDhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetDisableDhcpOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableDhcp) {
		return nil, false
	}
	return o.DisableDhcp, true
}

// HasDisableDhcp returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasDisableDhcp() bool {
	if o != nil && !IsNil(o.DisableDhcp) {
		return true
	}

	return false
}

// SetDisableDhcp gets a reference to the given bool and assigns it to the DisableDhcp field.
func (o *IpamsvcAddress) SetDisableDhcp(v bool) {
	o.DisableDhcp = &v
}

// GetDiscoveryAttrs returns the DiscoveryAttrs field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetDiscoveryAttrs() map[string]interface{} {
	if o == nil || IsNil(o.DiscoveryAttrs) {
		var ret map[string]interface{}
		return ret
	}
	return o.DiscoveryAttrs
}

// GetDiscoveryAttrsOk returns a tuple with the DiscoveryAttrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetDiscoveryAttrsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DiscoveryAttrs) {
		return map[string]interface{}{}, false
	}
	return o.DiscoveryAttrs, true
}

// HasDiscoveryAttrs returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasDiscoveryAttrs() bool {
	if o != nil && !IsNil(o.DiscoveryAttrs) {
		return true
	}

	return false
}

// SetDiscoveryAttrs gets a reference to the given map[string]interface{} and assigns it to the DiscoveryAttrs field.
func (o *IpamsvcAddress) SetDiscoveryAttrs(v map[string]interface{}) {
	o.DiscoveryAttrs = v
}

// GetDiscoveryMetadata returns the DiscoveryMetadata field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetDiscoveryMetadata() map[string]interface{} {
	if o == nil || IsNil(o.DiscoveryMetadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.DiscoveryMetadata
}

// GetDiscoveryMetadataOk returns a tuple with the DiscoveryMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetDiscoveryMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DiscoveryMetadata) {
		return map[string]interface{}{}, false
	}
	return o.DiscoveryMetadata, true
}

// HasDiscoveryMetadata returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasDiscoveryMetadata() bool {
	if o != nil && !IsNil(o.DiscoveryMetadata) {
		return true
	}

	return false
}

// SetDiscoveryMetadata gets a reference to the given map[string]interface{} and assigns it to the DiscoveryMetadata field.
func (o *IpamsvcAddress) SetDiscoveryMetadata(v map[string]interface{}) {
	o.DiscoveryMetadata = v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *IpamsvcAddress) SetHost(v string) {
	o.Host = &v
}

// GetHwaddr returns the Hwaddr field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetHwaddr() string {
	if o == nil || IsNil(o.Hwaddr) {
		var ret string
		return ret
	}
	return *o.Hwaddr
}

// GetHwaddrOk returns a tuple with the Hwaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetHwaddrOk() (*string, bool) {
	if o == nil || IsNil(o.Hwaddr) {
		return nil, false
	}
	return o.Hwaddr, true
}

// HasHwaddr returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasHwaddr() bool {
	if o != nil && !IsNil(o.Hwaddr) {
		return true
	}

	return false
}

// SetHwaddr gets a reference to the given string and assigns it to the Hwaddr field.
func (o *IpamsvcAddress) SetHwaddr(v string) {
	o.Hwaddr = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IpamsvcAddress) SetId(v string) {
	o.Id = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetInterface() string {
	if o == nil || IsNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *IpamsvcAddress) SetInterface(v string) {
	o.Interface = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetNames() []IpamsvcName {
	if o == nil || IsNil(o.Names) {
		var ret []IpamsvcName
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetNamesOk() ([]IpamsvcName, bool) {
	if o == nil || IsNil(o.Names) {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasNames() bool {
	if o != nil && !IsNil(o.Names) {
		return true
	}

	return false
}

// SetNames gets a reference to the given []IpamsvcName and assigns it to the Names field.
func (o *IpamsvcAddress) SetNames(v []IpamsvcName) {
	o.Names = v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *IpamsvcAddress) SetParent(v string) {
	o.Parent = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *IpamsvcAddress) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetRange() string {
	if o == nil || IsNil(o.Range) {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetRangeOk() (*string, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *IpamsvcAddress) SetRange(v string) {
	o.Range = &v
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetSpace() string {
	if o == nil || IsNil(o.Space) {
		var ret string
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetSpaceOk() (*string, bool) {
	if o == nil || IsNil(o.Space) {
		return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasSpace() bool {
	if o != nil && !IsNil(o.Space) {
		return true
	}

	return false
}

// SetSpace gets a reference to the given string and assigns it to the Space field.
func (o *IpamsvcAddress) SetSpace(v string) {
	o.Space = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IpamsvcAddress) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *IpamsvcAddress) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *IpamsvcAddress) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *IpamsvcAddress) GetUsage() []string {
	if o == nil || IsNil(o.Usage) {
		var ret []string
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAddress) GetUsageOk() ([]string, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *IpamsvcAddress) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []string and assigns it to the Usage field.
func (o *IpamsvcAddress) SetUsage(v []string) {
	o.Usage = v
}

func (o IpamsvcAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DhcpInfo) {
		toSerialize["dhcp_info"] = o.DhcpInfo
	}
	if !IsNil(o.DisableDhcp) {
		toSerialize["disable_dhcp"] = o.DisableDhcp
	}
	if !IsNil(o.DiscoveryAttrs) {
		toSerialize["discovery_attrs"] = o.DiscoveryAttrs
	}
	if !IsNil(o.DiscoveryMetadata) {
		toSerialize["discovery_metadata"] = o.DiscoveryMetadata
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Hwaddr) {
		toSerialize["hwaddr"] = o.Hwaddr
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !IsNil(o.Names) {
		toSerialize["names"] = o.Names
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !IsNil(o.Space) {
		toSerialize["space"] = o.Space
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

type NullableIpamsvcAddress struct {
	value *IpamsvcAddress
	isSet bool
}

func (v NullableIpamsvcAddress) Get() *IpamsvcAddress {
	return v.value
}

func (v *NullableIpamsvcAddress) Set(val *IpamsvcAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcAddress(val *IpamsvcAddress) *NullableIpamsvcAddress {
	return &NullableIpamsvcAddress{value: val, isSet: true}
}

func (v NullableIpamsvcAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
