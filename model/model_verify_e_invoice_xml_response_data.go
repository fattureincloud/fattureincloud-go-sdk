/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// VerifyEInvoiceXmlResponseData struct for VerifyEInvoiceXmlResponseData
type VerifyEInvoiceXmlResponseData struct {
	// Determine if the invoice XML is valid.
	Success NullableBool `json:"success,omitempty"`
}

// NewVerifyEInvoiceXmlResponseData instantiates a new VerifyEInvoiceXmlResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyEInvoiceXmlResponseData() *VerifyEInvoiceXmlResponseData {
	this := VerifyEInvoiceXmlResponseData{}
	return &this
}

// NewVerifyEInvoiceXmlResponseDataWithDefaults instantiates a new VerifyEInvoiceXmlResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyEInvoiceXmlResponseDataWithDefaults() *VerifyEInvoiceXmlResponseData {
	this := VerifyEInvoiceXmlResponseData{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerifyEInvoiceXmlResponseData) GetSuccess() bool {
	if o == nil || o.Success.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Success.Get()
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerifyEInvoiceXmlResponseData) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Success.Get(), o.Success.IsSet()
}

// HasSuccess returns a boolean if a field has been set.
func (o *VerifyEInvoiceXmlResponseData) HasSuccess() bool {
	if o != nil && o.Success.IsSet() {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given NullableBool and assigns it to the Success field.
func (o *VerifyEInvoiceXmlResponseData) SetSuccess(v bool) *VerifyEInvoiceXmlResponseData {
	o.Success.Set(&v)
	return o
}
// SetSuccessNil sets the value for Success to be an explicit nil
func (o *VerifyEInvoiceXmlResponseData) SetSuccessNil() *VerifyEInvoiceXmlResponseData {
	o.Success.Set(nil)
	return o
}

// UnsetSuccess ensures that no value is present for Success, not even an explicit nil
func (o *VerifyEInvoiceXmlResponseData) UnsetSuccess() {
	o.Success.Unset()
}

func (o VerifyEInvoiceXmlResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success.IsSet() {
		toSerialize["success"] = o.Success.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVerifyEInvoiceXmlResponseData struct {
	value *VerifyEInvoiceXmlResponseData
	isSet bool
}

func (v NullableVerifyEInvoiceXmlResponseData) Get() *VerifyEInvoiceXmlResponseData {
	return v.value
}

func (v *NullableVerifyEInvoiceXmlResponseData) Set(val *VerifyEInvoiceXmlResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyEInvoiceXmlResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyEInvoiceXmlResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyEInvoiceXmlResponseData(val *VerifyEInvoiceXmlResponseData) *NullableVerifyEInvoiceXmlResponseData {
	return &NullableVerifyEInvoiceXmlResponseData{value: val, isSet: true}
}

func (v NullableVerifyEInvoiceXmlResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyEInvoiceXmlResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


