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

// checks if the ListCurrenciesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCurrenciesResponse{}

// ListCurrenciesResponse 
type ListCurrenciesResponse struct {
Data []Currency `json:"data,omitempty"`
}

// NewListCurrenciesResponse instantiates a new ListCurrenciesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCurrenciesResponse() *ListCurrenciesResponse {
	this := ListCurrenciesResponse{}
	return &this
}

// NewListCurrenciesResponseWithDefaults instantiates a new ListCurrenciesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCurrenciesResponseWithDefaults() *ListCurrenciesResponse {
	this := ListCurrenciesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListCurrenciesResponse) GetData() []Currency {
	if o == nil {
		var ret []Currency
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListCurrenciesResponse) GetDataOk() ([]Currency, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListCurrenciesResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Currency and assigns it to the Data field.
func (o *ListCurrenciesResponse) SetData(v []Currency) *ListCurrenciesResponse {
	o.Data = v
        return o
}

func (o ListCurrenciesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCurrenciesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListCurrenciesResponse struct {
	value *ListCurrenciesResponse
	isSet bool
}

func (v NullableListCurrenciesResponse) Get() *ListCurrenciesResponse {
	return v.value
}

func (v *NullableListCurrenciesResponse) Set(val *ListCurrenciesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCurrenciesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCurrenciesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCurrenciesResponse(val *ListCurrenciesResponse) *NullableListCurrenciesResponse {
	return &NullableListCurrenciesResponse{value: val, isSet: true}
}

func (v NullableListCurrenciesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCurrenciesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


