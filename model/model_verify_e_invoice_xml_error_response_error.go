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

// checks if the VerifyEInvoiceXmlErrorResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyEInvoiceXmlErrorResponseError{}

// VerifyEInvoiceXmlErrorResponseError struct for VerifyEInvoiceXmlErrorResponseError
type VerifyEInvoiceXmlErrorResponseError struct {
Message NullableString `json:"message,omitempty"`
ValidationResult NullableVerifyEInvoiceXmlErrorResponseErrorValidationResult `json:"validation_result,omitempty"`
}

// NewVerifyEInvoiceXmlErrorResponseError instantiates a new VerifyEInvoiceXmlErrorResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyEInvoiceXmlErrorResponseError() *VerifyEInvoiceXmlErrorResponseError {
	this := VerifyEInvoiceXmlErrorResponseError{}
	return &this
}

// NewVerifyEInvoiceXmlErrorResponseErrorWithDefaults instantiates a new VerifyEInvoiceXmlErrorResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyEInvoiceXmlErrorResponseErrorWithDefaults() *VerifyEInvoiceXmlErrorResponseError {
	this := VerifyEInvoiceXmlErrorResponseError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerifyEInvoiceXmlErrorResponseError) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerifyEInvoiceXmlErrorResponseError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *VerifyEInvoiceXmlErrorResponseError) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *VerifyEInvoiceXmlErrorResponseError) SetMessage(v string) *VerifyEInvoiceXmlErrorResponseError {
	o.Message.Set(&v)
		return o
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *VerifyEInvoiceXmlErrorResponseError) SetMessageNil() *VerifyEInvoiceXmlErrorResponseError {
	o.Message.Set(nil)
	return o
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *VerifyEInvoiceXmlErrorResponseError) UnsetMessage() {
	o.Message.Unset()
}

// GetValidationResult returns the ValidationResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerifyEInvoiceXmlErrorResponseError) GetValidationResult() VerifyEInvoiceXmlErrorResponseErrorValidationResult {
	if o == nil || IsNil(o.ValidationResult.Get()) {
		var ret VerifyEInvoiceXmlErrorResponseErrorValidationResult
		return ret
	}
	return *o.ValidationResult.Get()
}

// GetValidationResultOk returns a tuple with the ValidationResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerifyEInvoiceXmlErrorResponseError) GetValidationResultOk() (*VerifyEInvoiceXmlErrorResponseErrorValidationResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationResult.Get(), o.ValidationResult.IsSet()
}

// HasValidationResult returns a boolean if a field has been set.
func (o *VerifyEInvoiceXmlErrorResponseError) HasValidationResult() bool {
	if o != nil && o.ValidationResult.IsSet() {
		return true
	}

	return false
}

// SetValidationResult gets a reference to the given NullableVerifyEInvoiceXmlErrorResponseErrorValidationResult and assigns it to the ValidationResult field.
func (o *VerifyEInvoiceXmlErrorResponseError) SetValidationResult(v VerifyEInvoiceXmlErrorResponseErrorValidationResult) *VerifyEInvoiceXmlErrorResponseError {
	o.ValidationResult.Set(&v)
		return o
}
// SetValidationResultNil sets the value for ValidationResult to be an explicit nil
func (o *VerifyEInvoiceXmlErrorResponseError) SetValidationResultNil() *VerifyEInvoiceXmlErrorResponseError {
	o.ValidationResult.Set(nil)
	return o
}

// UnsetValidationResult ensures that no value is present for ValidationResult, not even an explicit nil
func (o *VerifyEInvoiceXmlErrorResponseError) UnsetValidationResult() {
	o.ValidationResult.Unset()
}

func (o VerifyEInvoiceXmlErrorResponseError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyEInvoiceXmlErrorResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.ValidationResult.IsSet() {
		toSerialize["validation_result"] = o.ValidationResult.Get()
	}
	return toSerialize, nil
}

type NullableVerifyEInvoiceXmlErrorResponseError struct {
	value *VerifyEInvoiceXmlErrorResponseError
	isSet bool
}

func (v NullableVerifyEInvoiceXmlErrorResponseError) Get() *VerifyEInvoiceXmlErrorResponseError {
	return v.value
}

func (v *NullableVerifyEInvoiceXmlErrorResponseError) Set(val *VerifyEInvoiceXmlErrorResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyEInvoiceXmlErrorResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyEInvoiceXmlErrorResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyEInvoiceXmlErrorResponseError(val *VerifyEInvoiceXmlErrorResponseError) *NullableVerifyEInvoiceXmlErrorResponseError {
	return &NullableVerifyEInvoiceXmlErrorResponseError{value: val, isSet: true}
}

func (v NullableVerifyEInvoiceXmlErrorResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyEInvoiceXmlErrorResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


