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

// checks if the IssuedDocumentOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssuedDocumentOptions{}

// IssuedDocumentOptions struct for IssuedDocumentOptions
type IssuedDocumentOptions struct {
	// Fixes your last payment amount to match your document total
	FixPayments NullableBool `json:"fix_payments,omitempty"`
	// Original documents ids [only for join/transform]
	CreateFrom []string `json:"create_from,omitempty"`
	// Tranform a document [only for transform]
	Transform NullableBool `json:"transform,omitempty"`
	// Keep original document [only for transform]
	KeepCopy NullableBool `json:"keep_copy,omitempty"`
	// Join type [only for join]
	JoinType NullableString `json:"join_type,omitempty"`
}

// NewIssuedDocumentOptions instantiates a new IssuedDocumentOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuedDocumentOptions() *IssuedDocumentOptions {
	this := IssuedDocumentOptions{}
	return &this
}

// NewIssuedDocumentOptionsWithDefaults instantiates a new IssuedDocumentOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuedDocumentOptionsWithDefaults() *IssuedDocumentOptions {
	this := IssuedDocumentOptions{}
	return &this
}

// GetFixPayments returns the FixPayments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentOptions) GetFixPayments() bool {
	if o == nil || IsNil(o.FixPayments.Get()) {
		var ret bool
		return ret
	}
	return *o.FixPayments.Get()
}

// GetFixPaymentsOk returns a tuple with the FixPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentOptions) GetFixPaymentsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FixPayments.Get(), o.FixPayments.IsSet()
}

// HasFixPayments returns a boolean if a field has been set.
func (o *IssuedDocumentOptions) HasFixPayments() bool {
	if o != nil && o.FixPayments.IsSet() {
		return true
	}

	return false
}

// SetFixPayments gets a reference to the given NullableBool and assigns it to the FixPayments field.
func (o *IssuedDocumentOptions) SetFixPayments(v bool) *IssuedDocumentOptions {
	o.FixPayments.Set(&v)
	return o
}
// SetFixPaymentsNil sets the value for FixPayments to be an explicit nil
func (o *IssuedDocumentOptions) SetFixPaymentsNil() *IssuedDocumentOptions {
	o.FixPayments.Set(nil)
	return o
}

// UnsetFixPayments ensures that no value is present for FixPayments, not even an explicit nil
func (o *IssuedDocumentOptions) UnsetFixPayments() {
	o.FixPayments.Unset()
}

// GetCreateFrom returns the CreateFrom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentOptions) GetCreateFrom() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CreateFrom
}

// GetCreateFromOk returns a tuple with the CreateFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentOptions) GetCreateFromOk() ([]string, bool) {
	if o == nil || IsNil(o.CreateFrom) {
		return nil, false
	}
	return o.CreateFrom, true
}

// HasCreateFrom returns a boolean if a field has been set.
func (o *IssuedDocumentOptions) HasCreateFrom() bool {
	if o != nil && !IsNil(o.CreateFrom) {
		return true
	}

	return false
}

// SetCreateFrom gets a reference to the given []string and assigns it to the CreateFrom field.
func (o *IssuedDocumentOptions) SetCreateFrom(v []string) *IssuedDocumentOptions {
	o.CreateFrom = v
	return o
}

// GetTransform returns the Transform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentOptions) GetTransform() bool {
	if o == nil || IsNil(o.Transform.Get()) {
		var ret bool
		return ret
	}
	return *o.Transform.Get()
}

// GetTransformOk returns a tuple with the Transform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentOptions) GetTransformOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transform.Get(), o.Transform.IsSet()
}

// HasTransform returns a boolean if a field has been set.
func (o *IssuedDocumentOptions) HasTransform() bool {
	if o != nil && o.Transform.IsSet() {
		return true
	}

	return false
}

// SetTransform gets a reference to the given NullableBool and assigns it to the Transform field.
func (o *IssuedDocumentOptions) SetTransform(v bool) *IssuedDocumentOptions {
	o.Transform.Set(&v)
	return o
}
// SetTransformNil sets the value for Transform to be an explicit nil
func (o *IssuedDocumentOptions) SetTransformNil() *IssuedDocumentOptions {
	o.Transform.Set(nil)
	return o
}

