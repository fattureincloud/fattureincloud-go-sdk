/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.20
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
)

// CashbookEntryKind Cashbook kind.
type CashbookEntryKind string

// List of CashbookEntryKind
var CashbookEntryKinds = struct {
	CASHBOOK CashbookEntryKind
	ISSUED_DOCUMENT CashbookEntryKind
	RECEIVED_DOCUMENT CashbookEntryKind
	TAX CashbookEntryKind
	RECEIPT CashbookEntryKind
} {
	CASHBOOK: "cashbook",
	ISSUED_DOCUMENT: "issued_document",
	RECEIVED_DOCUMENT: "received_document",
	TAX: "tax",
	RECEIPT: "receipt",
}

// All allowed values of CashbookEntryKind enum
var AllowedCashbookEntryKindEnumValues = []CashbookEntryKind{
	"cashbook",
	"issued_document",
	"received_document",
	"tax",
	"receipt",
}

func (v *CashbookEntryKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CashbookEntryKind(value)
	for _, existing := range AllowedCashbookEntryKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CashbookEntryKind", value)
}

// NewCashbookEntryKindFromValue returns a pointer to a valid CashbookEntryKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCashbookEntryKindFromValue(v string) (*CashbookEntryKind, error) {
	ev := CashbookEntryKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CashbookEntryKind: valid values are %v", v, AllowedCashbookEntryKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CashbookEntryKind) IsValid() bool {
	for _, existing := range AllowedCashbookEntryKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CashbookEntryKind value
func (v CashbookEntryKind) Ptr() *CashbookEntryKind {
	return &v
}

type NullableCashbookEntryKind struct {
	value *CashbookEntryKind
	isSet bool
}

func (v NullableCashbookEntryKind) Get() *CashbookEntryKind {
	return v.value
}

func (v *NullableCashbookEntryKind) Set(val *CashbookEntryKind) {
	v.value = val
	v.isSet = true
}

func (v NullableCashbookEntryKind) IsSet() bool {
	return v.isSet
}

func (v *NullableCashbookEntryKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashbookEntryKind(val *CashbookEntryKind) *NullableCashbookEntryKind {
	return &NullableCashbookEntryKind{value: val, isSet: true}
}

func (v NullableCashbookEntryKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashbookEntryKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

