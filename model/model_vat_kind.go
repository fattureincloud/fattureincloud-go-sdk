/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.2
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
)

// VatKind Vat kind [esigibilità IVA]
type VatKind string

// List of VatKind
var VatKinds = struct {
	I VatKind
	D VatKind
	S VatKind
} {
	I: "I",
	D: "D",
	S: "S",
}

// All allowed values of VatKind enum
var AllowedVatKindEnumValues = []VatKind{
	"I",
	"D",
	"S",
}

func (v *VatKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VatKind(value)
	for _, existing := range AllowedVatKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VatKind", value)
}

// NewVatKindFromValue returns a pointer to a valid VatKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVatKindFromValue(v string) (*VatKind, error) {
	ev := VatKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VatKind: valid values are %v", v, AllowedVatKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VatKind) IsValid() bool {
	for _, existing := range AllowedVatKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VatKind value
func (v VatKind) Ptr() *VatKind {
	return &v
}

type NullableVatKind struct {
	value *VatKind
	isSet bool
}

func (v NullableVatKind) Get() *VatKind {
	return v.value
}

func (v *NullableVatKind) Set(val *VatKind) {
	v.value = val
	v.isSet = true
}

func (v NullableVatKind) IsSet() bool {
	return v.isSet
}

func (v *NullableVatKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVatKind(val *VatKind) *NullableVatKind {
	return &NullableVatKind{value: val, isSet: true}
}

func (v NullableVatKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVatKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

