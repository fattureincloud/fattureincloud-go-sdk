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

// checks if the ModifyVatTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyVatTypeRequest{}

// ModifyVatTypeRequest struct for ModifyVatTypeRequest
type ModifyVatTypeRequest struct {
	Data NullableVatType `json:"data,omitempty"`
}

// NewModifyVatTypeRequest instantiates a new ModifyVatTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyVatTypeRequest() *ModifyVatTypeRequest {
	this := ModifyVatTypeRequest{}
	return &this
}

// NewModifyVatTypeRequestWithDefaults instantiates a new ModifyVatTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyVatTypeRequestWithDefaults() *ModifyVatTypeRequest {
	this := ModifyVatTypeRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModifyVatTypeRequest) GetData() VatType {
	if o == nil || IsNil(o.Data.Get()) {
		var ret VatType
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModifyVatTypeRequest) GetDataOk() (*VatType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *ModifyVatTypeRequest) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableVatType and assigns it to the Data field.
func (o *ModifyVatTypeRequest) SetData(v VatType) *ModifyVatTypeRequest {
	o.Data.Set(&v)
	return o
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *ModifyVatTypeRequest) SetDataNil() *ModifyVatTypeRequest {
	o.Data.Set(nil)
	return o
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *ModifyVatTypeRequest) UnsetData() {
	o.Data.Unset()
}

func (o ModifyVatTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyVatTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return toSerialize, nil
}

type NullableModifyVatTypeRequest struct {
	value *ModifyVatTypeRequest
	isSet bool
}

func (v NullableModifyVatTypeRequest) Get() *ModifyVatTypeRequest {
	return v.value
}

func (v *NullableModifyVatTypeRequest) Set(val *ModifyVatTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyVatTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyVatTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyVatTypeRequest(val *ModifyVatTypeRequest) *NullableModifyVatTypeRequest {
	return &NullableModifyVatTypeRequest{value: val, isSet: true}
}

func (v NullableModifyVatTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyVatTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


