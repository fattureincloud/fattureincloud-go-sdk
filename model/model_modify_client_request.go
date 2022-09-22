/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.20
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// ModifyClientRequest 
type ModifyClientRequest struct {
	Data *Client `json:"data,omitempty"`
}

// NewModifyClientRequest instantiates a new ModifyClientRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyClientRequest() *ModifyClientRequest {
	this := ModifyClientRequest{}
	return &this
}

// NewModifyClientRequestWithDefaults instantiates a new ModifyClientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyClientRequestWithDefaults() *ModifyClientRequest {
	this := ModifyClientRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModifyClientRequest) GetData() Client {
	if o == nil || o.Data == nil {
		var ret Client
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyClientRequest) GetDataOk() (*Client, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModifyClientRequest) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given Client and assigns it to the Data field.
func (o *ModifyClientRequest) SetData(v Client) *ModifyClientRequest {
	o.Data = &v
	return o
}

func (o ModifyClientRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableModifyClientRequest struct {
	value *ModifyClientRequest
	isSet bool
}

func (v NullableModifyClientRequest) Get() *ModifyClientRequest {
	return v.value
}

func (v *NullableModifyClientRequest) Set(val *ModifyClientRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyClientRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyClientRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyClientRequest(val *ModifyClientRequest) *NullableModifyClientRequest {
	return &NullableModifyClientRequest{value: val, isSet: true}
}

func (v NullableModifyClientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyClientRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


