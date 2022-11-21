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

// ListRevenueCentersResponse 
type ListRevenueCentersResponse struct {
	Data []string `json:"data,omitempty"`
}

// NewListRevenueCentersResponse instantiates a new ListRevenueCentersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRevenueCentersResponse() *ListRevenueCentersResponse {
	this := ListRevenueCentersResponse{}
	return &this
}

// NewListRevenueCentersResponseWithDefaults instantiates a new ListRevenueCentersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRevenueCentersResponseWithDefaults() *ListRevenueCentersResponse {
	this := ListRevenueCentersResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListRevenueCentersResponse) GetData() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListRevenueCentersResponse) GetDataOk() ([]string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListRevenueCentersResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *ListRevenueCentersResponse) SetData(v []string) *ListRevenueCentersResponse {
	o.Data = v
	return o
}

func (o ListRevenueCentersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListRevenueCentersResponse struct {
	value *ListRevenueCentersResponse
	isSet bool
}

func (v NullableListRevenueCentersResponse) Get() *ListRevenueCentersResponse {
	return v.value
}

func (v *NullableListRevenueCentersResponse) Set(val *ListRevenueCentersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRevenueCentersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRevenueCentersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRevenueCentersResponse(val *ListRevenueCentersResponse) *NullableListRevenueCentersResponse {
	return &NullableListRevenueCentersResponse{value: val, isSet: true}
}

func (v NullableListRevenueCentersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRevenueCentersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


