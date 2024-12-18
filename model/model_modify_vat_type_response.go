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

// checks if the ModifyVatTypeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyVatTypeResponse{}

// ModifyVatTypeResponse struct for ModifyVatTypeResponse
type ModifyVatTypeResponse struct {
	Data NullableVatType `json:"data,omitempty"`
}

// NewModifyVatTypeResponse instantiates a new ModifyVatTypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyVatTypeResponse() *ModifyVatTypeResponse {
	this := ModifyVatTypeResponse{}
	return &this
}

// NewModifyVatTypeResponseWithDefaults instantiates a new ModifyVatTypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyVatTypeResponseWithDefaults() *ModifyVatTypeResponse {
	this := ModifyVatTypeResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModifyVatTypeResponse) GetData() VatType {
	if o == nil || IsNil(o.Data.Get()) {
		var ret VatType
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModifyVatTypeResponse) GetDataOk() (*VatType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *ModifyVatTypeResponse) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableVatType and assigns it to the Data field.
func (o *ModifyVatTypeResponse) SetData(v VatType) *ModifyVatTypeResponse {
	o.Data.Set(&v)
	return o
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *ModifyVatTypeResponse) SetDataNil() *ModifyVatTypeResponse {
	o.Data.Set(nil)
	return o
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *ModifyVatTypeResponse) UnsetData() {
	o.Data.Unset()
}

func (o ModifyVatTypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyVatTypeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return toSerialize, nil
}

type NullableModifyVatTypeResponse struct {
	value *ModifyVatTypeResponse
	isSet bool
}

func (v NullableModifyVatTypeResponse) Get() *ModifyVatTypeResponse {
	return v.value
}

func (v *NullableModifyVatTypeResponse) Set(val *ModifyVatTypeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyVatTypeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyVatTypeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyVatTypeResponse(val *ModifyVatTypeResponse) *NullableModifyVatTypeResponse {
	return &NullableModifyVatTypeResponse{value: val, isSet: true}
}

func (v NullableModifyVatTypeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyVatTypeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


