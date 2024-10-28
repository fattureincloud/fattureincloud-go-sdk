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

// checks if the SendEInvoiceResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendEInvoiceResponseData{}

// SendEInvoiceResponseData struct for SendEInvoiceResponseData
type SendEInvoiceResponseData struct {
	// Response message.
Name NullableString `json:"name,omitempty"`
	// E-invoice sent date.
Date NullableString `json:"date,omitempty"`
}

// NewSendEInvoiceResponseData instantiates a new SendEInvoiceResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendEInvoiceResponseData() *SendEInvoiceResponseData {
	this := SendEInvoiceResponseData{}
	return &this
}

// NewSendEInvoiceResponseDataWithDefaults instantiates a new SendEInvoiceResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendEInvoiceResponseDataWithDefaults() *SendEInvoiceResponseData {
	this := SendEInvoiceResponseData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendEInvoiceResponseData) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendEInvoiceResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SendEInvoiceResponseData) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SendEInvoiceResponseData) SetName(v string) *SendEInvoiceResponseData {
	o.Name.Set(&v)
        return o
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SendEInvoiceResponseData) SetNameNil() *SendEInvoiceResponseData {
	o.Name.Set(nil)
    return o
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SendEInvoiceResponseData) UnsetName() {
	o.Name.Unset()
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendEInvoiceResponseData) GetDate() string {
	if o == nil || IsNil(o.Date.Get()) {
		var ret string
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendEInvoiceResponseData) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *SendEInvoiceResponseData) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableString and assigns it to the Date field.
func (o *SendEInvoiceResponseData) SetDate(v string) *SendEInvoiceResponseData {
	o.Date.Set(&v)
        return o
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *SendEInvoiceResponseData) SetDateNil() *SendEInvoiceResponseData {
	o.Date.Set(nil)
    return o
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *SendEInvoiceResponseData) UnsetDate() {
	o.Date.Unset()
}

func (o SendEInvoiceResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendEInvoiceResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}
	return toSerialize, nil
}

type NullableSendEInvoiceResponseData struct {
	value *SendEInvoiceResponseData
	isSet bool
}

func (v NullableSendEInvoiceResponseData) Get() *SendEInvoiceResponseData {
	return v.value
}

func (v *NullableSendEInvoiceResponseData) Set(val *SendEInvoiceResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableSendEInvoiceResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableSendEInvoiceResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendEInvoiceResponseData(val *SendEInvoiceResponseData) *NullableSendEInvoiceResponseData {
	return &NullableSendEInvoiceResponseData{value: val, isSet: true}
}

func (v NullableSendEInvoiceResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendEInvoiceResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


