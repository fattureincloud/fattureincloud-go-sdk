/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.28
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the Receipt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Receipt{}

// Receipt 
type Receipt struct {
	// Receipt unique identifier.
	Id NullableInt32 `json:"id,omitempty"`
	// Receipt date.
	Date NullableString `json:"date,omitempty"`
	// Receipt number.
	Number NullableFloat32 `json:"number,omitempty"`
	// If it's null or empty string use the default numeration.
	Numeration NullableString `json:"numeration,omitempty"`
	// Total net amount.
	AmountNet NullableFloat32 `json:"amount_net,omitempty"`
	// Total vat amount.
	AmountVat NullableFloat32 `json:"amount_vat,omitempty"`
	// Total gross amount.
	AmountGross NullableFloat32 `json:"amount_gross,omitempty"`
	UseGrossPrices NullableBool `json:"use_gross_prices,omitempty"`
	Type *ReceiptType `json:"type,omitempty"`
	// Receipt description.
	Description NullableString `json:"description,omitempty"`
	// Revenue center.
	RcCenter NullableString `json:"rc_center,omitempty"`
	CreatedAt NullableString `json:"created_at,omitempty"`
	UpdatedAt NullableString `json:"updated_at,omitempty"`
	PaymentAccount NullablePaymentAccount `json:"payment_account,omitempty"`
	ItemsList []ReceiptItemsListItem `json:"items_list,omitempty"`
}

// NewReceipt instantiates a new Receipt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceipt() *Receipt {
	this := Receipt{}
	return &this
}

// NewReceiptWithDefaults instantiates a new Receipt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiptWithDefaults() *Receipt {
	this := Receipt{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receipt) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receipt) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Receipt) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *Receipt) SetId(v int32) *Receipt {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Receipt) SetIdNil() *Receipt {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Receipt) UnsetId() {
	o.Id.Unset()
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receipt) GetDate() string {
	if o == nil || IsNil(o.Date.Get()) {
		var ret string
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receipt) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *Receipt) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableString and assigns it to the Date field.
func (o *Receipt) SetDate(v string) *Receipt {
	o.Date.Set(&v)
	return o
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *Receipt) SetDateNil() *Receipt {
	o.Date.Set(nil)
	return o
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *Receipt) UnsetDate() {
	o.Date.Unset()
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receipt) GetNumber() float32 {
	if o == nil || IsNil(o.Number.Get()) {
		var ret float32
		return ret
	}
	return *o.Number.Get()
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receipt) GetNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Number.Get(), o.Number.IsSet()
}

// HasNumber returns a boolean if a field has been set.
func (o *Receipt) HasNumber() bool {
	if o != nil && o.Number.IsSet() {
		return true
	}

	return false
}

// SetNumber gets a reference to the given NullableFloat32 and assigns it to the Number field.
func (o *Receipt) SetNumber(v float32) *Receipt {
	o.Number.Set(&v)
	return o
}
// SetNumberNil sets the value for Number to be an explicit nil
func (o *Receipt) SetNumberNil() *Receipt {
	o.Number.Set(nil)
	return o
}

// UnsetNumber ensures that no value is present for Number, not even an explicit nil
func (o *Receipt) UnsetNumber() {
	o.Number.Unset()
}

// GetNumeration returns the Numeration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receipt) GetNumeration() string {
	if o == nil || IsNil(o.Numeration.Get()) {
		var ret string
		return ret
	}
	return *o.Numeration.Get()
}

// GetNumerationOk returns a tuple with the Numeration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receipt) GetNumerationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Numeration.Get(), o.Numeration.IsSet()
}

// HasNumeration returns a boolean if a field has been set.
func (o *Receipt) HasNumeration() bool {
	if o != nil && o.Numeration.IsSet() {
		return true
	}

	return false
}

// SetNumeration gets a reference to the given NullableString and assigns it to the Numeration field.
func (o *Receipt) SetNumeration(v string) *Receipt {
	o.Numeration.Set(&v)
	return o
}
// SetNumerationNil sets the value for Numeration to be an explicit nil
func (o *Receipt) SetNumerationNil() *Receipt {
	o.Numeration.Set(nil)
	return o
}

// UnsetNumeration ensures that no value is present for Numeration, not even an explicit nil
func (o *Receipt) UnsetNumeration() {
	o.Numeration.Unset()
}

// GetAmountNet returns the AmountNet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receipt) GetAmountNet() float32 {
	if o == nil || IsNil(o.AmountNet.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountNet.Get()
}

// GetAmountNetOk returns a tuple with the AmountNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receipt) GetAmountNetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountNet.Get(), o.AmountNet.IsSet()
}

// HasAmountNet returns a boolean if a field has been set.
func (o *Receipt) HasAmountNet() bool {
	if o != nil && o.AmountNet.IsSet() {
		return true
	}

	return false
}

// SetAmountNet gets a reference to the given NullableFloat32 and assigns it to the AmountNet field.
func (o *Receipt) SetAmountNet(v float32) *Receipt {
	o.AmountNet.Set(&v)
	return o
}
// SetAmountNetNil sets the value for AmountNet to be an explicit nil
func (o *Receipt) SetAmountNetNil() *Receipt {
	o.AmountNet.Set(nil)
	return o
}

