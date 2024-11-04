/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.3
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
)

// FattureInCloudPlanType Fatture in Cloud account plan type.
type FattureInCloudPlanType string

// List of FattureInCloudPlanType
var FattureInCloudPlanTypes = struct {
	TRIAL FattureInCloudPlanType
	STANDARD FattureInCloudPlanType
	PREMIUM FattureInCloudPlanType
	PREMIUM_PLUS FattureInCloudPlanType
	COMPLETE FattureInCloudPlanType
} {
			TRIAL: "trial",
			STANDARD: "standard",
			PREMIUM: "premium",
			PREMIUM_PLUS: "premium_plus",
			COMPLETE: "complete",
}

// All allowed values of FattureInCloudPlanType enum
var AllowedFattureInCloudPlanTypeEnumValues = []FattureInCloudPlanType{
	"trial",
	"standard",
	"premium",
	"premium_plus",
	"complete",
}

func (v *FattureInCloudPlanType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FattureInCloudPlanType(value)
	for _, existing := range AllowedFattureInCloudPlanTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FattureInCloudPlanType", value)
}

// NewFattureInCloudPlanTypeFromValue returns a pointer to a valid FattureInCloudPlanType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFattureInCloudPlanTypeFromValue(v string) (*FattureInCloudPlanType, error) {
	ev := FattureInCloudPlanType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FattureInCloudPlanType: valid values are %v", v, AllowedFattureInCloudPlanTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FattureInCloudPlanType) IsValid() bool {
	for _, existing := range AllowedFattureInCloudPlanTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FattureInCloudPlanType value
func (v FattureInCloudPlanType) Ptr() *FattureInCloudPlanType {
	return &v
}

type NullableFattureInCloudPlanType struct {
	value *FattureInCloudPlanType
	isSet bool
}

func (v NullableFattureInCloudPlanType) Get() *FattureInCloudPlanType {
	return v.value
}

func (v *NullableFattureInCloudPlanType) Set(val *FattureInCloudPlanType) {
	v.value = val
	v.isSet = true
}

func (v NullableFattureInCloudPlanType) IsSet() bool {
	return v.isSet
}

func (v *NullableFattureInCloudPlanType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFattureInCloudPlanType(val *FattureInCloudPlanType) *NullableFattureInCloudPlanType {
	return &NullableFattureInCloudPlanType{value: val, isSet: true}
}

func (v NullableFattureInCloudPlanType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFattureInCloudPlanType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

