/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.2
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ReceiptPreCreateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceiptPreCreateInfo{}

// ReceiptPreCreateInfo struct for ReceiptPreCreateInfo
type ReceiptPreCreateInfo struct {
Numerations *map[string]map[string]int32 `json:"numerations,omitempty"`
	// Receipt used numerations list
NumerationsList []string `json:"numerations_list,omitempty"`
	// Receipt used revenue centers list
RcCentersList []string `json:"rc_centers_list,omitempty"`
	// Payment accounts list
PaymentAccountsList []PaymentAccount `json:"payment_accounts_list,omitempty"`
	// Receipt categories list
CategoriesList []string `json:"categories_list,omitempty"`
	// Vat types list
VatTypesList []VatType `json:"vat_types_list,omitempty"`
}

// NewReceiptPreCreateInfo instantiates a new ReceiptPreCreateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceiptPreCreateInfo() *ReceiptPreCreateInfo {
	this := ReceiptPreCreateInfo{}
	return &this
}

// NewReceiptPreCreateInfoWithDefaults instantiates a new ReceiptPreCreateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiptPreCreateInfoWithDefaults() *ReceiptPreCreateInfo {
	this := ReceiptPreCreateInfo{}
	return &this
}

// GetNumerations returns the Numerations field value if set, zero value otherwise.
func (o *ReceiptPreCreateInfo) GetNumerations() map[string]map[string]int32 {
	if o == nil || IsNil(o.Numerations) {
		var ret map[string]map[string]int32
		return ret
	}
	return *o.Numerations
}

// GetNumerationsOk returns a tuple with the Numerations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPreCreateInfo) GetNumerationsOk() (*map[string]map[string]int32, bool) {
	if o == nil || IsNil(o.Numerations) {
		return nil, false
	}
	return o.Numerations, true
}

// HasNumerations returns a boolean if a field has been set.
func (o *ReceiptPreCreateInfo) HasNumerations() bool {
	if o != nil && !IsNil(o.Numerations) {
		return true
	}

	return false
}

// SetNumerations gets a reference to the given map[string]map[string]int32 and assigns it to the Numerations field.
func (o *ReceiptPreCreateInfo) SetNumerations(v map[string]map[string]int32) *ReceiptPreCreateInfo {
	o.Numerations = &v
        return o
}

// GetNumerationsList returns the NumerationsList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceiptPreCreateInfo) GetNumerationsList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NumerationsList
}

// GetNumerationsListOk returns a tuple with the NumerationsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceiptPreCreateInfo) GetNumerationsListOk() ([]string, bool) {
	if o == nil || IsNil(o.NumerationsList) {
		return nil, false
	}
	return o.NumerationsList, true
}

// HasNumerationsList returns a boolean if a field has been set.
func (o *ReceiptPreCreateInfo) HasNumerationsList() bool {
	if o != nil && !IsNil(o.NumerationsList) {
		return true
	}

	return false
}

// SetNumerationsList gets a reference to the given []string and assigns it to the NumerationsList field.
func (o *ReceiptPreCreateInfo) SetNumerationsList(v []string) *ReceiptPreCreateInfo {
	o.NumerationsList = v
        return o
}

// GetRcCentersList returns the RcCentersList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceiptPreCreateInfo) GetRcCentersList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RcCentersList
}

// GetRcCentersListOk returns a tuple with the RcCentersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceiptPreCreateInfo) GetRcCentersListOk() ([]string, bool) {
	if o == nil || IsNil(o.RcCentersList) {
		return nil, false
	}
	return o.RcCentersList, true
}

// HasRcCentersList returns a boolean if a field has been set.
func (o *ReceiptPreCreateInfo) HasRcCentersList() bool {
	if o != nil && !IsNil(o.RcCentersList) {
		return true
	}

	return false
}

// SetRcCentersList gets a reference to the given []string and assigns it to the RcCentersList field.
func (o *ReceiptPreCreateInfo) SetRcCentersList(v []string) *ReceiptPreCreateInfo {
	o.RcCentersList = v
        return o
}

// GetPaymentAccountsList returns the PaymentAccountsList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceiptPreCreateInfo) GetPaymentAccountsList() []PaymentAccount {
	if o == nil {
		var ret []PaymentAccount
		return ret
	}
	return o.PaymentAccountsList
}

