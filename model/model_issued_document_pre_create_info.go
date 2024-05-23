/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.33
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the IssuedDocumentPreCreateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssuedDocumentPreCreateInfo{}

// IssuedDocumentPreCreateInfo struct for IssuedDocumentPreCreateInfo
type IssuedDocumentPreCreateInfo struct {
	Numerations map[string]map[string]int32 `json:"numerations,omitempty"`
	DnNumerations map[string]map[string]int32 `json:"dn_numerations,omitempty"`
	DefaultValues NullableIssuedDocumentPreCreateInfoDefaultValues `json:"default_values,omitempty"`
	ExtraDataDefaultValues NullableIssuedDocumentPreCreateInfoExtraDataDefaultValues `json:"extra_data_default_values,omitempty"`
	ItemsDefaultValues NullableIssuedDocumentPreCreateInfoItemsDefaultValues `json:"items_default_values,omitempty"`
	// Countries list
	CountriesList []string `json:"countries_list,omitempty"`
	// Currencies list
	CurrenciesList []Currency `json:"currencies_list,omitempty"`
	// Document templates list
	TemplatesList []DocumentTemplate `json:"templates_list,omitempty"`
	// Delivery note templates list
	DnTemplatesList []DocumentTemplate `json:"dn_templates_list,omitempty"`
	// Accompanying invoice templates list
	AiTemplatesList []DocumentTemplate `json:"ai_templates_list,omitempty"`
	// Payment methods list
	PaymentMethodsList []PaymentMethod `json:"payment_methods_list,omitempty"`
	// Payment accounts list
	PaymentAccountsList []PaymentAccount `json:"payment_accounts_list,omitempty"`
	// Vat types list
	VatTypesList []VatType `json:"vat_types_list,omitempty"`
	// Languages list
	LanguagesList []Language `json:"languages_list,omitempty"`
}

// NewIssuedDocumentPreCreateInfo instantiates a new IssuedDocumentPreCreateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuedDocumentPreCreateInfo() *IssuedDocumentPreCreateInfo {
	this := IssuedDocumentPreCreateInfo{}
	return &this
}

// NewIssuedDocumentPreCreateInfoWithDefaults instantiates a new IssuedDocumentPreCreateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuedDocumentPreCreateInfoWithDefaults() *IssuedDocumentPreCreateInfo {
	this := IssuedDocumentPreCreateInfo{}
	return &this
}

// GetNumerations returns the Numerations field value if set, zero value otherwise.
func (o *IssuedDocumentPreCreateInfo) GetNumerations() map[string]map[string]int32 {
	if o == nil || IsNil(o.Numerations) {
		var ret map[string]map[string]int32
		return ret
	}
	return o.Numerations
}

// GetNumerationsOk returns a tuple with the Numerations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedDocumentPreCreateInfo) GetNumerationsOk() (map[string]map[string]int32, bool) {
	if o == nil || IsNil(o.Numerations) {
		return map[string]map[string]int32{}, false
	}
	return o.Numerations, true
}

// HasNumerations returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfo) HasNumerations() bool {
	if o != nil && !IsNil(o.Numerations) {
		return true
	}

	return false
}

// SetNumerations gets a reference to the given map[string]map[string]int32 and assigns it to the Numerations field.
func (o *IssuedDocumentPreCreateInfo) SetNumerations(v map[string]map[string]int32) *IssuedDocumentPreCreateInfo {
	o.Numerations = v
	return o
}

// GetDnNumerations returns the DnNumerations field value if set, zero value otherwise.
func (o *IssuedDocumentPreCreateInfo) GetDnNumerations() map[string]map[string]int32 {
	if o == nil || IsNil(o.DnNumerations) {
		var ret map[string]map[string]int32
		return ret
	}
	return o.DnNumerations
}

// GetDnNumerationsOk returns a tuple with the DnNumerations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedDocumentPreCreateInfo) GetDnNumerationsOk() (map[string]map[string]int32, bool) {
	if o == nil || IsNil(o.DnNumerations) {
		return map[string]map[string]int32{}, false
	}
	return o.DnNumerations, true
}

// HasDnNumerations returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfo) HasDnNumerations() bool {
	if o != nil && !IsNil(o.DnNumerations) {
		return true
	}

	return false
}

