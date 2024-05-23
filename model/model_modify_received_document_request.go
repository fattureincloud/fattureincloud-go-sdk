/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.33
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ModifyReceivedDocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyReceivedDocumentRequest{}

// ModifyReceivedDocumentRequest 
type ModifyReceivedDocumentRequest struct {
	Data *ReceivedDocument `json:"data,omitempty"`
}

// NewModifyReceivedDocumentRequest instantiates a new ModifyReceivedDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyReceivedDocumentRequest() *ModifyReceivedDocumentRequest {
	this := ModifyReceivedDocumentRequest{}
	return &this
}

// NewModifyReceivedDocumentRequestWithDefaults instantiates a new ModifyReceivedDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyReceivedDocumentRequestWithDefaults() *ModifyReceivedDocumentRequest {
	this := ModifyReceivedDocumentRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModifyReceivedDocumentRequest) GetData() ReceivedDocument {
	if o == nil || IsNil(o.Data) {
		var ret ReceivedDocument
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyReceivedDocumentRequest) GetDataOk() (*ReceivedDocument, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModifyReceivedDocumentRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ReceivedDocument and assigns it to the Data field.
func (o *ModifyReceivedDocumentRequest) SetData(v ReceivedDocument) *ModifyReceivedDocumentRequest {
	o.Data = &v
	return o
}

func (o ModifyReceivedDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyReceivedDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableModifyReceivedDocumentRequest struct {
	value *ModifyReceivedDocumentRequest
	isSet bool
}

func (v NullableModifyReceivedDocumentRequest) Get() *ModifyReceivedDocumentRequest {
	return v.value
}

func (v *NullableModifyReceivedDocumentRequest) Set(val *ModifyReceivedDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyReceivedDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyReceivedDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyReceivedDocumentRequest(val *ModifyReceivedDocumentRequest) *NullableModifyReceivedDocumentRequest {
	return &NullableModifyReceivedDocumentRequest{value: val, isSet: true}
}

func (v NullableModifyReceivedDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyReceivedDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


