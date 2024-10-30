/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.3
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ListTemplatesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTemplatesResponse{}

// ListTemplatesResponse 
type ListTemplatesResponse struct {
Data []DocumentTemplate `json:"data,omitempty"`
}

// NewListTemplatesResponse instantiates a new ListTemplatesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTemplatesResponse() *ListTemplatesResponse {
	this := ListTemplatesResponse{}
	return &this
}

// NewListTemplatesResponseWithDefaults instantiates a new ListTemplatesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTemplatesResponseWithDefaults() *ListTemplatesResponse {
	this := ListTemplatesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListTemplatesResponse) GetData() []DocumentTemplate {
	if o == nil {
		var ret []DocumentTemplate
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListTemplatesResponse) GetDataOk() ([]DocumentTemplate, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListTemplatesResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []DocumentTemplate and assigns it to the Data field.
func (o *ListTemplatesResponse) SetData(v []DocumentTemplate) *ListTemplatesResponse {
	o.Data = v
        return o
}

func (o ListTemplatesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTemplatesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListTemplatesResponse struct {
	value *ListTemplatesResponse
	isSet bool
}

func (v NullableListTemplatesResponse) Get() *ListTemplatesResponse {
	return v.value
}

func (v *NullableListTemplatesResponse) Set(val *ListTemplatesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListTemplatesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListTemplatesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTemplatesResponse(val *ListTemplatesResponse) *NullableListTemplatesResponse {
	return &NullableListTemplatesResponse{value: val, isSet: true}
}

func (v NullableListTemplatesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTemplatesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


