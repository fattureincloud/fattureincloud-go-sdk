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

// checks if the ListReceivedDocumentsResponsePage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListReceivedDocumentsResponsePage{}

// ListReceivedDocumentsResponsePage struct for ListReceivedDocumentsResponsePage
type ListReceivedDocumentsResponsePage struct {
	Data []ReceivedDocument `json:"data,omitempty"`
}

// NewListReceivedDocumentsResponsePage instantiates a new ListReceivedDocumentsResponsePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListReceivedDocumentsResponsePage() *ListReceivedDocumentsResponsePage {
	this := ListReceivedDocumentsResponsePage{}
	return &this
}

// NewListReceivedDocumentsResponsePageWithDefaults instantiates a new ListReceivedDocumentsResponsePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListReceivedDocumentsResponsePageWithDefaults() *ListReceivedDocumentsResponsePage {
	this := ListReceivedDocumentsResponsePage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListReceivedDocumentsResponsePage) GetData() []ReceivedDocument {
	if o == nil {
		var ret []ReceivedDocument
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListReceivedDocumentsResponsePage) GetDataOk() ([]ReceivedDocument, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListReceivedDocumentsResponsePage) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ReceivedDocument and assigns it to the Data field.
func (o *ListReceivedDocumentsResponsePage) SetData(v []ReceivedDocument) *ListReceivedDocumentsResponsePage {
	o.Data = v
	return o
}

func (o ListReceivedDocumentsResponsePage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListReceivedDocumentsResponsePage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListReceivedDocumentsResponsePage struct {
	value *ListReceivedDocumentsResponsePage
	isSet bool
}

func (v NullableListReceivedDocumentsResponsePage) Get() *ListReceivedDocumentsResponsePage {
	return v.value
}

func (v *NullableListReceivedDocumentsResponsePage) Set(val *ListReceivedDocumentsResponsePage) {
	v.value = val
	v.isSet = true
}

func (v NullableListReceivedDocumentsResponsePage) IsSet() bool {
	return v.isSet
}

func (v *NullableListReceivedDocumentsResponsePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListReceivedDocumentsResponsePage(val *ListReceivedDocumentsResponsePage) *NullableListReceivedDocumentsResponsePage {
	return &NullableListReceivedDocumentsResponsePage{value: val, isSet: true}
}

func (v NullableListReceivedDocumentsResponsePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListReceivedDocumentsResponsePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


