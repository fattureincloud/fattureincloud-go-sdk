/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.28
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ListReceivedDocumentCategoriesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListReceivedDocumentCategoriesResponse{}

// ListReceivedDocumentCategoriesResponse 
type ListReceivedDocumentCategoriesResponse struct {
	Data []string `json:"data,omitempty"`
}

// NewListReceivedDocumentCategoriesResponse instantiates a new ListReceivedDocumentCategoriesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListReceivedDocumentCategoriesResponse() *ListReceivedDocumentCategoriesResponse {
	this := ListReceivedDocumentCategoriesResponse{}
	return &this
}

// NewListReceivedDocumentCategoriesResponseWithDefaults instantiates a new ListReceivedDocumentCategoriesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListReceivedDocumentCategoriesResponseWithDefaults() *ListReceivedDocumentCategoriesResponse {
	this := ListReceivedDocumentCategoriesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListReceivedDocumentCategoriesResponse) GetData() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListReceivedDocumentCategoriesResponse) GetDataOk() ([]string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListReceivedDocumentCategoriesResponse) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *ListReceivedDocumentCategoriesResponse) SetData(v []string) *ListReceivedDocumentCategoriesResponse {
	o.Data = v
	return o
}

func (o ListReceivedDocumentCategoriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListReceivedDocumentCategoriesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListReceivedDocumentCategoriesResponse struct {
	value *ListReceivedDocumentCategoriesResponse
	isSet bool
}

func (v NullableListReceivedDocumentCategoriesResponse) Get() *ListReceivedDocumentCategoriesResponse {
	return v.value
}

func (v *NullableListReceivedDocumentCategoriesResponse) Set(val *ListReceivedDocumentCategoriesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListReceivedDocumentCategoriesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListReceivedDocumentCategoriesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListReceivedDocumentCategoriesResponse(val *ListReceivedDocumentCategoriesResponse) *NullableListReceivedDocumentCategoriesResponse {
	return &NullableListReceivedDocumentCategoriesResponse{value: val, isSet: true}
}

func (v NullableListReceivedDocumentCategoriesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListReceivedDocumentCategoriesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


