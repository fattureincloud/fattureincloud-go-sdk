/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fattureincloud

import (
	"encoding/json"
)

// CreateReceiptResponse 
type CreateReceiptResponse struct {
	Data *Receipt `json:"data,omitempty"`
}

// NewCreateReceiptResponse instantiates a new CreateReceiptResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReceiptResponse() *CreateReceiptResponse {
	this := CreateReceiptResponse{}
	return &this
}

// NewCreateReceiptResponseWithDefaults instantiates a new CreateReceiptResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReceiptResponseWithDefaults() *CreateReceiptResponse {
	this := CreateReceiptResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateReceiptResponse) GetData() Receipt {
	if o == nil || o.Data == nil {
		var ret Receipt
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateReceiptResponse) GetDataOk() (*Receipt, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateReceiptResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given Receipt and assigns it to the Data field.
func (o *CreateReceiptResponse) SetData(v Receipt) *CreateReceiptResponse {
	o.Data = &v
	return o
}

func (o CreateReceiptResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateReceiptResponse struct {
	value *CreateReceiptResponse
	isSet bool
}

func (v NullableCreateReceiptResponse) Get() *CreateReceiptResponse {
	return v.value
}

func (v *NullableCreateReceiptResponse) Set(val *CreateReceiptResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReceiptResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReceiptResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReceiptResponse(val *CreateReceiptResponse) *NullableCreateReceiptResponse {
	return &NullableCreateReceiptResponse{value: val, isSet: true}
}

func (v NullableCreateReceiptResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReceiptResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


