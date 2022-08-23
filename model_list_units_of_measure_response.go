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

// ListUnitsOfMeasureResponse 
type ListUnitsOfMeasureResponse struct {
	Data []string `json:"data,omitempty"`
}

// NewListUnitsOfMeasureResponse instantiates a new ListUnitsOfMeasureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUnitsOfMeasureResponse() *ListUnitsOfMeasureResponse {
	this := ListUnitsOfMeasureResponse{}
	return &this
}

// NewListUnitsOfMeasureResponseWithDefaults instantiates a new ListUnitsOfMeasureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUnitsOfMeasureResponseWithDefaults() *ListUnitsOfMeasureResponse {
	this := ListUnitsOfMeasureResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListUnitsOfMeasureResponse) GetData() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListUnitsOfMeasureResponse) GetDataOk() ([]string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListUnitsOfMeasureResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *ListUnitsOfMeasureResponse) SetData(v []string) *ListUnitsOfMeasureResponse {
	o.Data = v
	return o
}

func (o ListUnitsOfMeasureResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListUnitsOfMeasureResponse struct {
	value *ListUnitsOfMeasureResponse
	isSet bool
}

func (v NullableListUnitsOfMeasureResponse) Get() *ListUnitsOfMeasureResponse {
	return v.value
}

func (v *NullableListUnitsOfMeasureResponse) Set(val *ListUnitsOfMeasureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListUnitsOfMeasureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListUnitsOfMeasureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUnitsOfMeasureResponse(val *ListUnitsOfMeasureResponse) *NullableListUnitsOfMeasureResponse {
	return &NullableListUnitsOfMeasureResponse{value: val, isSet: true}
}

func (v NullableListUnitsOfMeasureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUnitsOfMeasureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


