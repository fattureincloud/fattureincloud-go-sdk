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

// checks if the EmailDataDefaultSenderEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailDataDefaultSenderEmail{}

// EmailDataDefaultSenderEmail Default sender email. (Other emails can be found in **sender_emails_list**)
type EmailDataDefaultSenderEmail struct {
	// Default sender email id
	Id NullableInt32 `json:"id,omitempty"`
	// Default sender email address
	Email NullableString `json:"email,omitempty"`
}

// NewEmailDataDefaultSenderEmail instantiates a new EmailDataDefaultSenderEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDataDefaultSenderEmail() *EmailDataDefaultSenderEmail {
	this := EmailDataDefaultSenderEmail{}
	return &this
}

// NewEmailDataDefaultSenderEmailWithDefaults instantiates a new EmailDataDefaultSenderEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDataDefaultSenderEmailWithDefaults() *EmailDataDefaultSenderEmail {
	this := EmailDataDefaultSenderEmail{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailDataDefaultSenderEmail) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailDataDefaultSenderEmail) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *EmailDataDefaultSenderEmail) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *EmailDataDefaultSenderEmail) SetId(v int32) *EmailDataDefaultSenderEmail {
	o.Id.Set(&v)
		return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *EmailDataDefaultSenderEmail) SetIdNil() *EmailDataDefaultSenderEmail {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *EmailDataDefaultSenderEmail) UnsetId() {
	o.Id.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailDataDefaultSenderEmail) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailDataDefaultSenderEmail) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *EmailDataDefaultSenderEmail) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *EmailDataDefaultSenderEmail) SetEmail(v string) *EmailDataDefaultSenderEmail {
	o.Email.Set(&v)
		return o
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *EmailDataDefaultSenderEmail) SetEmailNil() *EmailDataDefaultSenderEmail {
	o.Email.Set(nil)
	return o
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *EmailDataDefaultSenderEmail) UnsetEmail() {
	o.Email.Unset()
}

func (o EmailDataDefaultSenderEmail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailDataDefaultSenderEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	return toSerialize, nil
}

type NullableEmailDataDefaultSenderEmail struct {
	value *EmailDataDefaultSenderEmail
	isSet bool
}

func (v NullableEmailDataDefaultSenderEmail) Get() *EmailDataDefaultSenderEmail {
	return v.value
}

func (v *NullableEmailDataDefaultSenderEmail) Set(val *EmailDataDefaultSenderEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDataDefaultSenderEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDataDefaultSenderEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDataDefaultSenderEmail(val *EmailDataDefaultSenderEmail) *NullableEmailDataDefaultSenderEmail {
	return &NullableEmailDataDefaultSenderEmail{value: val, isSet: true}
}

func (v NullableEmailDataDefaultSenderEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDataDefaultSenderEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


