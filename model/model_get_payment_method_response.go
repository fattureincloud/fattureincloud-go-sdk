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

// checks if the GetPaymentMethodResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPaymentMethodResponse{}

// GetPaymentMethodResponse struct for GetPaymentMethodResponse
type GetPaymentMethodResponse struct {
	Data *PaymentMethod `json:"data,omitempty"`
}

// NewGetPaymentMethodResponse instantiates a new GetPaymentMethodResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPaymentMethodResponse() *GetPaymentMethodResponse {
	this := GetPaymentMethodResponse{}
	return &this
}

// NewGetPaymentMethodResponseWithDefaults instantiates a new GetPaymentMethodResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPaymentMethodResponseWithDefaults() *GetPaymentMethodResponse {
	this := GetPaymentMethodResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetPaymentMethodResponse) GetData() PaymentMethod {
	if o == nil || IsNil(o.Data) {
		var ret PaymentMethod
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentMethodResponse) GetDataOk() (*PaymentMethod, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetPaymentMethodResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PaymentMethod and assigns it to the Data field.
func (o *GetPaymentMethodResponse) SetData(v PaymentMethod) *GetPaymentMethodResponse {
	o.Data = &v
	return o
}

func (o GetPaymentMethodResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPaymentMethodResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetPaymentMethodResponse struct {
	value *GetPaymentMethodResponse
	isSet bool
}

func (v NullableGetPaymentMethodResponse) Get() *GetPaymentMethodResponse {
	return v.value
}

func (v *NullableGetPaymentMethodResponse) Set(val *GetPaymentMethodResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPaymentMethodResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPaymentMethodResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPaymentMethodResponse(val *GetPaymentMethodResponse) *NullableGetPaymentMethodResponse {
	return &NullableGetPaymentMethodResponse{value: val, isSet: true}
}

func (v NullableGetPaymentMethodResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPaymentMethodResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


