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

// checks if the SendEInvoiceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendEInvoiceResponse{}

// SendEInvoiceResponse struct for SendEInvoiceResponse
type SendEInvoiceResponse struct {
	Data NullableSendEInvoiceResponseData `json:"data,omitempty"`
}

// NewSendEInvoiceResponse instantiates a new SendEInvoiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendEInvoiceResponse() *SendEInvoiceResponse {
	this := SendEInvoiceResponse{}
	return &this
}

// NewSendEInvoiceResponseWithDefaults instantiates a new SendEInvoiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendEInvoiceResponseWithDefaults() *SendEInvoiceResponse {
	this := SendEInvoiceResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendEInvoiceResponse) GetData() SendEInvoiceResponseData {
	if o == nil || IsNil(o.Data.Get()) {
		var ret SendEInvoiceResponseData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendEInvoiceResponse) GetDataOk() (*SendEInvoiceResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *SendEInvoiceResponse) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableSendEInvoiceResponseData and assigns it to the Data field.
func (o *SendEInvoiceResponse) SetData(v SendEInvoiceResponseData) *SendEInvoiceResponse {
	o.Data.Set(&v)
	return o
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *SendEInvoiceResponse) SetDataNil() *SendEInvoiceResponse {
	o.Data.Set(nil)
	return o
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *SendEInvoiceResponse) UnsetData() {
	o.Data.Unset()
}

func (o SendEInvoiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendEInvoiceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return toSerialize, nil
}

type NullableSendEInvoiceResponse struct {
	value *SendEInvoiceResponse
	isSet bool
}

func (v NullableSendEInvoiceResponse) Get() *SendEInvoiceResponse {
	return v.value
}

func (v *NullableSendEInvoiceResponse) Set(val *SendEInvoiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSendEInvoiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSendEInvoiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendEInvoiceResponse(val *SendEInvoiceResponse) *NullableSendEInvoiceResponse {
	return &NullableSendEInvoiceResponse{value: val, isSet: true}
}

func (v NullableSendEInvoiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendEInvoiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


