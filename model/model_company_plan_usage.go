/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.32
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the CompanyPlanUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyPlanUsage{}

// CompanyPlanUsage struct for CompanyPlanUsage
type CompanyPlanUsage struct {
	// Plan limit
	Limit NullableFloat32 `json:"limit,omitempty"`
	// Plan usage
	Usage NullableFloat32 `json:"usage,omitempty"`
}

// NewCompanyPlanUsage instantiates a new CompanyPlanUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyPlanUsage() *CompanyPlanUsage {
	this := CompanyPlanUsage{}
	return &this
}

// NewCompanyPlanUsageWithDefaults instantiates a new CompanyPlanUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyPlanUsageWithDefaults() *CompanyPlanUsage {
	this := CompanyPlanUsage{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyPlanUsage) GetLimit() float32 {
	if o == nil || IsNil(o.Limit.Get()) {
		var ret float32
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyPlanUsage) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *CompanyPlanUsage) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableFloat32 and assigns it to the Limit field.
func (o *CompanyPlanUsage) SetLimit(v float32) *CompanyPlanUsage {
	o.Limit.Set(&v)
	return o
}
// SetLimitNil sets the value for Limit to be an explicit nil
func (o *CompanyPlanUsage) SetLimitNil() *CompanyPlanUsage {
	o.Limit.Set(nil)
	return o
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *CompanyPlanUsage) UnsetLimit() {
	o.Limit.Unset()
}

// GetUsage returns the Usage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyPlanUsage) GetUsage() float32 {
	if o == nil || IsNil(o.Usage.Get()) {
		var ret float32
		return ret
	}
	return *o.Usage.Get()
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyPlanUsage) GetUsageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usage.Get(), o.Usage.IsSet()
}

// HasUsage returns a boolean if a field has been set.
func (o *CompanyPlanUsage) HasUsage() bool {
	if o != nil && o.Usage.IsSet() {
		return true
	}

	return false
}

// SetUsage gets a reference to the given NullableFloat32 and assigns it to the Usage field.
func (o *CompanyPlanUsage) SetUsage(v float32) *CompanyPlanUsage {
	o.Usage.Set(&v)
	return o
}
// SetUsageNil sets the value for Usage to be an explicit nil
func (o *CompanyPlanUsage) SetUsageNil() *CompanyPlanUsage {
	o.Usage.Set(nil)
	return o
}

// UnsetUsage ensures that no value is present for Usage, not even an explicit nil
func (o *CompanyPlanUsage) UnsetUsage() {
	o.Usage.Unset()
}

func (o CompanyPlanUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyPlanUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	if o.Usage.IsSet() {
		toSerialize["usage"] = o.Usage.Get()
	}
	return toSerialize, nil
}

type NullableCompanyPlanUsage struct {
	value *CompanyPlanUsage
	isSet bool
}

func (v NullableCompanyPlanUsage) Get() *CompanyPlanUsage {
	return v.value
}

func (v *NullableCompanyPlanUsage) Set(val *CompanyPlanUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyPlanUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyPlanUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyPlanUsage(val *CompanyPlanUsage) *NullableCompanyPlanUsage {
	return &NullableCompanyPlanUsage{value: val, isSet: true}
}

func (v NullableCompanyPlanUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyPlanUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

