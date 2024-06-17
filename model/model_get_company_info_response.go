/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.0
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the GetCompanyInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCompanyInfoResponse{}

// GetCompanyInfoResponse 
type GetCompanyInfoResponse struct {
	Data *CompanyInfo `json:"data,omitempty"`
}

// NewGetCompanyInfoResponse instantiates a new GetCompanyInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCompanyInfoResponse() *GetCompanyInfoResponse {
	this := GetCompanyInfoResponse{}
	return &this
}

// NewGetCompanyInfoResponseWithDefaults instantiates a new GetCompanyInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCompanyInfoResponseWithDefaults() *GetCompanyInfoResponse {
	this := GetCompanyInfoResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetCompanyInfoResponse) GetData() CompanyInfo {
	if o == nil || IsNil(o.Data) {
		var ret CompanyInfo
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCompanyInfoResponse) GetDataOk() (*CompanyInfo, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetCompanyInfoResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CompanyInfo and assigns it to the Data field.
func (o *GetCompanyInfoResponse) SetData(v CompanyInfo) *GetCompanyInfoResponse {
	o.Data = &v
	return o
}

func (o GetCompanyInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCompanyInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetCompanyInfoResponse struct {
	value *GetCompanyInfoResponse
	isSet bool
}

func (v NullableGetCompanyInfoResponse) Get() *GetCompanyInfoResponse {
	return v.value
}

func (v *NullableGetCompanyInfoResponse) Set(val *GetCompanyInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCompanyInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCompanyInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCompanyInfoResponse(val *GetCompanyInfoResponse) *NullableGetCompanyInfoResponse {
	return &NullableGetCompanyInfoResponse{value: val, isSet: true}
}

func (v NullableGetCompanyInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCompanyInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


