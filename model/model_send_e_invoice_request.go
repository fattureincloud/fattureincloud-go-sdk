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

// checks if the SendEInvoiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendEInvoiceRequest{}

// SendEInvoiceRequest struct for SendEInvoiceRequest
type SendEInvoiceRequest struct {
	Data NullableSendEInvoiceRequestData `json:"data,omitempty"`
	Options NullableSendEInvoiceRequestOptions `json:"options,omitempty"`
}

// NewSendEInvoiceRequest instantiates a new SendEInvoiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendEInvoiceRequest() *SendEInvoiceRequest {
	this := SendEInvoiceRequest{}
	return &this
}

// NewSendEInvoiceRequestWithDefaults instantiates a new SendEInvoiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendEInvoiceRequestWithDefaults() *SendEInvoiceRequest {
	this := SendEInvoiceRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendEInvoiceRequest) GetData() SendEInvoiceRequestData {
	if o == nil || IsNil(o.Data.Get()) {
		var ret SendEInvoiceRequestData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendEInvoiceRequest) GetDataOk() (*SendEInvoiceRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *SendEInvoiceRequest) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableSendEInvoiceRequestData and assigns it to the Data field.
func (o *SendEInvoiceRequest) SetData(v SendEInvoiceRequestData) *SendEInvoiceRequest {
	o.Data.Set(&v)
	return o
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *SendEInvoiceRequest) SetDataNil() *SendEInvoiceRequest {
	o.Data.Set(nil)
	return o
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *SendEInvoiceRequest) UnsetData() {
	o.Data.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendEInvoiceRequest) GetOptions() SendEInvoiceRequestOptions {
	if o == nil || IsNil(o.Options.Get()) {
		var ret SendEInvoiceRequestOptions
		return ret
	}
	return *o.Options.Get()
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendEInvoiceRequest) GetOptionsOk() (*SendEInvoiceRequestOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options.Get(), o.Options.IsSet()
}

// HasOptions returns a boolean if a field has been set.
func (o *SendEInvoiceRequest) HasOptions() bool {
	if o != nil && o.Options.IsSet() {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NullableSendEInvoiceRequestOptions and assigns it to the Options field.
func (o *SendEInvoiceRequest) SetOptions(v SendEInvoiceRequestOptions) *SendEInvoiceRequest {
	o.Options.Set(&v)
	return o
}
// SetOptionsNil sets the value for Options to be an explicit nil
func (o *SendEInvoiceRequest) SetOptionsNil() *SendEInvoiceRequest {
	o.Options.Set(nil)
	return o
}

// UnsetOptions ensures that no value is present for Options, not even an explicit nil
func (o *SendEInvoiceRequest) UnsetOptions() {
	o.Options.Unset()
}

func (o SendEInvoiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendEInvoiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if o.Options.IsSet() {
		toSerialize["options"] = o.Options.Get()
	}
	return toSerialize, nil
}

type NullableSendEInvoiceRequest struct {
	value *SendEInvoiceRequest
	isSet bool
}

func (v NullableSendEInvoiceRequest) Get() *SendEInvoiceRequest {
	return v.value
}

func (v *NullableSendEInvoiceRequest) Set(val *SendEInvoiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendEInvoiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendEInvoiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendEInvoiceRequest(val *SendEInvoiceRequest) *NullableSendEInvoiceRequest {
	return &NullableSendEInvoiceRequest{value: val, isSet: true}
}

func (v NullableSendEInvoiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendEInvoiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


