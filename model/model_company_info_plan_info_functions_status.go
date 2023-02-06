/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.25
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// CompanyInfoPlanInfoFunctionsStatus struct for CompanyInfoPlanInfoFunctionsStatus
type CompanyInfoPlanInfoFunctionsStatus struct {
	TsDigital *FunctionStatus `json:"ts_digital,omitempty"`
	TsPay *FunctionStatus `json:"ts_pay,omitempty"`
}

// NewCompanyInfoPlanInfoFunctionsStatus instantiates a new CompanyInfoPlanInfoFunctionsStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyInfoPlanInfoFunctionsStatus() *CompanyInfoPlanInfoFunctionsStatus {
	this := CompanyInfoPlanInfoFunctionsStatus{}
	return &this
}

// NewCompanyInfoPlanInfoFunctionsStatusWithDefaults instantiates a new CompanyInfoPlanInfoFunctionsStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyInfoPlanInfoFunctionsStatusWithDefaults() *CompanyInfoPlanInfoFunctionsStatus {
	this := CompanyInfoPlanInfoFunctionsStatus{}
	return &this
}

// GetTsDigital returns the TsDigital field value if set, zero value otherwise.
func (o *CompanyInfoPlanInfoFunctionsStatus) GetTsDigital() FunctionStatus {
	if o == nil || isNil(o.TsDigital) {
		var ret FunctionStatus
		return ret
	}
	return *o.TsDigital
}

// GetTsDigitalOk returns a tuple with the TsDigital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyInfoPlanInfoFunctionsStatus) GetTsDigitalOk() (*FunctionStatus, bool) {
	if o == nil || isNil(o.TsDigital) {
		return nil, false
	}
	return o.TsDigital, true
}

// HasTsDigital returns a boolean if a field has been set.
func (o *CompanyInfoPlanInfoFunctionsStatus) HasTsDigital() bool {
	if o != nil && !isNil(o.TsDigital) {
		return true
	}

	return false
}

// SetTsDigital gets a reference to the given FunctionStatus and assigns it to the TsDigital field.
func (o *CompanyInfoPlanInfoFunctionsStatus) SetTsDigital(v FunctionStatus) *CompanyInfoPlanInfoFunctionsStatus {
	o.TsDigital = &v
	return o
}

// GetTsPay returns the TsPay field value if set, zero value otherwise.
func (o *CompanyInfoPlanInfoFunctionsStatus) GetTsPay() FunctionStatus {
	if o == nil || isNil(o.TsPay) {
		var ret FunctionStatus
		return ret
	}
	return *o.TsPay
}

// GetTsPayOk returns a tuple with the TsPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyInfoPlanInfoFunctionsStatus) GetTsPayOk() (*FunctionStatus, bool) {
	if o == nil || isNil(o.TsPay) {
		return nil, false
	}
	return o.TsPay, true
}

// HasTsPay returns a boolean if a field has been set.
func (o *CompanyInfoPlanInfoFunctionsStatus) HasTsPay() bool {
	if o != nil && !isNil(o.TsPay) {
		return true
	}

	return false
}

// SetTsPay gets a reference to the given FunctionStatus and assigns it to the TsPay field.
func (o *CompanyInfoPlanInfoFunctionsStatus) SetTsPay(v FunctionStatus) *CompanyInfoPlanInfoFunctionsStatus {
	o.TsPay = &v
	return o
}

func (o CompanyInfoPlanInfoFunctionsStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TsDigital) {
		toSerialize["ts_digital"] = o.TsDigital
	}
	if !isNil(o.TsPay) {
		toSerialize["ts_pay"] = o.TsPay
	}
	return json.Marshal(toSerialize)
}

type NullableCompanyInfoPlanInfoFunctionsStatus struct {
	value *CompanyInfoPlanInfoFunctionsStatus
	isSet bool
}

func (v NullableCompanyInfoPlanInfoFunctionsStatus) Get() *CompanyInfoPlanInfoFunctionsStatus {
	return v.value
}

func (v *NullableCompanyInfoPlanInfoFunctionsStatus) Set(val *CompanyInfoPlanInfoFunctionsStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyInfoPlanInfoFunctionsStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyInfoPlanInfoFunctionsStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyInfoPlanInfoFunctionsStatus(val *CompanyInfoPlanInfoFunctionsStatus) *NullableCompanyInfoPlanInfoFunctionsStatus {
	return &NullableCompanyInfoPlanInfoFunctionsStatus{value: val, isSet: true}
}

func (v NullableCompanyInfoPlanInfoFunctionsStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyInfoPlanInfoFunctionsStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


