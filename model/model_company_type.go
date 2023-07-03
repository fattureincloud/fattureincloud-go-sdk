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

// CompanyType Company type
type CompanyType string

// List of CompanyType
var CompanyTypes = struct {
	COMPANY CompanyType
	ACCOUNTANT CompanyType
} {
	COMPANY: "company",
	ACCOUNTANT: "accountant",
}

// All allowed values of CompanyType enum
var AllowedCompanyTypeEnumValues = []CompanyType{
	"company",
	"accountant",
}

func (v *CompanyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CompanyType(value)
	for _, existing := range AllowedCompanyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CompanyType", value)
}

// NewCompanyTypeFromValue returns a pointer to a valid CompanyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCompanyTypeFromValue(v string) (*CompanyType, error) {
	ev := CompanyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CompanyType: valid values are %v", v, AllowedCompanyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CompanyType) IsValid() bool {
	for _, existing := range AllowedCompanyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CompanyType value
func (v CompanyType) Ptr() *CompanyType {
	return &v
}

type NullableCompanyType struct {
	value *CompanyType
	isSet bool
}

func (v NullableCompanyType) Get() *CompanyType {
	return v.value
}

func (v *NullableCompanyType) Set(val *CompanyType) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyType) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyType(val *CompanyType) *NullableCompanyType {
	return &NullableCompanyType{value: val, isSet: true}
}

func (v NullableCompanyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

