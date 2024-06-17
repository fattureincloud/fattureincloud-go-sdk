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

// checks if the GetExistingIssuedDocumentTotalsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetExistingIssuedDocumentTotalsResponse{}

// GetExistingIssuedDocumentTotalsResponse 
type GetExistingIssuedDocumentTotalsResponse struct {
	Data *IssuedDocumentTotals `json:"data,omitempty"`
}

// NewGetExistingIssuedDocumentTotalsResponse instantiates a new GetExistingIssuedDocumentTotalsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExistingIssuedDocumentTotalsResponse() *GetExistingIssuedDocumentTotalsResponse {
	this := GetExistingIssuedDocumentTotalsResponse{}
	return &this
}

// NewGetExistingIssuedDocumentTotalsResponseWithDefaults instantiates a new GetExistingIssuedDocumentTotalsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExistingIssuedDocumentTotalsResponseWithDefaults() *GetExistingIssuedDocumentTotalsResponse {
	this := GetExistingIssuedDocumentTotalsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetExistingIssuedDocumentTotalsResponse) GetData() IssuedDocumentTotals {
	if o == nil || IsNil(o.Data) {
		var ret IssuedDocumentTotals
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExistingIssuedDocumentTotalsResponse) GetDataOk() (*IssuedDocumentTotals, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetExistingIssuedDocumentTotalsResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given IssuedDocumentTotals and assigns it to the Data field.
func (o *GetExistingIssuedDocumentTotalsResponse) SetData(v IssuedDocumentTotals) *GetExistingIssuedDocumentTotalsResponse {
	o.Data = &v
	return o
}

func (o GetExistingIssuedDocumentTotalsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExistingIssuedDocumentTotalsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetExistingIssuedDocumentTotalsResponse struct {
	value *GetExistingIssuedDocumentTotalsResponse
	isSet bool
}

func (v NullableGetExistingIssuedDocumentTotalsResponse) Get() *GetExistingIssuedDocumentTotalsResponse {
	return v.value
}

func (v *NullableGetExistingIssuedDocumentTotalsResponse) Set(val *GetExistingIssuedDocumentTotalsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExistingIssuedDocumentTotalsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExistingIssuedDocumentTotalsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExistingIssuedDocumentTotalsResponse(val *GetExistingIssuedDocumentTotalsResponse) *NullableGetExistingIssuedDocumentTotalsResponse {
	return &NullableGetExistingIssuedDocumentTotalsResponse{value: val, isSet: true}
}

func (v NullableGetExistingIssuedDocumentTotalsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExistingIssuedDocumentTotalsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


