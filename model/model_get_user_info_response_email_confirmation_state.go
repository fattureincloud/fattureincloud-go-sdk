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

// checks if the GetUserInfoResponseEmailConfirmationState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserInfoResponseEmailConfirmationState{}

// GetUserInfoResponseEmailConfirmationState struct for GetUserInfoResponseEmailConfirmationState
type GetUserInfoResponseEmailConfirmationState struct {
	NeedConfirmation NullableBool `json:"need_confirmation,omitempty"`
}

// NewGetUserInfoResponseEmailConfirmationState instantiates a new GetUserInfoResponseEmailConfirmationState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserInfoResponseEmailConfirmationState() *GetUserInfoResponseEmailConfirmationState {
	this := GetUserInfoResponseEmailConfirmationState{}
	return &this
}

// NewGetUserInfoResponseEmailConfirmationStateWithDefaults instantiates a new GetUserInfoResponseEmailConfirmationState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserInfoResponseEmailConfirmationStateWithDefaults() *GetUserInfoResponseEmailConfirmationState {
	this := GetUserInfoResponseEmailConfirmationState{}
	return &this
}

// GetNeedConfirmation returns the NeedConfirmation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUserInfoResponseEmailConfirmationState) GetNeedConfirmation() bool {
	if o == nil || IsNil(o.NeedConfirmation.Get()) {
		var ret bool
		return ret
	}
	return *o.NeedConfirmation.Get()
}

// GetNeedConfirmationOk returns a tuple with the NeedConfirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUserInfoResponseEmailConfirmationState) GetNeedConfirmationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NeedConfirmation.Get(), o.NeedConfirmation.IsSet()
}

// HasNeedConfirmation returns a boolean if a field has been set.
func (o *GetUserInfoResponseEmailConfirmationState) HasNeedConfirmation() bool {
	if o != nil && o.NeedConfirmation.IsSet() {
		return true
	}

	return false
}

// SetNeedConfirmation gets a reference to the given NullableBool and assigns it to the NeedConfirmation field.
func (o *GetUserInfoResponseEmailConfirmationState) SetNeedConfirmation(v bool) *GetUserInfoResponseEmailConfirmationState {
	o.NeedConfirmation.Set(&v)
	return o
}
// SetNeedConfirmationNil sets the value for NeedConfirmation to be an explicit nil
func (o *GetUserInfoResponseEmailConfirmationState) SetNeedConfirmationNil() *GetUserInfoResponseEmailConfirmationState {
	o.NeedConfirmation.Set(nil)
	return o
}

// UnsetNeedConfirmation ensures that no value is present for NeedConfirmation, not even an explicit nil
func (o *GetUserInfoResponseEmailConfirmationState) UnsetNeedConfirmation() {
	o.NeedConfirmation.Unset()
}

func (o GetUserInfoResponseEmailConfirmationState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserInfoResponseEmailConfirmationState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NeedConfirmation.IsSet() {
		toSerialize["need_confirmation"] = o.NeedConfirmation.Get()
	}
	return toSerialize, nil
}

type NullableGetUserInfoResponseEmailConfirmationState struct {
	value *GetUserInfoResponseEmailConfirmationState
	isSet bool
}

func (v NullableGetUserInfoResponseEmailConfirmationState) Get() *GetUserInfoResponseEmailConfirmationState {
	return v.value
}

func (v *NullableGetUserInfoResponseEmailConfirmationState) Set(val *GetUserInfoResponseEmailConfirmationState) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserInfoResponseEmailConfirmationState) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserInfoResponseEmailConfirmationState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserInfoResponseEmailConfirmationState(val *GetUserInfoResponseEmailConfirmationState) *NullableGetUserInfoResponseEmailConfirmationState {
	return &NullableGetUserInfoResponseEmailConfirmationState{value: val, isSet: true}
}

func (v NullableGetUserInfoResponseEmailConfirmationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserInfoResponseEmailConfirmationState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


