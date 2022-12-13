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

// TransformIssuedDocumentResponse struct for TransformIssuedDocumentResponse
type TransformIssuedDocumentResponse struct {
	Data *IssuedDocument `json:"data,omitempty"`
	Options *IssuedDocumentOptions `json:"options,omitempty"`
}

// NewTransformIssuedDocumentResponse instantiates a new TransformIssuedDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformIssuedDocumentResponse() *TransformIssuedDocumentResponse {
	this := TransformIssuedDocumentResponse{}
	return &this
}

// NewTransformIssuedDocumentResponseWithDefaults instantiates a new TransformIssuedDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformIssuedDocumentResponseWithDefaults() *TransformIssuedDocumentResponse {
	this := TransformIssuedDocumentResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TransformIssuedDocumentResponse) GetData() IssuedDocument {
	if o == nil || isNil(o.Data) {
		var ret IssuedDocument
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformIssuedDocumentResponse) GetDataOk() (*IssuedDocument, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TransformIssuedDocumentResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given IssuedDocument and assigns it to the Data field.
func (o *TransformIssuedDocumentResponse) SetData(v IssuedDocument) *TransformIssuedDocumentResponse {
	o.Data = &v
	return o
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *TransformIssuedDocumentResponse) GetOptions() IssuedDocumentOptions {
	if o == nil || isNil(o.Options) {
		var ret IssuedDocumentOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformIssuedDocumentResponse) GetOptionsOk() (*IssuedDocumentOptions, bool) {
	if o == nil || isNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *TransformIssuedDocumentResponse) HasOptions() bool {
	if o != nil && !isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given IssuedDocumentOptions and assigns it to the Options field.
func (o *TransformIssuedDocumentResponse) SetOptions(v IssuedDocumentOptions) *TransformIssuedDocumentResponse {
	o.Options = &v
	return o
}

func (o TransformIssuedDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableTransformIssuedDocumentResponse struct {
	value *TransformIssuedDocumentResponse
	isSet bool
}

func (v NullableTransformIssuedDocumentResponse) Get() *TransformIssuedDocumentResponse {
	return v.value
}

func (v *NullableTransformIssuedDocumentResponse) Set(val *TransformIssuedDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformIssuedDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformIssuedDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformIssuedDocumentResponse(val *TransformIssuedDocumentResponse) *NullableTransformIssuedDocumentResponse {
	return &NullableTransformIssuedDocumentResponse{value: val, isSet: true}
}

func (v NullableTransformIssuedDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformIssuedDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


