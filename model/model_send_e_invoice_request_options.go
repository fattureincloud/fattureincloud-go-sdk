/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.26
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the SendEInvoiceRequestOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendEInvoiceRequestOptions{}

// SendEInvoiceRequestOptions struct for SendEInvoiceRequestOptions
type SendEInvoiceRequestOptions struct {
	// If set to true the e-invoice will not be sent to the SDI.
	DryRun NullableBool `json:"dry_run,omitempty"`
}

// NewSendEInvoiceRequestOptions instantiates a new SendEInvoiceRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendEInvoiceRequestOptions() *SendEInvoiceRequestOptions {
	this := SendEInvoiceRequestOptions{}
	return &this
}

// NewSendEInvoiceRequestOptionsWithDefaults instantiates a new SendEInvoiceRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendEInvoiceRequestOptionsWithDefaults() *SendEInvoiceRequestOptions {
	this := SendEInvoiceRequestOptions{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendEInvoiceRequestOptions) GetDryRun() bool {
	if o == nil || isNil(o.DryRun.Get()) {
		var ret bool
		return ret
	}
	return *o.DryRun.Get()
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendEInvoiceRequestOptions) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DryRun.Get(), o.DryRun.IsSet()
}

// HasDryRun returns a boolean if a field has been set.
func (o *SendEInvoiceRequestOptions) HasDryRun() bool {
	if o != nil && o.DryRun.IsSet() {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given NullableBool and assigns it to the DryRun field.
func (o *SendEInvoiceRequestOptions) SetDryRun(v bool) *SendEInvoiceRequestOptions {
	o.DryRun.Set(&v)
	return o
}
// SetDryRunNil sets the value for DryRun to be an explicit nil
func (o *SendEInvoiceRequestOptions) SetDryRunNil() *SendEInvoiceRequestOptions {
	o.DryRun.Set(nil)
	return o
}

// UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
func (o *SendEInvoiceRequestOptions) UnsetDryRun() {
	o.DryRun.Unset()
}

func (o SendEInvoiceRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendEInvoiceRequestOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun.IsSet() {
		toSerialize["dry_run"] = o.DryRun.Get()
	}
	return toSerialize, nil
}

type NullableSendEInvoiceRequestOptions struct {
	value *SendEInvoiceRequestOptions
	isSet bool
}

func (v NullableSendEInvoiceRequestOptions) Get() *SendEInvoiceRequestOptions {
	return v.value
}

func (v *NullableSendEInvoiceRequestOptions) Set(val *SendEInvoiceRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSendEInvoiceRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSendEInvoiceRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendEInvoiceRequestOptions(val *SendEInvoiceRequestOptions) *NullableSendEInvoiceRequestOptions {
	return &NullableSendEInvoiceRequestOptions{value: val, isSet: true}
}

func (v NullableSendEInvoiceRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendEInvoiceRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

