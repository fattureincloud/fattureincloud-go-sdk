/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.21
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// ListCashbookEntriesResponse 
type ListCashbookEntriesResponse struct {
	Data []CashbookEntry `json:"data,omitempty"`
}

// NewListCashbookEntriesResponse instantiates a new ListCashbookEntriesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCashbookEntriesResponse() *ListCashbookEntriesResponse {
	this := ListCashbookEntriesResponse{}
	return &this
}

// NewListCashbookEntriesResponseWithDefaults instantiates a new ListCashbookEntriesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCashbookEntriesResponseWithDefaults() *ListCashbookEntriesResponse {
	this := ListCashbookEntriesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListCashbookEntriesResponse) GetData() []CashbookEntry {
	if o == nil {
		var ret []CashbookEntry
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListCashbookEntriesResponse) GetDataOk() ([]CashbookEntry, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListCashbookEntriesResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []CashbookEntry and assigns it to the Data field.
func (o *ListCashbookEntriesResponse) SetData(v []CashbookEntry) *ListCashbookEntriesResponse {
	o.Data = v
	return o
}

func (o ListCashbookEntriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListCashbookEntriesResponse struct {
	value *ListCashbookEntriesResponse
	isSet bool
}

func (v NullableListCashbookEntriesResponse) Get() *ListCashbookEntriesResponse {
	return v.value
}

func (v *NullableListCashbookEntriesResponse) Set(val *ListCashbookEntriesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCashbookEntriesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCashbookEntriesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCashbookEntriesResponse(val *ListCashbookEntriesResponse) *NullableListCashbookEntriesResponse {
	return &NullableListCashbookEntriesResponse{value: val, isSet: true}
}

func (v NullableListCashbookEntriesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCashbookEntriesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