// SetDnNumerations gets a reference to the given map[string]map[string]int32 and assigns it to the DnNumerations field.
func (o *IssuedDocumentPreCreateInfo) SetDnNumerations(v map[string]map[string]int32) *IssuedDocumentPreCreateInfo {
	o.DnNumerations = v
	return o
}

// GetDefaultValues returns the DefaultValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfo) GetDefaultValues() IssuedDocumentPreCreateInfoDefaultValues {
	if o == nil || IsNil(o.DefaultValues.Get()) {
		var ret IssuedDocumentPreCreateInfoDefaultValues
		return ret
	}
	return *o.DefaultValues.Get()
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfo) GetDefaultValuesOk() (*IssuedDocumentPreCreateInfoDefaultValues, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultValues.Get(), o.DefaultValues.IsSet()
}

// HasDefaultValues returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfo) HasDefaultValues() bool {
	if o != nil && o.DefaultValues.IsSet() {
		return true
	}

	return false
}

// SetDefaultValues gets a reference to the given NullableIssuedDocumentPreCreateInfoDefaultValues and assigns it to the DefaultValues field.
func (o *IssuedDocumentPreCreateInfo) SetDefaultValues(v IssuedDocumentPreCreateInfoDefaultValues) *IssuedDocumentPreCreateInfo {
	o.DefaultValues.Set(&v)
	return o
}
// SetDefaultValuesNil sets the value for DefaultValues to be an explicit nil
func (o *IssuedDocumentPreCreateInfo) SetDefaultValuesNil() *IssuedDocumentPreCreateInfo {
	o.DefaultValues.Set(nil)
	return o
}

// UnsetDefaultValues ensures that no value is present for DefaultValues, not even an explicit nil
func (o *IssuedDocumentPreCreateInfo) UnsetDefaultValues() {
	o.DefaultValues.Unset()
}

// GetExtraDataDefaultValues returns the ExtraDataDefaultValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfo) GetExtraDataDefaultValues() IssuedDocumentPreCreateInfoExtraDataDefaultValues {
	if o == nil || IsNil(o.ExtraDataDefaultValues.Get()) {
		var ret IssuedDocumentPreCreateInfoExtraDataDefaultValues
		return ret
	}
	return *o.ExtraDataDefaultValues.Get()
}

// GetExtraDataDefaultValuesOk returns a tuple with the ExtraDataDefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfo) GetExtraDataDefaultValuesOk() (*IssuedDocumentPreCreateInfoExtraDataDefaultValues, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraDataDefaultValues.Get(), o.ExtraDataDefaultValues.IsSet()
}

// HasExtraDataDefaultValues returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfo) HasExtraDataDefaultValues() bool {
	if o != nil && o.ExtraDataDefaultValues.IsSet() {
		return true
	}

	return false
}

// SetExtraDataDefaultValues gets a reference to the given NullableIssuedDocumentPreCreateInfoExtraDataDefaultValues and assigns it to the ExtraDataDefaultValues field.
func (o *IssuedDocumentPreCreateInfo) SetExtraDataDefaultValues(v IssuedDocumentPreCreateInfoExtraDataDefaultValues) *IssuedDocumentPreCreateInfo {
	o.ExtraDataDefaultValues.Set(&v)
	return o
}
// SetExtraDataDefaultValuesNil sets the value for ExtraDataDefaultValues to be an explicit nil
func (o *IssuedDocumentPreCreateInfo) SetExtraDataDefaultValuesNil() *IssuedDocumentPreCreateInfo {
	o.ExtraDataDefaultValues.Set(nil)
	return o
}

// UnsetExtraDataDefaultValues ensures that no value is present for ExtraDataDefaultValues, not even an explicit nil
func (o *IssuedDocumentPreCreateInfo) UnsetExtraDataDefaultValues() {
	o.ExtraDataDefaultValues.Unset()
}

// GetItemsDefaultValues returns the ItemsDefaultValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfo) GetItemsDefaultValues() IssuedDocumentPreCreateInfoItemsDefaultValues {
	if o == nil || IsNil(o.ItemsDefaultValues.Get()) {
		var ret IssuedDocumentPreCreateInfoItemsDefaultValues
		return ret
	}
	return *o.ItemsDefaultValues.Get()
}

