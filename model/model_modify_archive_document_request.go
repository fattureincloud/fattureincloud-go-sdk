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

// checks if the ModifyArchiveDocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyArchiveDocumentRequest{}

// ModifyArchiveDocumentRequest 
type ModifyArchiveDocumentRequest struct {
	Data *ArchiveDocument `json:"data,omitempty"`
}

// NewModifyArchiveDocumentRequest instantiates a new ModifyArchiveDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyArchiveDocumentRequest() *ModifyArchiveDocumentRequest {
	this := ModifyArchiveDocumentRequest{}
	return &this
}

// NewModifyArchiveDocumentRequestWithDefaults instantiates a new ModifyArchiveDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyArchiveDocumentRequestWithDefaults() *ModifyArchiveDocumentRequest {
	this := ModifyArchiveDocumentRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModifyArchiveDocumentRequest) GetData() ArchiveDocument {
	if o == nil || IsNil(o.Data) {
		var ret ArchiveDocument
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyArchiveDocumentRequest) GetDataOk() (*ArchiveDocument, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModifyArchiveDocumentRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ArchiveDocument and assigns it to the Data field.
func (o *ModifyArchiveDocumentRequest) SetData(v ArchiveDocument) *ModifyArchiveDocumentRequest {
	o.Data = &v
	return o
}

func (o ModifyArchiveDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyArchiveDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableModifyArchiveDocumentRequest struct {
	value *ModifyArchiveDocumentRequest
	isSet bool
}

func (v NullableModifyArchiveDocumentRequest) Get() *ModifyArchiveDocumentRequest {
	return v.value
}

func (v *NullableModifyArchiveDocumentRequest) Set(val *ModifyArchiveDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyArchiveDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyArchiveDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyArchiveDocumentRequest(val *ModifyArchiveDocumentRequest) *NullableModifyArchiveDocumentRequest {
	return &NullableModifyArchiveDocumentRequest{value: val, isSet: true}
}

func (v NullableModifyArchiveDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyArchiveDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


