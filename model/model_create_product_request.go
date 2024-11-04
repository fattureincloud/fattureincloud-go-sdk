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

// checks if the CreateProductRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProductRequest{}

// CreateProductRequest struct for CreateProductRequest
type CreateProductRequest struct {
Data *Product `json:"data,omitempty"`
}

// NewCreateProductRequest instantiates a new CreateProductRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProductRequest() *CreateProductRequest {
	this := CreateProductRequest{}
	return &this
}

// NewCreateProductRequestWithDefaults instantiates a new CreateProductRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProductRequestWithDefaults() *CreateProductRequest {
	this := CreateProductRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateProductRequest) GetData() Product {
	if o == nil || IsNil(o.Data) {
		var ret Product
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductRequest) GetDataOk() (*Product, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateProductRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Product and assigns it to the Data field.
func (o *CreateProductRequest) SetData(v Product) *CreateProductRequest {
	o.Data = &v
		return o
}

func (o CreateProductRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProductRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateProductRequest struct {
	value *CreateProductRequest
	isSet bool
}

func (v NullableCreateProductRequest) Get() *CreateProductRequest {
	return v.value
}

func (v *NullableCreateProductRequest) Set(val *CreateProductRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProductRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProductRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProductRequest(val *CreateProductRequest) *NullableCreateProductRequest {
	return &NullableCreateProductRequest{value: val, isSet: true}
}

func (v NullableCreateProductRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProductRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


