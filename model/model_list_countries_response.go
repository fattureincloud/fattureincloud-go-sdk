/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// ListCountriesResponse 
type ListCountriesResponse struct {
	Data []string `json:"data,omitempty"`
}

// NewListCountriesResponse instantiates a new ListCountriesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCountriesResponse() *ListCountriesResponse {
	this := ListCountriesResponse{}
	return &this
}

// NewListCountriesResponseWithDefaults instantiates a new ListCountriesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCountriesResponseWithDefaults() *ListCountriesResponse {
	this := ListCountriesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListCountriesResponse) GetData() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListCountriesResponse) GetDataOk() ([]string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListCountriesResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *ListCountriesResponse) SetData(v []string) *ListCountriesResponse {
	o.Data = v
	return o
}

func (o ListCountriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListCountriesResponse struct {
	value *ListCountriesResponse
	isSet bool
}

func (v NullableListCountriesResponse) Get() *ListCountriesResponse {
	return v.value
}

func (v *NullableListCountriesResponse) Set(val *ListCountriesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCountriesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCountriesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCountriesResponse(val *ListCountriesResponse) *NullableListCountriesResponse {
	return &NullableListCountriesResponse{value: val, isSet: true}
}

func (v NullableListCountriesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCountriesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


