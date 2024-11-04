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

// checks if the VerifyEInvoiceXmlErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyEInvoiceXmlErrorResponse{}

// VerifyEInvoiceXmlErrorResponse struct for VerifyEInvoiceXmlErrorResponse
type VerifyEInvoiceXmlErrorResponse struct {
Error NullableVerifyEInvoiceXmlErrorResponseError `json:"error,omitempty"`
Extra NullableVerifyEInvoiceXmlErrorResponseExtra `json:"extra,omitempty"`
}

// NewVerifyEInvoiceXmlErrorResponse instantiates a new VerifyEInvoiceXmlErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyEInvoiceXmlErrorResponse() *VerifyEInvoiceXmlErrorResponse {
	this := VerifyEInvoiceXmlErrorResponse{}
	return &this
}

// NewVerifyEInvoiceXmlErrorResponseWithDefaults instantiates a new VerifyEInvoiceXmlErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyEInvoiceXmlErrorResponseWithDefaults() *VerifyEInvoiceXmlErrorResponse {
	this := VerifyEInvoiceXmlErrorResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerifyEInvoiceXmlErrorResponse) GetError() VerifyEInvoiceXmlErrorResponseError {
	if o == nil || IsNil(o.Error.Get()) {
		var ret VerifyEInvoiceXmlErrorResponseError
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerifyEInvoiceXmlErrorResponse) GetErrorOk() (*VerifyEInvoiceXmlErrorResponseError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *VerifyEInvoiceXmlErrorResponse) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableVerifyEInvoiceXmlErrorResponseError and assigns it to the Error field.
func (o *VerifyEInvoiceXmlErrorResponse) SetError(v VerifyEInvoiceXmlErrorResponseError) *VerifyEInvoiceXmlErrorResponse {
	o.Error.Set(&v)
		return o
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *VerifyEInvoiceXmlErrorResponse) SetErrorNil() *VerifyEInvoiceXmlErrorResponse {
	o.Error.Set(nil)
	return o
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *VerifyEInvoiceXmlErrorResponse) UnsetError() {
	o.Error.Unset()
}

// GetExtra returns the Extra field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerifyEInvoiceXmlErrorResponse) GetExtra() VerifyEInvoiceXmlErrorResponseExtra {
	if o == nil || IsNil(o.Extra.Get()) {
		var ret VerifyEInvoiceXmlErrorResponseExtra
		return ret
	}
	return *o.Extra.Get()
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerifyEInvoiceXmlErrorResponse) GetExtraOk() (*VerifyEInvoiceXmlErrorResponseExtra, bool) {
	if o == nil {
		return nil, false
	}
	return o.Extra.Get(), o.Extra.IsSet()
}

// HasExtra returns a boolean if a field has been set.
func (o *VerifyEInvoiceXmlErrorResponse) HasExtra() bool {
	if o != nil && o.Extra.IsSet() {
		return true
	}

	return false
}

// SetExtra gets a reference to the given NullableVerifyEInvoiceXmlErrorResponseExtra and assigns it to the Extra field.
func (o *VerifyEInvoiceXmlErrorResponse) SetExtra(v VerifyEInvoiceXmlErrorResponseExtra) *VerifyEInvoiceXmlErrorResponse {
	o.Extra.Set(&v)
		return o
}
// SetExtraNil sets the value for Extra to be an explicit nil
func (o *VerifyEInvoiceXmlErrorResponse) SetExtraNil() *VerifyEInvoiceXmlErrorResponse {
	o.Extra.Set(nil)
	return o
}

// UnsetExtra ensures that no value is present for Extra, not even an explicit nil
func (o *VerifyEInvoiceXmlErrorResponse) UnsetExtra() {
	o.Extra.Unset()
}

func (o VerifyEInvoiceXmlErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyEInvoiceXmlErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.Extra.IsSet() {
		toSerialize["extra"] = o.Extra.Get()
	}
	return toSerialize, nil
}

type NullableVerifyEInvoiceXmlErrorResponse struct {
	value *VerifyEInvoiceXmlErrorResponse
	isSet bool
}

func (v NullableVerifyEInvoiceXmlErrorResponse) Get() *VerifyEInvoiceXmlErrorResponse {
	return v.value
}

func (v *NullableVerifyEInvoiceXmlErrorResponse) Set(val *VerifyEInvoiceXmlErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyEInvoiceXmlErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyEInvoiceXmlErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyEInvoiceXmlErrorResponse(val *VerifyEInvoiceXmlErrorResponse) *NullableVerifyEInvoiceXmlErrorResponse {
	return &NullableVerifyEInvoiceXmlErrorResponse{value: val, isSet: true}
}

func (v NullableVerifyEInvoiceXmlErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyEInvoiceXmlErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


