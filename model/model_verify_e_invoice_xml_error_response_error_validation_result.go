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

// checks if the VerifyEInvoiceXmlErrorResponseErrorValidationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyEInvoiceXmlErrorResponseErrorValidationResult{}

// VerifyEInvoiceXmlErrorResponseErrorValidationResult struct for VerifyEInvoiceXmlErrorResponseErrorValidationResult
type VerifyEInvoiceXmlErrorResponseErrorValidationResult struct {
	XmlErrors []string `json:"xml_errors,omitempty"`
}

// NewVerifyEInvoiceXmlErrorResponseErrorValidationResult instantiates a new VerifyEInvoiceXmlErrorResponseErrorValidationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyEInvoiceXmlErrorResponseErrorValidationResult() *VerifyEInvoiceXmlErrorResponseErrorValidationResult {
	this := VerifyEInvoiceXmlErrorResponseErrorValidationResult{}
	return &this
}

// NewVerifyEInvoiceXmlErrorResponseErrorValidationResultWithDefaults instantiates a new VerifyEInvoiceXmlErrorResponseErrorValidationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyEInvoiceXmlErrorResponseErrorValidationResultWithDefaults() *VerifyEInvoiceXmlErrorResponseErrorValidationResult {
	this := VerifyEInvoiceXmlErrorResponseErrorValidationResult{}
	return &this
}

// GetXmlErrors returns the XmlErrors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerifyEInvoiceXmlErrorResponseErrorValidationResult) GetXmlErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.XmlErrors
}

// GetXmlErrorsOk returns a tuple with the XmlErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerifyEInvoiceXmlErrorResponseErrorValidationResult) GetXmlErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.XmlErrors) {
		return nil, false
	}
	return o.XmlErrors, true
}

// HasXmlErrors returns a boolean if a field has been set.
func (o *VerifyEInvoiceXmlErrorResponseErrorValidationResult) HasXmlErrors() bool {
	if o != nil && IsNil(o.XmlErrors) {
		return true
	}

	return false
}

// SetXmlErrors gets a reference to the given []string and assigns it to the XmlErrors field.
func (o *VerifyEInvoiceXmlErrorResponseErrorValidationResult) SetXmlErrors(v []string) *VerifyEInvoiceXmlErrorResponseErrorValidationResult {
	o.XmlErrors = v
	return o
}

func (o VerifyEInvoiceXmlErrorResponseErrorValidationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyEInvoiceXmlErrorResponseErrorValidationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.XmlErrors != nil {
		toSerialize["xml_errors"] = o.XmlErrors
	}
	return toSerialize, nil
}

type NullableVerifyEInvoiceXmlErrorResponseErrorValidationResult struct {
	value *VerifyEInvoiceXmlErrorResponseErrorValidationResult
	isSet bool
}

func (v NullableVerifyEInvoiceXmlErrorResponseErrorValidationResult) Get() *VerifyEInvoiceXmlErrorResponseErrorValidationResult {
	return v.value
}

func (v *NullableVerifyEInvoiceXmlErrorResponseErrorValidationResult) Set(val *VerifyEInvoiceXmlErrorResponseErrorValidationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyEInvoiceXmlErrorResponseErrorValidationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyEInvoiceXmlErrorResponseErrorValidationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyEInvoiceXmlErrorResponseErrorValidationResult(val *VerifyEInvoiceXmlErrorResponseErrorValidationResult) *NullableVerifyEInvoiceXmlErrorResponseErrorValidationResult {
	return &NullableVerifyEInvoiceXmlErrorResponseErrorValidationResult{value: val, isSet: true}
}

func (v NullableVerifyEInvoiceXmlErrorResponseErrorValidationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyEInvoiceXmlErrorResponseErrorValidationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


