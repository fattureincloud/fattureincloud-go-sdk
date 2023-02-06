/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.25
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// CreateClientResponse struct for CreateClientResponse
type CreateClientResponse struct {
	Data *Client `json:"data,omitempty"`
}

// NewCreateClientResponse instantiates a new CreateClientResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClientResponse() *CreateClientResponse {
	this := CreateClientResponse{}
	return &this
}

// NewCreateClientResponseWithDefaults instantiates a new CreateClientResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClientResponseWithDefaults() *CreateClientResponse {
	this := CreateClientResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateClientResponse) GetData() Client {
	if o == nil || isNil(o.Data) {
		var ret Client
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClientResponse) GetDataOk() (*Client, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateClientResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Client and assigns it to the Data field.
func (o *CreateClientResponse) SetData(v Client) *CreateClientResponse {
	o.Data = &v
	return o
}

func (o CreateClientResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateClientResponse struct {
	value *CreateClientResponse
	isSet bool
}

func (v NullableCreateClientResponse) Get() *CreateClientResponse {
	return v.value
}

func (v *NullableCreateClientResponse) Set(val *CreateClientResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClientResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClientResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClientResponse(val *CreateClientResponse) *NullableCreateClientResponse {
	return &NullableCreateClientResponse{value: val, isSet: true}
}

func (v NullableCreateClientResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClientResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


