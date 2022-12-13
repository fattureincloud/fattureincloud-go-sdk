/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.24
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// VatItem struct for VatItem
type VatItem struct {
	AmountNet NullableFloat32 `json:"amount_net,omitempty"`
	AmountVat NullableFloat32 `json:"amount_vat,omitempty"`
}

// NewVatItem instantiates a new VatItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVatItem() *VatItem {
	this := VatItem{}
	return &this
}

// NewVatItemWithDefaults instantiates a new VatItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVatItemWithDefaults() *VatItem {
	this := VatItem{}
	return &this
}

// GetAmountNet returns the AmountNet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VatItem) GetAmountNet() float32 {
	if o == nil || isNil(o.AmountNet.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountNet.Get()
}

// GetAmountNetOk returns a tuple with the AmountNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VatItem) GetAmountNetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountNet.Get(), o.AmountNet.IsSet()
}

// HasAmountNet returns a boolean if a field has been set.
func (o *VatItem) HasAmountNet() bool {
	if o != nil && o.AmountNet.IsSet() {
		return true
	}

	return false
}

// SetAmountNet gets a reference to the given NullableFloat32 and assigns it to the AmountNet field.
func (o *VatItem) SetAmountNet(v float32) *VatItem {
	o.AmountNet.Set(&v)
	return o
}
// SetAmountNetNil sets the value for AmountNet to be an explicit nil
func (o *VatItem) SetAmountNetNil() *VatItem {
	o.AmountNet.Set(nil)
	return o
}

// UnsetAmountNet ensures that no value is present for AmountNet, not even an explicit nil
func (o *VatItem) UnsetAmountNet() {
	o.AmountNet.Unset()
}

// GetAmountVat returns the AmountVat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VatItem) GetAmountVat() float32 {
	if o == nil || isNil(o.AmountVat.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountVat.Get()
}

// GetAmountVatOk returns a tuple with the AmountVat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VatItem) GetAmountVatOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountVat.Get(), o.AmountVat.IsSet()
}

// HasAmountVat returns a boolean if a field has been set.
func (o *VatItem) HasAmountVat() bool {
	if o != nil && o.AmountVat.IsSet() {
		return true
	}

	return false
}

// SetAmountVat gets a reference to the given NullableFloat32 and assigns it to the AmountVat field.
func (o *VatItem) SetAmountVat(v float32) *VatItem {
	o.AmountVat.Set(&v)
	return o
}
// SetAmountVatNil sets the value for AmountVat to be an explicit nil
func (o *VatItem) SetAmountVatNil() *VatItem {
	o.AmountVat.Set(nil)
	return o
}

// UnsetAmountVat ensures that no value is present for AmountVat, not even an explicit nil
func (o *VatItem) UnsetAmountVat() {
	o.AmountVat.Unset()
}

func (o VatItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AmountNet.IsSet() {
		toSerialize["amount_net"] = o.AmountNet.Get()
	}
	if o.AmountVat.IsSet() {
		toSerialize["amount_vat"] = o.AmountVat.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVatItem struct {
	value *VatItem
	isSet bool
}

func (v NullableVatItem) Get() *VatItem {
	return v.value
}

func (v *NullableVatItem) Set(val *VatItem) {
	v.value = val
	v.isSet = true
}

func (v NullableVatItem) IsSet() bool {
	return v.isSet
}

func (v *NullableVatItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVatItem(val *VatItem) *NullableVatItem {
	return &NullableVatItem{value: val, isSet: true}
}

func (v NullableVatItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVatItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


