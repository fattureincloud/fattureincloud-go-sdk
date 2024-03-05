/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.32
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ReceiptItemsListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceiptItemsListItem{}

// ReceiptItemsListItem struct for ReceiptItemsListItem
type ReceiptItemsListItem struct {
	// Receipt item id
	Id NullableInt32 `json:"id,omitempty"`
	// Receipt item total net amount
	AmountNet NullableFloat32 `json:"amount_net,omitempty"`
	// Receipt item total gross amount
	AmountGross NullableFloat32 `json:"amount_gross,omitempty"`
	// Receipt item category
	Category NullableString `json:"category,omitempty"`
	Vat NullableVatType `json:"vat,omitempty"`
}

// NewReceiptItemsListItem instantiates a new ReceiptItemsListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceiptItemsListItem() *ReceiptItemsListItem {
	this := ReceiptItemsListItem{}
	return &this
}

// NewReceiptItemsListItemWithDefaults instantiates a new ReceiptItemsListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiptItemsListItemWithDefaults() *ReceiptItemsListItem {
	this := ReceiptItemsListItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceiptItemsListItem) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceiptItemsListItem) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ReceiptItemsListItem) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *ReceiptItemsListItem) SetId(v int32) *ReceiptItemsListItem {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ReceiptItemsListItem) SetIdNil() *ReceiptItemsListItem {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ReceiptItemsListItem) UnsetId() {
	o.Id.Unset()
}

// GetAmountNet returns the AmountNet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceiptItemsListItem) GetAmountNet() float32 {
	if o == nil || IsNil(o.AmountNet.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountNet.Get()
}

// GetAmountNetOk returns a tuple with the AmountNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceiptItemsListItem) GetAmountNetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountNet.Get(), o.AmountNet.IsSet()
}

// HasAmountNet returns a boolean if a field has been set.
func (o *ReceiptItemsListItem) HasAmountNet() bool {
	if o != nil && o.AmountNet.IsSet() {
		return true
	}

	return false
}

// SetAmountNet gets a reference to the given NullableFloat32 and assigns it to the AmountNet field.
func (o *ReceiptItemsListItem) SetAmountNet(v float32) *ReceiptItemsListItem {
	o.AmountNet.Set(&v)
	return o
}
// SetAmountNetNil sets the value for AmountNet to be an explicit nil
func (o *ReceiptItemsListItem) SetAmountNetNil() *ReceiptItemsListItem {
	o.AmountNet.Set(nil)
	return o
}

// UnsetAmountNet ensures that no value is present for AmountNet, not even an explicit nil
func (o *ReceiptItemsListItem) UnsetAmountNet() {
	o.AmountNet.Unset()
}

// GetAmountGross returns the AmountGross field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceiptItemsListItem) GetAmountGross() float32 {
	if o == nil || IsNil(o.AmountGross.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountGross.Get()
}

// GetAmountGrossOk returns a tuple with the AmountGross field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceiptItemsListItem) GetAmountGrossOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountGross.Get(), o.AmountGross.IsSet()
}

// HasAmountGross returns a boolean if a field has been set.
func (o *ReceiptItemsListItem) HasAmountGross() bool {
	if o != nil && o.AmountGross.IsSet() {
		return true
	}

	return false
}

// SetAmountGross gets a reference to the given NullableFloat32 and assigns it to the AmountGross field.
func (o *ReceiptItemsListItem) SetAmountGross(v float32) *ReceiptItemsListItem {
	o.AmountGross.Set(&v)
	return o
}
// SetAmountGrossNil sets the value for AmountGross to be an explicit nil
func (o *ReceiptItemsListItem) SetAmountGrossNil() *ReceiptItemsListItem {
	o.AmountGross.Set(nil)
	return o
}

// UnsetAmountGross ensures that no value is present for AmountGross, not even an explicit nil
func (o *ReceiptItemsListItem) UnsetAmountGross() {
	o.AmountGross.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceiptItemsListItem) GetCategory() string {
	if o == nil || IsNil(o.Category.Get()) {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceiptItemsListItem) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *ReceiptItemsListItem) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *ReceiptItemsListItem) SetCategory(v string) *ReceiptItemsListItem {
	o.Category.Set(&v)
	return o
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *ReceiptItemsListItem) SetCategoryNil() *ReceiptItemsListItem {
	o.Category.Set(nil)
	return o
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *ReceiptItemsListItem) UnsetCategory() {
	o.Category.Unset()
}

// GetVat returns the Vat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceiptItemsListItem) GetVat() VatType {
	if o == nil || IsNil(o.Vat.Get()) {
		var ret VatType
		return ret
	}
	return *o.Vat.Get()
}

// GetVatOk returns a tuple with the Vat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceiptItemsListItem) GetVatOk() (*VatType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vat.Get(), o.Vat.IsSet()
}

// HasVat returns a boolean if a field has been set.
func (o *ReceiptItemsListItem) HasVat() bool {
	if o != nil && o.Vat.IsSet() {
		return true
	}

	return false
}

// SetVat gets a reference to the given NullableVatType and assigns it to the Vat field.
func (o *ReceiptItemsListItem) SetVat(v VatType) *ReceiptItemsListItem {
	o.Vat.Set(&v)
	return o
}
// SetVatNil sets the value for Vat to be an explicit nil
func (o *ReceiptItemsListItem) SetVatNil() *ReceiptItemsListItem {
	o.Vat.Set(nil)
	return o
}

// UnsetVat ensures that no value is present for Vat, not even an explicit nil
func (o *ReceiptItemsListItem) UnsetVat() {
	o.Vat.Unset()
}

func (o ReceiptItemsListItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceiptItemsListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.AmountNet.IsSet() {
		toSerialize["amount_net"] = o.AmountNet.Get()
	}
	if o.AmountGross.IsSet() {
		toSerialize["amount_gross"] = o.AmountGross.Get()
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.Vat.IsSet() {
		toSerialize["vat"] = o.Vat.Get()
	}
	return toSerialize, nil
}

type NullableReceiptItemsListItem struct {
	value *ReceiptItemsListItem
	isSet bool
}

func (v NullableReceiptItemsListItem) Get() *ReceiptItemsListItem {
	return v.value
}

func (v *NullableReceiptItemsListItem) Set(val *ReceiptItemsListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiptItemsListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiptItemsListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiptItemsListItem(val *ReceiptItemsListItem) *NullableReceiptItemsListItem {
	return &NullableReceiptItemsListItem{value: val, isSet: true}
}

func (v NullableReceiptItemsListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiptItemsListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


