/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.27
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the Language type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Language{}

// Language struct for Language
type Language struct {
	// Language code.
	Code NullableString `json:"code,omitempty"`
	// Language extended name.
	Name NullableString `json:"name,omitempty"`
}

// NewLanguage instantiates a new Language object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguage() *Language {
	this := Language{}
	return &this
}

// NewLanguageWithDefaults instantiates a new Language object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguageWithDefaults() *Language {
	this := Language{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Language) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Language) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *Language) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *Language) SetCode(v string) *Language {
	o.Code.Set(&v)
	return o
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *Language) SetCodeNil() *Language {
	o.Code.Set(nil)
	return o
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *Language) UnsetCode() {
	o.Code.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Language) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Language) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Language) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Language) SetName(v string) *Language {
	o.Name.Set(&v)
	return o
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Language) SetNameNil() *Language {
	o.Name.Set(nil)
	return o
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Language) UnsetName() {
	o.Name.Unset()
}

func (o Language) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Language) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableLanguage struct {
	value *Language
	isSet bool
}

func (v NullableLanguage) Get() *Language {
	return v.value
}

func (v *NullableLanguage) Set(val *Language) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguage(val *Language) *NullableLanguage {
	return &NullableLanguage{value: val, isSet: true}
}

func (v NullableLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


