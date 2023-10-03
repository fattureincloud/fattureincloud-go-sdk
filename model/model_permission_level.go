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
	"fmt"
)

// PermissionLevel Permission level
type PermissionLevel string

// List of PermissionLevel
var PermissionLevels = struct {
	NONE PermissionLevel
	READ PermissionLevel
	WRITE PermissionLevel
	DETAILED PermissionLevel
} {
	NONE: "none",
	READ: "read",
	WRITE: "write",
	DETAILED: "detailed",
}

// All allowed values of PermissionLevel enum
var AllowedPermissionLevelEnumValues = []PermissionLevel{
	"none",
	"read",
	"write",
	"detailed",
}

func (v *PermissionLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PermissionLevel(value)
	for _, existing := range AllowedPermissionLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PermissionLevel", value)
}

// NewPermissionLevelFromValue returns a pointer to a valid PermissionLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPermissionLevelFromValue(v string) (*PermissionLevel, error) {
	ev := PermissionLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PermissionLevel: valid values are %v", v, AllowedPermissionLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PermissionLevel) IsValid() bool {
	for _, existing := range AllowedPermissionLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PermissionLevel value
func (v PermissionLevel) Ptr() *PermissionLevel {
	return &v
}

type NullablePermissionLevel struct {
	value *PermissionLevel
	isSet bool
}

func (v NullablePermissionLevel) Get() *PermissionLevel {
	return v.value
}

func (v *NullablePermissionLevel) Set(val *PermissionLevel) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionLevel) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionLevel(val *PermissionLevel) *NullablePermissionLevel {
	return &NullablePermissionLevel{value: val, isSet: true}
}

func (v NullablePermissionLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

