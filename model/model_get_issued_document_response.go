/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.31
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the GetIssuedDocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIssuedDocumentResponse{}

// GetIssuedDocumentResponse struct for GetIssuedDocumentResponse
type GetIssuedDocumentResponse struct {
	Data *IssuedDocument `json:"data,omitempty"`
}

// NewGetIssuedDocumentResponse instantiates a new GetIssuedDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIssuedDocumentResponse() *GetIssuedDocumentResponse {
	this := GetIssuedDocumentResponse{}
	return &this
}

// NewGetIssuedDocumentResponseWithDefaults instantiates a new GetIssuedDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIssuedDocumentResponseWithDefaults() *GetIssuedDocumentResponse {
	this := GetIssuedDocumentResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetIssuedDocumentResponse) GetData() IssuedDocument {
	if o == nil || IsNil(o.Data) {
		var ret IssuedDocument
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssuedDocumentResponse) GetDataOk() (*IssuedDocument, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetIssuedDocumentResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given IssuedDocument and assigns it to the Data field.
func (o *GetIssuedDocumentResponse) SetData(v IssuedDocument) *GetIssuedDocumentResponse {
	o.Data = &v
	return o
}

func (o GetIssuedDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIssuedDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetIssuedDocumentResponse struct {
	value *GetIssuedDocumentResponse
	isSet bool
}

func (v NullableGetIssuedDocumentResponse) Get() *GetIssuedDocumentResponse {
	return v.value
}

func (v *NullableGetIssuedDocumentResponse) Set(val *GetIssuedDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIssuedDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIssuedDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIssuedDocumentResponse(val *GetIssuedDocumentResponse) *NullableGetIssuedDocumentResponse {
	return &NullableGetIssuedDocumentResponse{value: val, isSet: true}
}

func (v NullableGetIssuedDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIssuedDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


