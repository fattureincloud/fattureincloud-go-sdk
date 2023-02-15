/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.26
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ModifyIssuedDocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyIssuedDocumentRequest{}

// ModifyIssuedDocumentRequest 
type ModifyIssuedDocumentRequest struct {
	Data *IssuedDocument `json:"data,omitempty"`
	Options *IssuedDocumentOptions `json:"options,omitempty"`
}

// NewModifyIssuedDocumentRequest instantiates a new ModifyIssuedDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyIssuedDocumentRequest() *ModifyIssuedDocumentRequest {
	this := ModifyIssuedDocumentRequest{}
	return &this
}

// NewModifyIssuedDocumentRequestWithDefaults instantiates a new ModifyIssuedDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyIssuedDocumentRequestWithDefaults() *ModifyIssuedDocumentRequest {
	this := ModifyIssuedDocumentRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModifyIssuedDocumentRequest) GetData() IssuedDocument {
	if o == nil || isNil(o.Data) {
		var ret IssuedDocument
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssuedDocumentRequest) GetDataOk() (*IssuedDocument, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModifyIssuedDocumentRequest) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given IssuedDocument and assigns it to the Data field.
func (o *ModifyIssuedDocumentRequest) SetData(v IssuedDocument) *ModifyIssuedDocumentRequest {
	o.Data = &v
	return o
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ModifyIssuedDocumentRequest) GetOptions() IssuedDocumentOptions {
	if o == nil || isNil(o.Options) {
		var ret IssuedDocumentOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssuedDocumentRequest) GetOptionsOk() (*IssuedDocumentOptions, bool) {
	if o == nil || isNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ModifyIssuedDocumentRequest) HasOptions() bool {
	if o != nil && !isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given IssuedDocumentOptions and assigns it to the Options field.
func (o *ModifyIssuedDocumentRequest) SetOptions(v IssuedDocumentOptions) *ModifyIssuedDocumentRequest {
	o.Options = &v
	return o
}

func (o ModifyIssuedDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyIssuedDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableModifyIssuedDocumentRequest struct {
	value *ModifyIssuedDocumentRequest
	isSet bool
}

func (v NullableModifyIssuedDocumentRequest) Get() *ModifyIssuedDocumentRequest {
	return v.value
}

func (v *NullableModifyIssuedDocumentRequest) Set(val *ModifyIssuedDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyIssuedDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyIssuedDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyIssuedDocumentRequest(val *ModifyIssuedDocumentRequest) *NullableModifyIssuedDocumentRequest {
	return &NullableModifyIssuedDocumentRequest{value: val, isSet: true}
}

func (v NullableModifyIssuedDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyIssuedDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


