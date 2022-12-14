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

// PaymentAccount 
type PaymentAccount struct {
	// Unique identifier
	Id NullableInt32 `json:"id,omitempty"`
	// Payment account name.
	Name NullableString `json:"name,omitempty"`
	Type *PaymentAccountType `json:"type,omitempty"`
	// Payment account iban.
	Iban NullableString `json:"iban,omitempty"`
	// Payment account sia.
	Sia NullableString `json:"sia,omitempty"`
	// Payment account cuc.
	Cuc NullableString `json:"cuc,omitempty"`
	// Determine if the payment method is virtual.
	Virtual NullableBool `json:"virtual,omitempty"`
}

// NewPaymentAccount instantiates a new PaymentAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentAccount() *PaymentAccount {
	this := PaymentAccount{}
	return &this
}

// NewPaymentAccountWithDefaults instantiates a new PaymentAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentAccountWithDefaults() *PaymentAccount {
	this := PaymentAccount{}
	var type_ PaymentAccountType = PaymentAccountTypes.STANDARD
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentAccount) GetId() int32 {
	if o == nil || isNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAccount) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *PaymentAccount) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *PaymentAccount) SetId(v int32) *PaymentAccount {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *PaymentAccount) SetIdNil() *PaymentAccount {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *PaymentAccount) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentAccount) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PaymentAccount) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PaymentAccount) SetName(v string) *PaymentAccount {
	o.Name.Set(&v)
	return o
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PaymentAccount) SetNameNil() *PaymentAccount {
	o.Name.Set(nil)
	return o
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PaymentAccount) UnsetName() {
	o.Name.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentAccount) GetType() PaymentAccountType {
	if o == nil || isNil(o.Type) {
		var ret PaymentAccountType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAccount) GetTypeOk() (*PaymentAccountType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentAccount) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PaymentAccountType and assigns it to the Type field.
func (o *PaymentAccount) SetType(v PaymentAccountType) *PaymentAccount {
	o.Type = &v
	return o
}

// GetIban returns the Iban field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentAccount) GetIban() string {
	if o == nil || isNil(o.Iban.Get()) {
		var ret string
		return ret
	}
	return *o.Iban.Get()
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAccount) GetIbanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Iban.Get(), o.Iban.IsSet()
}

// HasIban returns a boolean if a field has been set.
func (o *PaymentAccount) HasIban() bool {
	if o != nil && o.Iban.IsSet() {
		return true
	}

	return false
}

// SetIban gets a reference to the given NullableString and assigns it to the Iban field.
func (o *PaymentAccount) SetIban(v string) *PaymentAccount {
	o.Iban.Set(&v)
	return o
}
// SetIbanNil sets the value for Iban to be an explicit nil
func (o *PaymentAccount) SetIbanNil() *PaymentAccount {
	o.Iban.Set(nil)
	return o
}

// UnsetIban ensures that no value is present for Iban, not even an explicit nil
func (o *PaymentAccount) UnsetIban() {
	o.Iban.Unset()
}

// GetSia returns the Sia field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentAccount) GetSia() string {
	if o == nil || isNil(o.Sia.Get()) {
		var ret string
		return ret
	}
	return *o.Sia.Get()
}

// GetSiaOk returns a tuple with the Sia field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAccount) GetSiaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sia.Get(), o.Sia.IsSet()
}

// HasSia returns a boolean if a field has been set.
func (o *PaymentAccount) HasSia() bool {
	if o != nil && o.Sia.IsSet() {
		return true
	}

	return false
}

// SetSia gets a reference to the given NullableString and assigns it to the Sia field.
func (o *PaymentAccount) SetSia(v string) *PaymentAccount {
	o.Sia.Set(&v)
	return o
}
// SetSiaNil sets the value for Sia to be an explicit nil
func (o *PaymentAccount) SetSiaNil() *PaymentAccount {
	o.Sia.Set(nil)
	return o
}

// UnsetSia ensures that no value is present for Sia, not even an explicit nil
func (o *PaymentAccount) UnsetSia() {
	o.Sia.Unset()
}

// GetCuc returns the Cuc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentAccount) GetCuc() string {
	if o == nil || isNil(o.Cuc.Get()) {
		var ret string
		return ret
	}
	return *o.Cuc.Get()
}

// GetCucOk returns a tuple with the Cuc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAccount) GetCucOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cuc.Get(), o.Cuc.IsSet()
}

// HasCuc returns a boolean if a field has been set.
func (o *PaymentAccount) HasCuc() bool {
	if o != nil && o.Cuc.IsSet() {
		return true
	}

	return false
}

// SetCuc gets a reference to the given NullableString and assigns it to the Cuc field.
func (o *PaymentAccount) SetCuc(v string) *PaymentAccount {
	o.Cuc.Set(&v)
	return o
}
// SetCucNil sets the value for Cuc to be an explicit nil
func (o *PaymentAccount) SetCucNil() *PaymentAccount {
	o.Cuc.Set(nil)
	return o
}

// UnsetCuc ensures that no value is present for Cuc, not even an explicit nil
func (o *PaymentAccount) UnsetCuc() {
	o.Cuc.Unset()
}

// GetVirtual returns the Virtual field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentAccount) GetVirtual() bool {
	if o == nil || isNil(o.Virtual.Get()) {
		var ret bool
		return ret
	}
	return *o.Virtual.Get()
}

// GetVirtualOk returns a tuple with the Virtual field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAccount) GetVirtualOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Virtual.Get(), o.Virtual.IsSet()
}

// HasVirtual returns a boolean if a field has been set.
func (o *PaymentAccount) HasVirtual() bool {
	if o != nil && o.Virtual.IsSet() {
		return true
	}

	return false
}

// SetVirtual gets a reference to the given NullableBool and assigns it to the Virtual field.
func (o *PaymentAccount) SetVirtual(v bool) *PaymentAccount {
	o.Virtual.Set(&v)
	return o
}
// SetVirtualNil sets the value for Virtual to be an explicit nil
func (o *PaymentAccount) SetVirtualNil() *PaymentAccount {
	o.Virtual.Set(nil)
	return o
}

// UnsetVirtual ensures that no value is present for Virtual, not even an explicit nil
func (o *PaymentAccount) UnsetVirtual() {
	o.Virtual.Unset()
}

func (o PaymentAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Iban.IsSet() {
		toSerialize["iban"] = o.Iban.Get()
	}
	if o.Sia.IsSet() {
		toSerialize["sia"] = o.Sia.Get()
	}
	if o.Cuc.IsSet() {
		toSerialize["cuc"] = o.Cuc.Get()
	}
	if o.Virtual.IsSet() {
		toSerialize["virtual"] = o.Virtual.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentAccount struct {
	value *PaymentAccount
	isSet bool
}

func (v NullablePaymentAccount) Get() *PaymentAccount {
	return v.value
}

func (v *NullablePaymentAccount) Set(val *PaymentAccount) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentAccount) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentAccount(val *PaymentAccount) *NullablePaymentAccount {
	return &NullablePaymentAccount{value: val, isSet: true}
}

func (v NullablePaymentAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


