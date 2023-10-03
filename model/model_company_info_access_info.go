/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.30
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the CompanyInfoAccessInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyInfoAccessInfo{}

// CompanyInfoAccessInfo struct for CompanyInfoAccessInfo
type CompanyInfoAccessInfo struct {
	Role *UserCompanyRole `json:"role,omitempty"`
	Permissions *Permissions `json:"permissions,omitempty"`
	// Company activated through accountant
	ThroughAccountant NullableBool `json:"through_accountant,omitempty"`
}

// NewCompanyInfoAccessInfo instantiates a new CompanyInfoAccessInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyInfoAccessInfo() *CompanyInfoAccessInfo {
	this := CompanyInfoAccessInfo{}
	return &this
}

// NewCompanyInfoAccessInfoWithDefaults instantiates a new CompanyInfoAccessInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyInfoAccessInfoWithDefaults() *CompanyInfoAccessInfo {
	this := CompanyInfoAccessInfo{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *CompanyInfoAccessInfo) GetRole() UserCompanyRole {
	if o == nil || IsNil(o.Role) {
		var ret UserCompanyRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyInfoAccessInfo) GetRoleOk() (*UserCompanyRole, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *CompanyInfoAccessInfo) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given UserCompanyRole and assigns it to the Role field.
func (o *CompanyInfoAccessInfo) SetRole(v UserCompanyRole) *CompanyInfoAccessInfo {
	o.Role = &v
	return o
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *CompanyInfoAccessInfo) GetPermissions() Permissions {
	if o == nil || IsNil(o.Permissions) {
		var ret Permissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyInfoAccessInfo) GetPermissionsOk() (*Permissions, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CompanyInfoAccessInfo) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given Permissions and assigns it to the Permissions field.
func (o *CompanyInfoAccessInfo) SetPermissions(v Permissions) *CompanyInfoAccessInfo {
	o.Permissions = &v
	return o
}

// GetThroughAccountant returns the ThroughAccountant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyInfoAccessInfo) GetThroughAccountant() bool {
	if o == nil || IsNil(o.ThroughAccountant.Get()) {
		var ret bool
		return ret
	}
	return *o.ThroughAccountant.Get()
}

// GetThroughAccountantOk returns a tuple with the ThroughAccountant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyInfoAccessInfo) GetThroughAccountantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThroughAccountant.Get(), o.ThroughAccountant.IsSet()
}

// HasThroughAccountant returns a boolean if a field has been set.
func (o *CompanyInfoAccessInfo) HasThroughAccountant() bool {
	if o != nil && o.ThroughAccountant.IsSet() {
		return true
	}

	return false
}

// SetThroughAccountant gets a reference to the given NullableBool and assigns it to the ThroughAccountant field.
func (o *CompanyInfoAccessInfo) SetThroughAccountant(v bool) *CompanyInfoAccessInfo {
	o.ThroughAccountant.Set(&v)
	return o
}
// SetThroughAccountantNil sets the value for ThroughAccountant to be an explicit nil
func (o *CompanyInfoAccessInfo) SetThroughAccountantNil() *CompanyInfoAccessInfo {
	o.ThroughAccountant.Set(nil)
	return o
}

// UnsetThroughAccountant ensures that no value is present for ThroughAccountant, not even an explicit nil
func (o *CompanyInfoAccessInfo) UnsetThroughAccountant() {
	o.ThroughAccountant.Unset()
}

func (o CompanyInfoAccessInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyInfoAccessInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if o.ThroughAccountant.IsSet() {
		toSerialize["through_accountant"] = o.ThroughAccountant.Get()
	}
	return toSerialize, nil
}

type NullableCompanyInfoAccessInfo struct {
	value *CompanyInfoAccessInfo
	isSet bool
}

func (v NullableCompanyInfoAccessInfo) Get() *CompanyInfoAccessInfo {
	return v.value
}

func (v *NullableCompanyInfoAccessInfo) Set(val *CompanyInfoAccessInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyInfoAccessInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyInfoAccessInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyInfoAccessInfo(val *CompanyInfoAccessInfo) *NullableCompanyInfoAccessInfo {
	return &NullableCompanyInfoAccessInfo{value: val, isSet: true}
}

func (v NullableCompanyInfoAccessInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyInfoAccessInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


