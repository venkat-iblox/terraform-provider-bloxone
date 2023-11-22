/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_config

import (
	"encoding/json"
	"time"
)

// checks if the ConfigAuthZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigAuthZone{}

// ConfigAuthZone Authoritative zone.
type ConfigAuthZone struct {
	// Optional. Comment for zone configuration.
	Comment *string `json:"comment,omitempty"`
	// Time when the object has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.
	Disabled *bool `json:"disabled,omitempty"`
	// Optional. DNS primaries external to BloxOne DDI. Order is not significant.
	ExternalPrimaries []ConfigExternalPrimary `json:"external_primaries,omitempty"`
	// list of external providers for the auth zone.
	ExternalProviders []AuthZoneExternalProvider `json:"external_providers,omitempty"`
	// DNS secondaries external to BloxOne DDI. Order is not significant.
	ExternalSecondaries []ConfigExternalSecondary `json:"external_secondaries,omitempty"`
	// Zone FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation.
	Fqdn *string `json:"fqdn,omitempty"`
	// _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_.
	GssTsigEnabled *bool `json:"gss_tsig_enabled,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The list of the inheritance assigned hosts of the object.
	InheritanceAssignedHosts []Inheritance2AssignedHost `json:"inheritance_assigned_hosts,omitempty"`
	InheritanceSources       *ConfigAuthZoneInheritance `json:"inheritance_sources,omitempty"`
	// On-create-only. SOA serial is allowed to be set when the authoritative zone is created.
	InitialSoaSerial *int64 `json:"initial_soa_serial,omitempty"`
	// Optional. BloxOne DDI hosts acting as internal secondaries. Order is not significant.
	InternalSecondaries []ConfigInternalSecondary `json:"internal_secondaries,omitempty"`
	// Reverse zone network address in the following format: \"ip-address/cidr\". Defaults to empty.
	MappedSubnet *string `json:"mapped_subnet,omitempty"`
	// Zone mapping type. Allowed values:  * _forward_,  * _ipv4_reverse_.  * _ipv6_reverse_.  Defaults to forward.
	Mapping *string `json:"mapping,omitempty"`
	// Also notify all external secondary DNS servers if enabled.  Defaults to _false_.
	Notify *bool `json:"notify,omitempty"`
	// The resource identifier.
	Nsgs []string `json:"nsgs,omitempty"`
	// The resource identifier.
	Parent *string `json:"parent,omitempty"`
	// Primary type for an authoritative zone. Read only after creation. Allowed values:  * _external_: zone data owned by an external nameserver,  * _cloud_: zone data is owned by a BloxOne DDI host.
	PrimaryType *string `json:"primary_type,omitempty"`
	// Zone FQDN in punycode.
	ProtocolFqdn *string `json:"protocol_fqdn,omitempty"`
	// Optional. Clients must match this ACL to make authoritative queries. Also used for recursive queries if that ACL is unset.  Defaults to empty.
	QueryAcl []ConfigACLItem `json:"query_acl,omitempty"`
	// Tagging specifics.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Optional. Clients must match this ACL to receive zone transfers.
	TransferAcl []ConfigACLItem `json:"transfer_acl,omitempty"`
	// Optional. Specifies which hosts are allowed to submit Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.  Defaults to empty.
	UpdateAcl []ConfigACLItem `json:"update_acl,omitempty"`
	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Optional. Use default forwarders to resolve queries for subzones.  Defaults to _true_.
	UseForwardersForSubzones *bool `json:"use_forwarders_for_subzones,omitempty"`
	// The resource identifier.
	View *string `json:"view,omitempty"`
	// The list of an auth zone warnings.
	Warnings      []ConfigWarning      `json:"warnings,omitempty"`
	ZoneAuthority *ConfigZoneAuthority `json:"zone_authority,omitempty"`
}

// NewConfigAuthZone instantiates a new ConfigAuthZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigAuthZone() *ConfigAuthZone {
	this := ConfigAuthZone{}
	return &this
}

// NewConfigAuthZoneWithDefaults instantiates a new ConfigAuthZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigAuthZoneWithDefaults() *ConfigAuthZone {
	this := ConfigAuthZone{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ConfigAuthZone) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ConfigAuthZone) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ConfigAuthZone) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetExternalPrimaries returns the ExternalPrimaries field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetExternalPrimaries() []ConfigExternalPrimary {
	if o == nil || IsNil(o.ExternalPrimaries) {
		var ret []ConfigExternalPrimary
		return ret
	}
	return o.ExternalPrimaries
}

// GetExternalPrimariesOk returns a tuple with the ExternalPrimaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetExternalPrimariesOk() ([]ConfigExternalPrimary, bool) {
	if o == nil || IsNil(o.ExternalPrimaries) {
		return nil, false
	}
	return o.ExternalPrimaries, true
}

// HasExternalPrimaries returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasExternalPrimaries() bool {
	if o != nil && !IsNil(o.ExternalPrimaries) {
		return true
	}

	return false
}

// SetExternalPrimaries gets a reference to the given []ConfigExternalPrimary and assigns it to the ExternalPrimaries field.
func (o *ConfigAuthZone) SetExternalPrimaries(v []ConfigExternalPrimary) {
	o.ExternalPrimaries = v
}

// GetExternalProviders returns the ExternalProviders field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetExternalProviders() []AuthZoneExternalProvider {
	if o == nil || IsNil(o.ExternalProviders) {
		var ret []AuthZoneExternalProvider
		return ret
	}
	return o.ExternalProviders
}

// GetExternalProvidersOk returns a tuple with the ExternalProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetExternalProvidersOk() ([]AuthZoneExternalProvider, bool) {
	if o == nil || IsNil(o.ExternalProviders) {
		return nil, false
	}
	return o.ExternalProviders, true
}

// HasExternalProviders returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasExternalProviders() bool {
	if o != nil && !IsNil(o.ExternalProviders) {
		return true
	}

	return false
}

// SetExternalProviders gets a reference to the given []AuthZoneExternalProvider and assigns it to the ExternalProviders field.
func (o *ConfigAuthZone) SetExternalProviders(v []AuthZoneExternalProvider) {
	o.ExternalProviders = v
}

// GetExternalSecondaries returns the ExternalSecondaries field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetExternalSecondaries() []ConfigExternalSecondary {
	if o == nil || IsNil(o.ExternalSecondaries) {
		var ret []ConfigExternalSecondary
		return ret
	}
	return o.ExternalSecondaries
}

// GetExternalSecondariesOk returns a tuple with the ExternalSecondaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetExternalSecondariesOk() ([]ConfigExternalSecondary, bool) {
	if o == nil || IsNil(o.ExternalSecondaries) {
		return nil, false
	}
	return o.ExternalSecondaries, true
}

// HasExternalSecondaries returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasExternalSecondaries() bool {
	if o != nil && !IsNil(o.ExternalSecondaries) {
		return true
	}

	return false
}

// SetExternalSecondaries gets a reference to the given []ConfigExternalSecondary and assigns it to the ExternalSecondaries field.
func (o *ConfigAuthZone) SetExternalSecondaries(v []ConfigExternalSecondary) {
	o.ExternalSecondaries = v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *ConfigAuthZone) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetGssTsigEnabled returns the GssTsigEnabled field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetGssTsigEnabled() bool {
	if o == nil || IsNil(o.GssTsigEnabled) {
		var ret bool
		return ret
	}
	return *o.GssTsigEnabled
}

// GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetGssTsigEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.GssTsigEnabled) {
		return nil, false
	}
	return o.GssTsigEnabled, true
}

// HasGssTsigEnabled returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasGssTsigEnabled() bool {
	if o != nil && !IsNil(o.GssTsigEnabled) {
		return true
	}

	return false
}

// SetGssTsigEnabled gets a reference to the given bool and assigns it to the GssTsigEnabled field.
func (o *ConfigAuthZone) SetGssTsigEnabled(v bool) {
	o.GssTsigEnabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConfigAuthZone) SetId(v string) {
	o.Id = &v
}

// GetInheritanceAssignedHosts returns the InheritanceAssignedHosts field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetInheritanceAssignedHosts() []Inheritance2AssignedHost {
	if o == nil || IsNil(o.InheritanceAssignedHosts) {
		var ret []Inheritance2AssignedHost
		return ret
	}
	return o.InheritanceAssignedHosts
}

// GetInheritanceAssignedHostsOk returns a tuple with the InheritanceAssignedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetInheritanceAssignedHostsOk() ([]Inheritance2AssignedHost, bool) {
	if o == nil || IsNil(o.InheritanceAssignedHosts) {
		return nil, false
	}
	return o.InheritanceAssignedHosts, true
}

// HasInheritanceAssignedHosts returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasInheritanceAssignedHosts() bool {
	if o != nil && !IsNil(o.InheritanceAssignedHosts) {
		return true
	}

	return false
}

// SetInheritanceAssignedHosts gets a reference to the given []Inheritance2AssignedHost and assigns it to the InheritanceAssignedHosts field.
func (o *ConfigAuthZone) SetInheritanceAssignedHosts(v []Inheritance2AssignedHost) {
	o.InheritanceAssignedHosts = v
}

// GetInheritanceSources returns the InheritanceSources field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetInheritanceSources() ConfigAuthZoneInheritance {
	if o == nil || IsNil(o.InheritanceSources) {
		var ret ConfigAuthZoneInheritance
		return ret
	}
	return *o.InheritanceSources
}

// GetInheritanceSourcesOk returns a tuple with the InheritanceSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetInheritanceSourcesOk() (*ConfigAuthZoneInheritance, bool) {
	if o == nil || IsNil(o.InheritanceSources) {
		return nil, false
	}
	return o.InheritanceSources, true
}

// HasInheritanceSources returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasInheritanceSources() bool {
	if o != nil && !IsNil(o.InheritanceSources) {
		return true
	}

	return false
}

// SetInheritanceSources gets a reference to the given ConfigAuthZoneInheritance and assigns it to the InheritanceSources field.
func (o *ConfigAuthZone) SetInheritanceSources(v ConfigAuthZoneInheritance) {
	o.InheritanceSources = &v
}

// GetInitialSoaSerial returns the InitialSoaSerial field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetInitialSoaSerial() int64 {
	if o == nil || IsNil(o.InitialSoaSerial) {
		var ret int64
		return ret
	}
	return *o.InitialSoaSerial
}

// GetInitialSoaSerialOk returns a tuple with the InitialSoaSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetInitialSoaSerialOk() (*int64, bool) {
	if o == nil || IsNil(o.InitialSoaSerial) {
		return nil, false
	}
	return o.InitialSoaSerial, true
}

// HasInitialSoaSerial returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasInitialSoaSerial() bool {
	if o != nil && !IsNil(o.InitialSoaSerial) {
		return true
	}

	return false
}

// SetInitialSoaSerial gets a reference to the given int64 and assigns it to the InitialSoaSerial field.
func (o *ConfigAuthZone) SetInitialSoaSerial(v int64) {
	o.InitialSoaSerial = &v
}

// GetInternalSecondaries returns the InternalSecondaries field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetInternalSecondaries() []ConfigInternalSecondary {
	if o == nil || IsNil(o.InternalSecondaries) {
		var ret []ConfigInternalSecondary
		return ret
	}
	return o.InternalSecondaries
}

// GetInternalSecondariesOk returns a tuple with the InternalSecondaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetInternalSecondariesOk() ([]ConfigInternalSecondary, bool) {
	if o == nil || IsNil(o.InternalSecondaries) {
		return nil, false
	}
	return o.InternalSecondaries, true
}

// HasInternalSecondaries returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasInternalSecondaries() bool {
	if o != nil && !IsNil(o.InternalSecondaries) {
		return true
	}

	return false
}

// SetInternalSecondaries gets a reference to the given []ConfigInternalSecondary and assigns it to the InternalSecondaries field.
func (o *ConfigAuthZone) SetInternalSecondaries(v []ConfigInternalSecondary) {
	o.InternalSecondaries = v
}

// GetMappedSubnet returns the MappedSubnet field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetMappedSubnet() string {
	if o == nil || IsNil(o.MappedSubnet) {
		var ret string
		return ret
	}
	return *o.MappedSubnet
}

// GetMappedSubnetOk returns a tuple with the MappedSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetMappedSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.MappedSubnet) {
		return nil, false
	}
	return o.MappedSubnet, true
}

// HasMappedSubnet returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasMappedSubnet() bool {
	if o != nil && !IsNil(o.MappedSubnet) {
		return true
	}

	return false
}

// SetMappedSubnet gets a reference to the given string and assigns it to the MappedSubnet field.
func (o *ConfigAuthZone) SetMappedSubnet(v string) {
	o.MappedSubnet = &v
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetMapping() string {
	if o == nil || IsNil(o.Mapping) {
		var ret string
		return ret
	}
	return *o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetMappingOk() (*string, bool) {
	if o == nil || IsNil(o.Mapping) {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasMapping() bool {
	if o != nil && !IsNil(o.Mapping) {
		return true
	}

	return false
}

// SetMapping gets a reference to the given string and assigns it to the Mapping field.
func (o *ConfigAuthZone) SetMapping(v string) {
	o.Mapping = &v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetNotify() bool {
	if o == nil || IsNil(o.Notify) {
		var ret bool
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetNotifyOk() (*bool, bool) {
	if o == nil || IsNil(o.Notify) {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasNotify() bool {
	if o != nil && !IsNil(o.Notify) {
		return true
	}

	return false
}

// SetNotify gets a reference to the given bool and assigns it to the Notify field.
func (o *ConfigAuthZone) SetNotify(v bool) {
	o.Notify = &v
}

// GetNsgs returns the Nsgs field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetNsgs() []string {
	if o == nil || IsNil(o.Nsgs) {
		var ret []string
		return ret
	}
	return o.Nsgs
}

// GetNsgsOk returns a tuple with the Nsgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetNsgsOk() ([]string, bool) {
	if o == nil || IsNil(o.Nsgs) {
		return nil, false
	}
	return o.Nsgs, true
}

// HasNsgs returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasNsgs() bool {
	if o != nil && !IsNil(o.Nsgs) {
		return true
	}

	return false
}

// SetNsgs gets a reference to the given []string and assigns it to the Nsgs field.
func (o *ConfigAuthZone) SetNsgs(v []string) {
	o.Nsgs = v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *ConfigAuthZone) SetParent(v string) {
	o.Parent = &v
}

// GetPrimaryType returns the PrimaryType field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetPrimaryType() string {
	if o == nil || IsNil(o.PrimaryType) {
		var ret string
		return ret
	}
	return *o.PrimaryType
}

// GetPrimaryTypeOk returns a tuple with the PrimaryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetPrimaryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryType) {
		return nil, false
	}
	return o.PrimaryType, true
}

// HasPrimaryType returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasPrimaryType() bool {
	if o != nil && !IsNil(o.PrimaryType) {
		return true
	}

	return false
}

// SetPrimaryType gets a reference to the given string and assigns it to the PrimaryType field.
func (o *ConfigAuthZone) SetPrimaryType(v string) {
	o.PrimaryType = &v
}

// GetProtocolFqdn returns the ProtocolFqdn field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetProtocolFqdn() string {
	if o == nil || IsNil(o.ProtocolFqdn) {
		var ret string
		return ret
	}
	return *o.ProtocolFqdn
}

// GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetProtocolFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolFqdn) {
		return nil, false
	}
	return o.ProtocolFqdn, true
}

// HasProtocolFqdn returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasProtocolFqdn() bool {
	if o != nil && !IsNil(o.ProtocolFqdn) {
		return true
	}

	return false
}

// SetProtocolFqdn gets a reference to the given string and assigns it to the ProtocolFqdn field.
func (o *ConfigAuthZone) SetProtocolFqdn(v string) {
	o.ProtocolFqdn = &v
}

// GetQueryAcl returns the QueryAcl field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetQueryAcl() []ConfigACLItem {
	if o == nil || IsNil(o.QueryAcl) {
		var ret []ConfigACLItem
		return ret
	}
	return o.QueryAcl
}

// GetQueryAclOk returns a tuple with the QueryAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetQueryAclOk() ([]ConfigACLItem, bool) {
	if o == nil || IsNil(o.QueryAcl) {
		return nil, false
	}
	return o.QueryAcl, true
}

// HasQueryAcl returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasQueryAcl() bool {
	if o != nil && !IsNil(o.QueryAcl) {
		return true
	}

	return false
}

// SetQueryAcl gets a reference to the given []ConfigACLItem and assigns it to the QueryAcl field.
func (o *ConfigAuthZone) SetQueryAcl(v []ConfigACLItem) {
	o.QueryAcl = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *ConfigAuthZone) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetTransferAcl returns the TransferAcl field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetTransferAcl() []ConfigACLItem {
	if o == nil || IsNil(o.TransferAcl) {
		var ret []ConfigACLItem
		return ret
	}
	return o.TransferAcl
}

// GetTransferAclOk returns a tuple with the TransferAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetTransferAclOk() ([]ConfigACLItem, bool) {
	if o == nil || IsNil(o.TransferAcl) {
		return nil, false
	}
	return o.TransferAcl, true
}

// HasTransferAcl returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasTransferAcl() bool {
	if o != nil && !IsNil(o.TransferAcl) {
		return true
	}

	return false
}

// SetTransferAcl gets a reference to the given []ConfigACLItem and assigns it to the TransferAcl field.
func (o *ConfigAuthZone) SetTransferAcl(v []ConfigACLItem) {
	o.TransferAcl = v
}

// GetUpdateAcl returns the UpdateAcl field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetUpdateAcl() []ConfigACLItem {
	if o == nil || IsNil(o.UpdateAcl) {
		var ret []ConfigACLItem
		return ret
	}
	return o.UpdateAcl
}

// GetUpdateAclOk returns a tuple with the UpdateAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetUpdateAclOk() ([]ConfigACLItem, bool) {
	if o == nil || IsNil(o.UpdateAcl) {
		return nil, false
	}
	return o.UpdateAcl, true
}

// HasUpdateAcl returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasUpdateAcl() bool {
	if o != nil && !IsNil(o.UpdateAcl) {
		return true
	}

	return false
}

// SetUpdateAcl gets a reference to the given []ConfigACLItem and assigns it to the UpdateAcl field.
func (o *ConfigAuthZone) SetUpdateAcl(v []ConfigACLItem) {
	o.UpdateAcl = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ConfigAuthZone) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUseForwardersForSubzones returns the UseForwardersForSubzones field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetUseForwardersForSubzones() bool {
	if o == nil || IsNil(o.UseForwardersForSubzones) {
		var ret bool
		return ret
	}
	return *o.UseForwardersForSubzones
}

// GetUseForwardersForSubzonesOk returns a tuple with the UseForwardersForSubzones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetUseForwardersForSubzonesOk() (*bool, bool) {
	if o == nil || IsNil(o.UseForwardersForSubzones) {
		return nil, false
	}
	return o.UseForwardersForSubzones, true
}

// HasUseForwardersForSubzones returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasUseForwardersForSubzones() bool {
	if o != nil && !IsNil(o.UseForwardersForSubzones) {
		return true
	}

	return false
}

// SetUseForwardersForSubzones gets a reference to the given bool and assigns it to the UseForwardersForSubzones field.
func (o *ConfigAuthZone) SetUseForwardersForSubzones(v bool) {
	o.UseForwardersForSubzones = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetView() string {
	if o == nil || IsNil(o.View) {
		var ret string
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetViewOk() (*string, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given string and assigns it to the View field.
func (o *ConfigAuthZone) SetView(v string) {
	o.View = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetWarnings() []ConfigWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []ConfigWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetWarningsOk() ([]ConfigWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []ConfigWarning and assigns it to the Warnings field.
func (o *ConfigAuthZone) SetWarnings(v []ConfigWarning) {
	o.Warnings = v
}

// GetZoneAuthority returns the ZoneAuthority field value if set, zero value otherwise.
func (o *ConfigAuthZone) GetZoneAuthority() ConfigZoneAuthority {
	if o == nil || IsNil(o.ZoneAuthority) {
		var ret ConfigZoneAuthority
		return ret
	}
	return *o.ZoneAuthority
}

// GetZoneAuthorityOk returns a tuple with the ZoneAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigAuthZone) GetZoneAuthorityOk() (*ConfigZoneAuthority, bool) {
	if o == nil || IsNil(o.ZoneAuthority) {
		return nil, false
	}
	return o.ZoneAuthority, true
}

// HasZoneAuthority returns a boolean if a field has been set.
func (o *ConfigAuthZone) HasZoneAuthority() bool {
	if o != nil && !IsNil(o.ZoneAuthority) {
		return true
	}

	return false
}

// SetZoneAuthority gets a reference to the given ConfigZoneAuthority and assigns it to the ZoneAuthority field.
func (o *ConfigAuthZone) SetZoneAuthority(v ConfigZoneAuthority) {
	o.ZoneAuthority = &v
}

func (o ConfigAuthZone) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigAuthZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.ExternalPrimaries) {
		toSerialize["external_primaries"] = o.ExternalPrimaries
	}
	if !IsNil(o.ExternalProviders) {
		toSerialize["external_providers"] = o.ExternalProviders
	}
	if !IsNil(o.ExternalSecondaries) {
		toSerialize["external_secondaries"] = o.ExternalSecondaries
	}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.GssTsigEnabled) {
		toSerialize["gss_tsig_enabled"] = o.GssTsigEnabled
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InheritanceAssignedHosts) {
		toSerialize["inheritance_assigned_hosts"] = o.InheritanceAssignedHosts
	}
	if !IsNil(o.InheritanceSources) {
		toSerialize["inheritance_sources"] = o.InheritanceSources
	}
	if !IsNil(o.InitialSoaSerial) {
		toSerialize["initial_soa_serial"] = o.InitialSoaSerial
	}
	if !IsNil(o.InternalSecondaries) {
		toSerialize["internal_secondaries"] = o.InternalSecondaries
	}
	if !IsNil(o.MappedSubnet) {
		toSerialize["mapped_subnet"] = o.MappedSubnet
	}
	if !IsNil(o.Mapping) {
		toSerialize["mapping"] = o.Mapping
	}
	if !IsNil(o.Notify) {
		toSerialize["notify"] = o.Notify
	}
	if !IsNil(o.Nsgs) {
		toSerialize["nsgs"] = o.Nsgs
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.PrimaryType) {
		toSerialize["primary_type"] = o.PrimaryType
	}
	if !IsNil(o.ProtocolFqdn) {
		toSerialize["protocol_fqdn"] = o.ProtocolFqdn
	}
	if !IsNil(o.QueryAcl) {
		toSerialize["query_acl"] = o.QueryAcl
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TransferAcl) {
		toSerialize["transfer_acl"] = o.TransferAcl
	}
	if !IsNil(o.UpdateAcl) {
		toSerialize["update_acl"] = o.UpdateAcl
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UseForwardersForSubzones) {
		toSerialize["use_forwarders_for_subzones"] = o.UseForwardersForSubzones
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.ZoneAuthority) {
		toSerialize["zone_authority"] = o.ZoneAuthority
	}
	return toSerialize, nil
}

type NullableConfigAuthZone struct {
	value *ConfigAuthZone
	isSet bool
}

func (v NullableConfigAuthZone) Get() *ConfigAuthZone {
	return v.value
}

func (v *NullableConfigAuthZone) Set(val *ConfigAuthZone) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigAuthZone) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigAuthZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigAuthZone(val *ConfigAuthZone) *NullableConfigAuthZone {
	return &NullableConfigAuthZone{value: val, isSet: true}
}

func (v NullableConfigAuthZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigAuthZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
