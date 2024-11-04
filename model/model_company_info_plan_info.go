/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.3
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the CompanyInfoPlanInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyInfoPlanInfo{}

// CompanyInfoPlanInfo Company plan info
type CompanyInfoPlanInfo struct {
	Limits NullableCompanyInfoPlanInfoLimits `json:"limits,omitempty"`
	Functions NullableCompanyInfoPlanInfoFunctions `json:"functions,omitempty"`
	FunctionsStatus NullableCompanyInfoPlanInfoFunctionsStatus `json:"functions_status,omitempty"`
}

// NewCompanyInfoPlanInfo instantiates a new CompanyInfoPlanInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyInfoPlanInfo() *CompanyInfoPlanInfo {
	this := CompanyInfoPlanInfo{}
	return &this
}

// NewCompanyInfoPlanInfoWithDefaults instantiates a new CompanyInfoPlanInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyInfoPlanInfoWithDefaults() *CompanyInfoPlanInfo {
	this := CompanyInfoPlanInfo{}
	return &this
}

// GetLimits returns the Limits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyInfoPlanInfo) GetLimits() CompanyInfoPlanInfoLimits {
	if o == nil || IsNil(o.Limits.Get()) {
		var ret CompanyInfoPlanInfoLimits
		return ret
	}
	return *o.Limits.Get()
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyInfoPlanInfo) GetLimitsOk() (*CompanyInfoPlanInfoLimits, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limits.Get(), o.Limits.IsSet()
}

// HasLimits returns a boolean if a field has been set.
func (o *CompanyInfoPlanInfo) HasLimits() bool {
	if o != nil && o.Limits.IsSet() {
		return true
	}

	return false
}

// SetLimits gets a reference to the given NullableCompanyInfoPlanInfoLimits and assigns it to the Limits field.
func (o *CompanyInfoPlanInfo) SetLimits(v CompanyInfoPlanInfoLimits) *CompanyInfoPlanInfo {
	o.Limits.Set(&v)
		return o
}
// SetLimitsNil sets the value for Limits to be an explicit nil
func (o *CompanyInfoPlanInfo) SetLimitsNil() *CompanyInfoPlanInfo {
	o.Limits.Set(nil)
	return o
}

// UnsetLimits ensures that no value is present for Limits, not even an explicit nil
func (o *CompanyInfoPlanInfo) UnsetLimits() {
	o.Limits.Unset()
}

// GetFunctions returns the Functions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyInfoPlanInfo) GetFunctions() CompanyInfoPlanInfoFunctions {
	if o == nil || IsNil(o.Functions.Get()) {
		var ret CompanyInfoPlanInfoFunctions
		return ret
	}
	return *o.Functions.Get()
}

// GetFunctionsOk returns a tuple with the Functions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyInfoPlanInfo) GetFunctionsOk() (*CompanyInfoPlanInfoFunctions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Functions.Get(), o.Functions.IsSet()
}

// HasFunctions returns a boolean if a field has been set.
func (o *CompanyInfoPlanInfo) HasFunctions() bool {
	if o != nil && o.Functions.IsSet() {
		return true
	}

	return false
}

// SetFunctions gets a reference to the given NullableCompanyInfoPlanInfoFunctions and assigns it to the Functions field.
func (o *CompanyInfoPlanInfo) SetFunctions(v CompanyInfoPlanInfoFunctions) *CompanyInfoPlanInfo {
	o.Functions.Set(&v)
		return o
}
// SetFunctionsNil sets the value for Functions to be an explicit nil
func (o *CompanyInfoPlanInfo) SetFunctionsNil() *CompanyInfoPlanInfo {
	o.Functions.Set(nil)
	return o
}

// UnsetFunctions ensures that no value is present for Functions, not even an explicit nil
func (o *CompanyInfoPlanInfo) UnsetFunctions() {
	o.Functions.Unset()
}

// GetFunctionsStatus returns the FunctionsStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyInfoPlanInfo) GetFunctionsStatus() CompanyInfoPlanInfoFunctionsStatus {
	if o == nil || IsNil(o.FunctionsStatus.Get()) {
		var ret CompanyInfoPlanInfoFunctionsStatus
		return ret
	}
	return *o.FunctionsStatus.Get()
}

// GetFunctionsStatusOk returns a tuple with the FunctionsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyInfoPlanInfo) GetFunctionsStatusOk() (*CompanyInfoPlanInfoFunctionsStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.FunctionsStatus.Get(), o.FunctionsStatus.IsSet()
}

// HasFunctionsStatus returns a boolean if a field has been set.
func (o *CompanyInfoPlanInfo) HasFunctionsStatus() bool {
	if o != nil && o.FunctionsStatus.IsSet() {
		return true
	}

	return false
}

// SetFunctionsStatus gets a reference to the given NullableCompanyInfoPlanInfoFunctionsStatus and assigns it to the FunctionsStatus field.
func (o *CompanyInfoPlanInfo) SetFunctionsStatus(v CompanyInfoPlanInfoFunctionsStatus) *CompanyInfoPlanInfo {
	o.FunctionsStatus.Set(&v)
		return o
}
// SetFunctionsStatusNil sets the value for FunctionsStatus to be an explicit nil
func (o *CompanyInfoPlanInfo) SetFunctionsStatusNil() *CompanyInfoPlanInfo {
	o.FunctionsStatus.Set(nil)
	return o
}

// UnsetFunctionsStatus ensures that no value is present for FunctionsStatus, not even an explicit nil
func (o *CompanyInfoPlanInfo) UnsetFunctionsStatus() {
	o.FunctionsStatus.Unset()
}

func (o CompanyInfoPlanInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyInfoPlanInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Limits.IsSet() {
		toSerialize["limits"] = o.Limits.Get()
	}
	if o.Functions.IsSet() {
		toSerialize["functions"] = o.Functions.Get()
	}
	if o.FunctionsStatus.IsSet() {
		toSerialize["functions_status"] = o.FunctionsStatus.Get()
	}
	return toSerialize, nil
}

type NullableCompanyInfoPlanInfo struct {
	value *CompanyInfoPlanInfo
	isSet bool
}

func (v NullableCompanyInfoPlanInfo) Get() *CompanyInfoPlanInfo {
	return v.value
}

func (v *NullableCompanyInfoPlanInfo) Set(val *CompanyInfoPlanInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyInfoPlanInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyInfoPlanInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyInfoPlanInfo(val *CompanyInfoPlanInfo) *NullableCompanyInfoPlanInfo {
	return &NullableCompanyInfoPlanInfo{value: val, isSet: true}
}

func (v NullableCompanyInfoPlanInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyInfoPlanInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


