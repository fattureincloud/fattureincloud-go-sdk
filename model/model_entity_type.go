/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.28
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
)

// EntityType Entity type.
type EntityType string

// List of EntityType
var EntityTypes = struct {
	COMPANY EntityType
	PERSON EntityType
	PA EntityType
	CONDO EntityType
} {
	COMPANY: "company",
	PERSON: "person",
	PA: "pa",
	CONDO: "condo",
}

// All allowed values of EntityType enum
var AllowedEntityTypeEnumValues = []EntityType{
	"company",
	"person",
	"pa",
	"condo",
}

func (v *EntityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EntityType(value)
	for _, existing := range AllowedEntityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EntityType", value)
}

// NewEntityTypeFromValue returns a pointer to a valid EntityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEntityTypeFromValue(v string) (*EntityType, error) {
	ev := EntityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EntityType: valid values are %v", v, AllowedEntityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EntityType) IsValid() bool {
	for _, existing := range AllowedEntityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityType value
func (v EntityType) Ptr() *EntityType {
	return &v
}

type NullableEntityType struct {
	value *EntityType
	isSet bool
}

func (v NullableEntityType) Get() *EntityType {
	return v.value
}

func (v *NullableEntityType) Set(val *EntityType) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityType) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityType(val *EntityType) *NullableEntityType {
	return &NullableEntityType{value: val, isSet: true}
}

func (v NullableEntityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

