/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// ReceivedDocumentInfoItemsDefaultValues Default values for the document items.
type ReceivedDocumentInfoItemsDefaultValues struct {
	// Default vat value.
	Vat NullableFloat32 `json:"vat,omitempty"`
}

// NewReceivedDocumentInfoItemsDefaultValues instantiates a new ReceivedDocumentInfoItemsDefaultValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceivedDocumentInfoItemsDefaultValues() *ReceivedDocumentInfoItemsDefaultValues {
	this := ReceivedDocumentInfoItemsDefaultValues{}
	return &this
}

// NewReceivedDocumentInfoItemsDefaultValuesWithDefaults instantiates a new ReceivedDocumentInfoItemsDefaultValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceivedDocumentInfoItemsDefaultValuesWithDefaults() *ReceivedDocumentInfoItemsDefaultValues {
	this := ReceivedDocumentInfoItemsDefaultValues{}
	return &this
}

// GetVat returns the Vat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentInfoItemsDefaultValues) GetVat() float32 {
	if o == nil || o.Vat.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Vat.Get()
}

// GetVatOk returns a tuple with the Vat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentInfoItemsDefaultValues) GetVatOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vat.Get(), o.Vat.IsSet()
}

// HasVat returns a boolean if a field has been set.
func (o *ReceivedDocumentInfoItemsDefaultValues) HasVat() bool {
	if o != nil && o.Vat.IsSet() {
		return true
	}

	return false
}

// SetVat gets a reference to the given NullableFloat32 and assigns it to the Vat field.
func (o *ReceivedDocumentInfoItemsDefaultValues) SetVat(v float32) *ReceivedDocumentInfoItemsDefaultValues {
	o.Vat.Set(&v)
	return o
}
// SetVatNil sets the value for Vat to be an explicit nil
func (o *ReceivedDocumentInfoItemsDefaultValues) SetVatNil() *ReceivedDocumentInfoItemsDefaultValues {
	o.Vat.Set(nil)
	return o
}

// UnsetVat ensures that no value is present for Vat, not even an explicit nil
func (o *ReceivedDocumentInfoItemsDefaultValues) UnsetVat() {
	o.Vat.Unset()
}

func (o ReceivedDocumentInfoItemsDefaultValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Vat.IsSet() {
		toSerialize["vat"] = o.Vat.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableReceivedDocumentInfoItemsDefaultValues struct {
	value *ReceivedDocumentInfoItemsDefaultValues
	isSet bool
}

func (v NullableReceivedDocumentInfoItemsDefaultValues) Get() *ReceivedDocumentInfoItemsDefaultValues {
	return v.value
}

func (v *NullableReceivedDocumentInfoItemsDefaultValues) Set(val *ReceivedDocumentInfoItemsDefaultValues) {
	v.value = val
	v.isSet = true
}

func (v NullableReceivedDocumentInfoItemsDefaultValues) IsSet() bool {
	return v.isSet
}

func (v *NullableReceivedDocumentInfoItemsDefaultValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceivedDocumentInfoItemsDefaultValues(val *ReceivedDocumentInfoItemsDefaultValues) *NullableReceivedDocumentInfoItemsDefaultValues {
	return &NullableReceivedDocumentInfoItemsDefaultValues{value: val, isSet: true}
}

func (v NullableReceivedDocumentInfoItemsDefaultValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceivedDocumentInfoItemsDefaultValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


