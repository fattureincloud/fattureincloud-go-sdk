/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.26
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
)

// CashbookEntryType Cashbook type.
type CashbookEntryType string

// List of CashbookEntryType
var CashbookEntryTypes = struct {
	IN CashbookEntryType
	OUT CashbookEntryType
} {
	IN: "in",
	OUT: "out",
}

// All allowed values of CashbookEntryType enum
var AllowedCashbookEntryTypeEnumValues = []CashbookEntryType{
	"in",
	"out",
}

func (v *CashbookEntryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CashbookEntryType(value)
	for _, existing := range AllowedCashbookEntryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CashbookEntryType", value)
}

// NewCashbookEntryTypeFromValue returns a pointer to a valid CashbookEntryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCashbookEntryTypeFromValue(v string) (*CashbookEntryType, error) {
	ev := CashbookEntryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CashbookEntryType: valid values are %v", v, AllowedCashbookEntryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CashbookEntryType) IsValid() bool {
	for _, existing := range AllowedCashbookEntryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CashbookEntryType value
func (v CashbookEntryType) Ptr() *CashbookEntryType {
	return &v
}

type NullableCashbookEntryType struct {
	value *CashbookEntryType
	isSet bool
}

func (v NullableCashbookEntryType) Get() *CashbookEntryType {
	return v.value
}

func (v *NullableCashbookEntryType) Set(val *CashbookEntryType) {
	v.value = val
	v.isSet = true
}

func (v NullableCashbookEntryType) IsSet() bool {
	return v.isSet
}

func (v *NullableCashbookEntryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashbookEntryType(val *CashbookEntryType) *NullableCashbookEntryType {
	return &NullableCashbookEntryType{value: val, isSet: true}
}

func (v NullableCashbookEntryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashbookEntryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

