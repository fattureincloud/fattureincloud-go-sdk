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

// checks if the CreateIssuedDocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateIssuedDocumentResponse{}

// CreateIssuedDocumentResponse struct for CreateIssuedDocumentResponse
type CreateIssuedDocumentResponse struct {
	Data *IssuedDocument `json:"data,omitempty"`
}

// NewCreateIssuedDocumentResponse instantiates a new CreateIssuedDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIssuedDocumentResponse() *CreateIssuedDocumentResponse {
	this := CreateIssuedDocumentResponse{}
	return &this
}

// NewCreateIssuedDocumentResponseWithDefaults instantiates a new CreateIssuedDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIssuedDocumentResponseWithDefaults() *CreateIssuedDocumentResponse {
	this := CreateIssuedDocumentResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateIssuedDocumentResponse) GetData() IssuedDocument {
	if o == nil || IsNil(o.Data) {
		var ret IssuedDocument
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssuedDocumentResponse) GetDataOk() (*IssuedDocument, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateIssuedDocumentResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given IssuedDocument and assigns it to the Data field.
func (o *CreateIssuedDocumentResponse) SetData(v IssuedDocument) *CreateIssuedDocumentResponse {
	o.Data = &v
		return o
}

func (o CreateIssuedDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIssuedDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateIssuedDocumentResponse struct {
	value *CreateIssuedDocumentResponse
	isSet bool
}

func (v NullableCreateIssuedDocumentResponse) Get() *CreateIssuedDocumentResponse {
	return v.value
}

func (v *NullableCreateIssuedDocumentResponse) Set(val *CreateIssuedDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIssuedDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIssuedDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIssuedDocumentResponse(val *CreateIssuedDocumentResponse) *NullableCreateIssuedDocumentResponse {
	return &NullableCreateIssuedDocumentResponse{value: val, isSet: true}
}

func (v NullableCreateIssuedDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIssuedDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


