/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.21
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// ListDetailedCountriesResponse struct for ListDetailedCountriesResponse
type ListDetailedCountriesResponse struct {
	Data []DetailedCountry `json:"data,omitempty"`
}

// NewListDetailedCountriesResponse instantiates a new ListDetailedCountriesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDetailedCountriesResponse() *ListDetailedCountriesResponse {
	this := ListDetailedCountriesResponse{}
	return &this
}

// NewListDetailedCountriesResponseWithDefaults instantiates a new ListDetailedCountriesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDetailedCountriesResponseWithDefaults() *ListDetailedCountriesResponse {
	this := ListDetailedCountriesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListDetailedCountriesResponse) GetData() []DetailedCountry {
	if o == nil || o.Data == nil {
		var ret []DetailedCountry
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDetailedCountriesResponse) GetDataOk() ([]DetailedCountry, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListDetailedCountriesResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []DetailedCountry and assigns it to the Data field.
func (o *ListDetailedCountriesResponse) SetData(v []DetailedCountry) *ListDetailedCountriesResponse {
	o.Data = v
	return o
}

func (o ListDetailedCountriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListDetailedCountriesResponse struct {
	value *ListDetailedCountriesResponse
	isSet bool
}

func (v NullableListDetailedCountriesResponse) Get() *ListDetailedCountriesResponse {
	return v.value
}

func (v *NullableListDetailedCountriesResponse) Set(val *ListDetailedCountriesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDetailedCountriesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDetailedCountriesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDetailedCountriesResponse(val *ListDetailedCountriesResponse) *NullableListDetailedCountriesResponse {
	return &NullableListDetailedCountriesResponse{value: val, isSet: true}
}

func (v NullableListDetailedCountriesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDetailedCountriesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


