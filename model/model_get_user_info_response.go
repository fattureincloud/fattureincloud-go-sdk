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

// checks if the GetUserInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserInfoResponse{}

// GetUserInfoResponse 
type GetUserInfoResponse struct {
	Data *User `json:"data,omitempty"`
	Info NullableGetUserInfoResponseInfo `json:"info,omitempty"`
	EmailConfirmationState NullableGetUserInfoResponseEmailConfirmationState `json:"email_confirmation_state,omitempty"`
}

// NewGetUserInfoResponse instantiates a new GetUserInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserInfoResponse() *GetUserInfoResponse {
	this := GetUserInfoResponse{}
	return &this
}

// NewGetUserInfoResponseWithDefaults instantiates a new GetUserInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserInfoResponseWithDefaults() *GetUserInfoResponse {
	this := GetUserInfoResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetUserInfoResponse) GetData() User {
	if o == nil || IsNil(o.Data) {
		var ret User
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserInfoResponse) GetDataOk() (*User, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetUserInfoResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given User and assigns it to the Data field.
func (o *GetUserInfoResponse) SetData(v User) *GetUserInfoResponse {
	o.Data = &v
	return o
}

// GetInfo returns the Info field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUserInfoResponse) GetInfo() GetUserInfoResponseInfo {
	if o == nil || IsNil(o.Info.Get()) {
		var ret GetUserInfoResponseInfo
		return ret
	}
	return *o.Info.Get()
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUserInfoResponse) GetInfoOk() (*GetUserInfoResponseInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Info.Get(), o.Info.IsSet()
}

// HasInfo returns a boolean if a field has been set.
func (o *GetUserInfoResponse) HasInfo() bool {
	if o != nil && o.Info.IsSet() {
		return true
	}

	return false
}

// SetInfo gets a reference to the given NullableGetUserInfoResponseInfo and assigns it to the Info field.
func (o *GetUserInfoResponse) SetInfo(v GetUserInfoResponseInfo) *GetUserInfoResponse {
	o.Info.Set(&v)
	return o
}
// SetInfoNil sets the value for Info to be an explicit nil
func (o *GetUserInfoResponse) SetInfoNil() *GetUserInfoResponse {
	o.Info.Set(nil)
	return o
}

// UnsetInfo ensures that no value is present for Info, not even an explicit nil
func (o *GetUserInfoResponse) UnsetInfo() {
	o.Info.Unset()
}

// GetEmailConfirmationState returns the EmailConfirmationState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUserInfoResponse) GetEmailConfirmationState() GetUserInfoResponseEmailConfirmationState {
	if o == nil || IsNil(o.EmailConfirmationState.Get()) {
		var ret GetUserInfoResponseEmailConfirmationState
		return ret
	}
	return *o.EmailConfirmationState.Get()
}

// GetEmailConfirmationStateOk returns a tuple with the EmailConfirmationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUserInfoResponse) GetEmailConfirmationStateOk() (*GetUserInfoResponseEmailConfirmationState, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailConfirmationState.Get(), o.EmailConfirmationState.IsSet()
}

// HasEmailConfirmationState returns a boolean if a field has been set.
func (o *GetUserInfoResponse) HasEmailConfirmationState() bool {
	if o != nil && o.EmailConfirmationState.IsSet() {
		return true
	}

	return false
}

// SetEmailConfirmationState gets a reference to the given NullableGetUserInfoResponseEmailConfirmationState and assigns it to the EmailConfirmationState field.
func (o *GetUserInfoResponse) SetEmailConfirmationState(v GetUserInfoResponseEmailConfirmationState) *GetUserInfoResponse {
	o.EmailConfirmationState.Set(&v)
	return o
}
// SetEmailConfirmationStateNil sets the value for EmailConfirmationState to be an explicit nil
func (o *GetUserInfoResponse) SetEmailConfirmationStateNil() *GetUserInfoResponse {
	o.EmailConfirmationState.Set(nil)
	return o
}

// UnsetEmailConfirmationState ensures that no value is present for EmailConfirmationState, not even an explicit nil
func (o *GetUserInfoResponse) UnsetEmailConfirmationState() {
	o.EmailConfirmationState.Unset()
}

func (o GetUserInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Info.IsSet() {
		toSerialize["info"] = o.Info.Get()
	}
	if o.EmailConfirmationState.IsSet() {
		toSerialize["email_confirmation_state"] = o.EmailConfirmationState.Get()
	}
	return toSerialize, nil
}

type NullableGetUserInfoResponse struct {
	value *GetUserInfoResponse
	isSet bool
}

func (v NullableGetUserInfoResponse) Get() *GetUserInfoResponse {
	return v.value
}

func (v *NullableGetUserInfoResponse) Set(val *GetUserInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserInfoResponse(val *GetUserInfoResponse) *NullableGetUserInfoResponse {
	return &NullableGetUserInfoResponse{value: val, isSet: true}
}

func (v NullableGetUserInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


