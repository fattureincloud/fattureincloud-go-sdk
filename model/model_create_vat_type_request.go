/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.32
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the CreateVatTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVatTypeRequest{}

// CreateVatTypeRequest struct for CreateVatTypeRequest
type CreateVatTypeRequest struct {
	Data NullableVatType `json:"data,omitempty"`
}

// NewCreateVatTypeRequest instantiates a new CreateVatTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVatTypeRequest() *CreateVatTypeRequest {
	this := CreateVatTypeRequest{}
	return &this
}

// NewCreateVatTypeRequestWithDefaults instantiates a new CreateVatTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVatTypeRequestWithDefaults() *CreateVatTypeRequest {
	this := CreateVatTypeRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVatTypeRequest) GetData() VatType {
	if o == nil || IsNil(o.Data.Get()) {
		var ret VatType
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVatTypeRequest) GetDataOk() (*VatType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *CreateVatTypeRequest) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableVatType and assigns it to the Data field.
func (o *CreateVatTypeRequest) SetData(v VatType) *CreateVatTypeRequest {
	o.Data.Set(&v)
	return o
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *CreateVatTypeRequest) SetDataNil() *CreateVatTypeRequest {
	o.Data.Set(nil)
	return o
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *CreateVatTypeRequest) UnsetData() {
	o.Data.Unset()
}

func (o CreateVatTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVatTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return toSerialize, nil
}

type NullableCreateVatTypeRequest struct {
	value *CreateVatTypeRequest
	isSet bool
}

func (v NullableCreateVatTypeRequest) Get() *CreateVatTypeRequest {
	return v.value
}

func (v *NullableCreateVatTypeRequest) Set(val *CreateVatTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVatTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVatTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVatTypeRequest(val *CreateVatTypeRequest) *NullableCreateVatTypeRequest {
	return &NullableCreateVatTypeRequest{value: val, isSet: true}
}

func (v NullableCreateVatTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVatTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


