/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.32
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ListUserCompaniesResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUserCompaniesResponseData{}

// ListUserCompaniesResponseData struct for ListUserCompaniesResponseData
type ListUserCompaniesResponseData struct {
	Companies []Company `json:"companies,omitempty"`
}

// NewListUserCompaniesResponseData instantiates a new ListUserCompaniesResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUserCompaniesResponseData() *ListUserCompaniesResponseData {
	this := ListUserCompaniesResponseData{}
	return &this
}

// NewListUserCompaniesResponseDataWithDefaults instantiates a new ListUserCompaniesResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUserCompaniesResponseDataWithDefaults() *ListUserCompaniesResponseData {
	this := ListUserCompaniesResponseData{}
	return &this
}

// GetCompanies returns the Companies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListUserCompaniesResponseData) GetCompanies() []Company {
	if o == nil {
		var ret []Company
		return ret
	}
	return o.Companies
}

// GetCompaniesOk returns a tuple with the Companies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListUserCompaniesResponseData) GetCompaniesOk() ([]Company, bool) {
	if o == nil || IsNil(o.Companies) {
		return nil, false
	}
	return o.Companies, true
}

// HasCompanies returns a boolean if a field has been set.
func (o *ListUserCompaniesResponseData) HasCompanies() bool {
	if o != nil && IsNil(o.Companies) {
		return true
	}

	return false
}

// SetCompanies gets a reference to the given []Company and assigns it to the Companies field.
func (o *ListUserCompaniesResponseData) SetCompanies(v []Company) *ListUserCompaniesResponseData {
	o.Companies = v
	return o
}

func (o ListUserCompaniesResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUserCompaniesResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Companies != nil {
		toSerialize["companies"] = o.Companies
	}
	return toSerialize, nil
}

type NullableListUserCompaniesResponseData struct {
	value *ListUserCompaniesResponseData
	isSet bool
}

func (v NullableListUserCompaniesResponseData) Get() *ListUserCompaniesResponseData {
	return v.value
}

func (v *NullableListUserCompaniesResponseData) Set(val *ListUserCompaniesResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableListUserCompaniesResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableListUserCompaniesResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUserCompaniesResponseData(val *ListUserCompaniesResponseData) *NullableListUserCompaniesResponseData {
	return &NullableListUserCompaniesResponseData{value: val, isSet: true}
}

func (v NullableListUserCompaniesResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUserCompaniesResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


