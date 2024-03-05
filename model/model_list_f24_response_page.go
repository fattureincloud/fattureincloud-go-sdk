/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.32
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ListF24ResponsePage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListF24ResponsePage{}

// ListF24ResponsePage struct for ListF24ResponsePage
type ListF24ResponsePage struct {
	Data []F24 `json:"data,omitempty"`
}

// NewListF24ResponsePage instantiates a new ListF24ResponsePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListF24ResponsePage() *ListF24ResponsePage {
	this := ListF24ResponsePage{}
	return &this
}

// NewListF24ResponsePageWithDefaults instantiates a new ListF24ResponsePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListF24ResponsePageWithDefaults() *ListF24ResponsePage {
	this := ListF24ResponsePage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListF24ResponsePage) GetData() []F24 {
	if o == nil {
		var ret []F24
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListF24ResponsePage) GetDataOk() ([]F24, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListF24ResponsePage) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []F24 and assigns it to the Data field.
func (o *ListF24ResponsePage) SetData(v []F24) *ListF24ResponsePage {
	o.Data = v
	return o
}

func (o ListF24ResponsePage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListF24ResponsePage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListF24ResponsePage struct {
	value *ListF24ResponsePage
	isSet bool
}

func (v NullableListF24ResponsePage) Get() *ListF24ResponsePage {
	return v.value
}

func (v *NullableListF24ResponsePage) Set(val *ListF24ResponsePage) {
	v.value = val
	v.isSet = true
}

func (v NullableListF24ResponsePage) IsSet() bool {
	return v.isSet
}

func (v *NullableListF24ResponsePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListF24ResponsePage(val *ListF24ResponsePage) *NullableListF24ResponsePage {
	return &NullableListF24ResponsePage{value: val, isSet: true}
}

func (v NullableListF24ResponsePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListF24ResponsePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


