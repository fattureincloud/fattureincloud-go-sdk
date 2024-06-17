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

// checks if the GetEmailDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEmailDataResponse{}

// GetEmailDataResponse struct for GetEmailDataResponse
type GetEmailDataResponse struct {
	Data *EmailData `json:"data,omitempty"`
}

// NewGetEmailDataResponse instantiates a new GetEmailDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEmailDataResponse() *GetEmailDataResponse {
	this := GetEmailDataResponse{}
	return &this
}

// NewGetEmailDataResponseWithDefaults instantiates a new GetEmailDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEmailDataResponseWithDefaults() *GetEmailDataResponse {
	this := GetEmailDataResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetEmailDataResponse) GetData() EmailData {
	if o == nil || IsNil(o.Data) {
		var ret EmailData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEmailDataResponse) GetDataOk() (*EmailData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetEmailDataResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given EmailData and assigns it to the Data field.
func (o *GetEmailDataResponse) SetData(v EmailData) *GetEmailDataResponse {
	o.Data = &v
	return o
}

func (o GetEmailDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEmailDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetEmailDataResponse struct {
	value *GetEmailDataResponse
	isSet bool
}

func (v NullableGetEmailDataResponse) Get() *GetEmailDataResponse {
	return v.value
}

func (v *NullableGetEmailDataResponse) Set(val *GetEmailDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEmailDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEmailDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEmailDataResponse(val *GetEmailDataResponse) *NullableGetEmailDataResponse {
	return &NullableGetEmailDataResponse{value: val, isSet: true}
}

func (v NullableGetEmailDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEmailDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


