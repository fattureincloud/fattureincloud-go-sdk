/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.22
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// ModifyPaymentAccountRequest struct for ModifyPaymentAccountRequest
type ModifyPaymentAccountRequest struct {
	Data NullablePaymentAccount `json:"data,omitempty"`
}

// NewModifyPaymentAccountRequest instantiates a new ModifyPaymentAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyPaymentAccountRequest() *ModifyPaymentAccountRequest {
	this := ModifyPaymentAccountRequest{}
	return &this
}

// NewModifyPaymentAccountRequestWithDefaults instantiates a new ModifyPaymentAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyPaymentAccountRequestWithDefaults() *ModifyPaymentAccountRequest {
	this := ModifyPaymentAccountRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModifyPaymentAccountRequest) GetData() PaymentAccount {
	if o == nil || o.Data.Get() == nil {
		var ret PaymentAccount
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModifyPaymentAccountRequest) GetDataOk() (*PaymentAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *ModifyPaymentAccountRequest) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullablePaymentAccount and assigns it to the Data field.
func (o *ModifyPaymentAccountRequest) SetData(v PaymentAccount) *ModifyPaymentAccountRequest {
	o.Data.Set(&v)
	return o
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *ModifyPaymentAccountRequest) SetDataNil() *ModifyPaymentAccountRequest {
	o.Data.Set(nil)
	return o
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *ModifyPaymentAccountRequest) UnsetData() {
	o.Data.Unset()
}

func (o ModifyPaymentAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableModifyPaymentAccountRequest struct {
	value *ModifyPaymentAccountRequest
	isSet bool
}

func (v NullableModifyPaymentAccountRequest) Get() *ModifyPaymentAccountRequest {
	return v.value
}

func (v *NullableModifyPaymentAccountRequest) Set(val *ModifyPaymentAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyPaymentAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyPaymentAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyPaymentAccountRequest(val *ModifyPaymentAccountRequest) *NullableModifyPaymentAccountRequest {
	return &NullableModifyPaymentAccountRequest{value: val, isSet: true}
}

func (v NullableModifyPaymentAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyPaymentAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


