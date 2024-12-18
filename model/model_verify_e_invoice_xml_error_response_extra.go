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

// checks if the VerifyEInvoiceXmlErrorResponseExtra type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyEInvoiceXmlErrorResponseExtra{}

// VerifyEInvoiceXmlErrorResponseExtra struct for VerifyEInvoiceXmlErrorResponseExtra
type VerifyEInvoiceXmlErrorResponseExtra struct {
	Errors []string `json:"errors,omitempty"`
}

// NewVerifyEInvoiceXmlErrorResponseExtra instantiates a new VerifyEInvoiceXmlErrorResponseExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyEInvoiceXmlErrorResponseExtra() *VerifyEInvoiceXmlErrorResponseExtra {
	this := VerifyEInvoiceXmlErrorResponseExtra{}
	return &this
}

// NewVerifyEInvoiceXmlErrorResponseExtraWithDefaults instantiates a new VerifyEInvoiceXmlErrorResponseExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyEInvoiceXmlErrorResponseExtraWithDefaults() *VerifyEInvoiceXmlErrorResponseExtra {
	this := VerifyEInvoiceXmlErrorResponseExtra{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerifyEInvoiceXmlErrorResponseExtra) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerifyEInvoiceXmlErrorResponseExtra) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *VerifyEInvoiceXmlErrorResponseExtra) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *VerifyEInvoiceXmlErrorResponseExtra) SetErrors(v []string) *VerifyEInvoiceXmlErrorResponseExtra {
	o.Errors = v
	return o
}

func (o VerifyEInvoiceXmlErrorResponseExtra) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyEInvoiceXmlErrorResponseExtra) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableVerifyEInvoiceXmlErrorResponseExtra struct {
	value *VerifyEInvoiceXmlErrorResponseExtra
	isSet bool
}

func (v NullableVerifyEInvoiceXmlErrorResponseExtra) Get() *VerifyEInvoiceXmlErrorResponseExtra {
	return v.value
}

func (v *NullableVerifyEInvoiceXmlErrorResponseExtra) Set(val *VerifyEInvoiceXmlErrorResponseExtra) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyEInvoiceXmlErrorResponseExtra) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyEInvoiceXmlErrorResponseExtra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyEInvoiceXmlErrorResponseExtra(val *VerifyEInvoiceXmlErrorResponseExtra) *NullableVerifyEInvoiceXmlErrorResponseExtra {
	return &NullableVerifyEInvoiceXmlErrorResponseExtra{value: val, isSet: true}
}

func (v NullableVerifyEInvoiceXmlErrorResponseExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyEInvoiceXmlErrorResponseExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


