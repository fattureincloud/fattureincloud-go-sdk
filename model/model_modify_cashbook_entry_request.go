/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.31
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ModifyCashbookEntryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyCashbookEntryRequest{}

// ModifyCashbookEntryRequest 
type ModifyCashbookEntryRequest struct {
	Data *CashbookEntry `json:"data,omitempty"`
}

// NewModifyCashbookEntryRequest instantiates a new ModifyCashbookEntryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyCashbookEntryRequest() *ModifyCashbookEntryRequest {
	this := ModifyCashbookEntryRequest{}
	return &this
}

// NewModifyCashbookEntryRequestWithDefaults instantiates a new ModifyCashbookEntryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyCashbookEntryRequestWithDefaults() *ModifyCashbookEntryRequest {
	this := ModifyCashbookEntryRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModifyCashbookEntryRequest) GetData() CashbookEntry {
	if o == nil || IsNil(o.Data) {
		var ret CashbookEntry
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyCashbookEntryRequest) GetDataOk() (*CashbookEntry, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModifyCashbookEntryRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CashbookEntry and assigns it to the Data field.
func (o *ModifyCashbookEntryRequest) SetData(v CashbookEntry) *ModifyCashbookEntryRequest {
	o.Data = &v
	return o
}

func (o ModifyCashbookEntryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyCashbookEntryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableModifyCashbookEntryRequest struct {
	value *ModifyCashbookEntryRequest
	isSet bool
}

func (v NullableModifyCashbookEntryRequest) Get() *ModifyCashbookEntryRequest {
	return v.value
}

func (v *NullableModifyCashbookEntryRequest) Set(val *ModifyCashbookEntryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyCashbookEntryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyCashbookEntryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyCashbookEntryRequest(val *ModifyCashbookEntryRequest) *NullableModifyCashbookEntryRequest {
	return &NullableModifyCashbookEntryRequest{value: val, isSet: true}
}

func (v NullableModifyCashbookEntryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyCashbookEntryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


