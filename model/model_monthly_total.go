/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.26
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the MonthlyTotal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonthlyTotal{}

// MonthlyTotal 
type MonthlyTotal struct {
	// Monthly total net amount.
	Net NullableFloat32 `json:"net,omitempty"`
	// Monthly total gross amount.
	Gross NullableFloat32 `json:"gross,omitempty"`
	// Monthly total receipt number.
	Count NullableFloat32 `json:"count,omitempty"`
}

// NewMonthlyTotal instantiates a new MonthlyTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonthlyTotal() *MonthlyTotal {
	this := MonthlyTotal{}
	return &this
}

// NewMonthlyTotalWithDefaults instantiates a new MonthlyTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonthlyTotalWithDefaults() *MonthlyTotal {
	this := MonthlyTotal{}
	return &this
}

// GetNet returns the Net field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyTotal) GetNet() float32 {
	if o == nil || isNil(o.Net.Get()) {
		var ret float32
		return ret
	}
	return *o.Net.Get()
}

// GetNetOk returns a tuple with the Net field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonthlyTotal) GetNetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Net.Get(), o.Net.IsSet()
}

// HasNet returns a boolean if a field has been set.
func (o *MonthlyTotal) HasNet() bool {
	if o != nil && o.Net.IsSet() {
		return true
	}

	return false
}

// SetNet gets a reference to the given NullableFloat32 and assigns it to the Net field.
func (o *MonthlyTotal) SetNet(v float32) *MonthlyTotal {
	o.Net.Set(&v)
	return o
}
// SetNetNil sets the value for Net to be an explicit nil
func (o *MonthlyTotal) SetNetNil() *MonthlyTotal {
	o.Net.Set(nil)
	return o
}

// UnsetNet ensures that no value is present for Net, not even an explicit nil
func (o *MonthlyTotal) UnsetNet() {
	o.Net.Unset()
}

// GetGross returns the Gross field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyTotal) GetGross() float32 {
	if o == nil || isNil(o.Gross.Get()) {
		var ret float32
		return ret
	}
	return *o.Gross.Get()
}

// GetGrossOk returns a tuple with the Gross field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonthlyTotal) GetGrossOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gross.Get(), o.Gross.IsSet()
}

// HasGross returns a boolean if a field has been set.
func (o *MonthlyTotal) HasGross() bool {
	if o != nil && o.Gross.IsSet() {
		return true
	}

	return false
}

// SetGross gets a reference to the given NullableFloat32 and assigns it to the Gross field.
func (o *MonthlyTotal) SetGross(v float32) *MonthlyTotal {
	o.Gross.Set(&v)
	return o
}
// SetGrossNil sets the value for Gross to be an explicit nil
func (o *MonthlyTotal) SetGrossNil() *MonthlyTotal {
	o.Gross.Set(nil)
	return o
}

// UnsetGross ensures that no value is present for Gross, not even an explicit nil
func (o *MonthlyTotal) UnsetGross() {
	o.Gross.Unset()
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyTotal) GetCount() float32 {
	if o == nil || isNil(o.Count.Get()) {
		var ret float32
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonthlyTotal) GetCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *MonthlyTotal) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableFloat32 and assigns it to the Count field.
func (o *MonthlyTotal) SetCount(v float32) *MonthlyTotal {
	o.Count.Set(&v)
	return o
}
// SetCountNil sets the value for Count to be an explicit nil
func (o *MonthlyTotal) SetCountNil() *MonthlyTotal {
	o.Count.Set(nil)
	return o
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *MonthlyTotal) UnsetCount() {
	o.Count.Unset()
}

func (o MonthlyTotal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonthlyTotal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Net.IsSet() {
		toSerialize["net"] = o.Net.Get()
	}
	if o.Gross.IsSet() {
		toSerialize["gross"] = o.Gross.Get()
	}
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	return toSerialize, nil
}

type NullableMonthlyTotal struct {
	value *MonthlyTotal
	isSet bool
}

func (v NullableMonthlyTotal) Get() *MonthlyTotal {
	return v.value
}

func (v *NullableMonthlyTotal) Set(val *MonthlyTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableMonthlyTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableMonthlyTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonthlyTotal(val *MonthlyTotal) *NullableMonthlyTotal {
	return &NullableMonthlyTotal{value: val, isSet: true}
}

func (v NullableMonthlyTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonthlyTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


