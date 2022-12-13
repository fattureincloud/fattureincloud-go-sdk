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

// ModifyArchiveDocumentResponse 
type ModifyArchiveDocumentResponse struct {
	Data *ArchiveDocument `json:"data,omitempty"`
}

// NewModifyArchiveDocumentResponse instantiates a new ModifyArchiveDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyArchiveDocumentResponse() *ModifyArchiveDocumentResponse {
	this := ModifyArchiveDocumentResponse{}
	return &this
}

// NewModifyArchiveDocumentResponseWithDefaults instantiates a new ModifyArchiveDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyArchiveDocumentResponseWithDefaults() *ModifyArchiveDocumentResponse {
	this := ModifyArchiveDocumentResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModifyArchiveDocumentResponse) GetData() ArchiveDocument {
	if o == nil || isNil(o.Data) {
		var ret ArchiveDocument
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyArchiveDocumentResponse) GetDataOk() (*ArchiveDocument, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModifyArchiveDocumentResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ArchiveDocument and assigns it to the Data field.
func (o *ModifyArchiveDocumentResponse) SetData(v ArchiveDocument) *ModifyArchiveDocumentResponse {
	o.Data = &v
	return o
}

func (o ModifyArchiveDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableModifyArchiveDocumentResponse struct {
	value *ModifyArchiveDocumentResponse
	isSet bool
}

func (v NullableModifyArchiveDocumentResponse) Get() *ModifyArchiveDocumentResponse {
	return v.value
}

func (v *NullableModifyArchiveDocumentResponse) Set(val *ModifyArchiveDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyArchiveDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyArchiveDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyArchiveDocumentResponse(val *ModifyArchiveDocumentResponse) *NullableModifyArchiveDocumentResponse {
	return &NullableModifyArchiveDocumentResponse{value: val, isSet: true}
}

func (v NullableModifyArchiveDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyArchiveDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


