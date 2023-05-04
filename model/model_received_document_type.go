/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.27
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
)

// ReceivedDocumentType Received document type.
type ReceivedDocumentType string

// List of ReceivedDocumentType
var ReceivedDocumentTypes = struct {
	EXPENSE ReceivedDocumentType
	PASSIVE_CREDIT_NOTE ReceivedDocumentType
	PASSIVE_DELIVERY_NOTE ReceivedDocumentType
} {
	EXPENSE: "expense",
	PASSIVE_CREDIT_NOTE: "passive_credit_note",
	PASSIVE_DELIVERY_NOTE: "passive_delivery_note",
}

// All allowed values of ReceivedDocumentType enum
var AllowedReceivedDocumentTypeEnumValues = []ReceivedDocumentType{
	"expense",
	"passive_credit_note",
	"passive_delivery_note",
}

func (v *ReceivedDocumentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReceivedDocumentType(value)
	for _, existing := range AllowedReceivedDocumentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReceivedDocumentType", value)
}

// NewReceivedDocumentTypeFromValue returns a pointer to a valid ReceivedDocumentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReceivedDocumentTypeFromValue(v string) (*ReceivedDocumentType, error) {
	ev := ReceivedDocumentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReceivedDocumentType: valid values are %v", v, AllowedReceivedDocumentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReceivedDocumentType) IsValid() bool {
	for _, existing := range AllowedReceivedDocumentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReceivedDocumentType value
func (v ReceivedDocumentType) Ptr() *ReceivedDocumentType {
	return &v
}

type NullableReceivedDocumentType struct {
	value *ReceivedDocumentType
	isSet bool
}

func (v NullableReceivedDocumentType) Get() *ReceivedDocumentType {
	return v.value
}

func (v *NullableReceivedDocumentType) Set(val *ReceivedDocumentType) {
	v.value = val
	v.isSet = true
}

func (v NullableReceivedDocumentType) IsSet() bool {
	return v.isSet
}

func (v *NullableReceivedDocumentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceivedDocumentType(val *ReceivedDocumentType) *NullableReceivedDocumentType {
	return &NullableReceivedDocumentType{value: val, isSet: true}
}

func (v NullableReceivedDocumentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceivedDocumentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

