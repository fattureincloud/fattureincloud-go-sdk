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

// GetSupplierResponse struct for GetSupplierResponse
type GetSupplierResponse struct {
	Data *Supplier `json:"data,omitempty"`
}

// NewGetSupplierResponse instantiates a new GetSupplierResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSupplierResponse() *GetSupplierResponse {
	this := GetSupplierResponse{}
	return &this
}

// NewGetSupplierResponseWithDefaults instantiates a new GetSupplierResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSupplierResponseWithDefaults() *GetSupplierResponse {
	this := GetSupplierResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetSupplierResponse) GetData() Supplier {
	if o == nil || isNil(o.Data) {
		var ret Supplier
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSupplierResponse) GetDataOk() (*Supplier, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetSupplierResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Supplier and assigns it to the Data field.
func (o *GetSupplierResponse) SetData(v Supplier) *GetSupplierResponse {
	o.Data = &v
	return o
}

func (o GetSupplierResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetSupplierResponse struct {
	value *GetSupplierResponse
	isSet bool
}

func (v NullableGetSupplierResponse) Get() *GetSupplierResponse {
	return v.value
}

func (v *NullableGetSupplierResponse) Set(val *GetSupplierResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSupplierResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSupplierResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSupplierResponse(val *GetSupplierResponse) *NullableGetSupplierResponse {
	return &NullableGetSupplierResponse{value: val, isSet: true}
}

func (v NullableGetSupplierResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSupplierResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


