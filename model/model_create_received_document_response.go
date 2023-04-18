/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.27
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the CreateReceivedDocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateReceivedDocumentResponse{}

// CreateReceivedDocumentResponse 
type CreateReceivedDocumentResponse struct {
	Data *ReceivedDocument `json:"data,omitempty"`
}

// NewCreateReceivedDocumentResponse instantiates a new CreateReceivedDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReceivedDocumentResponse() *CreateReceivedDocumentResponse {
	this := CreateReceivedDocumentResponse{}
	return &this
}

// NewCreateReceivedDocumentResponseWithDefaults instantiates a new CreateReceivedDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReceivedDocumentResponseWithDefaults() *CreateReceivedDocumentResponse {
	this := CreateReceivedDocumentResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateReceivedDocumentResponse) GetData() ReceivedDocument {
	if o == nil || IsNil(o.Data) {
		var ret ReceivedDocument
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateReceivedDocumentResponse) GetDataOk() (*ReceivedDocument, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateReceivedDocumentResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ReceivedDocument and assigns it to the Data field.
func (o *CreateReceivedDocumentResponse) SetData(v ReceivedDocument) *CreateReceivedDocumentResponse {
	o.Data = &v
	return o
}

func (o CreateReceivedDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateReceivedDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateReceivedDocumentResponse struct {
	value *CreateReceivedDocumentResponse
	isSet bool
}

func (v NullableCreateReceivedDocumentResponse) Get() *CreateReceivedDocumentResponse {
	return v.value
}

func (v *NullableCreateReceivedDocumentResponse) Set(val *CreateReceivedDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReceivedDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReceivedDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReceivedDocumentResponse(val *CreateReceivedDocumentResponse) *NullableCreateReceivedDocumentResponse {
	return &NullableCreateReceivedDocumentResponse{value: val, isSet: true}
}

func (v NullableCreateReceivedDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReceivedDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


