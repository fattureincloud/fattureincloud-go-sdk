/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.23
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// ListArchiveCategoriesResponse struct for ListArchiveCategoriesResponse
type ListArchiveCategoriesResponse struct {
	Data []string `json:"data,omitempty"`
}

// NewListArchiveCategoriesResponse instantiates a new ListArchiveCategoriesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListArchiveCategoriesResponse() *ListArchiveCategoriesResponse {
	this := ListArchiveCategoriesResponse{}
	return &this
}

// NewListArchiveCategoriesResponseWithDefaults instantiates a new ListArchiveCategoriesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListArchiveCategoriesResponseWithDefaults() *ListArchiveCategoriesResponse {
	this := ListArchiveCategoriesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListArchiveCategoriesResponse) GetData() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListArchiveCategoriesResponse) GetDataOk() ([]string, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListArchiveCategoriesResponse) HasData() bool {
	if o != nil && isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *ListArchiveCategoriesResponse) SetData(v []string) *ListArchiveCategoriesResponse {
	o.Data = v
	return o
}

func (o ListArchiveCategoriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListArchiveCategoriesResponse struct {
	value *ListArchiveCategoriesResponse
	isSet bool
}

func (v NullableListArchiveCategoriesResponse) Get() *ListArchiveCategoriesResponse {
	return v.value
}

func (v *NullableListArchiveCategoriesResponse) Set(val *ListArchiveCategoriesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListArchiveCategoriesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListArchiveCategoriesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListArchiveCategoriesResponse(val *ListArchiveCategoriesResponse) *NullableListArchiveCategoriesResponse {
	return &NullableListArchiveCategoriesResponse{value: val, isSet: true}
}

func (v NullableListArchiveCategoriesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListArchiveCategoriesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