// UnsetAmountNet ensures that no value is present for AmountNet, not even an explicit nil
func (o *Receipt) UnsetAmountNet() {
	o.AmountNet.Unset()
}

// GetAmountVat returns the AmountVat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receipt) GetAmountVat() float32 {
	if o == nil || IsNil(o.AmountVat.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountVat.Get()
}

// GetAmountVatOk returns a tuple with the AmountVat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receipt) GetAmountVatOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountVat.Get(), o.AmountVat.IsSet()
}

// HasAmountVat returns a boolean if a field has been set.
func (o *Receipt) HasAmountVat() bool {
	if o != nil && o.AmountVat.IsSet() {
		return true
	}

	return false
}

// SetAmountVat gets a reference to the given NullableFloat32 and assigns it to the AmountVat field.
func (o *Receipt) SetAmountVat(v float32) *Receipt {
	o.AmountVat.Set(&v)
	return o
}
// SetAmountVatNil sets the value for AmountVat to be an explicit nil
func (o *Receipt) SetAmountVatNil() *Receipt {
	o.AmountVat.Set(nil)
	return o
}

// UnsetAmountVat ensures that no value is present for AmountVat, not even an explicit nil
func (o *Receipt) UnsetAmountVat() {
	o.AmountVat.Unset()
}

// GetAmountGross returns the AmountGross field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receipt) GetAmountGross() float32 {
	if o == nil || IsNil(o.AmountGross.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountGross.Get()
}

// GetAmountGrossOk returns a tuple with the AmountGross field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receipt) GetAmountGrossOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountGross.Get(), o.AmountGross.IsSet()
}

// HasAmountGross returns a boolean if a field has been set.
func (o *Receipt) HasAmountGross() bool {
	if o != nil && o.AmountGross.IsSet() {
		return true
	}

	return false
}

// SetAmountGross gets a reference to the given NullableFloat32 and assigns it to the AmountGross field.
func (o *Receipt) SetAmountGross(v float32) *Receipt {
	o.AmountGross.Set(&v)
	return o
}
// SetAmountGrossNil sets the value for AmountGross to be an explicit nil
func (o *Receipt) SetAmountGrossNil() *Receipt {
	o.AmountGross.Set(nil)
	return o
}

// UnsetAmountGross ensures that no value is present for AmountGross, not even an explicit nil
func (o *Receipt) UnsetAmountGross() {
	o.AmountGross.Unset()
}

// GetUseGrossPrices returns the UseGrossPrices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receipt) GetUseGrossPrices() bool {
	if o == nil || IsNil(o.UseGrossPrices.Get()) {
		var ret bool
		return ret
	}
	return *o.UseGrossPrices.Get()
}

// GetUseGrossPricesOk returns a tuple with the UseGrossPrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receipt) GetUseGrossPricesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseGrossPrices.Get(), o.UseGrossPrices.IsSet()
}

// HasUseGrossPrices returns a boolean if a field has been set.
func (o *Receipt) HasUseGrossPrices() bool {
	if o != nil && o.UseGrossPrices.IsSet() {
		return true
	}

	return false
}

// SetUseGrossPrices gets a reference to the given NullableBool and assigns it to the UseGrossPrices field.
func (o *Receipt) SetUseGrossPrices(v bool) *Receipt {
	o.UseGrossPrices.Set(&v)
	return o
}
// SetUseGrossPricesNil sets the value for UseGrossPrices to be an explicit nil
func (o *Receipt) SetUseGrossPricesNil() *Receipt {
	o.UseGrossPrices.Set(nil)
	return o
}

// UnsetUseGrossPrices ensures that no value is present for UseGrossPrices, not even an explicit nil
func (o *Receipt) UnsetUseGrossPrices() {
	o.UseGrossPrices.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Receipt) GetType() ReceiptType {
	if o == nil || IsNil(o.Type) {
		var ret ReceiptType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetTypeOk() (*ReceiptType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Receipt) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ReceiptType and assigns it to the Type field.
func (o *Receipt) SetType(v ReceiptType) *Receipt {
	o.Type = &v
	return o
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receipt) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receipt) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Receipt) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Receipt) SetDescription(v string) *Receipt {
	o.Description.Set(&v)
	return o
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Receipt) SetDescriptionNil() *Receipt {
	o.Description.Set(nil)
	return o
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Receipt) UnsetDescription() {
	o.Description.Unset()
}

// GetRcCenter returns the RcCenter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receipt) GetRcCenter() string {
	if o == nil || IsNil(o.RcCenter.Get()) {
		var ret string
		return ret
	}
	return *o.RcCenter.Get()
}

// GetRcCenterOk returns a tuple with the RcCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receipt) GetRcCenterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RcCenter.Get(), o.RcCenter.IsSet()
}

// HasRcCenter returns a boolean if a field has been set.
func (o *Receipt) HasRcCenter() bool {
	if o != nil && o.RcCenter.IsSet() {
		return true
	}

	return false
}

