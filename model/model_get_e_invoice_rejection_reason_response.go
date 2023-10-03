/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.30
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the GetEInvoiceRejectionReasonResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEInvoiceRejectionReasonResponse{}

// GetEInvoiceRejectionReasonResponse struct for GetEInvoiceRejectionReasonResponse
type GetEInvoiceRejectionReasonResponse struct {
	Data *EInvoiceRejectionReason `json:"data,omitempty"`
}

// NewGetEInvoiceRejectionReasonResponse instantiates a new GetEInvoiceRejectionReasonResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEInvoiceRejectionReasonResponse() *GetEInvoiceRejectionReasonResponse {
	this := GetEInvoiceRejectionReasonResponse{}
	return &this
}

// NewGetEInvoiceRejectionReasonResponseWithDefaults instantiates a new GetEInvoiceRejectionReasonResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEInvoiceRejectionReasonResponseWithDefaults() *GetEInvoiceRejectionReasonResponse {
	this := GetEInvoiceRejectionReasonResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetEInvoiceRejectionReasonResponse) GetData() EInvoiceRejectionReason {
	if o == nil || IsNil(o.Data) {
		var ret EInvoiceRejectionReason
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEInvoiceRejectionReasonResponse) GetDataOk() (*EInvoiceRejectionReason, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetEInvoiceRejectionReasonResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given EInvoiceRejectionReason and assigns it to the Data field.
func (o *GetEInvoiceRejectionReasonResponse) SetData(v EInvoiceRejectionReason) *GetEInvoiceRejectionReasonResponse {
	o.Data = &v
	return o
}

func (o GetEInvoiceRejectionReasonResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEInvoiceRejectionReasonResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetEInvoiceRejectionReasonResponse struct {
	value *GetEInvoiceRejectionReasonResponse
	isSet bool
}

func (v NullableGetEInvoiceRejectionReasonResponse) Get() *GetEInvoiceRejectionReasonResponse {
	return v.value
}

func (v *NullableGetEInvoiceRejectionReasonResponse) Set(val *GetEInvoiceRejectionReasonResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEInvoiceRejectionReasonResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEInvoiceRejectionReasonResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEInvoiceRejectionReasonResponse(val *GetEInvoiceRejectionReasonResponse) *NullableGetEInvoiceRejectionReasonResponse {
	return &NullableGetEInvoiceRejectionReasonResponse{value: val, isSet: true}
}

func (v NullableGetEInvoiceRejectionReasonResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEInvoiceRejectionReasonResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


