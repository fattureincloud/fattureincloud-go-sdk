/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.32
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
)

// OriginalDocumentType Issued document original document type
type OriginalDocumentType string

// List of OriginalDocumentType
var OriginalDocumentTypes = struct {
	ORDINE OriginalDocumentType
	CONTRATTO OriginalDocumentType
	CONVENZIONE OriginalDocumentType
} {
	ORDINE: "ordine",
	CONTRATTO: "contratto",
	CONVENZIONE: "convenzione",
}

// All allowed values of OriginalDocumentType enum
var AllowedOriginalDocumentTypeEnumValues = []OriginalDocumentType{
	"ordine",
	"contratto",
	"convenzione",
}

func (v *OriginalDocumentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OriginalDocumentType(value)
	for _, existing := range AllowedOriginalDocumentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OriginalDocumentType", value)
}

// NewOriginalDocumentTypeFromValue returns a pointer to a valid OriginalDocumentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOriginalDocumentTypeFromValue(v string) (*OriginalDocumentType, error) {
	ev := OriginalDocumentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OriginalDocumentType: valid values are %v", v, AllowedOriginalDocumentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OriginalDocumentType) IsValid() bool {
	for _, existing := range AllowedOriginalDocumentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OriginalDocumentType value
func (v OriginalDocumentType) Ptr() *OriginalDocumentType {
	return &v
}

type NullableOriginalDocumentType struct {
	value *OriginalDocumentType
	isSet bool
}

func (v NullableOriginalDocumentType) Get() *OriginalDocumentType {
	return v.value
}

func (v *NullableOriginalDocumentType) Set(val *OriginalDocumentType) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginalDocumentType) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginalDocumentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginalDocumentType(val *OriginalDocumentType) *NullableOriginalDocumentType {
	return &NullableOriginalDocumentType{value: val, isSet: true}
}

func (v NullableOriginalDocumentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginalDocumentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

