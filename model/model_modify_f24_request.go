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

// checks if the ModifyF24Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyF24Request{}

// ModifyF24Request 
type ModifyF24Request struct {
	Data *F24 `json:"data,omitempty"`
}

// NewModifyF24Request instantiates a new ModifyF24Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyF24Request() *ModifyF24Request {
	this := ModifyF24Request{}
	return &this
}

// NewModifyF24RequestWithDefaults instantiates a new ModifyF24Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyF24RequestWithDefaults() *ModifyF24Request {
	this := ModifyF24Request{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModifyF24Request) GetData() F24 {
	if o == nil || IsNil(o.Data) {
		var ret F24
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyF24Request) GetDataOk() (*F24, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModifyF24Request) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given F24 and assigns it to the Data field.
func (o *ModifyF24Request) SetData(v F24) *ModifyF24Request {
	o.Data = &v
	return o
}

func (o ModifyF24Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyF24Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableModifyF24Request struct {
	value *ModifyF24Request
	isSet bool
}

func (v NullableModifyF24Request) Get() *ModifyF24Request {
	return v.value
}

func (v *NullableModifyF24Request) Set(val *ModifyF24Request) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyF24Request) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyF24Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyF24Request(val *ModifyF24Request) *NullableModifyF24Request {
	return &NullableModifyF24Request{value: val, isSet: true}
}

func (v NullableModifyF24Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyF24Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


