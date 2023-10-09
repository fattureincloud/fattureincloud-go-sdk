/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.30
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the AttachmentData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentData{}

// AttachmentData struct for AttachmentData
type AttachmentData struct {
	// Uploaded attachment token.
	AttachmentToken NullableString `json:"attachment_token,omitempty"`
}

// NewAttachmentData instantiates a new AttachmentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentData() *AttachmentData {
	this := AttachmentData{}
	return &this
}

// NewAttachmentDataWithDefaults instantiates a new AttachmentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentDataWithDefaults() *AttachmentData {
	this := AttachmentData{}
	return &this
}

// GetAttachmentToken returns the AttachmentToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentData) GetAttachmentToken() string {
	if o == nil || IsNil(o.AttachmentToken.Get()) {
		var ret string
		return ret
	}
	return *o.AttachmentToken.Get()
}

// GetAttachmentTokenOk returns a tuple with the AttachmentToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentData) GetAttachmentTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachmentToken.Get(), o.AttachmentToken.IsSet()
}

// HasAttachmentToken returns a boolean if a field has been set.
func (o *AttachmentData) HasAttachmentToken() bool {
	if o != nil && o.AttachmentToken.IsSet() {
		return true
	}

	return false
}

// SetAttachmentToken gets a reference to the given NullableString and assigns it to the AttachmentToken field.
func (o *AttachmentData) SetAttachmentToken(v string) *AttachmentData {
	o.AttachmentToken.Set(&v)
	return o
}
// SetAttachmentTokenNil sets the value for AttachmentToken to be an explicit nil
func (o *AttachmentData) SetAttachmentTokenNil() *AttachmentData {
	o.AttachmentToken.Set(nil)
	return o
}

// UnsetAttachmentToken ensures that no value is present for AttachmentToken, not even an explicit nil
func (o *AttachmentData) UnsetAttachmentToken() {
	o.AttachmentToken.Unset()
}

func (o AttachmentData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AttachmentToken.IsSet() {
		toSerialize["attachment_token"] = o.AttachmentToken.Get()
	}
	return toSerialize, nil
}

type NullableAttachmentData struct {
	value *AttachmentData
	isSet bool
}

func (v NullableAttachmentData) Get() *AttachmentData {
	return v.value
}

func (v *NullableAttachmentData) Set(val *AttachmentData) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentData) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentData(val *AttachmentData) *NullableAttachmentData {
	return &NullableAttachmentData{value: val, isSet: true}
}

func (v NullableAttachmentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


