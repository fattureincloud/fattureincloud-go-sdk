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

// checks if the ListUserCompaniesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUserCompaniesResponse{}

// ListUserCompaniesResponse 
type ListUserCompaniesResponse struct {
	Data NullableListUserCompaniesResponseData `json:"data,omitempty"`
}

// NewListUserCompaniesResponse instantiates a new ListUserCompaniesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUserCompaniesResponse() *ListUserCompaniesResponse {
	this := ListUserCompaniesResponse{}
	return &this
}

// NewListUserCompaniesResponseWithDefaults instantiates a new ListUserCompaniesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUserCompaniesResponseWithDefaults() *ListUserCompaniesResponse {
	this := ListUserCompaniesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListUserCompaniesResponse) GetData() ListUserCompaniesResponseData {
	if o == nil || IsNil(o.Data.Get()) {
		var ret ListUserCompaniesResponseData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListUserCompaniesResponse) GetDataOk() (*ListUserCompaniesResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *ListUserCompaniesResponse) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableListUserCompaniesResponseData and assigns it to the Data field.
func (o *ListUserCompaniesResponse) SetData(v ListUserCompaniesResponseData) *ListUserCompaniesResponse {
	o.Data.Set(&v)
	return o
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *ListUserCompaniesResponse) SetDataNil() *ListUserCompaniesResponse {
	o.Data.Set(nil)
	return o
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *ListUserCompaniesResponse) UnsetData() {
	o.Data.Unset()
}

func (o ListUserCompaniesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUserCompaniesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return toSerialize, nil
}

type NullableListUserCompaniesResponse struct {
	value *ListUserCompaniesResponse
	isSet bool
}

func (v NullableListUserCompaniesResponse) Get() *ListUserCompaniesResponse {
	return v.value
}

func (v *NullableListUserCompaniesResponse) Set(val *ListUserCompaniesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListUserCompaniesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListUserCompaniesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUserCompaniesResponse(val *ListUserCompaniesResponse) *NullableListUserCompaniesResponse {
	return &NullableListUserCompaniesResponse{value: val, isSet: true}
}

func (v NullableListUserCompaniesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUserCompaniesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


