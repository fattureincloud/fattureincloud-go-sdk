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

// checks if the ReceivedDocumentInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceivedDocumentInfo{}

// ReceivedDocumentInfo struct for ReceivedDocumentInfo
type ReceivedDocumentInfo struct {
	DefaultValues NullableReceivedDocumentInfoDefaultValues `json:"default_values,omitempty"`
	ItemsDefaultValues NullableReceivedDocumentInfoItemsDefaultValues `json:"items_default_values,omitempty"`
	// Countries list
	CountriesList []string `json:"countries_list,omitempty"`
	// Currencies list
	CurrenciesList []Currency `json:"currencies_list,omitempty"`
	// Categories list
	CategoriesList []string `json:"categories_list,omitempty"`
	// Payments accounts list
	PaymentAccountsList []PaymentAccount `json:"payment_accounts_list,omitempty"`
	// Vat types list
	VatTypesList []VatType `json:"vat_types_list,omitempty"`
}

// NewReceivedDocumentInfo instantiates a new ReceivedDocumentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceivedDocumentInfo() *ReceivedDocumentInfo {
	this := ReceivedDocumentInfo{}
	return &this
}

// NewReceivedDocumentInfoWithDefaults instantiates a new ReceivedDocumentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceivedDocumentInfoWithDefaults() *ReceivedDocumentInfo {
	this := ReceivedDocumentInfo{}
	return &this
}

// GetDefaultValues returns the DefaultValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentInfo) GetDefaultValues() ReceivedDocumentInfoDefaultValues {
	if o == nil || IsNil(o.DefaultValues.Get()) {
		var ret ReceivedDocumentInfoDefaultValues
		return ret
	}
	return *o.DefaultValues.Get()
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentInfo) GetDefaultValuesOk() (*ReceivedDocumentInfoDefaultValues, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultValues.Get(), o.DefaultValues.IsSet()
}

// HasDefaultValues returns a boolean if a field has been set.
func (o *ReceivedDocumentInfo) HasDefaultValues() bool {
	if o != nil && o.DefaultValues.IsSet() {
		return true
	}

	return false
}

// SetDefaultValues gets a reference to the given NullableReceivedDocumentInfoDefaultValues and assigns it to the DefaultValues field.
func (o *ReceivedDocumentInfo) SetDefaultValues(v ReceivedDocumentInfoDefaultValues) *ReceivedDocumentInfo {
	o.DefaultValues.Set(&v)
	return o
}
// SetDefaultValuesNil sets the value for DefaultValues to be an explicit nil
func (o *ReceivedDocumentInfo) SetDefaultValuesNil() *ReceivedDocumentInfo {
	o.DefaultValues.Set(nil)
	return o
}

// UnsetDefaultValues ensures that no value is present for DefaultValues, not even an explicit nil
func (o *ReceivedDocumentInfo) UnsetDefaultValues() {
	o.DefaultValues.Unset()
}

// GetItemsDefaultValues returns the ItemsDefaultValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentInfo) GetItemsDefaultValues() ReceivedDocumentInfoItemsDefaultValues {
	if o == nil || IsNil(o.ItemsDefaultValues.Get()) {
		var ret ReceivedDocumentInfoItemsDefaultValues
		return ret
	}
	return *o.ItemsDefaultValues.Get()
}

// GetItemsDefaultValuesOk returns a tuple with the ItemsDefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentInfo) GetItemsDefaultValuesOk() (*ReceivedDocumentInfoItemsDefaultValues, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemsDefaultValues.Get(), o.ItemsDefaultValues.IsSet()
}

// HasItemsDefaultValues returns a boolean if a field has been set.
func (o *ReceivedDocumentInfo) HasItemsDefaultValues() bool {
	if o != nil && o.ItemsDefaultValues.IsSet() {
		return true
	}

	return false
}

// SetItemsDefaultValues gets a reference to the given NullableReceivedDocumentInfoItemsDefaultValues and assigns it to the ItemsDefaultValues field.
func (o *ReceivedDocumentInfo) SetItemsDefaultValues(v ReceivedDocumentInfoItemsDefaultValues) *ReceivedDocumentInfo {
	o.ItemsDefaultValues.Set(&v)
	return o
}
// SetItemsDefaultValuesNil sets the value for ItemsDefaultValues to be an explicit nil
func (o *ReceivedDocumentInfo) SetItemsDefaultValuesNil() *ReceivedDocumentInfo {
	o.ItemsDefaultValues.Set(nil)
	return o
}

// UnsetItemsDefaultValues ensures that no value is present for ItemsDefaultValues, not even an explicit nil
func (o *ReceivedDocumentInfo) UnsetItemsDefaultValues() {
	o.ItemsDefaultValues.Unset()
}

// GetCountriesList returns the CountriesList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentInfo) GetCountriesList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CountriesList
}

// GetCountriesListOk returns a tuple with the CountriesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentInfo) GetCountriesListOk() ([]string, bool) {
	if o == nil || IsNil(o.CountriesList) {
		return nil, false
	}
	return o.CountriesList, true
}

// HasCountriesList returns a boolean if a field has been set.
func (o *ReceivedDocumentInfo) HasCountriesList() bool {
	if o != nil && !IsNil(o.CountriesList) {
		return true
	}

	return false
}

