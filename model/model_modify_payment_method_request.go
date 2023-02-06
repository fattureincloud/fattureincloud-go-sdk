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

// ModifyPaymentMethodRequest struct for ModifyPaymentMethodRequest
type ModifyPaymentMethodRequest struct {
	Data *PaymentMethod `json:"data,omitempty"`
}

// NewModifyPaymentMethodRequest instantiates a new ModifyPaymentMethodRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyPaymentMethodRequest() *ModifyPaymentMethodRequest {
	this := ModifyPaymentMethodRequest{}
	return &this
}

// NewModifyPaymentMethodRequestWithDefaults instantiates a new ModifyPaymentMethodRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyPaymentMethodRequestWithDefaults() *ModifyPaymentMethodRequest {
	this := ModifyPaymentMethodRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModifyPaymentMethodRequest) GetData() PaymentMethod {
	if o == nil || isNil(o.Data) {
		var ret PaymentMethod
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyPaymentMethodRequest) GetDataOk() (*PaymentMethod, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModifyPaymentMethodRequest) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PaymentMethod and assigns it to the Data field.
func (o *ModifyPaymentMethodRequest) SetData(v PaymentMethod) *ModifyPaymentMethodRequest {
	o.Data = &v
	return o
}

func (o ModifyPaymentMethodRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableModifyPaymentMethodRequest struct {
	value *ModifyPaymentMethodRequest
	isSet bool
}

func (v NullableModifyPaymentMethodRequest) Get() *ModifyPaymentMethodRequest {
	return v.value
}

func (v *NullableModifyPaymentMethodRequest) Set(val *ModifyPaymentMethodRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyPaymentMethodRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyPaymentMethodRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyPaymentMethodRequest(val *ModifyPaymentMethodRequest) *NullableModifyPaymentMethodRequest {
	return &NullableModifyPaymentMethodRequest{value: val, isSet: true}
}

func (v NullableModifyPaymentMethodRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyPaymentMethodRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


