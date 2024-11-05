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

// checks if the ListLanguagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListLanguagesResponse{}

// ListLanguagesResponse 
type ListLanguagesResponse struct {
	Data []Language `json:"data,omitempty"`
}

// NewListLanguagesResponse instantiates a new ListLanguagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLanguagesResponse() *ListLanguagesResponse {
	this := ListLanguagesResponse{}
	return &this
}

// NewListLanguagesResponseWithDefaults instantiates a new ListLanguagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLanguagesResponseWithDefaults() *ListLanguagesResponse {
	this := ListLanguagesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListLanguagesResponse) GetData() []Language {
	if o == nil {
		var ret []Language
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListLanguagesResponse) GetDataOk() ([]Language, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListLanguagesResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Language and assigns it to the Data field.
func (o *ListLanguagesResponse) SetData(v []Language) *ListLanguagesResponse {
	o.Data = v
	return o
}

func (o ListLanguagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListLanguagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListLanguagesResponse struct {
	value *ListLanguagesResponse
	isSet bool
}

func (v NullableListLanguagesResponse) Get() *ListLanguagesResponse {
	return v.value
}

func (v *NullableListLanguagesResponse) Set(val *ListLanguagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListLanguagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListLanguagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLanguagesResponse(val *ListLanguagesResponse) *NullableListLanguagesResponse {
	return &NullableListLanguagesResponse{value: val, isSet: true}
}

func (v NullableListLanguagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLanguagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


