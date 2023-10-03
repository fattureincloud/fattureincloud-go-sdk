/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.30
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ListSuppliersResponsePage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSuppliersResponsePage{}

// ListSuppliersResponsePage struct for ListSuppliersResponsePage
type ListSuppliersResponsePage struct {
	Data []Supplier `json:"data,omitempty"`
}

// NewListSuppliersResponsePage instantiates a new ListSuppliersResponsePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSuppliersResponsePage() *ListSuppliersResponsePage {
	this := ListSuppliersResponsePage{}
	return &this
}

// NewListSuppliersResponsePageWithDefaults instantiates a new ListSuppliersResponsePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSuppliersResponsePageWithDefaults() *ListSuppliersResponsePage {
	this := ListSuppliersResponsePage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListSuppliersResponsePage) GetData() []Supplier {
	if o == nil {
		var ret []Supplier
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListSuppliersResponsePage) GetDataOk() ([]Supplier, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListSuppliersResponsePage) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Supplier and assigns it to the Data field.
func (o *ListSuppliersResponsePage) SetData(v []Supplier) *ListSuppliersResponsePage {
	o.Data = v
	return o
}

func (o ListSuppliersResponsePage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSuppliersResponsePage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListSuppliersResponsePage struct {
	value *ListSuppliersResponsePage
	isSet bool
}

func (v NullableListSuppliersResponsePage) Get() *ListSuppliersResponsePage {
	return v.value
}

func (v *NullableListSuppliersResponsePage) Set(val *ListSuppliersResponsePage) {
	v.value = val
	v.isSet = true
}

func (v NullableListSuppliersResponsePage) IsSet() bool {
	return v.isSet
}

func (v *NullableListSuppliersResponsePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSuppliersResponsePage(val *ListSuppliersResponsePage) *NullableListSuppliersResponsePage {
	return &NullableListSuppliersResponsePage{value: val, isSet: true}
}

func (v NullableListSuppliersResponsePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSuppliersResponsePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


