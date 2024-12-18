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

// checks if the ListPaymentAccountsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPaymentAccountsResponse{}

// ListPaymentAccountsResponse 
type ListPaymentAccountsResponse struct {
	Data []PaymentAccount `json:"data,omitempty"`
}

// NewListPaymentAccountsResponse instantiates a new ListPaymentAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPaymentAccountsResponse() *ListPaymentAccountsResponse {
	this := ListPaymentAccountsResponse{}
	return &this
}

// NewListPaymentAccountsResponseWithDefaults instantiates a new ListPaymentAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPaymentAccountsResponseWithDefaults() *ListPaymentAccountsResponse {
	this := ListPaymentAccountsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListPaymentAccountsResponse) GetData() []PaymentAccount {
	if o == nil {
		var ret []PaymentAccount
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListPaymentAccountsResponse) GetDataOk() ([]PaymentAccount, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListPaymentAccountsResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PaymentAccount and assigns it to the Data field.
func (o *ListPaymentAccountsResponse) SetData(v []PaymentAccount) *ListPaymentAccountsResponse {
	o.Data = v
	return o
}

func (o ListPaymentAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPaymentAccountsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListPaymentAccountsResponse struct {
	value *ListPaymentAccountsResponse
	isSet bool
}

func (v NullableListPaymentAccountsResponse) Get() *ListPaymentAccountsResponse {
	return v.value
}

func (v *NullableListPaymentAccountsResponse) Set(val *ListPaymentAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPaymentAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPaymentAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPaymentAccountsResponse(val *ListPaymentAccountsResponse) *NullableListPaymentAccountsResponse {
	return &NullableListPaymentAccountsResponse{value: val, isSet: true}
}

func (v NullableListPaymentAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPaymentAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