// GetItemsDefaultValuesOk returns a tuple with the ItemsDefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfo) GetItemsDefaultValuesOk() (*IssuedDocumentPreCreateInfoItemsDefaultValues, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemsDefaultValues.Get(), o.ItemsDefaultValues.IsSet()
}

// HasItemsDefaultValues returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfo) HasItemsDefaultValues() bool {
	if o != nil && o.ItemsDefaultValues.IsSet() {
		return true
	}

	return false
}

// SetItemsDefaultValues gets a reference to the given NullableIssuedDocumentPreCreateInfoItemsDefaultValues and assigns it to the ItemsDefaultValues field.
func (o *IssuedDocumentPreCreateInfo) SetItemsDefaultValues(v IssuedDocumentPreCreateInfoItemsDefaultValues) *IssuedDocumentPreCreateInfo {
	o.ItemsDefaultValues.Set(&v)
	return o
}
// SetItemsDefaultValuesNil sets the value for ItemsDefaultValues to be an explicit nil
func (o *IssuedDocumentPreCreateInfo) SetItemsDefaultValuesNil() *IssuedDocumentPreCreateInfo {
	o.ItemsDefaultValues.Set(nil)
	return o
}

// UnsetItemsDefaultValues ensures that no value is present for ItemsDefaultValues, not even an explicit nil
func (o *IssuedDocumentPreCreateInfo) UnsetItemsDefaultValues() {
	o.ItemsDefaultValues.Unset()
}

// GetCountriesList returns the CountriesList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfo) GetCountriesList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CountriesList
}

// GetCountriesListOk returns a tuple with the CountriesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfo) GetCountriesListOk() ([]string, bool) {
	if o == nil || IsNil(o.CountriesList) {
		return nil, false
	}
	return o.CountriesList, true
}

// HasCountriesList returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfo) HasCountriesList() bool {
	if o != nil && !IsNil(o.CountriesList) {
		return true
	}

	return false
}

// SetCountriesList gets a reference to the given []string and assigns it to the CountriesList field.
func (o *IssuedDocumentPreCreateInfo) SetCountriesList(v []string) *IssuedDocumentPreCreateInfo {
	o.CountriesList = v
	return o
}

// GetCurrenciesList returns the CurrenciesList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfo) GetCurrenciesList() []Currency {
	if o == nil {
		var ret []Currency
		return ret
	}
	return o.CurrenciesList
}

// GetCurrenciesListOk returns a tuple with the CurrenciesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfo) GetCurrenciesListOk() ([]Currency, bool) {
	if o == nil || IsNil(o.CurrenciesList) {
		return nil, false
	}
	return o.CurrenciesList, true
}

// HasCurrenciesList returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfo) HasCurrenciesList() bool {
	if o != nil && !IsNil(o.CurrenciesList) {
		return true
	}

	return false
}

// SetCurrenciesList gets a reference to the given []Currency and assigns it to the CurrenciesList field.
func (o *IssuedDocumentPreCreateInfo) SetCurrenciesList(v []Currency) *IssuedDocumentPreCreateInfo {
	o.CurrenciesList = v
	return o
}

// GetTemplatesList returns the TemplatesList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfo) GetTemplatesList() []DocumentTemplate {
	if o == nil {
		var ret []DocumentTemplate
		return ret
	}
	return o.TemplatesList
}

// GetTemplatesListOk returns a tuple with the TemplatesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfo) GetTemplatesListOk() ([]DocumentTemplate, bool) {
	if o == nil || IsNil(o.TemplatesList) {
		return nil, false
	}
	return o.TemplatesList, true
}

// HasTemplatesList returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfo) HasTemplatesList() bool {
	if o != nil && !IsNil(o.TemplatesList) {
		return true
	}

	return false
}

// SetTemplatesList gets a reference to the given []DocumentTemplate and assigns it to the TemplatesList field.
func (o *IssuedDocumentPreCreateInfo) SetTemplatesList(v []DocumentTemplate) *IssuedDocumentPreCreateInfo {
	o.TemplatesList = v
	return o
}