// SetRcCenter gets a reference to the given NullableString and assigns it to the RcCenter field.
func (o *Receipt) SetRcCenter(v string) *Receipt {
	o.RcCenter.Set(&v)
	return o
}
// SetRcCenterNil sets the value for RcCenter to be an explicit nil
func (o *Receipt) SetRcCenterNil() *Receipt {
	o.RcCenter.Set(nil)
	return o
}

// UnsetRcCenter ensures that no value is present for RcCenter, not even an explicit nil
func (o *Receipt) UnsetRcCenter() {
	o.RcCenter.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receipt) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receipt) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Receipt) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *Receipt) SetCreatedAt(v string) *Receipt {
	o.CreatedAt.Set(&v)
	return o
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Receipt) SetCreatedAtNil() *Receipt {
	o.CreatedAt.Set(nil)
	return o
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Receipt) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receipt) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receipt) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Receipt) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableString and assigns it to the UpdatedAt field.
func (o *Receipt) SetUpdatedAt(v string) *Receipt {
	o.UpdatedAt.Set(&v)
	return o
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *Receipt) SetUpdatedAtNil() *Receipt {
	o.UpdatedAt.Set(nil)
	return o
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *Receipt) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetPaymentAccount returns the PaymentAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receipt) GetPaymentAccount() PaymentAccount {
	if o == nil || IsNil(o.PaymentAccount.Get()) {
		var ret PaymentAccount
		return ret
	}
	return *o.PaymentAccount.Get()
}

// GetPaymentAccountOk returns a tuple with the PaymentAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receipt) GetPaymentAccountOk() (*PaymentAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentAccount.Get(), o.PaymentAccount.IsSet()
}

// HasPaymentAccount returns a boolean if a field has been set.
func (o *Receipt) HasPaymentAccount() bool {
	if o != nil && o.PaymentAccount.IsSet() {
		return true
	}

	return false
}

// SetPaymentAccount gets a reference to the given NullablePaymentAccount and assigns it to the PaymentAccount field.
func (o *Receipt) SetPaymentAccount(v PaymentAccount) *Receipt {
	o.PaymentAccount.Set(&v)
	return o
}
// SetPaymentAccountNil sets the value for PaymentAccount to be an explicit nil
func (o *Receipt) SetPaymentAccountNil() *Receipt {
	o.PaymentAccount.Set(nil)
	return o
}

// UnsetPaymentAccount ensures that no value is present for PaymentAccount, not even an explicit nil
func (o *Receipt) UnsetPaymentAccount() {
	o.PaymentAccount.Unset()
}

// GetItemsList returns the ItemsList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receipt) GetItemsList() []ReceiptItemsListItem {
	if o == nil {
		var ret []ReceiptItemsListItem
		return ret
	}
	return o.ItemsList
}

// GetItemsListOk returns a tuple with the ItemsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receipt) GetItemsListOk() ([]ReceiptItemsListItem, bool) {
	if o == nil || IsNil(o.ItemsList) {
		return nil, false
	}
	return o.ItemsList, true
}

// HasItemsList returns a boolean if a field has been set.
func (o *Receipt) HasItemsList() bool {
	if o != nil && IsNil(o.ItemsList) {
		return true
	}

	return false
}

// SetItemsList gets a reference to the given []ReceiptItemsListItem and assigns it to the ItemsList field.
func (o *Receipt) SetItemsList(v []ReceiptItemsListItem) *Receipt {
	o.ItemsList = v
	return o
}

func (o Receipt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Receipt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}
	if o.Number.IsSet() {
		toSerialize["number"] = o.Number.Get()
	}
	if o.Numeration.IsSet() {
		toSerialize["numeration"] = o.Numeration.Get()
	}
	if o.AmountNet.IsSet() {
		toSerialize["amount_net"] = o.AmountNet.Get()
	}
	if o.AmountVat.IsSet() {
		toSerialize["amount_vat"] = o.AmountVat.Get()
	}
	if o.AmountGross.IsSet() {
		toSerialize["amount_gross"] = o.AmountGross.Get()
	}
	if o.UseGrossPrices.IsSet() {
		toSerialize["use_gross_prices"] = o.UseGrossPrices.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.RcCenter.IsSet() {
		toSerialize["rc_center"] = o.RcCenter.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.PaymentAccount.IsSet() {
		toSerialize["payment_account"] = o.PaymentAccount.Get()
	}
	if o.ItemsList != nil {
		toSerialize["items_list"] = o.ItemsList
	}
	return toSerialize, nil
}

type NullableReceipt struct {
	value *Receipt
	isSet bool
}

func (v NullableReceipt) Get() *Receipt {
	return v.value
}

func (v *NullableReceipt) Set(val *Receipt) {
	v.value = val
	v.isSet = true
}

func (v NullableReceipt) IsSet() bool {
	return v.isSet
}

func (v *NullableReceipt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceipt(val *Receipt) *NullableReceipt {
	return &NullableReceipt{value: val, isSet: true}
}

func (v NullableReceipt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceipt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


