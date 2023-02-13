/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.26
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the GetNewIssuedDocumentTotalsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNewIssuedDocumentTotalsRequest{}

// GetNewIssuedDocumentTotalsRequest struct for GetNewIssuedDocumentTotalsRequest
type GetNewIssuedDocumentTotalsRequest struct {
	Data *IssuedDocument `json:"data,omitempty"`
}

// NewGetNewIssuedDocumentTotalsRequest instantiates a new GetNewIssuedDocumentTotalsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNewIssuedDocumentTotalsRequest() *GetNewIssuedDocumentTotalsRequest {
	this := GetNewIssuedDocumentTotalsRequest{}
	return &this
}

// NewGetNewIssuedDocumentTotalsRequestWithDefaults instantiates a new GetNewIssuedDocumentTotalsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNewIssuedDocumentTotalsRequestWithDefaults() *GetNewIssuedDocumentTotalsRequest {
	this := GetNewIssuedDocumentTotalsRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetNewIssuedDocumentTotalsRequest) GetData() IssuedDocument {
	if o == nil || isNil(o.Data) {
		var ret IssuedDocument
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNewIssuedDocumentTotalsRequest) GetDataOk() (*IssuedDocument, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetNewIssuedDocumentTotalsRequest) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given IssuedDocument and assigns it to the Data field.
func (o *GetNewIssuedDocumentTotalsRequest) SetData(v IssuedDocument) *GetNewIssuedDocumentTotalsRequest {
	o.Data = &v
	return o
}

func (o GetNewIssuedDocumentTotalsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNewIssuedDocumentTotalsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetNewIssuedDocumentTotalsRequest struct {
	value *GetNewIssuedDocumentTotalsRequest
	isSet bool
}

func (v NullableGetNewIssuedDocumentTotalsRequest) Get() *GetNewIssuedDocumentTotalsRequest {
	return v.value
}

func (v *NullableGetNewIssuedDocumentTotalsRequest) Set(val *GetNewIssuedDocumentTotalsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNewIssuedDocumentTotalsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNewIssuedDocumentTotalsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNewIssuedDocumentTotalsRequest(val *GetNewIssuedDocumentTotalsRequest) *NullableGetNewIssuedDocumentTotalsRequest {
	return &NullableGetNewIssuedDocumentTotalsRequest{value: val, isSet: true}
}

func (v NullableGetNewIssuedDocumentTotalsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNewIssuedDocumentTotalsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


