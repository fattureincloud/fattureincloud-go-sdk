/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.20
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// GetCashbookEntryResponse 
type GetCashbookEntryResponse struct {
	Data *CashbookEntry `json:"data,omitempty"`
}

// NewGetCashbookEntryResponse instantiates a new GetCashbookEntryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCashbookEntryResponse() *GetCashbookEntryResponse {
	this := GetCashbookEntryResponse{}
	return &this
}

// NewGetCashbookEntryResponseWithDefaults instantiates a new GetCashbookEntryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCashbookEntryResponseWithDefaults() *GetCashbookEntryResponse {
	this := GetCashbookEntryResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetCashbookEntryResponse) GetData() CashbookEntry {
	if o == nil || o.Data == nil {
		var ret CashbookEntry
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCashbookEntryResponse) GetDataOk() (*CashbookEntry, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetCashbookEntryResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CashbookEntry and assigns it to the Data field.
func (o *GetCashbookEntryResponse) SetData(v CashbookEntry) *GetCashbookEntryResponse {
	o.Data = &v
	return o
}

func (o GetCashbookEntryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetCashbookEntryResponse struct {
	value *GetCashbookEntryResponse
	isSet bool
}

func (v NullableGetCashbookEntryResponse) Get() *GetCashbookEntryResponse {
	return v.value
}

func (v *NullableGetCashbookEntryResponse) Set(val *GetCashbookEntryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCashbookEntryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCashbookEntryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCashbookEntryResponse(val *GetCashbookEntryResponse) *NullableGetCashbookEntryResponse {
	return &NullableGetCashbookEntryResponse{value: val, isSet: true}
}

func (v NullableGetCashbookEntryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCashbookEntryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


