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

// checks if the IssuedDocumentPreCreateInfoExtraDataDefaultValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssuedDocumentPreCreateInfoExtraDataDefaultValues{}

// IssuedDocumentPreCreateInfoExtraDataDefaultValues Issued document extra data default values
type IssuedDocumentPreCreateInfoExtraDataDefaultValues struct {
TsCommunication NullableBool `json:"ts_communication,omitempty"`
TsTipoSpesa NullableString `json:"ts_tipo_spesa,omitempty"`
TsFlagTipoSpesa NullableInt32 `json:"ts_flag_tipo_spesa,omitempty"`
TsPagamentoTracciato NullableBool `json:"ts_pagamento_tracciato,omitempty"`
}

// NewIssuedDocumentPreCreateInfoExtraDataDefaultValues instantiates a new IssuedDocumentPreCreateInfoExtraDataDefaultValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuedDocumentPreCreateInfoExtraDataDefaultValues() *IssuedDocumentPreCreateInfoExtraDataDefaultValues {
	this := IssuedDocumentPreCreateInfoExtraDataDefaultValues{}
	return &this
}

// NewIssuedDocumentPreCreateInfoExtraDataDefaultValuesWithDefaults instantiates a new IssuedDocumentPreCreateInfoExtraDataDefaultValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuedDocumentPreCreateInfoExtraDataDefaultValuesWithDefaults() *IssuedDocumentPreCreateInfoExtraDataDefaultValues {
	this := IssuedDocumentPreCreateInfoExtraDataDefaultValues{}
	return &this
}

// GetTsCommunication returns the TsCommunication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) GetTsCommunication() bool {
	if o == nil || IsNil(o.TsCommunication.Get()) {
		var ret bool
		return ret
	}
	return *o.TsCommunication.Get()
}

// GetTsCommunicationOk returns a tuple with the TsCommunication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) GetTsCommunicationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TsCommunication.Get(), o.TsCommunication.IsSet()
}

// HasTsCommunication returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) HasTsCommunication() bool {
	if o != nil && o.TsCommunication.IsSet() {
		return true
	}

	return false
}

// SetTsCommunication gets a reference to the given NullableBool and assigns it to the TsCommunication field.
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) SetTsCommunication(v bool) *IssuedDocumentPreCreateInfoExtraDataDefaultValues {
	o.TsCommunication.Set(&v)
        return o
}
// SetTsCommunicationNil sets the value for TsCommunication to be an explicit nil
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) SetTsCommunicationNil() *IssuedDocumentPreCreateInfoExtraDataDefaultValues {
	o.TsCommunication.Set(nil)
    return o
}

// UnsetTsCommunication ensures that no value is present for TsCommunication, not even an explicit nil
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) UnsetTsCommunication() {
	o.TsCommunication.Unset()
}

// GetTsTipoSpesa returns the TsTipoSpesa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) GetTsTipoSpesa() string {
	if o == nil || IsNil(o.TsTipoSpesa.Get()) {
		var ret string
		return ret
	}
	return *o.TsTipoSpesa.Get()
}

// GetTsTipoSpesaOk returns a tuple with the TsTipoSpesa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) GetTsTipoSpesaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TsTipoSpesa.Get(), o.TsTipoSpesa.IsSet()
}

// HasTsTipoSpesa returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) HasTsTipoSpesa() bool {
	if o != nil && o.TsTipoSpesa.IsSet() {
		return true
	}

	return false
}

// SetTsTipoSpesa gets a reference to the given NullableString and assigns it to the TsTipoSpesa field.
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) SetTsTipoSpesa(v string) *IssuedDocumentPreCreateInfoExtraDataDefaultValues {
	o.TsTipoSpesa.Set(&v)
        return o
}
// SetTsTipoSpesaNil sets the value for TsTipoSpesa to be an explicit nil
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) SetTsTipoSpesaNil() *IssuedDocumentPreCreateInfoExtraDataDefaultValues {
	o.TsTipoSpesa.Set(nil)
    return o
}

// UnsetTsTipoSpesa ensures that no value is present for TsTipoSpesa, not even an explicit nil
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) UnsetTsTipoSpesa() {
	o.TsTipoSpesa.Unset()
}

// GetTsFlagTipoSpesa returns the TsFlagTipoSpesa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) GetTsFlagTipoSpesa() int32 {
	if o == nil || IsNil(o.TsFlagTipoSpesa.Get()) {
		var ret int32
		return ret
	}
	return *o.TsFlagTipoSpesa.Get()
}

// GetTsFlagTipoSpesaOk returns a tuple with the TsFlagTipoSpesa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) GetTsFlagTipoSpesaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TsFlagTipoSpesa.Get(), o.TsFlagTipoSpesa.IsSet()
}

