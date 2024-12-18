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

// checks if the ModifyReceiptResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyReceiptResponse{}

// ModifyReceiptResponse 
type ModifyReceiptResponse struct {
	Data *Receipt `json:"data,omitempty"`
}

// NewModifyReceiptResponse instantiates a new ModifyReceiptResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyReceiptResponse() *ModifyReceiptResponse {
	this := ModifyReceiptResponse{}
	return &this
}

// NewModifyReceiptResponseWithDefaults instantiates a new ModifyReceiptResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyReceiptResponseWithDefaults() *ModifyReceiptResponse {
	this := ModifyReceiptResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModifyReceiptResponse) GetData() Receipt {
	if o == nil || IsNil(o.Data) {
		var ret Receipt
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyReceiptResponse) GetDataOk() (*Receipt, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModifyReceiptResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Receipt and assigns it to the Data field.
func (o *ModifyReceiptResponse) SetData(v Receipt) *ModifyReceiptResponse {
	o.Data = &v
	return o
}

func (o ModifyReceiptResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyReceiptResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableModifyReceiptResponse struct {
	value *ModifyReceiptResponse
	isSet bool
}

func (v NullableModifyReceiptResponse) Get() *ModifyReceiptResponse {
	return v.value
}

func (v *NullableModifyReceiptResponse) Set(val *ModifyReceiptResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyReceiptResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyReceiptResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyReceiptResponse(val *ModifyReceiptResponse) *NullableModifyReceiptResponse {
	return &NullableModifyReceiptResponse{value: val, isSet: true}
}

func (v NullableModifyReceiptResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyReceiptResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


