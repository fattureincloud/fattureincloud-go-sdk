/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.21
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// GetUserInfoResponseInfo struct for GetUserInfoResponseInfo
type GetUserInfoResponseInfo struct {
	NeedMarketingConsentsConfirmation NullableBool `json:"need_marketing_consents_confirmation,omitempty"`
	NeedPasswordChange NullableBool `json:"need_password_change,omitempty"`
	NeedTermsOfServiceConfirmation NullableBool `json:"need_terms_of_service_confirmation,omitempty"`
}

// NewGetUserInfoResponseInfo instantiates a new GetUserInfoResponseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserInfoResponseInfo() *GetUserInfoResponseInfo {
	this := GetUserInfoResponseInfo{}
	return &this
}

// NewGetUserInfoResponseInfoWithDefaults instantiates a new GetUserInfoResponseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserInfoResponseInfoWithDefaults() *GetUserInfoResponseInfo {
	this := GetUserInfoResponseInfo{}
	return &this
}

// GetNeedMarketingConsentsConfirmation returns the NeedMarketingConsentsConfirmation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUserInfoResponseInfo) GetNeedMarketingConsentsConfirmation() bool {
	if o == nil || o.NeedMarketingConsentsConfirmation.Get() == nil {
		var ret bool
		return ret
	}
	return *o.NeedMarketingConsentsConfirmation.Get()
}

// GetNeedMarketingConsentsConfirmationOk returns a tuple with the NeedMarketingConsentsConfirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUserInfoResponseInfo) GetNeedMarketingConsentsConfirmationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NeedMarketingConsentsConfirmation.Get(), o.NeedMarketingConsentsConfirmation.IsSet()
}

// HasNeedMarketingConsentsConfirmation returns a boolean if a field has been set.
func (o *GetUserInfoResponseInfo) HasNeedMarketingConsentsConfirmation() bool {
	if o != nil && o.NeedMarketingConsentsConfirmation.IsSet() {
		return true
	}

	return false
}

// SetNeedMarketingConsentsConfirmation gets a reference to the given NullableBool and assigns it to the NeedMarketingConsentsConfirmation field.
func (o *GetUserInfoResponseInfo) SetNeedMarketingConsentsConfirmation(v bool) *GetUserInfoResponseInfo {
	o.NeedMarketingConsentsConfirmation.Set(&v)
	return o
}
// SetNeedMarketingConsentsConfirmationNil sets the value for NeedMarketingConsentsConfirmation to be an explicit nil
func (o *GetUserInfoResponseInfo) SetNeedMarketingConsentsConfirmationNil() *GetUserInfoResponseInfo {
	o.NeedMarketingConsentsConfirmation.Set(nil)
	return o
}

// UnsetNeedMarketingConsentsConfirmation ensures that no value is present for NeedMarketingConsentsConfirmation, not even an explicit nil
func (o *GetUserInfoResponseInfo) UnsetNeedMarketingConsentsConfirmation() {
	o.NeedMarketingConsentsConfirmation.Unset()
}

// GetNeedPasswordChange returns the NeedPasswordChange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUserInfoResponseInfo) GetNeedPasswordChange() bool {
	if o == nil || o.NeedPasswordChange.Get() == nil {
		var ret bool
		return ret
	}
	return *o.NeedPasswordChange.Get()
}

// GetNeedPasswordChangeOk returns a tuple with the NeedPasswordChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUserInfoResponseInfo) GetNeedPasswordChangeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NeedPasswordChange.Get(), o.NeedPasswordChange.IsSet()
}

// HasNeedPasswordChange returns a boolean if a field has been set.
func (o *GetUserInfoResponseInfo) HasNeedPasswordChange() bool {
	if o != nil && o.NeedPasswordChange.IsSet() {
		return true
	}

	return false
}