// UnsetTransform ensures that no value is present for Transform, not even an explicit nil
func (o *IssuedDocumentOptions) UnsetTransform() {
	o.Transform.Unset()
}

// GetKeepCopy returns the KeepCopy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentOptions) GetKeepCopy() bool {
	if o == nil || IsNil(o.KeepCopy.Get()) {
		var ret bool
		return ret
	}
	return *o.KeepCopy.Get()
}

// GetKeepCopyOk returns a tuple with the KeepCopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentOptions) GetKeepCopyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeepCopy.Get(), o.KeepCopy.IsSet()
}

// HasKeepCopy returns a boolean if a field has been set.
func (o *IssuedDocumentOptions) HasKeepCopy() bool {
	if o != nil && o.KeepCopy.IsSet() {
		return true
	}

	return false
}

// SetKeepCopy gets a reference to the given NullableBool and assigns it to the KeepCopy field.
func (o *IssuedDocumentOptions) SetKeepCopy(v bool) *IssuedDocumentOptions {
	o.KeepCopy.Set(&v)
	return o
}
// SetKeepCopyNil sets the value for KeepCopy to be an explicit nil
func (o *IssuedDocumentOptions) SetKeepCopyNil() *IssuedDocumentOptions {
	o.KeepCopy.Set(nil)
	return o
}

// UnsetKeepCopy ensures that no value is present for KeepCopy, not even an explicit nil
func (o *IssuedDocumentOptions) UnsetKeepCopy() {
	o.KeepCopy.Unset()
}

// GetJoinType returns the JoinType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentOptions) GetJoinType() string {
	if o == nil || IsNil(o.JoinType.Get()) {
		var ret string
		return ret
	}
	return *o.JoinType.Get()
}

// GetJoinTypeOk returns a tuple with the JoinType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentOptions) GetJoinTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JoinType.Get(), o.JoinType.IsSet()
}

// HasJoinType returns a boolean if a field has been set.
func (o *IssuedDocumentOptions) HasJoinType() bool {
	if o != nil && o.JoinType.IsSet() {
		return true
	}

	return false
}

// SetJoinType gets a reference to the given NullableString and assigns it to the JoinType field.
func (o *IssuedDocumentOptions) SetJoinType(v string) *IssuedDocumentOptions {
	o.JoinType.Set(&v)
	return o
}
// SetJoinTypeNil sets the value for JoinType to be an explicit nil
func (o *IssuedDocumentOptions) SetJoinTypeNil() *IssuedDocumentOptions {
	o.JoinType.Set(nil)
	return o
}

// UnsetJoinType ensures that no value is present for JoinType, not even an explicit nil
func (o *IssuedDocumentOptions) UnsetJoinType() {
	o.JoinType.Unset()
}

func (o IssuedDocumentOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssuedDocumentOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FixPayments.IsSet() {
		toSerialize["fix_payments"] = o.FixPayments.Get()
	}
	if o.CreateFrom != nil {
		toSerialize["create_from"] = o.CreateFrom
	}
	if o.Transform.IsSet() {
		toSerialize["transform"] = o.Transform.Get()
	}
	if o.KeepCopy.IsSet() {
		toSerialize["keep_copy"] = o.KeepCopy.Get()
	}
	if o.JoinType.IsSet() {
		toSerialize["join_type"] = o.JoinType.Get()
	}
	return toSerialize, nil
}

type NullableIssuedDocumentOptions struct {
	value *IssuedDocumentOptions
	isSet bool
}

func (v NullableIssuedDocumentOptions) Get() *IssuedDocumentOptions {
	return v.value
}

func (v *NullableIssuedDocumentOptions) Set(val *IssuedDocumentOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuedDocumentOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuedDocumentOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuedDocumentOptions(val *IssuedDocumentOptions) *NullableIssuedDocumentOptions {
	return &NullableIssuedDocumentOptions{value: val, isSet: true}
}

func (v NullableIssuedDocumentOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuedDocumentOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


