/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.29
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ListIssuedDocumentsResponsePage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListIssuedDocumentsResponsePage{}

// ListIssuedDocumentsResponsePage struct for ListIssuedDocumentsResponsePage
type ListIssuedDocumentsResponsePage struct {
	Data []IssuedDocument `json:"data,omitempty"`
}

// NewListIssuedDocumentsResponsePage instantiates a new ListIssuedDocumentsResponsePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListIssuedDocumentsResponsePage() *ListIssuedDocumentsResponsePage {
	this := ListIssuedDocumentsResponsePage{}
	return &this
}

// NewListIssuedDocumentsResponsePageWithDefaults instantiates a new ListIssuedDocumentsResponsePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIssuedDocumentsResponsePageWithDefaults() *ListIssuedDocumentsResponsePage {
	this := ListIssuedDocumentsResponsePage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListIssuedDocumentsResponsePage) GetData() []IssuedDocument {
	if o == nil {
		var ret []IssuedDocument
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListIssuedDocumentsResponsePage) GetDataOk() ([]IssuedDocument, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListIssuedDocumentsResponsePage) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []IssuedDocument and assigns it to the Data field.
func (o *ListIssuedDocumentsResponsePage) SetData(v []IssuedDocument) *ListIssuedDocumentsResponsePage {
	o.Data = v
	return o
}

func (o ListIssuedDocumentsResponsePage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListIssuedDocumentsResponsePage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListIssuedDocumentsResponsePage struct {
	value *ListIssuedDocumentsResponsePage
	isSet bool
}

func (v NullableListIssuedDocumentsResponsePage) Get() *ListIssuedDocumentsResponsePage {
	return v.value
}

func (v *NullableListIssuedDocumentsResponsePage) Set(val *ListIssuedDocumentsResponsePage) {
	v.value = val
	v.isSet = true
}

func (v NullableListIssuedDocumentsResponsePage) IsSet() bool {
	return v.isSet
}

func (v *NullableListIssuedDocumentsResponsePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListIssuedDocumentsResponsePage(val *ListIssuedDocumentsResponsePage) *NullableListIssuedDocumentsResponsePage {
	return &NullableListIssuedDocumentsResponsePage{value: val, isSet: true}
}

func (v NullableListIssuedDocumentsResponsePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListIssuedDocumentsResponsePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


