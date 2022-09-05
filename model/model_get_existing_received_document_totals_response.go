/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// GetExistingReceivedDocumentTotalsResponse 
type GetExistingReceivedDocumentTotalsResponse struct {
	Data *ReceivedDocumentTotals `json:"data,omitempty"`
}

// NewGetExistingReceivedDocumentTotalsResponse instantiates a new GetExistingReceivedDocumentTotalsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExistingReceivedDocumentTotalsResponse() *GetExistingReceivedDocumentTotalsResponse {
	this := GetExistingReceivedDocumentTotalsResponse{}
	return &this
}

// NewGetExistingReceivedDocumentTotalsResponseWithDefaults instantiates a new GetExistingReceivedDocumentTotalsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExistingReceivedDocumentTotalsResponseWithDefaults() *GetExistingReceivedDocumentTotalsResponse {
	this := GetExistingReceivedDocumentTotalsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetExistingReceivedDocumentTotalsResponse) GetData() ReceivedDocumentTotals {
	if o == nil || o.Data == nil {
		var ret ReceivedDocumentTotals
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExistingReceivedDocumentTotalsResponse) GetDataOk() (*ReceivedDocumentTotals, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetExistingReceivedDocumentTotalsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ReceivedDocumentTotals and assigns it to the Data field.
func (o *GetExistingReceivedDocumentTotalsResponse) SetData(v ReceivedDocumentTotals) *GetExistingReceivedDocumentTotalsResponse {
	o.Data = &v
	return o
}

func (o GetExistingReceivedDocumentTotalsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetExistingReceivedDocumentTotalsResponse struct {
	value *GetExistingReceivedDocumentTotalsResponse
	isSet bool
}

func (v NullableGetExistingReceivedDocumentTotalsResponse) Get() *GetExistingReceivedDocumentTotalsResponse {
	return v.value
}

func (v *NullableGetExistingReceivedDocumentTotalsResponse) Set(val *GetExistingReceivedDocumentTotalsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExistingReceivedDocumentTotalsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExistingReceivedDocumentTotalsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExistingReceivedDocumentTotalsResponse(val *GetExistingReceivedDocumentTotalsResponse) *NullableGetExistingReceivedDocumentTotalsResponse {
	return &NullableGetExistingReceivedDocumentTotalsResponse{value: val, isSet: true}
}

func (v NullableGetExistingReceivedDocumentTotalsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExistingReceivedDocumentTotalsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