// SetNeedPasswordChange gets a reference to the given NullableBool and assigns it to the NeedPasswordChange field.
func (o *GetUserInfoResponseInfo) SetNeedPasswordChange(v bool) *GetUserInfoResponseInfo {
	o.NeedPasswordChange.Set(&v)
	return o
}
// SetNeedPasswordChangeNil sets the value for NeedPasswordChange to be an explicit nil
func (o *GetUserInfoResponseInfo) SetNeedPasswordChangeNil() *GetUserInfoResponseInfo {
	o.NeedPasswordChange.Set(nil)
	return o
}

// UnsetNeedPasswordChange ensures that no value is present for NeedPasswordChange, not even an explicit nil
func (o *GetUserInfoResponseInfo) UnsetNeedPasswordChange() {
	o.NeedPasswordChange.Unset()
}

// GetNeedTermsOfServiceConfirmation returns the NeedTermsOfServiceConfirmation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUserInfoResponseInfo) GetNeedTermsOfServiceConfirmation() bool {
	if o == nil || o.NeedTermsOfServiceConfirmation.Get() == nil {
		var ret bool
		return ret
	}
	return *o.NeedTermsOfServiceConfirmation.Get()
}

// GetNeedTermsOfServiceConfirmationOk returns a tuple with the NeedTermsOfServiceConfirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUserInfoResponseInfo) GetNeedTermsOfServiceConfirmationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NeedTermsOfServiceConfirmation.Get(), o.NeedTermsOfServiceConfirmation.IsSet()
}

// HasNeedTermsOfServiceConfirmation returns a boolean if a field has been set.
func (o *GetUserInfoResponseInfo) HasNeedTermsOfServiceConfirmation() bool {
	if o != nil && o.NeedTermsOfServiceConfirmation.IsSet() {
		return true
	}

	return false
}

// SetNeedTermsOfServiceConfirmation gets a reference to the given NullableBool and assigns it to the NeedTermsOfServiceConfirmation field.
func (o *GetUserInfoResponseInfo) SetNeedTermsOfServiceConfirmation(v bool) *GetUserInfoResponseInfo {
	o.NeedTermsOfServiceConfirmation.Set(&v)
	return o
}
// SetNeedTermsOfServiceConfirmationNil sets the value for NeedTermsOfServiceConfirmation to be an explicit nil
func (o *GetUserInfoResponseInfo) SetNeedTermsOfServiceConfirmationNil() *GetUserInfoResponseInfo {
	o.NeedTermsOfServiceConfirmation.Set(nil)
	return o
}

// UnsetNeedTermsOfServiceConfirmation ensures that no value is present for NeedTermsOfServiceConfirmation, not even an explicit nil
func (o *GetUserInfoResponseInfo) UnsetNeedTermsOfServiceConfirmation() {
	o.NeedTermsOfServiceConfirmation.Unset()
}

func (o GetUserInfoResponseInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NeedMarketingConsentsConfirmation.IsSet() {
		toSerialize["need_marketing_consents_confirmation"] = o.NeedMarketingConsentsConfirmation.Get()
	}
	if o.NeedPasswordChange.IsSet() {
		toSerialize["need_password_change"] = o.NeedPasswordChange.Get()
	}
	if o.NeedTermsOfServiceConfirmation.IsSet() {
		toSerialize["need_terms_of_service_confirmation"] = o.NeedTermsOfServiceConfirmation.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGetUserInfoResponseInfo struct {
	value *GetUserInfoResponseInfo
	isSet bool
}

func (v NullableGetUserInfoResponseInfo) Get() *GetUserInfoResponseInfo {
	return v.value
}

func (v *NullableGetUserInfoResponseInfo) Set(val *GetUserInfoResponseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserInfoResponseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserInfoResponseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserInfoResponseInfo(val *GetUserInfoResponseInfo) *NullableGetUserInfoResponseInfo {
	return &NullableGetUserInfoResponseInfo{value: val, isSet: true}
}

func (v NullableGetUserInfoResponseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserInfoResponseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


