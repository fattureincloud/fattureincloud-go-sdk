/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.22
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// CreateF24Response 
type CreateF24Response struct {
	Data *F24 `json:"data,omitempty"`
}

// NewCreateF24Response instantiates a new CreateF24Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateF24Response() *CreateF24Response {
	this := CreateF24Response{}
	return &this
}

// NewCreateF24ResponseWithDefaults instantiates a new CreateF24Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateF24ResponseWithDefaults() *CreateF24Response {
	this := CreateF24Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateF24Response) GetData() F24 {
	if o == nil || o.Data == nil {
		var ret F24
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateF24Response) GetDataOk() (*F24, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateF24Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given F24 and assigns it to the Data field.
func (o *CreateF24Response) SetData(v F24) *CreateF24Response {
	o.Data = &v
	return o
}

func (o CreateF24Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateF24Response struct {
	value *CreateF24Response
	isSet bool
}

func (v NullableCreateF24Response) Get() *CreateF24Response {
	return v.value
}

func (v *NullableCreateF24Response) Set(val *CreateF24Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateF24Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateF24Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateF24Response(val *CreateF24Response) *NullableCreateF24Response {
	return &NullableCreateF24Response{value: val, isSet: true}
}

func (v NullableCreateF24Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateF24Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


