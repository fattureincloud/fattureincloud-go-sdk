/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.3
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the CreateArchiveDocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateArchiveDocumentResponse{}

// CreateArchiveDocumentResponse 
type CreateArchiveDocumentResponse struct {
	Data *ArchiveDocument `json:"data,omitempty"`
}

// NewCreateArchiveDocumentResponse instantiates a new CreateArchiveDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateArchiveDocumentResponse() *CreateArchiveDocumentResponse {
	this := CreateArchiveDocumentResponse{}
	return &this
}

// NewCreateArchiveDocumentResponseWithDefaults instantiates a new CreateArchiveDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateArchiveDocumentResponseWithDefaults() *CreateArchiveDocumentResponse {
	this := CreateArchiveDocumentResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateArchiveDocumentResponse) GetData() ArchiveDocument {
	if o == nil || IsNil(o.Data) {
		var ret ArchiveDocument
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateArchiveDocumentResponse) GetDataOk() (*ArchiveDocument, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateArchiveDocumentResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ArchiveDocument and assigns it to the Data field.
func (o *CreateArchiveDocumentResponse) SetData(v ArchiveDocument) *CreateArchiveDocumentResponse {
	o.Data = &v
	return o
}

func (o CreateArchiveDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateArchiveDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateArchiveDocumentResponse struct {
	value *CreateArchiveDocumentResponse
	isSet bool
}

func (v NullableCreateArchiveDocumentResponse) Get() *CreateArchiveDocumentResponse {
	return v.value
}

func (v *NullableCreateArchiveDocumentResponse) Set(val *CreateArchiveDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateArchiveDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateArchiveDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateArchiveDocumentResponse(val *CreateArchiveDocumentResponse) *NullableCreateArchiveDocumentResponse {
	return &NullableCreateArchiveDocumentResponse{value: val, isSet: true}
}

func (v NullableCreateArchiveDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateArchiveDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


