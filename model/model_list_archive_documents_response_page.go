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

// ListArchiveDocumentsResponsePage struct for ListArchiveDocumentsResponsePage
type ListArchiveDocumentsResponsePage struct {
	Data []ArchiveDocument `json:"data,omitempty"`
}

// NewListArchiveDocumentsResponsePage instantiates a new ListArchiveDocumentsResponsePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListArchiveDocumentsResponsePage() *ListArchiveDocumentsResponsePage {
	this := ListArchiveDocumentsResponsePage{}
	return &this
}

// NewListArchiveDocumentsResponsePageWithDefaults instantiates a new ListArchiveDocumentsResponsePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListArchiveDocumentsResponsePageWithDefaults() *ListArchiveDocumentsResponsePage {
	this := ListArchiveDocumentsResponsePage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListArchiveDocumentsResponsePage) GetData() []ArchiveDocument {
	if o == nil {
		var ret []ArchiveDocument
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListArchiveDocumentsResponsePage) GetDataOk() ([]ArchiveDocument, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListArchiveDocumentsResponsePage) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ArchiveDocument and assigns it to the Data field.
func (o *ListArchiveDocumentsResponsePage) SetData(v []ArchiveDocument) *ListArchiveDocumentsResponsePage {
	o.Data = v
	return o
}

func (o ListArchiveDocumentsResponsePage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListArchiveDocumentsResponsePage struct {
	value *ListArchiveDocumentsResponsePage
	isSet bool
}

func (v NullableListArchiveDocumentsResponsePage) Get() *ListArchiveDocumentsResponsePage {
	return v.value
}

func (v *NullableListArchiveDocumentsResponsePage) Set(val *ListArchiveDocumentsResponsePage) {
	v.value = val
	v.isSet = true
}

func (v NullableListArchiveDocumentsResponsePage) IsSet() bool {
	return v.isSet
}

func (v *NullableListArchiveDocumentsResponsePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListArchiveDocumentsResponsePage(val *ListArchiveDocumentsResponsePage) *NullableListArchiveDocumentsResponsePage {
	return &NullableListArchiveDocumentsResponsePage{value: val, isSet: true}
}

func (v NullableListArchiveDocumentsResponsePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListArchiveDocumentsResponsePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


