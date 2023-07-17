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

// checks if the IssuedDocumentTotals type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssuedDocumentTotals{}

// IssuedDocumentTotals struct for IssuedDocumentTotals
type IssuedDocumentTotals struct {
	// Issued document total net amount
	AmountNet NullableFloat32 `json:"amount_net,omitempty"`
	// Issued document rivalsa amount
	AmountRivalsa NullableFloat32 `json:"amount_rivalsa,omitempty"`
	// Issued document net amount with rivalsa
	AmountNetWithRivalsa NullableFloat32 `json:"amount_net_with_rivalsa,omitempty"`
	// Issued document cassa amount
	AmountCassa NullableFloat32 `json:"amount_cassa,omitempty"`
	// Issued document taxable amount
	TaxableAmount NullableFloat32 `json:"taxable_amount,omitempty"`
	// Issued document not taxable amount
	NotTaxableAmount NullableFloat32 `json:"not_taxable_amount,omitempty"`
	// Issued document total vat amount
	AmountVat NullableFloat32 `json:"amount_vat,omitempty"`
	// Issued document total gross amount
	AmountGross NullableFloat32 `json:"amount_gross,omitempty"`
	// Issued document Taxable withholding tax amount
	TaxableAmountWithholdingTax NullableFloat32 `json:"taxable_amount_withholding_tax,omitempty"`
	// Issued document withholding tax amount
	AmountWithholdingTax NullableFloat32 `json:"amount_withholding_tax,omitempty"`
	// Issued document other withholding tax taxable amount
	TaxableAmountOtherWithholdingTax NullableFloat32 `json:"taxable_amount_other_withholding_tax,omitempty"`
	// Issued document other withholding tax amount
	AmountOtherWithholdingTax NullableFloat32 `json:"amount_other_withholding_tax,omitempty"`
	// Issued document stamp duty value [0 if not present].
	StampDuty NullableFloat32 `json:"stamp_duty,omitempty"`
	// Issued document total amount due
	AmountDue NullableFloat32 `json:"amount_due,omitempty"`
	// Is enasarco maximal excedeed
	IsEnasarcoMaximalExceeded NullableBool `json:"is_enasarco_maximal_exceeded,omitempty"`
	// Issued document payments sum
	PaymentsSum NullableFloat32 `json:"payments_sum,omitempty"`
	VatList map[string]VatItem `json:"vat_list,omitempty"`
}

// NewIssuedDocumentTotals instantiates a new IssuedDocumentTotals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuedDocumentTotals() *IssuedDocumentTotals {
	this := IssuedDocumentTotals{}
	return &this
}

// NewIssuedDocumentTotalsWithDefaults instantiates a new IssuedDocumentTotals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuedDocumentTotalsWithDefaults() *IssuedDocumentTotals {
	this := IssuedDocumentTotals{}
	return &this
}

// GetAmountNet returns the AmountNet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentTotals) GetAmountNet() float32 {
	if o == nil || IsNil(o.AmountNet.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountNet.Get()
}

// GetAmountNetOk returns a tuple with the AmountNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentTotals) GetAmountNetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountNet.Get(), o.AmountNet.IsSet()
}

// HasAmountNet returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasAmountNet() bool {
	if o != nil && o.AmountNet.IsSet() {
		return true
	}

	return false
}

// SetAmountNet gets a reference to the given NullableFloat32 and assigns it to the AmountNet field.
func (o *IssuedDocumentTotals) SetAmountNet(v float32) *IssuedDocumentTotals {
	o.AmountNet.Set(&v)
	return o
}
// SetAmountNetNil sets the value for AmountNet to be an explicit nil
func (o *IssuedDocumentTotals) SetAmountNetNil() *IssuedDocumentTotals {
	o.AmountNet.Set(nil)
	return o
}

// UnsetAmountNet ensures that no value is present for AmountNet, not even an explicit nil
func (o *IssuedDocumentTotals) UnsetAmountNet() {
	o.AmountNet.Unset()
}

// GetAmountRivalsa returns the AmountRivalsa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentTotals) GetAmountRivalsa() float32 {
	if o == nil || IsNil(o.AmountRivalsa.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountRivalsa.Get()
}

