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

// checks if the DocumentTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentTemplate{}

// DocumentTemplate struct for DocumentTemplate
type DocumentTemplate struct {
	// Template id
	Id NullableInt32 `json:"id,omitempty"`
	// Template name
	Name NullableString `json:"name,omitempty"`
	// Template type
	Type NullableString `json:"type,omitempty"`
}

// NewDocumentTemplate instantiates a new DocumentTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentTemplate() *DocumentTemplate {
	this := DocumentTemplate{}
	return &this
}

// NewDocumentTemplateWithDefaults instantiates a new DocumentTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentTemplateWithDefaults() *DocumentTemplate {
	this := DocumentTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentTemplate) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentTemplate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DocumentTemplate) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *DocumentTemplate) SetId(v int32) *DocumentTemplate {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *DocumentTemplate) SetIdNil() *DocumentTemplate {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DocumentTemplate) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentTemplate) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DocumentTemplate) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DocumentTemplate) SetName(v string) *DocumentTemplate {
	o.Name.Set(&v)
	return o
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DocumentTemplate) SetNameNil() *DocumentTemplate {
	o.Name.Set(nil)
	return o
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DocumentTemplate) UnsetName() {
	o.Name.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentTemplate) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentTemplate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *DocumentTemplate) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *DocumentTemplate) SetType(v string) *DocumentTemplate {
	o.Type.Set(&v)
	return o
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *DocumentTemplate) SetTypeNil() *DocumentTemplate {
	o.Type.Set(nil)
	return o
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *DocumentTemplate) UnsetType() {
	o.Type.Unset()
}

func (o DocumentTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	return toSerialize, nil
}

type NullableDocumentTemplate struct {
	value *DocumentTemplate
	isSet bool
}

func (v NullableDocumentTemplate) Get() *DocumentTemplate {
	return v.value
}

func (v *NullableDocumentTemplate) Set(val *DocumentTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentTemplate(val *DocumentTemplate) *NullableDocumentTemplate {
	return &NullableDocumentTemplate{value: val, isSet: true}
}

func (v NullableDocumentTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


