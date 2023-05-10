/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.28
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ModifySupplierResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifySupplierResponse{}

// ModifySupplierResponse struct for ModifySupplierResponse
type ModifySupplierResponse struct {
	Data *Supplier `json:"data,omitempty"`
}

// NewModifySupplierResponse instantiates a new ModifySupplierResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifySupplierResponse() *ModifySupplierResponse {
	this := ModifySupplierResponse{}
	return &this
}

// NewModifySupplierResponseWithDefaults instantiates a new ModifySupplierResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifySupplierResponseWithDefaults() *ModifySupplierResponse {
	this := ModifySupplierResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModifySupplierResponse) GetData() Supplier {
	if o == nil || IsNil(o.Data) {
		var ret Supplier
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifySupplierResponse) GetDataOk() (*Supplier, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModifySupplierResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Supplier and assigns it to the Data field.
func (o *ModifySupplierResponse) SetData(v Supplier) *ModifySupplierResponse {
	o.Data = &v
	return o
}

func (o ModifySupplierResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifySupplierResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableModifySupplierResponse struct {
	value *ModifySupplierResponse
	isSet bool
}

func (v NullableModifySupplierResponse) Get() *ModifySupplierResponse {
	return v.value
}

func (v *NullableModifySupplierResponse) Set(val *ModifySupplierResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifySupplierResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifySupplierResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifySupplierResponse(val *ModifySupplierResponse) *NullableModifySupplierResponse {
	return &NullableModifySupplierResponse{value: val, isSet: true}
}

func (v NullableModifySupplierResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifySupplierResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


