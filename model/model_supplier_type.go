/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.23
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
)

// SupplierType Supplier type.
type SupplierType string

// List of SupplierType
var SupplierTypes = struct {
	COMPANY SupplierType
	PERSON SupplierType
	PA SupplierType
	CONDO SupplierType
} {
	COMPANY: "company",
	PERSON: "person",
	PA: "pa",
	CONDO: "condo",
}

// All allowed values of SupplierType enum
var AllowedSupplierTypeEnumValues = []SupplierType{
	"company",
	"person",
	"pa",
	"condo",
}

func (v *SupplierType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SupplierType(value)
	for _, existing := range AllowedSupplierTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SupplierType", value)
}

// NewSupplierTypeFromValue returns a pointer to a valid SupplierType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSupplierTypeFromValue(v string) (*SupplierType, error) {
	ev := SupplierType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SupplierType: valid values are %v", v, AllowedSupplierTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SupplierType) IsValid() bool {
	for _, existing := range AllowedSupplierTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SupplierType value
func (v SupplierType) Ptr() *SupplierType {
	return &v
}

type NullableSupplierType struct {
	value *SupplierType
	isSet bool
}

func (v NullableSupplierType) Get() *SupplierType {
	return v.value
}

func (v *NullableSupplierType) Set(val *SupplierType) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplierType) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplierType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplierType(val *SupplierType) *NullableSupplierType {
	return &NullableSupplierType{value: val, isSet: true}
}

func (v NullableSupplierType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplierType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

