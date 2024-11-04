/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.3
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ReceivedDocumentTotals type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceivedDocumentTotals{}

// ReceivedDocumentTotals struct for ReceivedDocumentTotals
type ReceivedDocumentTotals struct {
	// Received document total net amount
	AmountNet NullableFloat32 `json:"amount_net,omitempty"`
	// Received document total vat amount
	AmountVat NullableFloat32 `json:"amount_vat,omitempty"`
	// Received document total gross amount
	AmountGross NullableFloat32 `json:"amount_gross,omitempty"`
	// Received document withholding tax amount
	AmountWithholdingTax NullableFloat32 `json:"amount_withholding_tax,omitempty"`
	// Received document other withholding tax amount
	AmountOtherWithholdingTax NullableFloat32 `json:"amount_other_withholding_tax,omitempty"`
	// Received document total amount due
	AmountDue NullableFloat32 `json:"amount_due,omitempty"`
	// Received document payments sum
	PaymentsSum NullableFloat32 `json:"payments_sum,omitempty"`
}

// NewReceivedDocumentTotals instantiates a new ReceivedDocumentTotals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceivedDocumentTotals() *ReceivedDocumentTotals {
	this := ReceivedDocumentTotals{}
	return &this
}

// NewReceivedDocumentTotalsWithDefaults instantiates a new ReceivedDocumentTotals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceivedDocumentTotalsWithDefaults() *ReceivedDocumentTotals {
	this := ReceivedDocumentTotals{}
	return &this
}

// GetAmountNet returns the AmountNet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentTotals) GetAmountNet() float32 {
	if o == nil || IsNil(o.AmountNet.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountNet.Get()
}

// GetAmountNetOk returns a tuple with the AmountNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentTotals) GetAmountNetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountNet.Get(), o.AmountNet.IsSet()
}

// HasAmountNet returns a boolean if a field has been set.
func (o *ReceivedDocumentTotals) HasAmountNet() bool {
	if o != nil && o.AmountNet.IsSet() {
		return true
	}

	return false
}

// SetAmountNet gets a reference to the given NullableFloat32 and assigns it to the AmountNet field.
func (o *ReceivedDocumentTotals) SetAmountNet(v float32) *ReceivedDocumentTotals {
	o.AmountNet.Set(&v)
		return o
}
// SetAmountNetNil sets the value for AmountNet to be an explicit nil
func (o *ReceivedDocumentTotals) SetAmountNetNil() *ReceivedDocumentTotals {
	o.AmountNet.Set(nil)
	return o
}

// UnsetAmountNet ensures that no value is present for AmountNet, not even an explicit nil
func (o *ReceivedDocumentTotals) UnsetAmountNet() {
	o.AmountNet.Unset()
}

// GetAmountVat returns the AmountVat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentTotals) GetAmountVat() float32 {
	if o == nil || IsNil(o.AmountVat.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountVat.Get()
}

// GetAmountVatOk returns a tuple with the AmountVat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentTotals) GetAmountVatOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountVat.Get(), o.AmountVat.IsSet()
}

// HasAmountVat returns a boolean if a field has been set.
func (o *ReceivedDocumentTotals) HasAmountVat() bool {
	if o != nil && o.AmountVat.IsSet() {
		return true
	}

	return false
}

// SetAmountVat gets a reference to the given NullableFloat32 and assigns it to the AmountVat field.
func (o *ReceivedDocumentTotals) SetAmountVat(v float32) *ReceivedDocumentTotals {
	o.AmountVat.Set(&v)
		return o
}
// SetAmountVatNil sets the value for AmountVat to be an explicit nil
func (o *ReceivedDocumentTotals) SetAmountVatNil() *ReceivedDocumentTotals {
	o.AmountVat.Set(nil)
	return o
}

// UnsetAmountVat ensures that no value is present for AmountVat, not even an explicit nil
func (o *ReceivedDocumentTotals) UnsetAmountVat() {
	o.AmountVat.Unset()
}

// GetAmountGross returns the AmountGross field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentTotals) GetAmountGross() float32 {
	if o == nil || IsNil(o.AmountGross.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountGross.Get()
}

// GetAmountGrossOk returns a tuple with the AmountGross field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentTotals) GetAmountGrossOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountGross.Get(), o.AmountGross.IsSet()
}