// GetAmountRivalsaOk returns a tuple with the AmountRivalsa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentTotals) GetAmountRivalsaOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountRivalsa.Get(), o.AmountRivalsa.IsSet()
}

// HasAmountRivalsa returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasAmountRivalsa() bool {
	if o != nil && o.AmountRivalsa.IsSet() {
		return true
	}

	return false
}

// SetAmountRivalsa gets a reference to the given NullableFloat32 and assigns it to the AmountRivalsa field.
func (o *IssuedDocumentTotals) SetAmountRivalsa(v float32) *IssuedDocumentTotals {
	o.AmountRivalsa.Set(&v)
	return o
}
// SetAmountRivalsaNil sets the value for AmountRivalsa to be an explicit nil
func (o *IssuedDocumentTotals) SetAmountRivalsaNil() *IssuedDocumentTotals {
	o.AmountRivalsa.Set(nil)
	return o
}

// UnsetAmountRivalsa ensures that no value is present for AmountRivalsa, not even an explicit nil
func (o *IssuedDocumentTotals) UnsetAmountRivalsa() {
	o.AmountRivalsa.Unset()
}

// GetAmountNetWithRivalsa returns the AmountNetWithRivalsa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentTotals) GetAmountNetWithRivalsa() float32 {
	if o == nil || IsNil(o.AmountNetWithRivalsa.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountNetWithRivalsa.Get()
}

// GetAmountNetWithRivalsaOk returns a tuple with the AmountNetWithRivalsa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentTotals) GetAmountNetWithRivalsaOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountNetWithRivalsa.Get(), o.AmountNetWithRivalsa.IsSet()
}

// HasAmountNetWithRivalsa returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasAmountNetWithRivalsa() bool {
	if o != nil && o.AmountNetWithRivalsa.IsSet() {
		return true
	}

	return false
}

// SetAmountNetWithRivalsa gets a reference to the given NullableFloat32 and assigns it to the AmountNetWithRivalsa field.
func (o *IssuedDocumentTotals) SetAmountNetWithRivalsa(v float32) *IssuedDocumentTotals {
	o.AmountNetWithRivalsa.Set(&v)
	return o
}
// SetAmountNetWithRivalsaNil sets the value for AmountNetWithRivalsa to be an explicit nil
func (o *IssuedDocumentTotals) SetAmountNetWithRivalsaNil() *IssuedDocumentTotals {
	o.AmountNetWithRivalsa.Set(nil)
	return o
}

// UnsetAmountNetWithRivalsa ensures that no value is present for AmountNetWithRivalsa, not even an explicit nil
func (o *IssuedDocumentTotals) UnsetAmountNetWithRivalsa() {
	o.AmountNetWithRivalsa.Unset()
}

// GetAmountCassa returns the AmountCassa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentTotals) GetAmountCassa() float32 {
	if o == nil || IsNil(o.AmountCassa.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountCassa.Get()
}

// GetAmountCassaOk returns a tuple with the AmountCassa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentTotals) GetAmountCassaOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountCassa.Get(), o.AmountCassa.IsSet()
}

// HasAmountCassa returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasAmountCassa() bool {
	if o != nil && o.AmountCassa.IsSet() {
		return true
	}

	return false
}

// SetAmountCassa gets a reference to the given NullableFloat32 and assigns it to the AmountCassa field.
func (o *IssuedDocumentTotals) SetAmountCassa(v float32) *IssuedDocumentTotals {
	o.AmountCassa.Set(&v)
	return o
}
// SetAmountCassaNil sets the value for AmountCassa to be an explicit nil
func (o *IssuedDocumentTotals) SetAmountCassaNil() *IssuedDocumentTotals {
	o.AmountCassa.Set(nil)
	return o
}

// UnsetAmountCassa ensures that no value is present for AmountCassa, not even an explicit nil
func (o *IssuedDocumentTotals) UnsetAmountCassa() {
	o.AmountCassa.Unset()
}

// GetTaxableAmount returns the TaxableAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentTotals) GetTaxableAmount() float32 {
	if o == nil || IsNil(o.TaxableAmount.Get()) {
		var ret float32
		return ret
	}
	return *o.TaxableAmount.Get()
}

// GetTaxableAmountOk returns a tuple with the TaxableAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentTotals) GetTaxableAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxableAmount.Get(), o.TaxableAmount.IsSet()
}

