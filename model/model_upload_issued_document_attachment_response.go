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

// checks if the UploadIssuedDocumentAttachmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadIssuedDocumentAttachmentResponse{}

// UploadIssuedDocumentAttachmentResponse struct for UploadIssuedDocumentAttachmentResponse
type UploadIssuedDocumentAttachmentResponse struct {
	Data *AttachmentData `json:"data,omitempty"`
}

// NewUploadIssuedDocumentAttachmentResponse instantiates a new UploadIssuedDocumentAttachmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadIssuedDocumentAttachmentResponse() *UploadIssuedDocumentAttachmentResponse {
	this := UploadIssuedDocumentAttachmentResponse{}
	return &this
}

// NewUploadIssuedDocumentAttachmentResponseWithDefaults instantiates a new UploadIssuedDocumentAttachmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadIssuedDocumentAttachmentResponseWithDefaults() *UploadIssuedDocumentAttachmentResponse {
	this := UploadIssuedDocumentAttachmentResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UploadIssuedDocumentAttachmentResponse) GetData() AttachmentData {
	if o == nil || IsNil(o.Data) {
		var ret AttachmentData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadIssuedDocumentAttachmentResponse) GetDataOk() (*AttachmentData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UploadIssuedDocumentAttachmentResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AttachmentData and assigns it to the Data field.
func (o *UploadIssuedDocumentAttachmentResponse) SetData(v AttachmentData) *UploadIssuedDocumentAttachmentResponse {
	o.Data = &v
	return o
}

func (o UploadIssuedDocumentAttachmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadIssuedDocumentAttachmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUploadIssuedDocumentAttachmentResponse struct {
	value *UploadIssuedDocumentAttachmentResponse
	isSet bool
}

func (v NullableUploadIssuedDocumentAttachmentResponse) Get() *UploadIssuedDocumentAttachmentResponse {
	return v.value
}

func (v *NullableUploadIssuedDocumentAttachmentResponse) Set(val *UploadIssuedDocumentAttachmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadIssuedDocumentAttachmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadIssuedDocumentAttachmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadIssuedDocumentAttachmentResponse(val *UploadIssuedDocumentAttachmentResponse) *NullableUploadIssuedDocumentAttachmentResponse {
	return &NullableUploadIssuedDocumentAttachmentResponse{value: val, isSet: true}
}

func (v NullableUploadIssuedDocumentAttachmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadIssuedDocumentAttachmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