// SetCountriesList gets a reference to the given []string and assigns it to the CountriesList field.
func (o *ReceivedDocumentInfo) SetCountriesList(v []string) *ReceivedDocumentInfo {
	o.CountriesList = v
	return o
}

// GetCurrenciesList returns the CurrenciesList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentInfo) GetCurrenciesList() []Currency {
	if o == nil {
		var ret []Currency
		return ret
	}
	return o.CurrenciesList
}

// GetCurrenciesListOk returns a tuple with the CurrenciesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentInfo) GetCurrenciesListOk() ([]Currency, bool) {
	if o == nil || IsNil(o.CurrenciesList) {
		return nil, false
	}
	return o.CurrenciesList, true
}

// HasCurrenciesList returns a boolean if a field has been set.
func (o *ReceivedDocumentInfo) HasCurrenciesList() bool {
	if o != nil && !IsNil(o.CurrenciesList) {
		return true
	}

	return false
}

// SetCurrenciesList gets a reference to the given []Currency and assigns it to the CurrenciesList field.
func (o *ReceivedDocumentInfo) SetCurrenciesList(v []Currency) *ReceivedDocumentInfo {
	o.CurrenciesList = v
	return o
}

// GetCategoriesList returns the CategoriesList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentInfo) GetCategoriesList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CategoriesList
}

// GetCategoriesListOk returns a tuple with the CategoriesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentInfo) GetCategoriesListOk() ([]string, bool) {
	if o == nil || IsNil(o.CategoriesList) {
		return nil, false
	}
	return o.CategoriesList, true
}

// HasCategoriesList returns a boolean if a field has been set.
func (o *ReceivedDocumentInfo) HasCategoriesList() bool {
	if o != nil && !IsNil(o.CategoriesList) {
		return true
	}

	return false
}

// SetCategoriesList gets a reference to the given []string and assigns it to the CategoriesList field.
func (o *ReceivedDocumentInfo) SetCategoriesList(v []string) *ReceivedDocumentInfo {
	o.CategoriesList = v
	return o
}

// GetPaymentAccountsList returns the PaymentAccountsList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentInfo) GetPaymentAccountsList() []PaymentAccount {
	if o == nil {
		var ret []PaymentAccount
		return ret
	}
	return o.PaymentAccountsList
}

// GetPaymentAccountsListOk returns a tuple with the PaymentAccountsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentInfo) GetPaymentAccountsListOk() ([]PaymentAccount, bool) {
	if o == nil || IsNil(o.PaymentAccountsList) {
		return nil, false
	}
	return o.PaymentAccountsList, true
}

// HasPaymentAccountsList returns a boolean if a field has been set.
func (o *ReceivedDocumentInfo) HasPaymentAccountsList() bool {
	if o != nil && !IsNil(o.PaymentAccountsList) {
		return true
	}

	return false
}

// SetPaymentAccountsList gets a reference to the given []PaymentAccount and assigns it to the PaymentAccountsList field.
func (o *ReceivedDocumentInfo) SetPaymentAccountsList(v []PaymentAccount) *ReceivedDocumentInfo {
	o.PaymentAccountsList = v
	return o
}

// GetVatTypesList returns the VatTypesList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentInfo) GetVatTypesList() []VatType {
	if o == nil {
		var ret []VatType
		return ret
	}
	return o.VatTypesList
}

// GetVatTypesListOk returns a tuple with the VatTypesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentInfo) GetVatTypesListOk() ([]VatType, bool) {
	if o == nil || IsNil(o.VatTypesList) {
		return nil, false
	}
	return o.VatTypesList, true
}

// HasVatTypesList returns a boolean if a field has been set.
func (o *ReceivedDocumentInfo) HasVatTypesList() bool {
	if o != nil && !IsNil(o.VatTypesList) {
		return true
	}

	return false
}

// SetVatTypesList gets a reference to the given []VatType and assigns it to the VatTypesList field.
func (o *ReceivedDocumentInfo) SetVatTypesList(v []VatType) *ReceivedDocumentInfo {
	o.VatTypesList = v
	return o
}

func (o ReceivedDocumentInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceivedDocumentInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultValues.IsSet() {
		toSerialize["default_values"] = o.DefaultValues.Get()
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
	if o.CategoriesList != nil {
		toSerialize["categories_list"] = o.CategoriesList
	}
	if o.PaymentAccountsList != nil {
		toSerialize["payment_accounts_list"] = o.PaymentAccountsList
	}
	if o.VatTypesList != nil {
		toSerialize["vat_types_list"] = o.VatTypesList
	}
	return toSerialize, nil
}

type NullableReceivedDocumentInfo struct {
	value *ReceivedDocumentInfo
	isSet bool
}

func (v NullableReceivedDocumentInfo) Get() *ReceivedDocumentInfo {
	return v.value
}

func (v *NullableReceivedDocumentInfo) Set(val *ReceivedDocumentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableReceivedDocumentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableReceivedDocumentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceivedDocumentInfo(val *ReceivedDocumentInfo) *NullableReceivedDocumentInfo {
	return &NullableReceivedDocumentInfo{value: val, isSet: true}
}

func (v NullableReceivedDocumentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceivedDocumentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


