/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.30
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the GetIssuedDocumentPreCreateInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIssuedDocumentPreCreateInfoResponse{}

// GetIssuedDocumentPreCreateInfoResponse struct for GetIssuedDocumentPreCreateInfoResponse
type GetIssuedDocumentPreCreateInfoResponse struct {
	Data *IssuedDocumentPreCreateInfo `json:"data,omitempty"`
}

// NewGetIssuedDocumentPreCreateInfoResponse instantiates a new GetIssuedDocumentPreCreateInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIssuedDocumentPreCreateInfoResponse() *GetIssuedDocumentPreCreateInfoResponse {
	this := GetIssuedDocumentPreCreateInfoResponse{}
	return &this
}

// NewGetIssuedDocumentPreCreateInfoResponseWithDefaults instantiates a new GetIssuedDocumentPreCreateInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIssuedDocumentPreCreateInfoResponseWithDefaults() *GetIssuedDocumentPreCreateInfoResponse {
	this := GetIssuedDocumentPreCreateInfoResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetIssuedDocumentPreCreateInfoResponse) GetData() IssuedDocumentPreCreateInfo {
	if o == nil || IsNil(o.Data) {
		var ret IssuedDocumentPreCreateInfo
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssuedDocumentPreCreateInfoResponse) GetDataOk() (*IssuedDocumentPreCreateInfo, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetIssuedDocumentPreCreateInfoResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given IssuedDocumentPreCreateInfo and assigns it to the Data field.
func (o *GetIssuedDocumentPreCreateInfoResponse) SetData(v IssuedDocumentPreCreateInfo) *GetIssuedDocumentPreCreateInfoResponse {
	o.Data = &v
	return o
}

func (o GetIssuedDocumentPreCreateInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIssuedDocumentPreCreateInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetIssuedDocumentPreCreateInfoResponse struct {
	value *GetIssuedDocumentPreCreateInfoResponse
	isSet bool
}

func (v NullableGetIssuedDocumentPreCreateInfoResponse) Get() *GetIssuedDocumentPreCreateInfoResponse {
	return v.value
}

func (v *NullableGetIssuedDocumentPreCreateInfoResponse) Set(val *GetIssuedDocumentPreCreateInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIssuedDocumentPreCreateInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIssuedDocumentPreCreateInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIssuedDocumentPreCreateInfoResponse(val *GetIssuedDocumentPreCreateInfoResponse) *NullableGetIssuedDocumentPreCreateInfoResponse {
	return &NullableGetIssuedDocumentPreCreateInfoResponse{value: val, isSet: true}
}

func (v NullableGetIssuedDocumentPreCreateInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIssuedDocumentPreCreateInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


