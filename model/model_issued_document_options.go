/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.20
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// IssuedDocumentOptions struct for IssuedDocumentOptions
type IssuedDocumentOptions struct {
	// Fixes your last payment amount to match your document total
	FixPayments NullableBool `json:"fix_payments,omitempty"`
}

// NewIssuedDocumentOptions instantiates a new IssuedDocumentOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuedDocumentOptions() *IssuedDocumentOptions {
	this := IssuedDocumentOptions{}
	return &this
}

// NewIssuedDocumentOptionsWithDefaults instantiates a new IssuedDocumentOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuedDocumentOptionsWithDefaults() *IssuedDocumentOptions {
	this := IssuedDocumentOptions{}
	return &this
}

// GetFixPayments returns the FixPayments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentOptions) GetFixPayments() bool {
	if o == nil || o.FixPayments.Get() == nil {
		var ret bool
		return ret
	}
	return *o.FixPayments.Get()
}

// GetFixPaymentsOk returns a tuple with the FixPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentOptions) GetFixPaymentsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FixPayments.Get(), o.FixPayments.IsSet()
}

// HasFixPayments returns a boolean if a field has been set.
func (o *IssuedDocumentOptions) HasFixPayments() bool {
	if o != nil && o.FixPayments.IsSet() {
		return true
	}

	return false
}

// SetFixPayments gets a reference to the given NullableBool and assigns it to the FixPayments field.
func (o *IssuedDocumentOptions) SetFixPayments(v bool) *IssuedDocumentOptions {
	o.FixPayments.Set(&v)
	return o
}
// SetFixPaymentsNil sets the value for FixPayments to be an explicit nil
func (o *IssuedDocumentOptions) SetFixPaymentsNil() *IssuedDocumentOptions {
	o.FixPayments.Set(nil)
	return o
}

// UnsetFixPayments ensures that no value is present for FixPayments, not even an explicit nil
func (o *IssuedDocumentOptions) UnsetFixPayments() {
	o.FixPayments.Unset()
}

func (o IssuedDocumentOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FixPayments.IsSet() {
		toSerialize["fix_payments"] = o.FixPayments.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIssuedDocumentOptions struct {
	value *IssuedDocumentOptions
	isSet bool
}

func (v NullableIssuedDocumentOptions) Get() *IssuedDocumentOptions {
	return v.value
}

func (v *NullableIssuedDocumentOptions) Set(val *IssuedDocumentOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuedDocumentOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuedDocumentOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuedDocumentOptions(val *IssuedDocumentOptions) *NullableIssuedDocumentOptions {
	return &NullableIssuedDocumentOptions{value: val, isSet: true}
}

func (v NullableIssuedDocumentOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuedDocumentOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


