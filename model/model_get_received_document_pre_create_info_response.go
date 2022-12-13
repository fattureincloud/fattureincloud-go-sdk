/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.24
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// GetReceivedDocumentPreCreateInfoResponse 
type GetReceivedDocumentPreCreateInfoResponse struct {
	Data *ReceivedDocumentInfo `json:"data,omitempty"`
}

// NewGetReceivedDocumentPreCreateInfoResponse instantiates a new GetReceivedDocumentPreCreateInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReceivedDocumentPreCreateInfoResponse() *GetReceivedDocumentPreCreateInfoResponse {
	this := GetReceivedDocumentPreCreateInfoResponse{}
	return &this
}

// NewGetReceivedDocumentPreCreateInfoResponseWithDefaults instantiates a new GetReceivedDocumentPreCreateInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReceivedDocumentPreCreateInfoResponseWithDefaults() *GetReceivedDocumentPreCreateInfoResponse {
	this := GetReceivedDocumentPreCreateInfoResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetReceivedDocumentPreCreateInfoResponse) GetData() ReceivedDocumentInfo {
	if o == nil || isNil(o.Data) {
		var ret ReceivedDocumentInfo
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReceivedDocumentPreCreateInfoResponse) GetDataOk() (*ReceivedDocumentInfo, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetReceivedDocumentPreCreateInfoResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ReceivedDocumentInfo and assigns it to the Data field.
func (o *GetReceivedDocumentPreCreateInfoResponse) SetData(v ReceivedDocumentInfo) *GetReceivedDocumentPreCreateInfoResponse {
	o.Data = &v
	return o
}

func (o GetReceivedDocumentPreCreateInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetReceivedDocumentPreCreateInfoResponse struct {
	value *GetReceivedDocumentPreCreateInfoResponse
	isSet bool
}

func (v NullableGetReceivedDocumentPreCreateInfoResponse) Get() *GetReceivedDocumentPreCreateInfoResponse {
	return v.value
}

func (v *NullableGetReceivedDocumentPreCreateInfoResponse) Set(val *GetReceivedDocumentPreCreateInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReceivedDocumentPreCreateInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReceivedDocumentPreCreateInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReceivedDocumentPreCreateInfoResponse(val *GetReceivedDocumentPreCreateInfoResponse) *NullableGetReceivedDocumentPreCreateInfoResponse {
	return &NullableGetReceivedDocumentPreCreateInfoResponse{value: val, isSet: true}
}

func (v NullableGetReceivedDocumentPreCreateInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReceivedDocumentPreCreateInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


