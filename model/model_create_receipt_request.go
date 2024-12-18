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

// checks if the CreateReceiptRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateReceiptRequest{}

// CreateReceiptRequest 
type CreateReceiptRequest struct {
	Data *Receipt `json:"data,omitempty"`
	// If true, the number is autocompleted progressively.
	AutocompleteNumber NullableBool `json:"autocomplete_number,omitempty"`
}

// NewCreateReceiptRequest instantiates a new CreateReceiptRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReceiptRequest() *CreateReceiptRequest {
	this := CreateReceiptRequest{}
	return &this
}

// NewCreateReceiptRequestWithDefaults instantiates a new CreateReceiptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReceiptRequestWithDefaults() *CreateReceiptRequest {
	this := CreateReceiptRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateReceiptRequest) GetData() Receipt {
	if o == nil || IsNil(o.Data) {
		var ret Receipt
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateReceiptRequest) GetDataOk() (*Receipt, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateReceiptRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Receipt and assigns it to the Data field.
func (o *CreateReceiptRequest) SetData(v Receipt) *CreateReceiptRequest {
	o.Data = &v
	return o
}

// GetAutocompleteNumber returns the AutocompleteNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateReceiptRequest) GetAutocompleteNumber() bool {
	if o == nil || IsNil(o.AutocompleteNumber.Get()) {
		var ret bool
		return ret
	}
	return *o.AutocompleteNumber.Get()
}

// GetAutocompleteNumberOk returns a tuple with the AutocompleteNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateReceiptRequest) GetAutocompleteNumberOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutocompleteNumber.Get(), o.AutocompleteNumber.IsSet()
}

// HasAutocompleteNumber returns a boolean if a field has been set.
func (o *CreateReceiptRequest) HasAutocompleteNumber() bool {
	if o != nil && o.AutocompleteNumber.IsSet() {
		return true
	}

	return false
}

// SetAutocompleteNumber gets a reference to the given NullableBool and assigns it to the AutocompleteNumber field.
func (o *CreateReceiptRequest) SetAutocompleteNumber(v bool) *CreateReceiptRequest {
	o.AutocompleteNumber.Set(&v)
	return o
}
// SetAutocompleteNumberNil sets the value for AutocompleteNumber to be an explicit nil
func (o *CreateReceiptRequest) SetAutocompleteNumberNil() *CreateReceiptRequest {
	o.AutocompleteNumber.Set(nil)
	return o
}

// UnsetAutocompleteNumber ensures that no value is present for AutocompleteNumber, not even an explicit nil
func (o *CreateReceiptRequest) UnsetAutocompleteNumber() {
	o.AutocompleteNumber.Unset()
}

func (o CreateReceiptRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateReceiptRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.AutocompleteNumber.IsSet() {
		toSerialize["autocomplete_number"] = o.AutocompleteNumber.Get()
	}
	return toSerialize, nil
}

type NullableCreateReceiptRequest struct {
	value *CreateReceiptRequest
	isSet bool
}

func (v NullableCreateReceiptRequest) Get() *CreateReceiptRequest {
	return v.value
}

func (v *NullableCreateReceiptRequest) Set(val *CreateReceiptRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReceiptRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReceiptRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReceiptRequest(val *CreateReceiptRequest) *NullableCreateReceiptRequest {
	return &NullableCreateReceiptRequest{value: val, isSet: true}
}

func (v NullableCreateReceiptRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReceiptRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


