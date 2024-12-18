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

// checks if the IssuedDocumentPreCreateInfoDefaultValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssuedDocumentPreCreateInfoDefaultValues{}

// IssuedDocumentPreCreateInfoDefaultValues Issued document default values
type IssuedDocumentPreCreateInfoDefaultValues struct {
	DefaultTemplate *DocumentTemplate `json:"default_template,omitempty"`
	DnTemplate *DocumentTemplate `json:"dn_template,omitempty"`
	AiTemplate *DocumentTemplate `json:"ai_template,omitempty"`
	// Default notes.
	Notes NullableString `json:"notes,omitempty"`
	// Default rivalsa percentage.
	Rivalsa NullableFloat32 `json:"rivalsa,omitempty"`
	// Default cassa percentage.
	Cassa NullableFloat32 `json:"cassa,omitempty"`
	// Default withholding tax percentage.
	WithholdingTax NullableFloat32 `json:"withholding_tax,omitempty"`
	// Default withholding tax taxable percentage.
	WithholdingTaxTaxable NullableFloat32 `json:"withholding_tax_taxable,omitempty"`
	// Default other withholding tax percentage.
	OtherWithholdingTax NullableFloat32 `json:"other_withholding_tax,omitempty"`
	// Use gross price by default.
	UseGrossPrices NullableBool `json:"use_gross_prices,omitempty"`
	PaymentMethod *PaymentMethod `json:"payment_method,omitempty"`
}

// NewIssuedDocumentPreCreateInfoDefaultValues instantiates a new IssuedDocumentPreCreateInfoDefaultValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuedDocumentPreCreateInfoDefaultValues() *IssuedDocumentPreCreateInfoDefaultValues {
	this := IssuedDocumentPreCreateInfoDefaultValues{}
	return &this
}

// NewIssuedDocumentPreCreateInfoDefaultValuesWithDefaults instantiates a new IssuedDocumentPreCreateInfoDefaultValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuedDocumentPreCreateInfoDefaultValuesWithDefaults() *IssuedDocumentPreCreateInfoDefaultValues {
	this := IssuedDocumentPreCreateInfoDefaultValues{}
	return &this
}

// GetDefaultTemplate returns the DefaultTemplate field value if set, zero value otherwise.
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetDefaultTemplate() DocumentTemplate {
	if o == nil || IsNil(o.DefaultTemplate) {
		var ret DocumentTemplate
		return ret
	}
	return *o.DefaultTemplate
}

// GetDefaultTemplateOk returns a tuple with the DefaultTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetDefaultTemplateOk() (*DocumentTemplate, bool) {
	if o == nil || IsNil(o.DefaultTemplate) {
		return nil, false
	}
	return o.DefaultTemplate, true
}

// HasDefaultTemplate returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfoDefaultValues) HasDefaultTemplate() bool {
	if o != nil && !IsNil(o.DefaultTemplate) {
		return true
	}

	return false
}

// SetDefaultTemplate gets a reference to the given DocumentTemplate and assigns it to the DefaultTemplate field.
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetDefaultTemplate(v DocumentTemplate) *IssuedDocumentPreCreateInfoDefaultValues {
	o.DefaultTemplate = &v
	return o
}

// GetDnTemplate returns the DnTemplate field value if set, zero value otherwise.
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetDnTemplate() DocumentTemplate {
	if o == nil || IsNil(o.DnTemplate) {
		var ret DocumentTemplate
		return ret
	}
	return *o.DnTemplate
}

// GetDnTemplateOk returns a tuple with the DnTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetDnTemplateOk() (*DocumentTemplate, bool) {
	if o == nil || IsNil(o.DnTemplate) {
		return nil, false
	}
	return o.DnTemplate, true
}

// HasDnTemplate returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfoDefaultValues) HasDnTemplate() bool {
	if o != nil && !IsNil(o.DnTemplate) {
		return true
	}

	return false
}

// SetDnTemplate gets a reference to the given DocumentTemplate and assigns it to the DnTemplate field.
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetDnTemplate(v DocumentTemplate) *IssuedDocumentPreCreateInfoDefaultValues {
	o.DnTemplate = &v
	return o
}