// HasAmountGross returns a boolean if a field has been set.
func (o *ReceivedDocumentTotals) HasAmountGross() bool {
	if o != nil && o.AmountGross.IsSet() {
		return true
	}

	return false
}

// SetAmountGross gets a reference to the given NullableFloat32 and assigns it to the AmountGross field.
func (o *ReceivedDocumentTotals) SetAmountGross(v float32) *ReceivedDocumentTotals {
	o.AmountGross.Set(&v)
		return o
}
// SetAmountGrossNil sets the value for AmountGross to be an explicit nil
func (o *ReceivedDocumentTotals) SetAmountGrossNil() *ReceivedDocumentTotals {
	o.AmountGross.Set(nil)
	return o
}

// UnsetAmountGross ensures that no value is present for AmountGross, not even an explicit nil
func (o *ReceivedDocumentTotals) UnsetAmountGross() {
	o.AmountGross.Unset()
}

// GetAmountWithholdingTax returns the AmountWithholdingTax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentTotals) GetAmountWithholdingTax() float32 {
	if o == nil || IsNil(o.AmountWithholdingTax.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountWithholdingTax.Get()
}

// GetAmountWithholdingTaxOk returns a tuple with the AmountWithholdingTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentTotals) GetAmountWithholdingTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountWithholdingTax.Get(), o.AmountWithholdingTax.IsSet()
}

// HasAmountWithholdingTax returns a boolean if a field has been set.
func (o *ReceivedDocumentTotals) HasAmountWithholdingTax() bool {
	if o != nil && o.AmountWithholdingTax.IsSet() {
		return true
	}

	return false
}

// SetAmountWithholdingTax gets a reference to the given NullableFloat32 and assigns it to the AmountWithholdingTax field.
func (o *ReceivedDocumentTotals) SetAmountWithholdingTax(v float32) *ReceivedDocumentTotals {
	o.AmountWithholdingTax.Set(&v)
		return o
}
// SetAmountWithholdingTaxNil sets the value for AmountWithholdingTax to be an explicit nil
func (o *ReceivedDocumentTotals) SetAmountWithholdingTaxNil() *ReceivedDocumentTotals {
	o.AmountWithholdingTax.Set(nil)
	return o
}

// UnsetAmountWithholdingTax ensures that no value is present for AmountWithholdingTax, not even an explicit nil
func (o *ReceivedDocumentTotals) UnsetAmountWithholdingTax() {
	o.AmountWithholdingTax.Unset()
}

// GetAmountOtherWithholdingTax returns the AmountOtherWithholdingTax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentTotals) GetAmountOtherWithholdingTax() float32 {
	if o == nil || IsNil(o.AmountOtherWithholdingTax.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountOtherWithholdingTax.Get()
}

// GetAmountOtherWithholdingTaxOk returns a tuple with the AmountOtherWithholdingTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentTotals) GetAmountOtherWithholdingTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountOtherWithholdingTax.Get(), o.AmountOtherWithholdingTax.IsSet()
}

// HasAmountOtherWithholdingTax returns a boolean if a field has been set.
func (o *ReceivedDocumentTotals) HasAmountOtherWithholdingTax() bool {
	if o != nil && o.AmountOtherWithholdingTax.IsSet() {
		return true
	}

	return false
}

// SetAmountOtherWithholdingTax gets a reference to the given NullableFloat32 and assigns it to the AmountOtherWithholdingTax field.
func (o *ReceivedDocumentTotals) SetAmountOtherWithholdingTax(v float32) *ReceivedDocumentTotals {
	o.AmountOtherWithholdingTax.Set(&v)
		return o
}
// SetAmountOtherWithholdingTaxNil sets the value for AmountOtherWithholdingTax to be an explicit nil
func (o *ReceivedDocumentTotals) SetAmountOtherWithholdingTaxNil() *ReceivedDocumentTotals {
	o.AmountOtherWithholdingTax.Set(nil)
	return o
}

// UnsetAmountOtherWithholdingTax ensures that no value is present for AmountOtherWithholdingTax, not even an explicit nil
func (o *ReceivedDocumentTotals) UnsetAmountOtherWithholdingTax() {
	o.AmountOtherWithholdingTax.Unset()
}

// GetAmountDue returns the AmountDue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentTotals) GetAmountDue() float32 {
	if o == nil || IsNil(o.AmountDue.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountDue.Get()
}