// HasTsFlagTipoSpesa returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) HasTsFlagTipoSpesa() bool {
	if o != nil && o.TsFlagTipoSpesa.IsSet() {
		return true
	}

	return false
}

// SetTsFlagTipoSpesa gets a reference to the given NullableInt32 and assigns it to the TsFlagTipoSpesa field.
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) SetTsFlagTipoSpesa(v int32) *IssuedDocumentPreCreateInfoExtraDataDefaultValues {
	o.TsFlagTipoSpesa.Set(&v)
        return o
}
// SetTsFlagTipoSpesaNil sets the value for TsFlagTipoSpesa to be an explicit nil
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) SetTsFlagTipoSpesaNil() *IssuedDocumentPreCreateInfoExtraDataDefaultValues {
	o.TsFlagTipoSpesa.Set(nil)
    return o
}

// UnsetTsFlagTipoSpesa ensures that no value is present for TsFlagTipoSpesa, not even an explicit nil
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) UnsetTsFlagTipoSpesa() {
	o.TsFlagTipoSpesa.Unset()
}

// GetTsPagamentoTracciato returns the TsPagamentoTracciato field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) GetTsPagamentoTracciato() bool {
	if o == nil || IsNil(o.TsPagamentoTracciato.Get()) {
		var ret bool
		return ret
	}
	return *o.TsPagamentoTracciato.Get()
}

// GetTsPagamentoTracciatoOk returns a tuple with the TsPagamentoTracciato field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) GetTsPagamentoTracciatoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TsPagamentoTracciato.Get(), o.TsPagamentoTracciato.IsSet()
}

// HasTsPagamentoTracciato returns a boolean if a field has been set.
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) HasTsPagamentoTracciato() bool {
	if o != nil && o.TsPagamentoTracciato.IsSet() {
		return true
	}

	return false
}

// SetTsPagamentoTracciato gets a reference to the given NullableBool and assigns it to the TsPagamentoTracciato field.
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) SetTsPagamentoTracciato(v bool) *IssuedDocumentPreCreateInfoExtraDataDefaultValues {
	o.TsPagamentoTracciato.Set(&v)
        return o
}
// SetTsPagamentoTracciatoNil sets the value for TsPagamentoTracciato to be an explicit nil
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) SetTsPagamentoTracciatoNil() *IssuedDocumentPreCreateInfoExtraDataDefaultValues {
	o.TsPagamentoTracciato.Set(nil)
    return o
}

// UnsetTsPagamentoTracciato ensures that no value is present for TsPagamentoTracciato, not even an explicit nil
func (o *IssuedDocumentPreCreateInfoExtraDataDefaultValues) UnsetTsPagamentoTracciato() {
	o.TsPagamentoTracciato.Unset()
}

func (o IssuedDocumentPreCreateInfoExtraDataDefaultValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssuedDocumentPreCreateInfoExtraDataDefaultValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TsCommunication.IsSet() {
		toSerialize["ts_communication"] = o.TsCommunication.Get()
	}
	if o.TsTipoSpesa.IsSet() {
		toSerialize["ts_tipo_spesa"] = o.TsTipoSpesa.Get()
	}
	if o.TsFlagTipoSpesa.IsSet() {
		toSerialize["ts_flag_tipo_spesa"] = o.TsFlagTipoSpesa.Get()
	}
	if o.TsPagamentoTracciato.IsSet() {
		toSerialize["ts_pagamento_tracciato"] = o.TsPagamentoTracciato.Get()
	}
	return toSerialize, nil
}

type NullableIssuedDocumentPreCreateInfoExtraDataDefaultValues struct {
	value *IssuedDocumentPreCreateInfoExtraDataDefaultValues
	isSet bool
}

func (v NullableIssuedDocumentPreCreateInfoExtraDataDefaultValues) Get() *IssuedDocumentPreCreateInfoExtraDataDefaultValues {
	return v.value
}

func (v *NullableIssuedDocumentPreCreateInfoExtraDataDefaultValues) Set(val *IssuedDocumentPreCreateInfoExtraDataDefaultValues) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuedDocumentPreCreateInfoExtraDataDefaultValues) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuedDocumentPreCreateInfoExtraDataDefaultValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuedDocumentPreCreateInfoExtraDataDefaultValues(val *IssuedDocumentPreCreateInfoExtraDataDefaultValues) *NullableIssuedDocumentPreCreateInfoExtraDataDefaultValues {
	return &NullableIssuedDocumentPreCreateInfoExtraDataDefaultValues{value: val, isSet: true}
}

func (v NullableIssuedDocumentPreCreateInfoExtraDataDefaultValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuedDocumentPreCreateInfoExtraDataDefaultValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