// GetAiTemplate returns the AiTemplate field value if set, zero value otherwise.
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetAiTemplate() DocumentTemplate {
	if o == nil || IsNil(o.AiTemplate) {
		var ret DocumentTemplate
		return ret
	}
	return *o.AiTemplate
}

// GetAiTemplateOk returns a tuple with the AiTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetAiTemplateOk() (*DocumentTemplate, bool) {
	if o == nil || IsNil(o.AiTemplate) {
		return nil, false
	}
	return o.AiTemplate, true
}

// HasAiTemplate returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfoDefaultValues) HasAiTemplate() bool {
	if o != nil && !IsNil(o.AiTemplate) {
		return true
	}

	return false
}

// SetAiTemplate gets a reference to the given DocumentTemplate and assigns it to the AiTemplate field.
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetAiTemplate(v DocumentTemplate) *IssuedDocumentPreCreateInfoDefaultValues {
	o.AiTemplate = &v
	return o
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetNotes() string {
	if o == nil || IsNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfoDefaultValues) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetNotes(v string) *IssuedDocumentPreCreateInfoDefaultValues {
	o.Notes.Set(&v)
	return o
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetNotesNil() *IssuedDocumentPreCreateInfoDefaultValues {
	o.Notes.Set(nil)
	return o
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *IssuedDocumentPreCreateInfoDefaultValues) UnsetNotes() {
	o.Notes.Unset()
}

// GetRivalsa returns the Rivalsa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetRivalsa() float32 {
	if o == nil || IsNil(o.Rivalsa.Get()) {
		var ret float32
		return ret
	}
	return *o.Rivalsa.Get()
}

// GetRivalsaOk returns a tuple with the Rivalsa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetRivalsaOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rivalsa.Get(), o.Rivalsa.IsSet()
}

// HasRivalsa returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfoDefaultValues) HasRivalsa() bool {
	if o != nil && o.Rivalsa.IsSet() {
		return true
	}

	return false
}

// SetRivalsa gets a reference to the given NullableFloat32 and assigns it to the Rivalsa field.
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetRivalsa(v float32) *IssuedDocumentPreCreateInfoDefaultValues {
	o.Rivalsa.Set(&v)
	return o
}
// SetRivalsaNil sets the value for Rivalsa to be an explicit nil
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetRivalsaNil() *IssuedDocumentPreCreateInfoDefaultValues {
	o.Rivalsa.Set(nil)
	return o
}

// UnsetRivalsa ensures that no value is present for Rivalsa, not even an explicit nil
func (o *IssuedDocumentPreCreateInfoDefaultValues) UnsetRivalsa() {
	o.Rivalsa.Unset()
}

// GetCassa returns the Cassa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetCassa() float32 {
	if o == nil || IsNil(o.Cassa.Get()) {
		var ret float32
		return ret
	}
	return *o.Cassa.Get()
}

// GetCassaOk returns a tuple with the Cassa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetCassaOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cassa.Get(), o.Cassa.IsSet()
}

// HasCassa returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfoDefaultValues) HasCassa() bool {
	if o != nil && o.Cassa.IsSet() {
		return true
	}

	return false
}

// SetCassa gets a reference to the given NullableFloat32 and assigns it to the Cassa field.
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetCassa(v float32) *IssuedDocumentPreCreateInfoDefaultValues {
	o.Cassa.Set(&v)
	return o
}
// SetCassaNil sets the value for Cassa to be an explicit nil
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetCassaNil() *IssuedDocumentPreCreateInfoDefaultValues {
	o.Cassa.Set(nil)
	return o
}

// UnsetCassa ensures that no value is present for Cassa, not even an explicit nil
func (o *IssuedDocumentPreCreateInfoDefaultValues) UnsetCassa() {
	o.Cassa.Unset()
}

// GetWithholdingTax returns the WithholdingTax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetWithholdingTax() float32 {
	if o == nil || IsNil(o.WithholdingTax.Get()) {
		var ret float32
		return ret
	}
	return *o.WithholdingTax.Get()
}

// GetWithholdingTaxOk returns a tuple with the WithholdingTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetWithholdingTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithholdingTax.Get(), o.WithholdingTax.IsSet()
}

