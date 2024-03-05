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

// PaymentMethodType Payment method type
type PaymentMethodType string

// List of PaymentMethodType
var PaymentMethodTypes = struct {
	STANDARD PaymentMethodType
	RIBA PaymentMethodType
} {
	STANDARD: "standard",
	RIBA: "riba",
}

// All allowed values of PaymentMethodType enum
var AllowedPaymentMethodTypeEnumValues = []PaymentMethodType{
	"standard",
	"riba",
}

func (v *PaymentMethodType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentMethodType(value)
	for _, existing := range AllowedPaymentMethodTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentMethodType", value)
}

// NewPaymentMethodTypeFromValue returns a pointer to a valid PaymentMethodType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentMethodTypeFromValue(v string) (*PaymentMethodType, error) {
	ev := PaymentMethodType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentMethodType: valid values are %v", v, AllowedPaymentMethodTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentMethodType) IsValid() bool {
	for _, existing := range AllowedPaymentMethodTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentMethodType value
func (v PaymentMethodType) Ptr() *PaymentMethodType {
	return &v
}

type NullablePaymentMethodType struct {
	value *PaymentMethodType
	isSet bool
}

func (v NullablePaymentMethodType) Get() *PaymentMethodType {
	return v.value
}

func (v *NullablePaymentMethodType) Set(val *PaymentMethodType) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodType) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodType(val *PaymentMethodType) *NullablePaymentMethodType {
	return &NullablePaymentMethodType{value: val, isSet: true}
}

func (v NullablePaymentMethodType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

