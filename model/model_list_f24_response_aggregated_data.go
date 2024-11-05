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

// checks if the ListF24ResponseAggregatedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListF24ResponseAggregatedData{}

// ListF24ResponseAggregatedData struct for ListF24ResponseAggregatedData
type ListF24ResponseAggregatedData struct {
	// Total amount.
	Amount NullableFloat32 `json:"amount,omitempty"`
}

// NewListF24ResponseAggregatedData instantiates a new ListF24ResponseAggregatedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListF24ResponseAggregatedData() *ListF24ResponseAggregatedData {
	this := ListF24ResponseAggregatedData{}
	return &this
}

// NewListF24ResponseAggregatedDataWithDefaults instantiates a new ListF24ResponseAggregatedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListF24ResponseAggregatedDataWithDefaults() *ListF24ResponseAggregatedData {
	this := ListF24ResponseAggregatedData{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListF24ResponseAggregatedData) GetAmount() float32 {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret float32
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListF24ResponseAggregatedData) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *ListF24ResponseAggregatedData) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat32 and assigns it to the Amount field.
func (o *ListF24ResponseAggregatedData) SetAmount(v float32) *ListF24ResponseAggregatedData {
	o.Amount.Set(&v)
	return o
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *ListF24ResponseAggregatedData) SetAmountNil() *ListF24ResponseAggregatedData {
	o.Amount.Set(nil)
	return o
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *ListF24ResponseAggregatedData) UnsetAmount() {
	o.Amount.Unset()
}

func (o ListF24ResponseAggregatedData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListF24ResponseAggregatedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	return toSerialize, nil
}

type NullableListF24ResponseAggregatedData struct {
	value *ListF24ResponseAggregatedData
	isSet bool
}

func (v NullableListF24ResponseAggregatedData) Get() *ListF24ResponseAggregatedData {
	return v.value
}

func (v *NullableListF24ResponseAggregatedData) Set(val *ListF24ResponseAggregatedData) {
	v.value = val
	v.isSet = true
}

func (v NullableListF24ResponseAggregatedData) IsSet() bool {
	return v.isSet
}

func (v *NullableListF24ResponseAggregatedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListF24ResponseAggregatedData(val *ListF24ResponseAggregatedData) *NullableListF24ResponseAggregatedData {
	return &NullableListF24ResponseAggregatedData{value: val, isSet: true}
}

func (v NullableListF24ResponseAggregatedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListF24ResponseAggregatedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


