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

// checks if the CreatePaymentMethodResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePaymentMethodResponse{}

// CreatePaymentMethodResponse struct for CreatePaymentMethodResponse
type CreatePaymentMethodResponse struct {
Data *PaymentMethod `json:"data,omitempty"`
}

// NewCreatePaymentMethodResponse instantiates a new CreatePaymentMethodResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePaymentMethodResponse() *CreatePaymentMethodResponse {
	this := CreatePaymentMethodResponse{}
	return &this
}

// NewCreatePaymentMethodResponseWithDefaults instantiates a new CreatePaymentMethodResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePaymentMethodResponseWithDefaults() *CreatePaymentMethodResponse {
	this := CreatePaymentMethodResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreatePaymentMethodResponse) GetData() PaymentMethod {
	if o == nil || IsNil(o.Data) {
		var ret PaymentMethod
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentMethodResponse) GetDataOk() (*PaymentMethod, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreatePaymentMethodResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PaymentMethod and assigns it to the Data field.
func (o *CreatePaymentMethodResponse) SetData(v PaymentMethod) *CreatePaymentMethodResponse {
	o.Data = &v
		return o
}

func (o CreatePaymentMethodResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePaymentMethodResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreatePaymentMethodResponse struct {
	value *CreatePaymentMethodResponse
	isSet bool
}

func (v NullableCreatePaymentMethodResponse) Get() *CreatePaymentMethodResponse {
	return v.value
}

func (v *NullableCreatePaymentMethodResponse) Set(val *CreatePaymentMethodResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePaymentMethodResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePaymentMethodResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePaymentMethodResponse(val *CreatePaymentMethodResponse) *NullableCreatePaymentMethodResponse {
	return &NullableCreatePaymentMethodResponse{value: val, isSet: true}
}

func (v NullableCreatePaymentMethodResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePaymentMethodResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


