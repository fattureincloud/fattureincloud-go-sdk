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

// checks if the ModifyClientResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyClientResponse{}

// ModifyClientResponse struct for ModifyClientResponse
type ModifyClientResponse struct {
	Data *Client `json:"data,omitempty"`
}

// NewModifyClientResponse instantiates a new ModifyClientResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyClientResponse() *ModifyClientResponse {
	this := ModifyClientResponse{}
	return &this
}

// NewModifyClientResponseWithDefaults instantiates a new ModifyClientResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyClientResponseWithDefaults() *ModifyClientResponse {
	this := ModifyClientResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModifyClientResponse) GetData() Client {
	if o == nil || IsNil(o.Data) {
		var ret Client
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyClientResponse) GetDataOk() (*Client, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModifyClientResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Client and assigns it to the Data field.
func (o *ModifyClientResponse) SetData(v Client) *ModifyClientResponse {
	o.Data = &v
		return o
}

func (o ModifyClientResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyClientResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableModifyClientResponse struct {
	value *ModifyClientResponse
	isSet bool
}

func (v NullableModifyClientResponse) Get() *ModifyClientResponse {
	return v.value
}

func (v *NullableModifyClientResponse) Set(val *ModifyClientResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyClientResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyClientResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyClientResponse(val *ModifyClientResponse) *NullableModifyClientResponse {
	return &NullableModifyClientResponse{value: val, isSet: true}
}

func (v NullableModifyClientResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyClientResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