// HasTaxableAmount returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasTaxableAmount() bool {
	if o != nil && o.TaxableAmount.IsSet() {
		return true
	}

	return false
}

// SetTaxableAmount gets a reference to the given NullableFloat32 and assigns it to the TaxableAmount field.
func (o *IssuedDocumentTotals) SetTaxableAmount(v float32) *IssuedDocumentTotals {
	o.TaxableAmount.Set(&v)
	return o
}
// SetTaxableAmountNil sets the value for TaxableAmount to be an explicit nil
func (o *IssuedDocumentTotals) SetTaxableAmountNil() *IssuedDocumentTotals {
	o.TaxableAmount.Set(nil)
	return o
}

// UnsetTaxableAmount ensures that no value is present for TaxableAmount, not even an explicit nil
func (o *IssuedDocumentTotals) UnsetTaxableAmount() {
	o.TaxableAmount.Unset()
}

// GetNotTaxableAmount returns the NotTaxableAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentTotals) GetNotTaxableAmount() float32 {
	if o == nil || IsNil(o.NotTaxableAmount.Get()) {
		var ret float32
		return ret
	}
	return *o.NotTaxableAmount.Get()
}

// GetNotTaxableAmountOk returns a tuple with the NotTaxableAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentTotals) GetNotTaxableAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotTaxableAmount.Get(), o.NotTaxableAmount.IsSet()
}

// HasNotTaxableAmount returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasNotTaxableAmount() bool {
	if o != nil && o.NotTaxableAmount.IsSet() {
		return true
	}

	return false
}

// SetNotTaxableAmount gets a reference to the given NullableFloat32 and assigns it to the NotTaxableAmount field.
func (o *IssuedDocumentTotals) SetNotTaxableAmount(v float32) *IssuedDocumentTotals {
	o.NotTaxableAmount.Set(&v)
	return o
}
// SetNotTaxableAmountNil sets the value for NotTaxableAmount to be an explicit nil
func (o *IssuedDocumentTotals) SetNotTaxableAmountNil() *IssuedDocumentTotals {
	o.NotTaxableAmount.Set(nil)
	return o
}

// UnsetNotTaxableAmount ensures that no value is present for NotTaxableAmount, not even an explicit nil
func (o *IssuedDocumentTotals) UnsetNotTaxableAmount() {
	o.NotTaxableAmount.Unset()
}

// GetAmountVat returns the AmountVat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentTotals) GetAmountVat() float32 {
	if o == nil || IsNil(o.AmountVat.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountVat.Get()
}

// GetAmountVatOk returns a tuple with the AmountVat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentTotals) GetAmountVatOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountVat.Get(), o.AmountVat.IsSet()
}

// HasAmountVat returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasAmountVat() bool {
	if o != nil && o.AmountVat.IsSet() {
		return true
	}

	return false
}

// SetAmountVat gets a reference to the given NullableFloat32 and assigns it to the AmountVat field.
func (o *IssuedDocumentTotals) SetAmountVat(v float32) *IssuedDocumentTotals {
	o.AmountVat.Set(&v)
	return o
}
// SetAmountVatNil sets the value for AmountVat to be an explicit nil
func (o *IssuedDocumentTotals) SetAmountVatNil() *IssuedDocumentTotals {
	o.AmountVat.Set(nil)
	return o
}

// UnsetAmountVat ensures that no value is present for AmountVat, not even an explicit nil
func (o *IssuedDocumentTotals) UnsetAmountVat() {
	o.AmountVat.Unset()
}

// GetAmountGross returns the AmountGross field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentTotals) GetAmountGross() float32 {
	if o == nil || IsNil(o.AmountGross.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountGross.Get()
}

// GetAmountGrossOk returns a tuple with the AmountGross field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentTotals) GetAmountGrossOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountGross.Get(), o.AmountGross.IsSet()
}

// HasAmountGross returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasAmountGross() bool {
	if o != nil && o.AmountGross.IsSet() {
		return true
	}

	return false
}