// GetAmountDueOk returns a tuple with the AmountDue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentTotals) GetAmountDueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountDue.Get(), o.AmountDue.IsSet()
}

// HasAmountDue returns a boolean if a field has been set.
func (o *ReceivedDocumentTotals) HasAmountDue() bool {
	if o != nil && o.AmountDue.IsSet() {
		return true
	}

	return false
}

// SetAmountDue gets a reference to the given NullableFloat32 and assigns it to the AmountDue field.
func (o *ReceivedDocumentTotals) SetAmountDue(v float32) *ReceivedDocumentTotals {
	o.AmountDue.Set(&v)
		return o
}
// SetAmountDueNil sets the value for AmountDue to be an explicit nil
func (o *ReceivedDocumentTotals) SetAmountDueNil() *ReceivedDocumentTotals {
	o.AmountDue.Set(nil)
	return o
}

// UnsetAmountDue ensures that no value is present for AmountDue, not even an explicit nil
func (o *ReceivedDocumentTotals) UnsetAmountDue() {
	o.AmountDue.Unset()
}

// GetPaymentsSum returns the PaymentsSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentTotals) GetPaymentsSum() float32 {
	if o == nil || IsNil(o.PaymentsSum.Get()) {
		var ret float32
		return ret
	}
	return *o.PaymentsSum.Get()
}

// GetPaymentsSumOk returns a tuple with the PaymentsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentTotals) GetPaymentsSumOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentsSum.Get(), o.PaymentsSum.IsSet()
}

// HasPaymentsSum returns a boolean if a field has been set.
func (o *ReceivedDocumentTotals) HasPaymentsSum() bool {
	if o != nil && o.PaymentsSum.IsSet() {
		return true
	}

	return false
}

// SetPaymentsSum gets a reference to the given NullableFloat32 and assigns it to the PaymentsSum field.
func (o *ReceivedDocumentTotals) SetPaymentsSum(v float32) *ReceivedDocumentTotals {
	o.PaymentsSum.Set(&v)
		return o
}
// SetPaymentsSumNil sets the value for PaymentsSum to be an explicit nil
func (o *ReceivedDocumentTotals) SetPaymentsSumNil() *ReceivedDocumentTotals {
	o.PaymentsSum.Set(nil)
	return o
}

// UnsetPaymentsSum ensures that no value is present for PaymentsSum, not even an explicit nil
func (o *ReceivedDocumentTotals) UnsetPaymentsSum() {
	o.PaymentsSum.Unset()
}

func (o ReceivedDocumentTotals) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceivedDocumentTotals) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AmountNet.IsSet() {
		toSerialize["amount_net"] = o.AmountNet.Get()
	}
	if o.AmountVat.IsSet() {
		toSerialize["amount_vat"] = o.AmountVat.Get()
	}
	if o.AmountGross.IsSet() {
		toSerialize["amount_gross"] = o.AmountGross.Get()
	}
	if o.AmountWithholdingTax.IsSet() {
		toSerialize["amount_withholding_tax"] = o.AmountWithholdingTax.Get()
	}
	if o.AmountOtherWithholdingTax.IsSet() {
		toSerialize["amount_other_withholding_tax"] = o.AmountOtherWithholdingTax.Get()
	}
	if o.AmountDue.IsSet() {
		toSerialize["amount_due"] = o.AmountDue.Get()
	}
	if o.PaymentsSum.IsSet() {
		toSerialize["payments_sum"] = o.PaymentsSum.Get()
	}
	return toSerialize, nil
}

type NullableReceivedDocumentTotals struct {
	value *ReceivedDocumentTotals
	isSet bool
}

func (v NullableReceivedDocumentTotals) Get() *ReceivedDocumentTotals {
	return v.value
}

func (v *NullableReceivedDocumentTotals) Set(val *ReceivedDocumentTotals) {
	v.value = val
	v.isSet = true
}

func (v NullableReceivedDocumentTotals) IsSet() bool {
	return v.isSet
}

func (v *NullableReceivedDocumentTotals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceivedDocumentTotals(val *ReceivedDocumentTotals) *NullableReceivedDocumentTotals {
	return &NullableReceivedDocumentTotals{value: val, isSet: true}
}

func (v NullableReceivedDocumentTotals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceivedDocumentTotals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


