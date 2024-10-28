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

// checks if the GetTaxProfileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTaxProfileResponse{}

// GetTaxProfileResponse struct for GetTaxProfileResponse
type GetTaxProfileResponse struct {
Data *TaxProfile `json:"data,omitempty"`
}

// NewGetTaxProfileResponse instantiates a new GetTaxProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTaxProfileResponse() *GetTaxProfileResponse {
	this := GetTaxProfileResponse{}
	return &this
}

// NewGetTaxProfileResponseWithDefaults instantiates a new GetTaxProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTaxProfileResponseWithDefaults() *GetTaxProfileResponse {
	this := GetTaxProfileResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetTaxProfileResponse) GetData() TaxProfile {
	if o == nil || IsNil(o.Data) {
		var ret TaxProfile
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTaxProfileResponse) GetDataOk() (*TaxProfile, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetTaxProfileResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given TaxProfile and assigns it to the Data field.
func (o *GetTaxProfileResponse) SetData(v TaxProfile) *GetTaxProfileResponse {
	o.Data = &v
        return o
}

func (o GetTaxProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTaxProfileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetTaxProfileResponse struct {
	value *GetTaxProfileResponse
	isSet bool
}

func (v NullableGetTaxProfileResponse) Get() *GetTaxProfileResponse {
	return v.value
}

func (v *NullableGetTaxProfileResponse) Set(val *GetTaxProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTaxProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTaxProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTaxProfileResponse(val *GetTaxProfileResponse) *NullableGetTaxProfileResponse {
	return &NullableGetTaxProfileResponse{value: val, isSet: true}
}

func (v NullableGetTaxProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTaxProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


