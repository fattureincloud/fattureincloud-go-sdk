/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.2
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the CreateClientRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateClientRequest{}

// CreateClientRequest struct for CreateClientRequest
type CreateClientRequest struct {
	Data *Client `json:"data,omitempty"`
}

// NewCreateClientRequest instantiates a new CreateClientRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClientRequest() *CreateClientRequest {
	this := CreateClientRequest{}
	return &this
}

// NewCreateClientRequestWithDefaults instantiates a new CreateClientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClientRequestWithDefaults() *CreateClientRequest {
	this := CreateClientRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateClientRequest) GetData() Client {
	if o == nil || IsNil(o.Data) {
		var ret Client
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetDataOk() (*Client, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateClientRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Client and assigns it to the Data field.
func (o *CreateClientRequest) SetData(v Client) *CreateClientRequest {
	o.Data = &v
	return o
}

func (o CreateClientRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateClientRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateClientRequest struct {
	value *CreateClientRequest
	isSet bool
}

func (v NullableCreateClientRequest) Get() *CreateClientRequest {
	return v.value
}

func (v *NullableCreateClientRequest) Set(val *CreateClientRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClientRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClientRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClientRequest(val *CreateClientRequest) *NullableCreateClientRequest {
	return &NullableCreateClientRequest{value: val, isSet: true}
}

func (v NullableCreateClientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClientRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


