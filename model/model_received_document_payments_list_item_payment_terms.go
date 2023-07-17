/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.29
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ReceivedDocumentPaymentsListItemPaymentTerms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceivedDocumentPaymentsListItemPaymentTerms{}

// ReceivedDocumentPaymentsListItemPaymentTerms struct for ReceivedDocumentPaymentsListItemPaymentTerms
type ReceivedDocumentPaymentsListItemPaymentTerms struct {
	// Received document payment number of days by which the payment must be made
	Days NullableInt32 `json:"days,omitempty"`
	Type *PaymentTermsType `json:"type,omitempty"`
}

// NewReceivedDocumentPaymentsListItemPaymentTerms instantiates a new ReceivedDocumentPaymentsListItemPaymentTerms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceivedDocumentPaymentsListItemPaymentTerms() *ReceivedDocumentPaymentsListItemPaymentTerms {
	this := ReceivedDocumentPaymentsListItemPaymentTerms{}
	return &this
}

// NewReceivedDocumentPaymentsListItemPaymentTermsWithDefaults instantiates a new ReceivedDocumentPaymentsListItemPaymentTerms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceivedDocumentPaymentsListItemPaymentTermsWithDefaults() *ReceivedDocumentPaymentsListItemPaymentTerms {
	this := ReceivedDocumentPaymentsListItemPaymentTerms{}
	var type_ PaymentTermsType = PaymentTermsTypes.STANDARD
	this.Type = &type_
	return &this
}

// GetDays returns the Days field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentPaymentsListItemPaymentTerms) GetDays() int32 {
	if o == nil || IsNil(o.Days.Get()) {
		var ret int32
		return ret
	}
	return *o.Days.Get()
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentPaymentsListItemPaymentTerms) GetDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Days.Get(), o.Days.IsSet()
}

// HasDays returns a boolean if a field has been set.
func (o *ReceivedDocumentPaymentsListItemPaymentTerms) HasDays() bool {
	if o != nil && o.Days.IsSet() {
		return true
	}

	return false
}

// SetDays gets a reference to the given NullableInt32 and assigns it to the Days field.
func (o *ReceivedDocumentPaymentsListItemPaymentTerms) SetDays(v int32) *ReceivedDocumentPaymentsListItemPaymentTerms {
	o.Days.Set(&v)
	return o
}
// SetDaysNil sets the value for Days to be an explicit nil
func (o *ReceivedDocumentPaymentsListItemPaymentTerms) SetDaysNil() *ReceivedDocumentPaymentsListItemPaymentTerms {
	o.Days.Set(nil)
	return o
}

// UnsetDays ensures that no value is present for Days, not even an explicit nil
func (o *ReceivedDocumentPaymentsListItemPaymentTerms) UnsetDays() {
	o.Days.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReceivedDocumentPaymentsListItemPaymentTerms) GetType() PaymentTermsType {
	if o == nil || IsNil(o.Type) {
		var ret PaymentTermsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceivedDocumentPaymentsListItemPaymentTerms) GetTypeOk() (*PaymentTermsType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReceivedDocumentPaymentsListItemPaymentTerms) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PaymentTermsType and assigns it to the Type field.
func (o *ReceivedDocumentPaymentsListItemPaymentTerms) SetType(v PaymentTermsType) *ReceivedDocumentPaymentsListItemPaymentTerms {
	o.Type = &v
	return o
}

func (o ReceivedDocumentPaymentsListItemPaymentTerms) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceivedDocumentPaymentsListItemPaymentTerms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Days.IsSet() {
		toSerialize["days"] = o.Days.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableReceivedDocumentPaymentsListItemPaymentTerms struct {
	value *ReceivedDocumentPaymentsListItemPaymentTerms
	isSet bool
}

func (v NullableReceivedDocumentPaymentsListItemPaymentTerms) Get() *ReceivedDocumentPaymentsListItemPaymentTerms {
	return v.value
}

func (v *NullableReceivedDocumentPaymentsListItemPaymentTerms) Set(val *ReceivedDocumentPaymentsListItemPaymentTerms) {
	v.value = val
	v.isSet = true
}

func (v NullableReceivedDocumentPaymentsListItemPaymentTerms) IsSet() bool {
	return v.isSet
}

func (v *NullableReceivedDocumentPaymentsListItemPaymentTerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceivedDocumentPaymentsListItemPaymentTerms(val *ReceivedDocumentPaymentsListItemPaymentTerms) *NullableReceivedDocumentPaymentsListItemPaymentTerms {
	return &NullableReceivedDocumentPaymentsListItemPaymentTerms{value: val, isSet: true}
}

func (v NullableReceivedDocumentPaymentsListItemPaymentTerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceivedDocumentPaymentsListItemPaymentTerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


