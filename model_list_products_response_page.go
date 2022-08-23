/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fattureincloud

import (
	"encoding/json"
)

// ListProductsResponsePage struct for ListProductsResponsePage
type ListProductsResponsePage struct {
	Data []Product `json:"data,omitempty"`
}

// NewListProductsResponsePage instantiates a new ListProductsResponsePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListProductsResponsePage() *ListProductsResponsePage {
	this := ListProductsResponsePage{}
	return &this
}

// NewListProductsResponsePageWithDefaults instantiates a new ListProductsResponsePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListProductsResponsePageWithDefaults() *ListProductsResponsePage {
	this := ListProductsResponsePage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListProductsResponsePage) GetData() []Product {
	if o == nil {
		var ret []Product
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListProductsResponsePage) GetDataOk() ([]Product, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListProductsResponsePage) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Product and assigns it to the Data field.
func (o *ListProductsResponsePage) SetData(v []Product) *ListProductsResponsePage {
	o.Data = v
	return o
}

func (o ListProductsResponsePage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListProductsResponsePage struct {
	value *ListProductsResponsePage
	isSet bool
}

func (v NullableListProductsResponsePage) Get() *ListProductsResponsePage {
	return v.value
}

func (v *NullableListProductsResponsePage) Set(val *ListProductsResponsePage) {
	v.value = val
	v.isSet = true
}

func (v NullableListProductsResponsePage) IsSet() bool {
	return v.isSet
}

func (v *NullableListProductsResponsePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListProductsResponsePage(val *ListProductsResponsePage) *NullableListProductsResponsePage {
	return &NullableListProductsResponsePage{value: val, isSet: true}
}

func (v NullableListProductsResponsePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListProductsResponsePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


