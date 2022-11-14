/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.21
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// JoinIssuedDocumentsResponse struct for JoinIssuedDocumentsResponse
type JoinIssuedDocumentsResponse struct {
	Data *IssuedDocument `json:"data,omitempty"`
	Options *IssuedDocumentOptions `json:"options,omitempty"`
}

// NewJoinIssuedDocumentsResponse instantiates a new JoinIssuedDocumentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJoinIssuedDocumentsResponse() *JoinIssuedDocumentsResponse {
	this := JoinIssuedDocumentsResponse{}
	return &this
}

// NewJoinIssuedDocumentsResponseWithDefaults instantiates a new JoinIssuedDocumentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJoinIssuedDocumentsResponseWithDefaults() *JoinIssuedDocumentsResponse {
	this := JoinIssuedDocumentsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *JoinIssuedDocumentsResponse) GetData() IssuedDocument {
	if o == nil || o.Data == nil {
		var ret IssuedDocument
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoinIssuedDocumentsResponse) GetDataOk() (*IssuedDocument, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *JoinIssuedDocumentsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given IssuedDocument and assigns it to the Data field.
func (o *JoinIssuedDocumentsResponse) SetData(v IssuedDocument) *JoinIssuedDocumentsResponse {
	o.Data = &v
	return o
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *JoinIssuedDocumentsResponse) GetOptions() IssuedDocumentOptions {
	if o == nil || o.Options == nil {
		var ret IssuedDocumentOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoinIssuedDocumentsResponse) GetOptionsOk() (*IssuedDocumentOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *JoinIssuedDocumentsResponse) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given IssuedDocumentOptions and assigns it to the Options field.
func (o *JoinIssuedDocumentsResponse) SetOptions(v IssuedDocumentOptions) *JoinIssuedDocumentsResponse {
	o.Options = &v
	return o
}

func (o JoinIssuedDocumentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableJoinIssuedDocumentsResponse struct {
	value *JoinIssuedDocumentsResponse
	isSet bool
}

func (v NullableJoinIssuedDocumentsResponse) Get() *JoinIssuedDocumentsResponse {
	return v.value
}

func (v *NullableJoinIssuedDocumentsResponse) Set(val *JoinIssuedDocumentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJoinIssuedDocumentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJoinIssuedDocumentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJoinIssuedDocumentsResponse(val *JoinIssuedDocumentsResponse) *NullableJoinIssuedDocumentsResponse {
	return &NullableJoinIssuedDocumentsResponse{value: val, isSet: true}
}

func (v NullableJoinIssuedDocumentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJoinIssuedDocumentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


