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

// checks if the ListF24ResponseAggregation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListF24ResponseAggregation{}

// ListF24ResponseAggregation struct for ListF24ResponseAggregation
type ListF24ResponseAggregation struct {
	AggregatedData *ListF24ResponseAggregatedData `json:"aggregated_data,omitempty"`
}

// NewListF24ResponseAggregation instantiates a new ListF24ResponseAggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListF24ResponseAggregation() *ListF24ResponseAggregation {
	this := ListF24ResponseAggregation{}
	return &this
}

// NewListF24ResponseAggregationWithDefaults instantiates a new ListF24ResponseAggregation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListF24ResponseAggregationWithDefaults() *ListF24ResponseAggregation {
	this := ListF24ResponseAggregation{}
	return &this
}

// GetAggregatedData returns the AggregatedData field value if set, zero value otherwise.
func (o *ListF24ResponseAggregation) GetAggregatedData() ListF24ResponseAggregatedData {
	if o == nil || IsNil(o.AggregatedData) {
		var ret ListF24ResponseAggregatedData
		return ret
	}
	return *o.AggregatedData
}

// GetAggregatedDataOk returns a tuple with the AggregatedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListF24ResponseAggregation) GetAggregatedDataOk() (*ListF24ResponseAggregatedData, bool) {
	if o == nil || IsNil(o.AggregatedData) {
		return nil, false
	}
	return o.AggregatedData, true
}

// HasAggregatedData returns a boolean if a field has been set.
func (o *ListF24ResponseAggregation) HasAggregatedData() bool {
	if o != nil && !IsNil(o.AggregatedData) {
		return true
	}

	return false
}

// SetAggregatedData gets a reference to the given ListF24ResponseAggregatedData and assigns it to the AggregatedData field.
func (o *ListF24ResponseAggregation) SetAggregatedData(v ListF24ResponseAggregatedData) *ListF24ResponseAggregation {
	o.AggregatedData = &v
	return o
}

func (o ListF24ResponseAggregation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListF24ResponseAggregation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AggregatedData) {
		toSerialize["aggregated_data"] = o.AggregatedData
	}
	return toSerialize, nil
}

type NullableListF24ResponseAggregation struct {
	value *ListF24ResponseAggregation
	isSet bool
}

func (v NullableListF24ResponseAggregation) Get() *ListF24ResponseAggregation {
	return v.value
}

func (v *NullableListF24ResponseAggregation) Set(val *ListF24ResponseAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableListF24ResponseAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableListF24ResponseAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListF24ResponseAggregation(val *ListF24ResponseAggregation) *NullableListF24ResponseAggregation {
	return &NullableListF24ResponseAggregation{value: val, isSet: true}
}

func (v NullableListF24ResponseAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListF24ResponseAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


