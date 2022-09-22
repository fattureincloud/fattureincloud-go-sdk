/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.20
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// GetArchiveDocumentResponse 
type GetArchiveDocumentResponse struct {
	Data *ArchiveDocument `json:"data,omitempty"`
}

// NewGetArchiveDocumentResponse instantiates a new GetArchiveDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetArchiveDocumentResponse() *GetArchiveDocumentResponse {
	this := GetArchiveDocumentResponse{}
	return &this
}

// NewGetArchiveDocumentResponseWithDefaults instantiates a new GetArchiveDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetArchiveDocumentResponseWithDefaults() *GetArchiveDocumentResponse {
	this := GetArchiveDocumentResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetArchiveDocumentResponse) GetData() ArchiveDocument {
	if o == nil || o.Data == nil {
		var ret ArchiveDocument
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetArchiveDocumentResponse) GetDataOk() (*ArchiveDocument, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetArchiveDocumentResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ArchiveDocument and assigns it to the Data field.
func (o *GetArchiveDocumentResponse) SetData(v ArchiveDocument) *GetArchiveDocumentResponse {
	o.Data = &v
	return o
}

func (o GetArchiveDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetArchiveDocumentResponse struct {
	value *GetArchiveDocumentResponse
	isSet bool
}

func (v NullableGetArchiveDocumentResponse) Get() *GetArchiveDocumentResponse {
	return v.value
}

func (v *NullableGetArchiveDocumentResponse) Set(val *GetArchiveDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetArchiveDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetArchiveDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetArchiveDocumentResponse(val *GetArchiveDocumentResponse) *NullableGetArchiveDocumentResponse {
	return &NullableGetArchiveDocumentResponse{value: val, isSet: true}
}

func (v NullableGetArchiveDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetArchiveDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


