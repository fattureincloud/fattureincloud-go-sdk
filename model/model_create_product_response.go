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

// CreateProductResponse struct for CreateProductResponse
type CreateProductResponse struct {
	Data *Product `json:"data,omitempty"`
}

// NewCreateProductResponse instantiates a new CreateProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProductResponse() *CreateProductResponse {
	this := CreateProductResponse{}
	return &this
}

// NewCreateProductResponseWithDefaults instantiates a new CreateProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProductResponseWithDefaults() *CreateProductResponse {
	this := CreateProductResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateProductResponse) GetData() Product {
	if o == nil || isNil(o.Data) {
		var ret Product
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductResponse) GetDataOk() (*Product, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateProductResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Product and assigns it to the Data field.
func (o *CreateProductResponse) SetData(v Product) *CreateProductResponse {
	o.Data = &v
	return o
}

func (o CreateProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateProductResponse struct {
	value *CreateProductResponse
	isSet bool
}

func (v NullableCreateProductResponse) Get() *CreateProductResponse {
	return v.value
}

func (v *NullableCreateProductResponse) Set(val *CreateProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProductResponse(val *CreateProductResponse) *NullableCreateProductResponse {
	return &NullableCreateProductResponse{value: val, isSet: true}
}

func (v NullableCreateProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