// SetAmountGross gets a reference to the given NullableFloat32 and assigns it to the AmountGross field.
func (o *IssuedDocumentTotals) SetAmountGross(v float32) *IssuedDocumentTotals {
	o.AmountGross.Set(&v)
	return o
}
// SetAmountGrossNil sets the value for AmountGross to be an explicit nil
func (o *IssuedDocumentTotals) SetAmountGrossNil() *IssuedDocumentTotals {
	o.AmountGross.Set(nil)
	return o
}

// UnsetAmountGross ensures that no value is present for AmountGross, not even an explicit nil
func (o *IssuedDocumentTotals) UnsetAmountGross() {
	o.AmountGross.Unset()
}

// GetTaxableAmountWithholdingTax returns the TaxableAmountWithholdingTax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentTotals) GetTaxableAmountWithholdingTax() float32 {
	if o == nil || IsNil(o.TaxableAmountWithholdingTax.Get()) {
		var ret float32
		return ret
	}
	return *o.TaxableAmountWithholdingTax.Get()
}

// GetTaxableAmountWithholdingTaxOk returns a tuple with the TaxableAmountWithholdingTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentTotals) GetTaxableAmountWithholdingTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxableAmountWithholdingTax.Get(), o.TaxableAmountWithholdingTax.IsSet()
}

// HasTaxableAmountWithholdingTax returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasTaxableAmountWithholdingTax() bool {
	if o != nil && o.TaxableAmountWithholdingTax.IsSet() {
		return true
	}

	return false
}

// SetTaxableAmountWithholdingTax gets a reference to the given NullableFloat32 and assigns it to the TaxableAmountWithholdingTax field.
func (o *IssuedDocumentTotals) SetTaxableAmountWithholdingTax(v float32) *IssuedDocumentTotals {
	o.TaxableAmountWithholdingTax.Set(&v)
	return o
}
// SetTaxableAmountWithholdingTaxNil sets the value for TaxableAmountWithholdingTax to be an explicit nil
func (o *IssuedDocumentTotals) SetTaxableAmountWithholdingTaxNil() *IssuedDocumentTotals {
	o.TaxableAmountWithholdingTax.Set(nil)
	return o
}

// UnsetTaxableAmountWithholdingTax ensures that no value is present for TaxableAmountWithholdingTax, not even an explicit nil
func (o *IssuedDocumentTotals) UnsetTaxableAmountWithholdingTax() {
	o.TaxableAmountWithholdingTax.Unset()
}

// GetAmountWithholdingTax returns the AmountWithholdingTax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentTotals) GetAmountWithholdingTax() float32 {
	if o == nil || IsNil(o.AmountWithholdingTax.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountWithholdingTax.Get()
}

// GetAmountWithholdingTaxOk returns a tuple with the AmountWithholdingTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentTotals) GetAmountWithholdingTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountWithholdingTax.Get(), o.AmountWithholdingTax.IsSet()
}

// HasAmountWithholdingTax returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasAmountWithholdingTax() bool {
	if o != nil && o.AmountWithholdingTax.IsSet() {
		return true
	}

	return false
}

// SetAmountWithholdingTax gets a reference to the given NullableFloat32 and assigns it to the AmountWithholdingTax field.
func (o *IssuedDocumentTotals) SetAmountWithholdingTax(v float32) *IssuedDocumentTotals {
	o.AmountWithholdingTax.Set(&v)
	return o
}
// SetAmountWithholdingTaxNil sets the value for AmountWithholdingTax to be an explicit nil
func (o *IssuedDocumentTotals) SetAmountWithholdingTaxNil() *IssuedDocumentTotals {
	o.AmountWithholdingTax.Set(nil)
	return o
}

// UnsetAmountWithholdingTax ensures that no value is present for AmountWithholdingTax, not even an explicit nil
func (o *IssuedDocumentTotals) UnsetAmountWithholdingTax() {
	o.AmountWithholdingTax.Unset()
}

// GetTaxableAmountOtherWithholdingTax returns the TaxableAmountOtherWithholdingTax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentTotals) GetTaxableAmountOtherWithholdingTax() float32 {
	if o == nil || IsNil(o.TaxableAmountOtherWithholdingTax.Get()) {
		var ret float32
		return ret
	}
	return *o.TaxableAmountOtherWithholdingTax.Get()
}

