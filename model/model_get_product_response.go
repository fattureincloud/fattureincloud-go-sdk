/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.27
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the GetProductResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductResponse{}

// GetProductResponse struct for GetProductResponse
type GetProductResponse struct {
	Data *Product `json:"data,omitempty"`
}

// NewGetProductResponse instantiates a new GetProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductResponse() *GetProductResponse {
	this := GetProductResponse{}
	return &this
}

// NewGetProductResponseWithDefaults instantiates a new GetProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductResponseWithDefaults() *GetProductResponse {
	this := GetProductResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetProductResponse) GetData() Product {
	if o == nil || IsNil(o.Data) {
		var ret Product
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetDataOk() (*Product, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetProductResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Product and assigns it to the Data field.
func (o *GetProductResponse) SetData(v Product) *GetProductResponse {
	o.Data = &v
	return o
}

func (o GetProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetProductResponse struct {
	value *GetProductResponse
	isSet bool
}

func (v NullableGetProductResponse) Get() *GetProductResponse {
	return v.value
}

func (v *NullableGetProductResponse) Set(val *GetProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductResponse(val *GetProductResponse) *NullableGetProductResponse {
	return &NullableGetProductResponse{value: val, isSet: true}
}

func (v NullableGetProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


