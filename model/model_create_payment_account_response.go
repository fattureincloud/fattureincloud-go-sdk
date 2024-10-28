/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.2
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the CreatePaymentAccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePaymentAccountResponse{}

// CreatePaymentAccountResponse struct for CreatePaymentAccountResponse
type CreatePaymentAccountResponse struct {
Data NullablePaymentAccount `json:"data,omitempty"`
}

// NewCreatePaymentAccountResponse instantiates a new CreatePaymentAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePaymentAccountResponse() *CreatePaymentAccountResponse {
	this := CreatePaymentAccountResponse{}
	return &this
}

// NewCreatePaymentAccountResponseWithDefaults instantiates a new CreatePaymentAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePaymentAccountResponseWithDefaults() *CreatePaymentAccountResponse {
	this := CreatePaymentAccountResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePaymentAccountResponse) GetData() PaymentAccount {
	if o == nil || IsNil(o.Data.Get()) {
		var ret PaymentAccount
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePaymentAccountResponse) GetDataOk() (*PaymentAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *CreatePaymentAccountResponse) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullablePaymentAccount and assigns it to the Data field.
func (o *CreatePaymentAccountResponse) SetData(v PaymentAccount) *CreatePaymentAccountResponse {
	o.Data.Set(&v)
        return o
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *CreatePaymentAccountResponse) SetDataNil() *CreatePaymentAccountResponse {
	o.Data.Set(nil)
    return o
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *CreatePaymentAccountResponse) UnsetData() {
	o.Data.Unset()
}

func (o CreatePaymentAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePaymentAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return toSerialize, nil
}

type NullableCreatePaymentAccountResponse struct {
	value *CreatePaymentAccountResponse
	isSet bool
}

func (v NullableCreatePaymentAccountResponse) Get() *CreatePaymentAccountResponse {
	return v.value
}

func (v *NullableCreatePaymentAccountResponse) Set(val *CreatePaymentAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePaymentAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePaymentAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePaymentAccountResponse(val *CreatePaymentAccountResponse) *NullableCreatePaymentAccountResponse {
	return &NullableCreatePaymentAccountResponse{value: val, isSet: true}
}

func (v NullableCreatePaymentAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePaymentAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


