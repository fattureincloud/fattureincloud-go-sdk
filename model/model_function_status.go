/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.28
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the FunctionStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionStatus{}

// FunctionStatus struct for FunctionStatus
type FunctionStatus struct {
	Active NullableBool `json:"active,omitempty"`
}

// NewFunctionStatus instantiates a new FunctionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionStatus() *FunctionStatus {
	this := FunctionStatus{}
	return &this
}

// NewFunctionStatusWithDefaults instantiates a new FunctionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionStatusWithDefaults() *FunctionStatus {
	this := FunctionStatus{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FunctionStatus) GetActive() bool {
	if o == nil || IsNil(o.Active.Get()) {
		var ret bool
		return ret
	}
	return *o.Active.Get()
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FunctionStatus) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Active.Get(), o.Active.IsSet()
}

// HasActive returns a boolean if a field has been set.
func (o *FunctionStatus) HasActive() bool {
	if o != nil && o.Active.IsSet() {
		return true
	}

	return false
}

// SetActive gets a reference to the given NullableBool and assigns it to the Active field.
func (o *FunctionStatus) SetActive(v bool) *FunctionStatus {
	o.Active.Set(&v)
	return o
}
// SetActiveNil sets the value for Active to be an explicit nil
func (o *FunctionStatus) SetActiveNil() *FunctionStatus {
	o.Active.Set(nil)
	return o
}

// UnsetActive ensures that no value is present for Active, not even an explicit nil
func (o *FunctionStatus) UnsetActive() {
	o.Active.Unset()
}

func (o FunctionStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Active.IsSet() {
		toSerialize["active"] = o.Active.Get()
	}
	return toSerialize, nil
}

type NullableFunctionStatus struct {
	value *FunctionStatus
	isSet bool
}

func (v NullableFunctionStatus) Get() *FunctionStatus {
	return v.value
}

func (v *NullableFunctionStatus) Set(val *FunctionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionStatus(val *FunctionStatus) *NullableFunctionStatus {
	return &NullableFunctionStatus{value: val, isSet: true}
}

func (v NullableFunctionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