// GetDnTemplatesList returns the DnTemplatesList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfo) GetDnTemplatesList() []DocumentTemplate {
	if o == nil {
		var ret []DocumentTemplate
		return ret
	}
	return o.DnTemplatesList
}

// GetDnTemplatesListOk returns a tuple with the DnTemplatesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfo) GetDnTemplatesListOk() ([]DocumentTemplate, bool) {
	if o == nil || IsNil(o.DnTemplatesList) {
		return nil, false
	}
	return o.DnTemplatesList, true
}

// HasDnTemplatesList returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfo) HasDnTemplatesList() bool {
	if o != nil && !IsNil(o.DnTemplatesList) {
		return true
	}

	return false
}

// SetDnTemplatesList gets a reference to the given []DocumentTemplate and assigns it to the DnTemplatesList field.
func (o *IssuedDocumentPreCreateInfo) SetDnTemplatesList(v []DocumentTemplate) *IssuedDocumentPreCreateInfo {
	o.DnTemplatesList = v
	return o
}

// GetAiTemplatesList returns the AiTemplatesList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfo) GetAiTemplatesList() []DocumentTemplate {
	if o == nil {
		var ret []DocumentTemplate
		return ret
	}
	return o.AiTemplatesList
}

// GetAiTemplatesListOk returns a tuple with the AiTemplatesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfo) GetAiTemplatesListOk() ([]DocumentTemplate, bool) {
	if o == nil || IsNil(o.AiTemplatesList) {
		return nil, false
	}
	return o.AiTemplatesList, true
}

// HasAiTemplatesList returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfo) HasAiTemplatesList() bool {
	if o != nil && !IsNil(o.AiTemplatesList) {
		return true
	}

	return false
}

// SetAiTemplatesList gets a reference to the given []DocumentTemplate and assigns it to the AiTemplatesList field.
func (o *IssuedDocumentPreCreateInfo) SetAiTemplatesList(v []DocumentTemplate) *IssuedDocumentPreCreateInfo {
	o.AiTemplatesList = v
	return o
}

// GetPaymentMethodsList returns the PaymentMethodsList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfo) GetPaymentMethodsList() []PaymentMethod {
	if o == nil {
		var ret []PaymentMethod
		return ret
	}
	return o.PaymentMethodsList
}

// GetPaymentMethodsListOk returns a tuple with the PaymentMethodsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfo) GetPaymentMethodsListOk() ([]PaymentMethod, bool) {
	if o == nil || IsNil(o.PaymentMethodsList) {
		return nil, false
	}
	return o.PaymentMethodsList, true
}

// HasPaymentMethodsList returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfo) HasPaymentMethodsList() bool {
	if o != nil && !IsNil(o.PaymentMethodsList) {
		return true
	}

	return false
}

// SetPaymentMethodsList gets a reference to the given []PaymentMethod and assigns it to the PaymentMethodsList field.
func (o *IssuedDocumentPreCreateInfo) SetPaymentMethodsList(v []PaymentMethod) *IssuedDocumentPreCreateInfo {
	o.PaymentMethodsList = v
	return o
}

// GetPaymentAccountsList returns the PaymentAccountsList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfo) GetPaymentAccountsList() []PaymentAccount {
	if o == nil {
		var ret []PaymentAccount
		return ret
	}
	return o.PaymentAccountsList
}

// GetPaymentAccountsListOk returns a tuple with the PaymentAccountsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfo) GetPaymentAccountsListOk() ([]PaymentAccount, bool) {
	if o == nil || IsNil(o.PaymentAccountsList) {
		return nil, false
	}
	return o.PaymentAccountsList, true
}

// HasPaymentAccountsList returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfo) HasPaymentAccountsList() bool {
	if o != nil && !IsNil(o.PaymentAccountsList) {
		return true
	}

	return false
}

// SetPaymentAccountsList gets a reference to the given []PaymentAccount and assigns it to the PaymentAccountsList field.
func (o *IssuedDocumentPreCreateInfo) SetPaymentAccountsList(v []PaymentAccount) *IssuedDocumentPreCreateInfo {
	o.PaymentAccountsList = v
	return o
}

