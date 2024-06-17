/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.0
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the IssuedDocumentPreCreateInfoItemsDefaultValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssuedDocumentPreCreateInfoItemsDefaultValues{}

// IssuedDocumentPreCreateInfoItemsDefaultValues Issued document items default values
type IssuedDocumentPreCreateInfoItemsDefaultValues struct {
	Vat NullableVatType `json:"vat,omitempty"`
}

// NewIssuedDocumentPreCreateInfoItemsDefaultValues instantiates a new IssuedDocumentPreCreateInfoItemsDefaultValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuedDocumentPreCreateInfoItemsDefaultValues() *IssuedDocumentPreCreateInfoItemsDefaultValues {
	this := IssuedDocumentPreCreateInfoItemsDefaultValues{}
	return &this
}

// NewIssuedDocumentPreCreateInfoItemsDefaultValuesWithDefaults instantiates a new IssuedDocumentPreCreateInfoItemsDefaultValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuedDocumentPreCreateInfoItemsDefaultValuesWithDefaults() *IssuedDocumentPreCreateInfoItemsDefaultValues {
	this := IssuedDocumentPreCreateInfoItemsDefaultValues{}
	return &this
}

// GetVat returns the Vat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfoItemsDefaultValues) GetVat() VatType {
	if o == nil || IsNil(o.Vat.Get()) {
		var ret VatType
		return ret
	}
	return *o.Vat.Get()
}

// GetVatOk returns a tuple with the Vat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfoItemsDefaultValues) GetVatOk() (*VatType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vat.Get(), o.Vat.IsSet()
}

// HasVat returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfoItemsDefaultValues) HasVat() bool {
	if o != nil && o.Vat.IsSet() {
		return true
	}

	return false
}

// SetVat gets a reference to the given NullableVatType and assigns it to the Vat field.
func (o *IssuedDocumentPreCreateInfoItemsDefaultValues) SetVat(v VatType) *IssuedDocumentPreCreateInfoItemsDefaultValues {
	o.Vat.Set(&v)
	return o
}
// SetVatNil sets the value for Vat to be an explicit nil
func (o *IssuedDocumentPreCreateInfoItemsDefaultValues) SetVatNil() *IssuedDocumentPreCreateInfoItemsDefaultValues {
	o.Vat.Set(nil)
	return o
}

// UnsetVat ensures that no value is present for Vat, not even an explicit nil
func (o *IssuedDocumentPreCreateInfoItemsDefaultValues) UnsetVat() {
	o.Vat.Unset()
}

func (o IssuedDocumentPreCreateInfoItemsDefaultValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssuedDocumentPreCreateInfoItemsDefaultValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Vat.IsSet() {
		toSerialize["vat"] = o.Vat.Get()
	}
	return toSerialize, nil
}

type NullableIssuedDocumentPreCreateInfoItemsDefaultValues struct {
	value *IssuedDocumentPreCreateInfoItemsDefaultValues
	isSet bool
}

func (v NullableIssuedDocumentPreCreateInfoItemsDefaultValues) Get() *IssuedDocumentPreCreateInfoItemsDefaultValues {
	return v.value
}

func (v *NullableIssuedDocumentPreCreateInfoItemsDefaultValues) Set(val *IssuedDocumentPreCreateInfoItemsDefaultValues) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuedDocumentPreCreateInfoItemsDefaultValues) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuedDocumentPreCreateInfoItemsDefaultValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuedDocumentPreCreateInfoItemsDefaultValues(val *IssuedDocumentPreCreateInfoItemsDefaultValues) *NullableIssuedDocumentPreCreateInfoItemsDefaultValues {
	return &NullableIssuedDocumentPreCreateInfoItemsDefaultValues{value: val, isSet: true}
}

func (v NullableIssuedDocumentPreCreateInfoItemsDefaultValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuedDocumentPreCreateInfoItemsDefaultValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


