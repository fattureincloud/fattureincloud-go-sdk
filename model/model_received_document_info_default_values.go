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

// checks if the ReceivedDocumentInfoDefaultValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceivedDocumentInfoDefaultValues{}

// ReceivedDocumentInfoDefaultValues Default values for the document.
type ReceivedDocumentInfoDefaultValues struct {
	Detailed NullableBool `json:"detailed,omitempty"`
}

// NewReceivedDocumentInfoDefaultValues instantiates a new ReceivedDocumentInfoDefaultValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceivedDocumentInfoDefaultValues() *ReceivedDocumentInfoDefaultValues {
	this := ReceivedDocumentInfoDefaultValues{}
	return &this
}

// NewReceivedDocumentInfoDefaultValuesWithDefaults instantiates a new ReceivedDocumentInfoDefaultValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceivedDocumentInfoDefaultValuesWithDefaults() *ReceivedDocumentInfoDefaultValues {
	this := ReceivedDocumentInfoDefaultValues{}
	return &this
}

// GetDetailed returns the Detailed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentInfoDefaultValues) GetDetailed() bool {
	if o == nil || isNil(o.Detailed.Get()) {
		var ret bool
		return ret
	}
	return *o.Detailed.Get()
}

// GetDetailedOk returns a tuple with the Detailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentInfoDefaultValues) GetDetailedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detailed.Get(), o.Detailed.IsSet()
}

// HasDetailed returns a boolean if a field has been set.
func (o *ReceivedDocumentInfoDefaultValues) HasDetailed() bool {
	if o != nil && o.Detailed.IsSet() {
		return true
	}

	return false
}

// SetDetailed gets a reference to the given NullableBool and assigns it to the Detailed field.
func (o *ReceivedDocumentInfoDefaultValues) SetDetailed(v bool) *ReceivedDocumentInfoDefaultValues {
	o.Detailed.Set(&v)
	return o
}
// SetDetailedNil sets the value for Detailed to be an explicit nil
func (o *ReceivedDocumentInfoDefaultValues) SetDetailedNil() *ReceivedDocumentInfoDefaultValues {
	o.Detailed.Set(nil)
	return o
}

// UnsetDetailed ensures that no value is present for Detailed, not even an explicit nil
func (o *ReceivedDocumentInfoDefaultValues) UnsetDetailed() {
	o.Detailed.Unset()
}

func (o ReceivedDocumentInfoDefaultValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceivedDocumentInfoDefaultValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Detailed.IsSet() {
		toSerialize["detailed"] = o.Detailed.Get()
	}
	return toSerialize, nil
}

type NullableReceivedDocumentInfoDefaultValues struct {
	value *ReceivedDocumentInfoDefaultValues
	isSet bool
}

func (v NullableReceivedDocumentInfoDefaultValues) Get() *ReceivedDocumentInfoDefaultValues {
	return v.value
}

func (v *NullableReceivedDocumentInfoDefaultValues) Set(val *ReceivedDocumentInfoDefaultValues) {
	v.value = val
	v.isSet = true
}

func (v NullableReceivedDocumentInfoDefaultValues) IsSet() bool {
	return v.isSet
}

func (v *NullableReceivedDocumentInfoDefaultValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceivedDocumentInfoDefaultValues(val *ReceivedDocumentInfoDefaultValues) *NullableReceivedDocumentInfoDefaultValues {
	return &NullableReceivedDocumentInfoDefaultValues{value: val, isSet: true}
}

func (v NullableReceivedDocumentInfoDefaultValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceivedDocumentInfoDefaultValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


