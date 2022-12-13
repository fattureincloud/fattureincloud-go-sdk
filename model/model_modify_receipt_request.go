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

// ModifyReceiptRequest 
type ModifyReceiptRequest struct {
	Data *Receipt `json:"data,omitempty"`
}

// NewModifyReceiptRequest instantiates a new ModifyReceiptRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyReceiptRequest() *ModifyReceiptRequest {
	this := ModifyReceiptRequest{}
	return &this
}

// NewModifyReceiptRequestWithDefaults instantiates a new ModifyReceiptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyReceiptRequestWithDefaults() *ModifyReceiptRequest {
	this := ModifyReceiptRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModifyReceiptRequest) GetData() Receipt {
	if o == nil || isNil(o.Data) {
		var ret Receipt
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyReceiptRequest) GetDataOk() (*Receipt, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModifyReceiptRequest) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Receipt and assigns it to the Data field.
func (o *ModifyReceiptRequest) SetData(v Receipt) *ModifyReceiptRequest {
	o.Data = &v
	return o
}

func (o ModifyReceiptRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableModifyReceiptRequest struct {
	value *ModifyReceiptRequest
	isSet bool
}

func (v NullableModifyReceiptRequest) Get() *ModifyReceiptRequest {
	return v.value
}

func (v *NullableModifyReceiptRequest) Set(val *ModifyReceiptRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyReceiptRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyReceiptRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyReceiptRequest(val *ModifyReceiptRequest) *NullableModifyReceiptRequest {
	return &NullableModifyReceiptRequest{value: val, isSet: true}
}

func (v NullableModifyReceiptRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyReceiptRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