// HasWithholdingTax returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfoDefaultValues) HasWithholdingTax() bool {
	if o != nil && o.WithholdingTax.IsSet() {
		return true
	}

	return false
}

// SetWithholdingTax gets a reference to the given NullableFloat32 and assigns it to the WithholdingTax field.
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetWithholdingTax(v float32) *IssuedDocumentPreCreateInfoDefaultValues {
	o.WithholdingTax.Set(&v)
	return o
}
// SetWithholdingTaxNil sets the value for WithholdingTax to be an explicit nil
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetWithholdingTaxNil() *IssuedDocumentPreCreateInfoDefaultValues {
	o.WithholdingTax.Set(nil)
	return o
}

// UnsetWithholdingTax ensures that no value is present for WithholdingTax, not even an explicit nil
func (o *IssuedDocumentPreCreateInfoDefaultValues) UnsetWithholdingTax() {
	o.WithholdingTax.Unset()
}

// GetWithholdingTaxTaxable returns the WithholdingTaxTaxable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetWithholdingTaxTaxable() float32 {
	if o == nil || IsNil(o.WithholdingTaxTaxable.Get()) {
		var ret float32
		return ret
	}
	return *o.WithholdingTaxTaxable.Get()
}

// GetWithholdingTaxTaxableOk returns a tuple with the WithholdingTaxTaxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetWithholdingTaxTaxableOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithholdingTaxTaxable.Get(), o.WithholdingTaxTaxable.IsSet()
}

// HasWithholdingTaxTaxable returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfoDefaultValues) HasWithholdingTaxTaxable() bool {
	if o != nil && o.WithholdingTaxTaxable.IsSet() {
		return true
	}

	return false
}

// SetWithholdingTaxTaxable gets a reference to the given NullableFloat32 and assigns it to the WithholdingTaxTaxable field.
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetWithholdingTaxTaxable(v float32) *IssuedDocumentPreCreateInfoDefaultValues {
	o.WithholdingTaxTaxable.Set(&v)
	return o
}
// SetWithholdingTaxTaxableNil sets the value for WithholdingTaxTaxable to be an explicit nil
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetWithholdingTaxTaxableNil() *IssuedDocumentPreCreateInfoDefaultValues {
	o.WithholdingTaxTaxable.Set(nil)
	return o
}

// UnsetWithholdingTaxTaxable ensures that no value is present for WithholdingTaxTaxable, not even an explicit nil
func (o *IssuedDocumentPreCreateInfoDefaultValues) UnsetWithholdingTaxTaxable() {
	o.WithholdingTaxTaxable.Unset()
}

// GetOtherWithholdingTax returns the OtherWithholdingTax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetOtherWithholdingTax() float32 {
	if o == nil || IsNil(o.OtherWithholdingTax.Get()) {
		var ret float32
		return ret
	}
	return *o.OtherWithholdingTax.Get()
}

// GetOtherWithholdingTaxOk returns a tuple with the OtherWithholdingTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetOtherWithholdingTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OtherWithholdingTax.Get(), o.OtherWithholdingTax.IsSet()
}

// HasOtherWithholdingTax returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfoDefaultValues) HasOtherWithholdingTax() bool {
	if o != nil && o.OtherWithholdingTax.IsSet() {
		return true
	}

	return false
}

// SetOtherWithholdingTax gets a reference to the given NullableFloat32 and assigns it to the OtherWithholdingTax field.
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetOtherWithholdingTax(v float32) *IssuedDocumentPreCreateInfoDefaultValues {
	o.OtherWithholdingTax.Set(&v)
	return o
}
// SetOtherWithholdingTaxNil sets the value for OtherWithholdingTax to be an explicit nil
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetOtherWithholdingTaxNil() *IssuedDocumentPreCreateInfoDefaultValues {
	o.OtherWithholdingTax.Set(nil)
	return o
}

// UnsetOtherWithholdingTax ensures that no value is present for OtherWithholdingTax, not even an explicit nil
func (o *IssuedDocumentPreCreateInfoDefaultValues) UnsetOtherWithholdingTax() {
	o.OtherWithholdingTax.Unset()
}

