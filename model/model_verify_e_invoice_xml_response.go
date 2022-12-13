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

// VerifyEInvoiceXmlResponse struct for VerifyEInvoiceXmlResponse
type VerifyEInvoiceXmlResponse struct {
	Data NullableVerifyEInvoiceXmlResponseData `json:"data,omitempty"`
}

// NewVerifyEInvoiceXmlResponse instantiates a new VerifyEInvoiceXmlResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyEInvoiceXmlResponse() *VerifyEInvoiceXmlResponse {
	this := VerifyEInvoiceXmlResponse{}
	return &this
}

// NewVerifyEInvoiceXmlResponseWithDefaults instantiates a new VerifyEInvoiceXmlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyEInvoiceXmlResponseWithDefaults() *VerifyEInvoiceXmlResponse {
	this := VerifyEInvoiceXmlResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerifyEInvoiceXmlResponse) GetData() VerifyEInvoiceXmlResponseData {
	if o == nil || isNil(o.Data.Get()) {
		var ret VerifyEInvoiceXmlResponseData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerifyEInvoiceXmlResponse) GetDataOk() (*VerifyEInvoiceXmlResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *VerifyEInvoiceXmlResponse) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableVerifyEInvoiceXmlResponseData and assigns it to the Data field.
func (o *VerifyEInvoiceXmlResponse) SetData(v VerifyEInvoiceXmlResponseData) *VerifyEInvoiceXmlResponse {
	o.Data.Set(&v)
	return o
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *VerifyEInvoiceXmlResponse) SetDataNil() *VerifyEInvoiceXmlResponse {
	o.Data.Set(nil)
	return o
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *VerifyEInvoiceXmlResponse) UnsetData() {
	o.Data.Unset()
}

func (o VerifyEInvoiceXmlResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVerifyEInvoiceXmlResponse struct {
	value *VerifyEInvoiceXmlResponse
	isSet bool
}

func (v NullableVerifyEInvoiceXmlResponse) Get() *VerifyEInvoiceXmlResponse {
	return v.value
}

func (v *NullableVerifyEInvoiceXmlResponse) Set(val *VerifyEInvoiceXmlResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyEInvoiceXmlResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyEInvoiceXmlResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyEInvoiceXmlResponse(val *VerifyEInvoiceXmlResponse) *NullableVerifyEInvoiceXmlResponse {
	return &NullableVerifyEInvoiceXmlResponse{value: val, isSet: true}
}

func (v NullableVerifyEInvoiceXmlResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyEInvoiceXmlResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