// GetTaxableAmountOtherWithholdingTaxOk returns a tuple with the TaxableAmountOtherWithholdingTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentTotals) GetTaxableAmountOtherWithholdingTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxableAmountOtherWithholdingTax.Get(), o.TaxableAmountOtherWithholdingTax.IsSet()
}

// HasTaxableAmountOtherWithholdingTax returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasTaxableAmountOtherWithholdingTax() bool {
	if o != nil && o.TaxableAmountOtherWithholdingTax.IsSet() {
		return true
	}

	return false
}

// SetTaxableAmountOtherWithholdingTax gets a reference to the given NullableFloat32 and assigns it to the TaxableAmountOtherWithholdingTax field.
func (o *IssuedDocumentTotals) SetTaxableAmountOtherWithholdingTax(v float32) *IssuedDocumentTotals {
	o.TaxableAmountOtherWithholdingTax.Set(&v)
	return o
}
// SetTaxableAmountOtherWithholdingTaxNil sets the value for TaxableAmountOtherWithholdingTax to be an explicit nil
func (o *IssuedDocumentTotals) SetTaxableAmountOtherWithholdingTaxNil() *IssuedDocumentTotals {
	o.TaxableAmountOtherWithholdingTax.Set(nil)
	return o
}

// UnsetTaxableAmountOtherWithholdingTax ensures that no value is present for TaxableAmountOtherWithholdingTax, not even an explicit nil
func (o *IssuedDocumentTotals) UnsetTaxableAmountOtherWithholdingTax() {
	o.TaxableAmountOtherWithholdingTax.Unset()
}

// GetAmountOtherWithholdingTax returns the AmountOtherWithholdingTax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentTotals) GetAmountOtherWithholdingTax() float32 {
	if o == nil || IsNil(o.AmountOtherWithholdingTax.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountOtherWithholdingTax.Get()
}

// GetAmountOtherWithholdingTaxOk returns a tuple with the AmountOtherWithholdingTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentTotals) GetAmountOtherWithholdingTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountOtherWithholdingTax.Get(), o.AmountOtherWithholdingTax.IsSet()
}

// HasAmountOtherWithholdingTax returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasAmountOtherWithholdingTax() bool {
	if o != nil && o.AmountOtherWithholdingTax.IsSet() {
		return true
	}

	return false
}

// SetAmountOtherWithholdingTax gets a reference to the given NullableFloat32 and assigns it to the AmountOtherWithholdingTax field.
func (o *IssuedDocumentTotals) SetAmountOtherWithholdingTax(v float32) *IssuedDocumentTotals {
	o.AmountOtherWithholdingTax.Set(&v)
	return o
}
// SetAmountOtherWithholdingTaxNil sets the value for AmountOtherWithholdingTax to be an explicit nil
func (o *IssuedDocumentTotals) SetAmountOtherWithholdingTaxNil() *IssuedDocumentTotals {
	o.AmountOtherWithholdingTax.Set(nil)
	return o
}

// UnsetAmountOtherWithholdingTax ensures that no value is present for AmountOtherWithholdingTax, not even an explicit nil
func (o *IssuedDocumentTotals) UnsetAmountOtherWithholdingTax() {
	o.AmountOtherWithholdingTax.Unset()
}

// GetStampDuty returns the StampDuty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentTotals) GetStampDuty() float32 {
	if o == nil || IsNil(o.StampDuty.Get()) {
		var ret float32
		return ret
	}
	return *o.StampDuty.Get()
}

// GetStampDutyOk returns a tuple with the StampDuty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentTotals) GetStampDutyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StampDuty.Get(), o.StampDuty.IsSet()
}

// HasStampDuty returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasStampDuty() bool {
	if o != nil && o.StampDuty.IsSet() {
		return true
	}

	return false
}

// SetStampDuty gets a reference to the given NullableFloat32 and assigns it to the StampDuty field.
func (o *IssuedDocumentTotals) SetStampDuty(v float32) *IssuedDocumentTotals {
	o.StampDuty.Set(&v)
	return o
}
// SetStampDutyNil sets the value for StampDuty to be an explicit nil
func (o *IssuedDocumentTotals) SetStampDutyNil() *IssuedDocumentTotals {
	o.StampDuty.Set(nil)
	return o
}

// UnsetStampDuty ensures that no value is present for StampDuty, not even an explicit nil
func (o *IssuedDocumentTotals) UnsetStampDuty() {
	o.StampDuty.Unset()
}