// GetVatTypesList returns the VatTypesList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfo) GetVatTypesList() []VatType {
	if o == nil {
		var ret []VatType
		return ret
	}
	return o.VatTypesList
}

// GetVatTypesListOk returns a tuple with the VatTypesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfo) GetVatTypesListOk() ([]VatType, bool) {
	if o == nil || IsNil(o.VatTypesList) {
		return nil, false
	}
	return o.VatTypesList, true
}

// HasVatTypesList returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfo) HasVatTypesList() bool {
	if o != nil && !IsNil(o.VatTypesList) {
		return true
	}

	return false
}

// SetVatTypesList gets a reference to the given []VatType and assigns it to the VatTypesList field.
func (o *IssuedDocumentPreCreateInfo) SetVatTypesList(v []VatType) *IssuedDocumentPreCreateInfo {
	o.VatTypesList = v
	return o
}

// GetLanguagesList returns the LanguagesList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfo) GetLanguagesList() []Language {
	if o == nil {
		var ret []Language
		return ret
	}
	return o.LanguagesList
}

// GetLanguagesListOk returns a tuple with the LanguagesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfo) GetLanguagesListOk() ([]Language, bool) {
	if o == nil || IsNil(o.LanguagesList) {
		return nil, false
	}
	return o.LanguagesList, true
}

// HasLanguagesList returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfo) HasLanguagesList() bool {
	if o != nil && !IsNil(o.LanguagesList) {
		return true
	}

	return false
}

// SetLanguagesList gets a reference to the given []Language and assigns it to the LanguagesList field.
func (o *IssuedDocumentPreCreateInfo) SetLanguagesList(v []Language) *IssuedDocumentPreCreateInfo {
	o.LanguagesList = v
	return o
}

func (o IssuedDocumentPreCreateInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssuedDocumentPreCreateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Numerations) {
		toSerialize["numerations"] = o.Numerations
	}
	if !IsNil(o.DnNumerations) {
		toSerialize["dn_numerations"] = o.DnNumerations
	}
	if o.DefaultValues.IsSet() {
		toSerialize["default_values"] = o.DefaultValues.Get()
	}
	if o.ExtraDataDefaultValues.IsSet() {
		toSerialize["extra_data_default_values"] = o.ExtraDataDefaultValues.Get()
	}
	if o.ItemsDefaultValues.IsSet() {
		toSerialize["items_default_values"] = o.ItemsDefaultValues.Get()
	}
	if o.CountriesList != nil {
		toSerialize["countries_list"] = o.CountriesList
	}
	if o.CurrenciesList != nil {
		toSerialize["currencies_list"] = o.CurrenciesList
	}
	if o.TemplatesList != nil {
		toSerialize["templates_list"] = o.TemplatesList
	}
	if o.DnTemplatesList != nil {
		toSerialize["dn_templates_list"] = o.DnTemplatesList
	}
	if o.AiTemplatesList != nil {
		toSerialize["ai_templates_list"] = o.AiTemplatesList
	}
	if o.PaymentMethodsList != nil {
		toSerialize["payment_methods_list"] = o.PaymentMethodsList
	}
	if o.PaymentAccountsList != nil {
		toSerialize["payment_accounts_list"] = o.PaymentAccountsList
	}
	if o.VatTypesList != nil {
		toSerialize["vat_types_list"] = o.VatTypesList
	}
	if o.LanguagesList != nil {
		toSerialize["languages_list"] = o.LanguagesList
	}
	return toSerialize, nil
}

type NullableIssuedDocumentPreCreateInfo struct {
	value *IssuedDocumentPreCreateInfo
	isSet bool
}

func (v NullableIssuedDocumentPreCreateInfo) Get() *IssuedDocumentPreCreateInfo {
	return v.value
}

func (v *NullableIssuedDocumentPreCreateInfo) Set(val *IssuedDocumentPreCreateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuedDocumentPreCreateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuedDocumentPreCreateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuedDocumentPreCreateInfo(val *IssuedDocumentPreCreateInfo) *NullableIssuedDocumentPreCreateInfo {
	return &NullableIssuedDocumentPreCreateInfo{value: val, isSet: true}
}

func (v NullableIssuedDocumentPreCreateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuedDocumentPreCreateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