// GetPaymentAccountsListOk returns a tuple with the PaymentAccountsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceiptPreCreateInfo) GetPaymentAccountsListOk() ([]PaymentAccount, bool) {
	if o == nil || IsNil(o.PaymentAccountsList) {
		return nil, false
	}
	return o.PaymentAccountsList, true
}

// HasPaymentAccountsList returns a boolean if a field has been set.
func (o *ReceiptPreCreateInfo) HasPaymentAccountsList() bool {
	if o != nil && !IsNil(o.PaymentAccountsList) {
		return true
	}

	return false
}

// SetPaymentAccountsList gets a reference to the given []PaymentAccount and assigns it to the PaymentAccountsList field.
func (o *ReceiptPreCreateInfo) SetPaymentAccountsList(v []PaymentAccount) *ReceiptPreCreateInfo {
	o.PaymentAccountsList = v
        return o
}

// GetCategoriesList returns the CategoriesList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceiptPreCreateInfo) GetCategoriesList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CategoriesList
}

// GetCategoriesListOk returns a tuple with the CategoriesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceiptPreCreateInfo) GetCategoriesListOk() ([]string, bool) {
	if o == nil || IsNil(o.CategoriesList) {
		return nil, false
	}
	return o.CategoriesList, true
}

// HasCategoriesList returns a boolean if a field has been set.
func (o *ReceiptPreCreateInfo) HasCategoriesList() bool {
	if o != nil && !IsNil(o.CategoriesList) {
		return true
	}

	return false
}

// SetCategoriesList gets a reference to the given []string and assigns it to the CategoriesList field.
func (o *ReceiptPreCreateInfo) SetCategoriesList(v []string) *ReceiptPreCreateInfo {
	o.CategoriesList = v
        return o
}

// GetVatTypesList returns the VatTypesList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceiptPreCreateInfo) GetVatTypesList() []VatType {
	if o == nil {
		var ret []VatType
		return ret
	}
	return o.VatTypesList
}

// GetVatTypesListOk returns a tuple with the VatTypesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceiptPreCreateInfo) GetVatTypesListOk() ([]VatType, bool) {
	if o == nil || IsNil(o.VatTypesList) {
		return nil, false
	}
	return o.VatTypesList, true
}

// HasVatTypesList returns a boolean if a field has been set.
func (o *ReceiptPreCreateInfo) HasVatTypesList() bool {
	if o != nil && !IsNil(o.VatTypesList) {
		return true
	}

	return false
}

// SetVatTypesList gets a reference to the given []VatType and assigns it to the VatTypesList field.
func (o *ReceiptPreCreateInfo) SetVatTypesList(v []VatType) *ReceiptPreCreateInfo {
	o.VatTypesList = v
        return o
}

func (o ReceiptPreCreateInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceiptPreCreateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Numerations) {
		toSerialize["numerations"] = o.Numerations
	}
	if o.NumerationsList != nil {
		toSerialize["numerations_list"] = o.NumerationsList
	}
	if o.RcCentersList != nil {
		toSerialize["rc_centers_list"] = o.RcCentersList
	}
	if o.PaymentAccountsList != nil {
		toSerialize["payment_accounts_list"] = o.PaymentAccountsList
	}
	if o.CategoriesList != nil {
		toSerialize["categories_list"] = o.CategoriesList
	}
	if o.VatTypesList != nil {
		toSerialize["vat_types_list"] = o.VatTypesList
	}
	return toSerialize, nil
}

type NullableReceiptPreCreateInfo struct {
	value *ReceiptPreCreateInfo
	isSet bool
}

func (v NullableReceiptPreCreateInfo) Get() *ReceiptPreCreateInfo {
	return v.value
}

func (v *NullableReceiptPreCreateInfo) Set(val *ReceiptPreCreateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiptPreCreateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiptPreCreateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiptPreCreateInfo(val *ReceiptPreCreateInfo) *NullableReceiptPreCreateInfo {
	return &NullableReceiptPreCreateInfo{value: val, isSet: true}
}

func (v NullableReceiptPreCreateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiptPreCreateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