// GetAmountDue returns the AmountDue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentTotals) GetAmountDue() float32 {
	if o == nil || IsNil(o.AmountDue.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountDue.Get()
}

// GetAmountDueOk returns a tuple with the AmountDue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentTotals) GetAmountDueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountDue.Get(), o.AmountDue.IsSet()
}

// HasAmountDue returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasAmountDue() bool {
	if o != nil && o.AmountDue.IsSet() {
		return true
	}

	return false
}

// SetAmountDue gets a reference to the given NullableFloat32 and assigns it to the AmountDue field.
func (o *IssuedDocumentTotals) SetAmountDue(v float32) *IssuedDocumentTotals {
	o.AmountDue.Set(&v)
	return o
}
// SetAmountDueNil sets the value for AmountDue to be an explicit nil
func (o *IssuedDocumentTotals) SetAmountDueNil() *IssuedDocumentTotals {
	o.AmountDue.Set(nil)
	return o
}

// UnsetAmountDue ensures that no value is present for AmountDue, not even an explicit nil
func (o *IssuedDocumentTotals) UnsetAmountDue() {
	o.AmountDue.Unset()
}

// GetIsEnasarcoMaximalExceeded returns the IsEnasarcoMaximalExceeded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentTotals) GetIsEnasarcoMaximalExceeded() bool {
	if o == nil || IsNil(o.IsEnasarcoMaximalExceeded.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEnasarcoMaximalExceeded.Get()
}

// GetIsEnasarcoMaximalExceededOk returns a tuple with the IsEnasarcoMaximalExceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentTotals) GetIsEnasarcoMaximalExceededOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEnasarcoMaximalExceeded.Get(), o.IsEnasarcoMaximalExceeded.IsSet()
}

// HasIsEnasarcoMaximalExceeded returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasIsEnasarcoMaximalExceeded() bool {
	if o != nil && o.IsEnasarcoMaximalExceeded.IsSet() {
		return true
	}

	return false
}

// SetIsEnasarcoMaximalExceeded gets a reference to the given NullableBool and assigns it to the IsEnasarcoMaximalExceeded field.
func (o *IssuedDocumentTotals) SetIsEnasarcoMaximalExceeded(v bool) *IssuedDocumentTotals {
	o.IsEnasarcoMaximalExceeded.Set(&v)
	return o
}
// SetIsEnasarcoMaximalExceededNil sets the value for IsEnasarcoMaximalExceeded to be an explicit nil
func (o *IssuedDocumentTotals) SetIsEnasarcoMaximalExceededNil() *IssuedDocumentTotals {
	o.IsEnasarcoMaximalExceeded.Set(nil)
	return o
}

// UnsetIsEnasarcoMaximalExceeded ensures that no value is present for IsEnasarcoMaximalExceeded, not even an explicit nil
func (o *IssuedDocumentTotals) UnsetIsEnasarcoMaximalExceeded() {
	o.IsEnasarcoMaximalExceeded.Unset()
}

// GetPaymentsSum returns the PaymentsSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentTotals) GetPaymentsSum() float32 {
	if o == nil || IsNil(o.PaymentsSum.Get()) {
		var ret float32
		return ret
	}
	return *o.PaymentsSum.Get()
}

// GetPaymentsSumOk returns a tuple with the PaymentsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentTotals) GetPaymentsSumOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentsSum.Get(), o.PaymentsSum.IsSet()
}

// HasPaymentsSum returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasPaymentsSum() bool {
	if o != nil && o.PaymentsSum.IsSet() {
		return true
	}

	return false
}

// SetPaymentsSum gets a reference to the given NullableFloat32 and assigns it to the PaymentsSum field.
func (o *IssuedDocumentTotals) SetPaymentsSum(v float32) *IssuedDocumentTotals {
	o.PaymentsSum.Set(&v)
	return o
}
// SetPaymentsSumNil sets the value for PaymentsSum to be an explicit nil
func (o *IssuedDocumentTotals) SetPaymentsSumNil() *IssuedDocumentTotals {
	o.PaymentsSum.Set(nil)
	return o
}

