/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.22
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// GetNewReceivedDocumentTotalsResponse 
type GetNewReceivedDocumentTotalsResponse struct {
	Data *ReceivedDocumentTotals `json:"data,omitempty"`
}

// NewGetNewReceivedDocumentTotalsResponse instantiates a new GetNewReceivedDocumentTotalsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNewReceivedDocumentTotalsResponse() *GetNewReceivedDocumentTotalsResponse {
	this := GetNewReceivedDocumentTotalsResponse{}
	return &this
}

// NewGetNewReceivedDocumentTotalsResponseWithDefaults instantiates a new GetNewReceivedDocumentTotalsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNewReceivedDocumentTotalsResponseWithDefaults() *GetNewReceivedDocumentTotalsResponse {
	this := GetNewReceivedDocumentTotalsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetNewReceivedDocumentTotalsResponse) GetData() ReceivedDocumentTotals {
	if o == nil || o.Data == nil {
		var ret ReceivedDocumentTotals
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNewReceivedDocumentTotalsResponse) GetDataOk() (*ReceivedDocumentTotals, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetNewReceivedDocumentTotalsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ReceivedDocumentTotals and assigns it to the Data field.
func (o *GetNewReceivedDocumentTotalsResponse) SetData(v ReceivedDocumentTotals) *GetNewReceivedDocumentTotalsResponse {
	o.Data = &v
	return o
}

func (o GetNewReceivedDocumentTotalsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetNewReceivedDocumentTotalsResponse struct {
	value *GetNewReceivedDocumentTotalsResponse
	isSet bool
}

func (v NullableGetNewReceivedDocumentTotalsResponse) Get() *GetNewReceivedDocumentTotalsResponse {
	return v.value
}

func (v *NullableGetNewReceivedDocumentTotalsResponse) Set(val *GetNewReceivedDocumentTotalsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNewReceivedDocumentTotalsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNewReceivedDocumentTotalsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNewReceivedDocumentTotalsResponse(val *GetNewReceivedDocumentTotalsResponse) *NullableGetNewReceivedDocumentTotalsResponse {
	return &NullableGetNewReceivedDocumentTotalsResponse{value: val, isSet: true}
}

func (v NullableGetNewReceivedDocumentTotalsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNewReceivedDocumentTotalsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


