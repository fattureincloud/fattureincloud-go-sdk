/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.2
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the CreateIssuedDocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateIssuedDocumentRequest{}

// CreateIssuedDocumentRequest struct for CreateIssuedDocumentRequest
type CreateIssuedDocumentRequest struct {
	Data *IssuedDocument `json:"data,omitempty"`
	Options *IssuedDocumentOptions `json:"options,omitempty"`
}

// NewCreateIssuedDocumentRequest instantiates a new CreateIssuedDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIssuedDocumentRequest() *CreateIssuedDocumentRequest {
	this := CreateIssuedDocumentRequest{}
	return &this
}

// NewCreateIssuedDocumentRequestWithDefaults instantiates a new CreateIssuedDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIssuedDocumentRequestWithDefaults() *CreateIssuedDocumentRequest {
	this := CreateIssuedDocumentRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateIssuedDocumentRequest) GetData() IssuedDocument {
	if o == nil || IsNil(o.Data) {
		var ret IssuedDocument
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssuedDocumentRequest) GetDataOk() (*IssuedDocument, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateIssuedDocumentRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given IssuedDocument and assigns it to the Data field.
func (o *CreateIssuedDocumentRequest) SetData(v IssuedDocument) *CreateIssuedDocumentRequest {
	o.Data = &v
	return o
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CreateIssuedDocumentRequest) GetOptions() IssuedDocumentOptions {
	if o == nil || IsNil(o.Options) {
		var ret IssuedDocumentOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssuedDocumentRequest) GetOptionsOk() (*IssuedDocumentOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CreateIssuedDocumentRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given IssuedDocumentOptions and assigns it to the Options field.
func (o *CreateIssuedDocumentRequest) SetOptions(v IssuedDocumentOptions) *CreateIssuedDocumentRequest {
	o.Options = &v
	return o
}

func (o CreateIssuedDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIssuedDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableCreateIssuedDocumentRequest struct {
	value *CreateIssuedDocumentRequest
	isSet bool
}

func (v NullableCreateIssuedDocumentRequest) Get() *CreateIssuedDocumentRequest {
	return v.value
}

func (v *NullableCreateIssuedDocumentRequest) Set(val *CreateIssuedDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIssuedDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIssuedDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIssuedDocumentRequest(val *CreateIssuedDocumentRequest) *NullableCreateIssuedDocumentRequest {
	return &NullableCreateIssuedDocumentRequest{value: val, isSet: true}
}

func (v NullableCreateIssuedDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIssuedDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