// UnsetPaymentsSum ensures that no value is present for PaymentsSum, not even an explicit nil
func (o *IssuedDocumentTotals) UnsetPaymentsSum() {
	o.PaymentsSum.Unset()
}

// GetVatList returns the VatList field value if set, zero value otherwise.
func (o *IssuedDocumentTotals) GetVatList() map[string]VatItem {
	if o == nil || IsNil(o.VatList) {
		var ret map[string]VatItem
		return ret
	}
	return o.VatList
}

// GetVatListOk returns a tuple with the VatList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedDocumentTotals) GetVatListOk() (map[string]VatItem, bool) {
	if o == nil || IsNil(o.VatList) {
		return map[string]VatItem{}, false
	}
	return o.VatList, true
}

// HasVatList returns a boolean if a field has been set.
func (o *IssuedDocumentTotals) HasVatList() bool {
	if o != nil && !IsNil(o.VatList) {
		return true
	}

	return false
}

// SetVatList gets a reference to the given map[string]VatItem and assigns it to the VatList field.
func (o *IssuedDocumentTotals) SetVatList(v map[string]VatItem) *IssuedDocumentTotals {
	o.VatList = v
	return o
}

func (o IssuedDocumentTotals) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssuedDocumentTotals) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AmountNet.IsSet() {
		toSerialize["amount_net"] = o.AmountNet.Get()
	}
	if o.AmountRivalsa.IsSet() {
		toSerialize["amount_rivalsa"] = o.AmountRivalsa.Get()
	}
	if o.AmountNetWithRivalsa.IsSet() {
		toSerialize["amount_net_with_rivalsa"] = o.AmountNetWithRivalsa.Get()
	}
	if o.AmountCassa.IsSet() {
		toSerialize["amount_cassa"] = o.AmountCassa.Get()
	}
	if o.TaxableAmount.IsSet() {
		toSerialize["taxable_amount"] = o.TaxableAmount.Get()
	}
	if o.NotTaxableAmount.IsSet() {
		toSerialize["not_taxable_amount"] = o.NotTaxableAmount.Get()
	}
	if o.AmountVat.IsSet() {
		toSerialize["amount_vat"] = o.AmountVat.Get()
	}
	if o.AmountGross.IsSet() {
		toSerialize["amount_gross"] = o.AmountGross.Get()
	}
	if o.TaxableAmountWithholdingTax.IsSet() {
		toSerialize["taxable_amount_withholding_tax"] = o.TaxableAmountWithholdingTax.Get()
	}
	if o.AmountWithholdingTax.IsSet() {
		toSerialize["amount_withholding_tax"] = o.AmountWithholdingTax.Get()
	}
	if o.TaxableAmountOtherWithholdingTax.IsSet() {
		toSerialize["taxable_amount_other_withholding_tax"] = o.TaxableAmountOtherWithholdingTax.Get()
	}
	if o.AmountOtherWithholdingTax.IsSet() {
		toSerialize["amount_other_withholding_tax"] = o.AmountOtherWithholdingTax.Get()
	}
	if o.StampDuty.IsSet() {
		toSerialize["stamp_duty"] = o.StampDuty.Get()
	}
	if o.AmountDue.IsSet() {
		toSerialize["amount_due"] = o.AmountDue.Get()
	}
	if o.IsEnasarcoMaximalExceeded.IsSet() {
		toSerialize["is_enasarco_maximal_exceeded"] = o.IsEnasarcoMaximalExceeded.Get()
	}
	if o.PaymentsSum.IsSet() {
		toSerialize["payments_sum"] = o.PaymentsSum.Get()
	}
	if !IsNil(o.VatList) {
		toSerialize["vat_list"] = o.VatList
	}
	return toSerialize, nil
}

type NullableIssuedDocumentTotals struct {
	value *IssuedDocumentTotals
	isSet bool
}

func (v NullableIssuedDocumentTotals) Get() *IssuedDocumentTotals {
	return v.value
}

func (v *NullableIssuedDocumentTotals) Set(val *IssuedDocumentTotals) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuedDocumentTotals) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuedDocumentTotals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuedDocumentTotals(val *IssuedDocumentTotals) *NullableIssuedDocumentTotals {
	return &NullableIssuedDocumentTotals{value: val, isSet: true}
}

func (v NullableIssuedDocumentTotals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuedDocumentTotals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


