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

// PaymentAccountType Payment account type
type PaymentAccountType string

// List of PaymentAccountType
var PaymentAccountTypes = struct {
    STANDARD PaymentAccountType
    BANK PaymentAccountType
} {
            STANDARD: "standard",
            BANK: "bank",
}

// All allowed values of PaymentAccountType enum
var AllowedPaymentAccountTypeEnumValues = []PaymentAccountType{
	"standard",
	"bank",
}

func (v *PaymentAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentAccountType(value)
	for _, existing := range AllowedPaymentAccountTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentAccountType", value)
}

// NewPaymentAccountTypeFromValue returns a pointer to a valid PaymentAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentAccountTypeFromValue(v string) (*PaymentAccountType, error) {
	ev := PaymentAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentAccountType: valid values are %v", v, AllowedPaymentAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentAccountType) IsValid() bool {
	for _, existing := range AllowedPaymentAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentAccountType value
func (v PaymentAccountType) Ptr() *PaymentAccountType {
	return &v
}

type NullablePaymentAccountType struct {
	value *PaymentAccountType
	isSet bool
}

func (v NullablePaymentAccountType) Get() *PaymentAccountType {
	return v.value
}

func (v *NullablePaymentAccountType) Set(val *PaymentAccountType) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentAccountType) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentAccountType(val *PaymentAccountType) *NullablePaymentAccountType {
	return &NullablePaymentAccountType{value: val, isSet: true}
}

func (v NullablePaymentAccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

