/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.25
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// UploadArchiveAttachmentResponse struct for UploadArchiveAttachmentResponse
type UploadArchiveAttachmentResponse struct {
	Data *AttachmentData `json:"data,omitempty"`
}

// NewUploadArchiveAttachmentResponse instantiates a new UploadArchiveAttachmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadArchiveAttachmentResponse() *UploadArchiveAttachmentResponse {
	this := UploadArchiveAttachmentResponse{}
	return &this
}

// NewUploadArchiveAttachmentResponseWithDefaults instantiates a new UploadArchiveAttachmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadArchiveAttachmentResponseWithDefaults() *UploadArchiveAttachmentResponse {
	this := UploadArchiveAttachmentResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UploadArchiveAttachmentResponse) GetData() AttachmentData {
	if o == nil || isNil(o.Data) {
		var ret AttachmentData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadArchiveAttachmentResponse) GetDataOk() (*AttachmentData, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UploadArchiveAttachmentResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AttachmentData and assigns it to the Data field.
func (o *UploadArchiveAttachmentResponse) SetData(v AttachmentData) *UploadArchiveAttachmentResponse {
	o.Data = &v
	return o
}

func (o UploadArchiveAttachmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableUploadArchiveAttachmentResponse struct {
	value *UploadArchiveAttachmentResponse
	isSet bool
}

func (v NullableUploadArchiveAttachmentResponse) Get() *UploadArchiveAttachmentResponse {
	return v.value
}

func (v *NullableUploadArchiveAttachmentResponse) Set(val *UploadArchiveAttachmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadArchiveAttachmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadArchiveAttachmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadArchiveAttachmentResponse(val *UploadArchiveAttachmentResponse) *NullableUploadArchiveAttachmentResponse {
	return &NullableUploadArchiveAttachmentResponse{value: val, isSet: true}
}

func (v NullableUploadArchiveAttachmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadArchiveAttachmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


