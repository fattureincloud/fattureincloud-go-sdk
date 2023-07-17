/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.29
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the SendEInvoiceRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendEInvoiceRequestData{}

// SendEInvoiceRequestData struct for SendEInvoiceRequestData
type SendEInvoiceRequestData struct {
	// Value of TipoCassa used (optional, override the company default value).
	CassaType NullableString `json:"cassa_type,omitempty"`
	// Value of CausalePagamento used (optional, override the company default value).
	WithholdingTaxCausal NullableString `json:"withholding_tax_causal,omitempty"`
}

// NewSendEInvoiceRequestData instantiates a new SendEInvoiceRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendEInvoiceRequestData() *SendEInvoiceRequestData {
	this := SendEInvoiceRequestData{}
	return &this
}

// NewSendEInvoiceRequestDataWithDefaults instantiates a new SendEInvoiceRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendEInvoiceRequestDataWithDefaults() *SendEInvoiceRequestData {
	this := SendEInvoiceRequestData{}
	return &this
}

// GetCassaType returns the CassaType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendEInvoiceRequestData) GetCassaType() string {
	if o == nil || IsNil(o.CassaType.Get()) {
		var ret string
		return ret
	}
	return *o.CassaType.Get()
}

// GetCassaTypeOk returns a tuple with the CassaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendEInvoiceRequestData) GetCassaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CassaType.Get(), o.CassaType.IsSet()
}

// HasCassaType returns a boolean if a field has been set.
func (o *SendEInvoiceRequestData) HasCassaType() bool {
	if o != nil && o.CassaType.IsSet() {
		return true
	}

	return false
}

// SetCassaType gets a reference to the given NullableString and assigns it to the CassaType field.
func (o *SendEInvoiceRequestData) SetCassaType(v string) *SendEInvoiceRequestData {
	o.CassaType.Set(&v)
	return o
}
// SetCassaTypeNil sets the value for CassaType to be an explicit nil
func (o *SendEInvoiceRequestData) SetCassaTypeNil() *SendEInvoiceRequestData {
	o.CassaType.Set(nil)
	return o
}

// UnsetCassaType ensures that no value is present for CassaType, not even an explicit nil
func (o *SendEInvoiceRequestData) UnsetCassaType() {
	o.CassaType.Unset()
}

// GetWithholdingTaxCausal returns the WithholdingTaxCausal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendEInvoiceRequestData) GetWithholdingTaxCausal() string {
	if o == nil || IsNil(o.WithholdingTaxCausal.Get()) {
		var ret string
		return ret
	}
	return *o.WithholdingTaxCausal.Get()
}

// GetWithholdingTaxCausalOk returns a tuple with the WithholdingTaxCausal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendEInvoiceRequestData) GetWithholdingTaxCausalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithholdingTaxCausal.Get(), o.WithholdingTaxCausal.IsSet()
}

// HasWithholdingTaxCausal returns a boolean if a field has been set.
func (o *SendEInvoiceRequestData) HasWithholdingTaxCausal() bool {
	if o != nil && o.WithholdingTaxCausal.IsSet() {
		return true
	}

	return false
}

// SetWithholdingTaxCausal gets a reference to the given NullableString and assigns it to the WithholdingTaxCausal field.
func (o *SendEInvoiceRequestData) SetWithholdingTaxCausal(v string) *SendEInvoiceRequestData {
	o.WithholdingTaxCausal.Set(&v)
	return o
}
// SetWithholdingTaxCausalNil sets the value for WithholdingTaxCausal to be an explicit nil
func (o *SendEInvoiceRequestData) SetWithholdingTaxCausalNil() *SendEInvoiceRequestData {
	o.WithholdingTaxCausal.Set(nil)
	return o
}

// UnsetWithholdingTaxCausal ensures that no value is present for WithholdingTaxCausal, not even an explicit nil
func (o *SendEInvoiceRequestData) UnsetWithholdingTaxCausal() {
	o.WithholdingTaxCausal.Unset()
}

func (o SendEInvoiceRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendEInvoiceRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CassaType.IsSet() {
		toSerialize["cassa_type"] = o.CassaType.Get()
	}
	if o.WithholdingTaxCausal.IsSet() {
		toSerialize["withholding_tax_causal"] = o.WithholdingTaxCausal.Get()
	}
	return toSerialize, nil
}

type NullableSendEInvoiceRequestData struct {
	value *SendEInvoiceRequestData
	isSet bool
}

func (v NullableSendEInvoiceRequestData) Get() *SendEInvoiceRequestData {
	return v.value
}

func (v *NullableSendEInvoiceRequestData) Set(val *SendEInvoiceRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSendEInvoiceRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSendEInvoiceRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendEInvoiceRequestData(val *SendEInvoiceRequestData) *NullableSendEInvoiceRequestData {
	return &NullableSendEInvoiceRequestData{value: val, isSet: true}
}

func (v NullableSendEInvoiceRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendEInvoiceRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