// GetUseGrossPrices returns the UseGrossPrices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetUseGrossPrices() bool {
	if o == nil || IsNil(o.UseGrossPrices.Get()) {
		var ret bool
		return ret
	}
	return *o.UseGrossPrices.Get()
}

// GetUseGrossPricesOk returns a tuple with the UseGrossPrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetUseGrossPricesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseGrossPrices.Get(), o.UseGrossPrices.IsSet()
}

// HasUseGrossPrices returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfoDefaultValues) HasUseGrossPrices() bool {
	if o != nil && o.UseGrossPrices.IsSet() {
		return true
	}

	return false
}

// SetUseGrossPrices gets a reference to the given NullableBool and assigns it to the UseGrossPrices field.
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetUseGrossPrices(v bool) *IssuedDocumentPreCreateInfoDefaultValues {
	o.UseGrossPrices.Set(&v)
	return o
}
// SetUseGrossPricesNil sets the value for UseGrossPrices to be an explicit nil
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetUseGrossPricesNil() *IssuedDocumentPreCreateInfoDefaultValues {
	o.UseGrossPrices.Set(nil)
	return o
}

// UnsetUseGrossPrices ensures that no value is present for UseGrossPrices, not even an explicit nil
func (o *IssuedDocumentPreCreateInfoDefaultValues) UnsetUseGrossPrices() {
	o.UseGrossPrices.Unset()
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetPaymentMethod() PaymentMethod {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret PaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedDocumentPreCreateInfoDefaultValues) GetPaymentMethodOk() (*PaymentMethod, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfoDefaultValues) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given PaymentMethod and assigns it to the PaymentMethod field.
func (o *IssuedDocumentPreCreateInfoDefaultValues) SetPaymentMethod(v PaymentMethod) *IssuedDocumentPreCreateInfoDefaultValues {
	o.PaymentMethod = &v
	return o
}

func (o IssuedDocumentPreCreateInfoDefaultValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssuedDocumentPreCreateInfoDefaultValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultTemplate) {
		toSerialize["default_template"] = o.DefaultTemplate
	}
	if !IsNil(o.DnTemplate) {
		toSerialize["dn_template"] = o.DnTemplate
	}
	if !IsNil(o.AiTemplate) {
		toSerialize["ai_template"] = o.AiTemplate
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if o.Rivalsa.IsSet() {
		toSerialize["rivalsa"] = o.Rivalsa.Get()
	}
	if o.Cassa.IsSet() {
		toSerialize["cassa"] = o.Cassa.Get()
	}
	if o.WithholdingTax.IsSet() {
		toSerialize["withholding_tax"] = o.WithholdingTax.Get()
	}
	if o.WithholdingTaxTaxable.IsSet() {
		toSerialize["withholding_tax_taxable"] = o.WithholdingTaxTaxable.Get()
	}
	if o.OtherWithholdingTax.IsSet() {
		toSerialize["other_withholding_tax"] = o.OtherWithholdingTax.Get()
	}
	if o.UseGrossPrices.IsSet() {
		toSerialize["use_gross_prices"] = o.UseGrossPrices.Get()
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	return toSerialize, nil
}

type NullableIssuedDocumentPreCreateInfoDefaultValues struct {
	value *IssuedDocumentPreCreateInfoDefaultValues
	isSet bool
}

func (v NullableIssuedDocumentPreCreateInfoDefaultValues) Get() *IssuedDocumentPreCreateInfoDefaultValues {
	return v.value
}

func (v *NullableIssuedDocumentPreCreateInfoDefaultValues) Set(val *IssuedDocumentPreCreateInfoDefaultValues) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuedDocumentPreCreateInfoDefaultValues) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuedDocumentPreCreateInfoDefaultValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuedDocumentPreCreateInfoDefaultValues(val *IssuedDocumentPreCreateInfoDefaultValues) *NullableIssuedDocumentPreCreateInfoDefaultValues {
	return &NullableIssuedDocumentPreCreateInfoDefaultValues{value: val, isSet: true}
}

func (v NullableIssuedDocumentPreCreateInfoDefaultValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuedDocumentPreCreateInfoDefaultValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


