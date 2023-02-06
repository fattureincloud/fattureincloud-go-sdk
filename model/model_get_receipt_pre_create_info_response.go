/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.25
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// GetReceiptPreCreateInfoResponse 
type GetReceiptPreCreateInfoResponse struct {
	Data *ReceiptPreCreateInfo `json:"data,omitempty"`
}

// NewGetReceiptPreCreateInfoResponse instantiates a new GetReceiptPreCreateInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReceiptPreCreateInfoResponse() *GetReceiptPreCreateInfoResponse {
	this := GetReceiptPreCreateInfoResponse{}
	return &this
}

// NewGetReceiptPreCreateInfoResponseWithDefaults instantiates a new GetReceiptPreCreateInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReceiptPreCreateInfoResponseWithDefaults() *GetReceiptPreCreateInfoResponse {
	this := GetReceiptPreCreateInfoResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetReceiptPreCreateInfoResponse) GetData() ReceiptPreCreateInfo {
	if o == nil || isNil(o.Data) {
		var ret ReceiptPreCreateInfo
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReceiptPreCreateInfoResponse) GetDataOk() (*ReceiptPreCreateInfo, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetReceiptPreCreateInfoResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ReceiptPreCreateInfo and assigns it to the Data field.
func (o *GetReceiptPreCreateInfoResponse) SetData(v ReceiptPreCreateInfo) *GetReceiptPreCreateInfoResponse {
	o.Data = &v
	return o
}

func (o GetReceiptPreCreateInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetReceiptPreCreateInfoResponse struct {
	value *GetReceiptPreCreateInfoResponse
	isSet bool
}

func (v NullableGetReceiptPreCreateInfoResponse) Get() *GetReceiptPreCreateInfoResponse {
	return v.value
}

func (v *NullableGetReceiptPreCreateInfoResponse) Set(val *GetReceiptPreCreateInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReceiptPreCreateInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReceiptPreCreateInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReceiptPreCreateInfoResponse(val *GetReceiptPreCreateInfoResponse) *NullableGetReceiptPreCreateInfoResponse {
	return &NullableGetReceiptPreCreateInfoResponse{value: val, isSet: true}
}

func (v NullableGetReceiptPreCreateInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReceiptPreCreateInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


