/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.23
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// PaymentMethodDetails struct for PaymentMethodDetails
type PaymentMethodDetails struct {
	// Details title.
	Title NullableString `json:"title,omitempty"`
	// Details description.
	Description NullableString `json:"description,omitempty"`
}

// NewPaymentMethodDetails instantiates a new PaymentMethodDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodDetails() *PaymentMethodDetails {
	this := PaymentMethodDetails{}
	return &this
}

// NewPaymentMethodDetailsWithDefaults instantiates a new PaymentMethodDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodDetailsWithDefaults() *PaymentMethodDetails {
	this := PaymentMethodDetails{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodDetails) GetTitle() string {
	if o == nil || isNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodDetails) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *PaymentMethodDetails) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *PaymentMethodDetails) SetTitle(v string) *PaymentMethodDetails {
	o.Title.Set(&v)
	return o
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *PaymentMethodDetails) SetTitleNil() *PaymentMethodDetails {
	o.Title.Set(nil)
	return o
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *PaymentMethodDetails) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodDetails) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodDetails) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentMethodDetails) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PaymentMethodDetails) SetDescription(v string) *PaymentMethodDetails {
	o.Description.Set(&v)
	return o
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PaymentMethodDetails) SetDescriptionNil() *PaymentMethodDetails {
	o.Description.Set(nil)
	return o
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PaymentMethodDetails) UnsetDescription() {
	o.Description.Unset()
}

func (o PaymentMethodDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethodDetails struct {
	value *PaymentMethodDetails
	isSet bool
}

func (v NullablePaymentMethodDetails) Get() *PaymentMethodDetails {
	return v.value
}

func (v *NullablePaymentMethodDetails) Set(val *PaymentMethodDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodDetails(val *PaymentMethodDetails) *NullablePaymentMethodDetails {
	return &NullablePaymentMethodDetails{value: val, isSet: true}
}

func (v NullablePaymentMethodDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


