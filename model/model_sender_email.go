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

// checks if the SenderEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SenderEmail{}

// SenderEmail struct for SenderEmail
type SenderEmail struct {
	// Id
	Id NullableInt32 `json:"id,omitempty"`
	// Email address
	Email NullableString `json:"email,omitempty"`
}

// NewSenderEmail instantiates a new SenderEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSenderEmail() *SenderEmail {
	this := SenderEmail{}
	return &this
}

// NewSenderEmailWithDefaults instantiates a new SenderEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSenderEmailWithDefaults() *SenderEmail {
	this := SenderEmail{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SenderEmail) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SenderEmail) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *SenderEmail) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *SenderEmail) SetId(v int32) *SenderEmail {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *SenderEmail) SetIdNil() *SenderEmail {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *SenderEmail) UnsetId() {
	o.Id.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SenderEmail) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SenderEmail) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *SenderEmail) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *SenderEmail) SetEmail(v string) *SenderEmail {
	o.Email.Set(&v)
	return o
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *SenderEmail) SetEmailNil() *SenderEmail {
	o.Email.Set(nil)
	return o
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *SenderEmail) UnsetEmail() {
	o.Email.Unset()
}

func (o SenderEmail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SenderEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	return toSerialize, nil
}

type NullableSenderEmail struct {
	value *SenderEmail
	isSet bool
}

func (v NullableSenderEmail) Get() *SenderEmail {
	return v.value
}

func (v *NullableSenderEmail) Set(val *SenderEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableSenderEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableSenderEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSenderEmail(val *SenderEmail) *NullableSenderEmail {
	return &NullableSenderEmail{value: val, isSet: true}
}

func (v NullableSenderEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSenderEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


