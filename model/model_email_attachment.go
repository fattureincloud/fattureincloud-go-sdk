/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.28
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the EmailAttachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailAttachment{}

// EmailAttachment struct for EmailAttachment
type EmailAttachment struct {
	// Email attachment filename.
	Filename *string `json:"filename,omitempty"`
	// Email attachment url.
	Url *string `json:"url,omitempty"`
}

// NewEmailAttachment instantiates a new EmailAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailAttachment() *EmailAttachment {
	this := EmailAttachment{}
	return &this
}

// NewEmailAttachmentWithDefaults instantiates a new EmailAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailAttachmentWithDefaults() *EmailAttachment {
	this := EmailAttachment{}
	return &this
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *EmailAttachment) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAttachment) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *EmailAttachment) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *EmailAttachment) SetFilename(v string) *EmailAttachment {
	o.Filename = &v
	return o
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *EmailAttachment) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAttachment) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *EmailAttachment) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *EmailAttachment) SetUrl(v string) *EmailAttachment {
	o.Url = &v
	return o
}

func (o EmailAttachment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailAttachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableEmailAttachment struct {
	value *EmailAttachment
	isSet bool
}

func (v NullableEmailAttachment) Get() *EmailAttachment {
	return v.value
}

func (v *NullableEmailAttachment) Set(val *EmailAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailAttachment(val *EmailAttachment) *NullableEmailAttachment {
	return &NullableEmailAttachment{value: val, isSet: true}
}

func (v NullableEmailAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


