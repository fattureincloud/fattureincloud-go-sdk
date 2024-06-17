/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.0
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ListVatTypesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListVatTypesResponse{}

// ListVatTypesResponse 
type ListVatTypesResponse struct {
	Data []VatType `json:"data,omitempty"`
}

// NewListVatTypesResponse instantiates a new ListVatTypesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVatTypesResponse() *ListVatTypesResponse {
	this := ListVatTypesResponse{}
	return &this
}

// NewListVatTypesResponseWithDefaults instantiates a new ListVatTypesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVatTypesResponseWithDefaults() *ListVatTypesResponse {
	this := ListVatTypesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListVatTypesResponse) GetData() []VatType {
	if o == nil {
		var ret []VatType
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListVatTypesResponse) GetDataOk() ([]VatType, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListVatTypesResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []VatType and assigns it to the Data field.
func (o *ListVatTypesResponse) SetData(v []VatType) *ListVatTypesResponse {
	o.Data = v
	return o
}

func (o ListVatTypesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListVatTypesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListVatTypesResponse struct {
	value *ListVatTypesResponse
	isSet bool
}

func (v NullableListVatTypesResponse) Get() *ListVatTypesResponse {
	return v.value
}

func (v *NullableListVatTypesResponse) Set(val *ListVatTypesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListVatTypesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListVatTypesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListVatTypesResponse(val *ListVatTypesResponse) *NullableListVatTypesResponse {
	return &NullableListVatTypesResponse{value: val, isSet: true}
}

func (v NullableListVatTypesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListVatTypesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


