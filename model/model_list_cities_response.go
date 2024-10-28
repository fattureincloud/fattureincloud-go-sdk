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

// checks if the ListCitiesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCitiesResponse{}

// ListCitiesResponse 
type ListCitiesResponse struct {
Data []City `json:"data,omitempty"`
}

// NewListCitiesResponse instantiates a new ListCitiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCitiesResponse() *ListCitiesResponse {
	this := ListCitiesResponse{}
	return &this
}

// NewListCitiesResponseWithDefaults instantiates a new ListCitiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCitiesResponseWithDefaults() *ListCitiesResponse {
	this := ListCitiesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListCitiesResponse) GetData() []City {
	if o == nil {
		var ret []City
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListCitiesResponse) GetDataOk() ([]City, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListCitiesResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []City and assigns it to the Data field.
func (o *ListCitiesResponse) SetData(v []City) *ListCitiesResponse {
	o.Data = v
        return o
}

func (o ListCitiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCitiesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListCitiesResponse struct {
	value *ListCitiesResponse
	isSet bool
}

func (v NullableListCitiesResponse) Get() *ListCitiesResponse {
	return v.value
}

func (v *NullableListCitiesResponse) Set(val *ListCitiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCitiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCitiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCitiesResponse(val *ListCitiesResponse) *NullableListCitiesResponse {
	return &NullableListCitiesResponse{value: val, isSet: true}
}

func (v NullableListCitiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCitiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


