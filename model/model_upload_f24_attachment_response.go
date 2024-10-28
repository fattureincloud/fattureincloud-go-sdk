/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.2
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the UploadF24AttachmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadF24AttachmentResponse{}

// UploadF24AttachmentResponse struct for UploadF24AttachmentResponse
type UploadF24AttachmentResponse struct {
Data *AttachmentData `json:"data,omitempty"`
}

// NewUploadF24AttachmentResponse instantiates a new UploadF24AttachmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadF24AttachmentResponse() *UploadF24AttachmentResponse {
	this := UploadF24AttachmentResponse{}
	return &this
}

// NewUploadF24AttachmentResponseWithDefaults instantiates a new UploadF24AttachmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadF24AttachmentResponseWithDefaults() *UploadF24AttachmentResponse {
	this := UploadF24AttachmentResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UploadF24AttachmentResponse) GetData() AttachmentData {
	if o == nil || IsNil(o.Data) {
		var ret AttachmentData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadF24AttachmentResponse) GetDataOk() (*AttachmentData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UploadF24AttachmentResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AttachmentData and assigns it to the Data field.
func (o *UploadF24AttachmentResponse) SetData(v AttachmentData) *UploadF24AttachmentResponse {
	o.Data = &v
        return o
}

func (o UploadF24AttachmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadF24AttachmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUploadF24AttachmentResponse struct {
	value *UploadF24AttachmentResponse
	isSet bool
}

func (v NullableUploadF24AttachmentResponse) Get() *UploadF24AttachmentResponse {
	return v.value
}

func (v *NullableUploadF24AttachmentResponse) Set(val *UploadF24AttachmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadF24AttachmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadF24AttachmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadF24AttachmentResponse(val *UploadF24AttachmentResponse) *NullableUploadF24AttachmentResponse {
	return &NullableUploadF24AttachmentResponse{value: val, isSet: true}
}

func (v NullableUploadF24AttachmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadF24AttachmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


