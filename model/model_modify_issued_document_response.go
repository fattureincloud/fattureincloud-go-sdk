/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.24
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// ModifyIssuedDocumentResponse struct for ModifyIssuedDocumentResponse
type ModifyIssuedDocumentResponse struct {
	Data *IssuedDocument `json:"data,omitempty"`
}

// NewModifyIssuedDocumentResponse instantiates a new ModifyIssuedDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyIssuedDocumentResponse() *ModifyIssuedDocumentResponse {
	this := ModifyIssuedDocumentResponse{}
	return &this
}

// NewModifyIssuedDocumentResponseWithDefaults instantiates a new ModifyIssuedDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyIssuedDocumentResponseWithDefaults() *ModifyIssuedDocumentResponse {
	this := ModifyIssuedDocumentResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModifyIssuedDocumentResponse) GetData() IssuedDocument {
	if o == nil || isNil(o.Data) {
		var ret IssuedDocument
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssuedDocumentResponse) GetDataOk() (*IssuedDocument, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModifyIssuedDocumentResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given IssuedDocument and assigns it to the Data field.
func (o *ModifyIssuedDocumentResponse) SetData(v IssuedDocument) *ModifyIssuedDocumentResponse {
	o.Data = &v
	return o
}

func (o ModifyIssuedDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableModifyIssuedDocumentResponse struct {
	value *ModifyIssuedDocumentResponse
	isSet bool
}

func (v NullableModifyIssuedDocumentResponse) Get() *ModifyIssuedDocumentResponse {
	return v.value
}

func (v *NullableModifyIssuedDocumentResponse) Set(val *ModifyIssuedDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyIssuedDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyIssuedDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyIssuedDocumentResponse(val *ModifyIssuedDocumentResponse) *NullableModifyIssuedDocumentResponse {
	return &NullableModifyIssuedDocumentResponse{value: val, isSet: true}
}

func (v NullableModifyIssuedDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyIssuedDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


