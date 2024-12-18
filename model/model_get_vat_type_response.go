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

// checks if the GetVatTypeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVatTypeResponse{}

// GetVatTypeResponse struct for GetVatTypeResponse
type GetVatTypeResponse struct {
	Data NullableVatType `json:"data,omitempty"`
}

// NewGetVatTypeResponse instantiates a new GetVatTypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVatTypeResponse() *GetVatTypeResponse {
	this := GetVatTypeResponse{}
	return &this
}

// NewGetVatTypeResponseWithDefaults instantiates a new GetVatTypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVatTypeResponseWithDefaults() *GetVatTypeResponse {
	this := GetVatTypeResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetVatTypeResponse) GetData() VatType {
	if o == nil || IsNil(o.Data.Get()) {
		var ret VatType
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetVatTypeResponse) GetDataOk() (*VatType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *GetVatTypeResponse) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableVatType and assigns it to the Data field.
func (o *GetVatTypeResponse) SetData(v VatType) *GetVatTypeResponse {
	o.Data.Set(&v)
	return o
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *GetVatTypeResponse) SetDataNil() *GetVatTypeResponse {
	o.Data.Set(nil)
	return o
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *GetVatTypeResponse) UnsetData() {
	o.Data.Unset()
}

func (o GetVatTypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVatTypeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return toSerialize, nil
}

type NullableGetVatTypeResponse struct {
	value *GetVatTypeResponse
	isSet bool
}

func (v NullableGetVatTypeResponse) Get() *GetVatTypeResponse {
	return v.value
}

func (v *NullableGetVatTypeResponse) Set(val *GetVatTypeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVatTypeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVatTypeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVatTypeResponse(val *GetVatTypeResponse) *NullableGetVatTypeResponse {
	return &NullableGetVatTypeResponse{value: val, isSet: true}
}

func (v NullableGetVatTypeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVatTypeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


