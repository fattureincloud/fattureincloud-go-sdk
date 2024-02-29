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

// PaymentTermsType Payment terms type
type PaymentTermsType string

// List of PaymentTermsType
var PaymentTermsTypes = struct {
	STANDARD PaymentTermsType
	END_OF_MONTH PaymentTermsType
} {
	STANDARD: "standard",
	END_OF_MONTH: "end_of_month",
}

// All allowed values of PaymentTermsType enum
var AllowedPaymentTermsTypeEnumValues = []PaymentTermsType{
	"standard",
	"end_of_month",
}

func (v *PaymentTermsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentTermsType(value)
	for _, existing := range AllowedPaymentTermsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentTermsType", value)
}

// NewPaymentTermsTypeFromValue returns a pointer to a valid PaymentTermsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentTermsTypeFromValue(v string) (*PaymentTermsType, error) {
	ev := PaymentTermsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentTermsType: valid values are %v", v, AllowedPaymentTermsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentTermsType) IsValid() bool {
	for _, existing := range AllowedPaymentTermsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentTermsType value
func (v PaymentTermsType) Ptr() *PaymentTermsType {
	return &v
}

type NullablePaymentTermsType struct {
	value *PaymentTermsType
	isSet bool
}

func (v NullablePaymentTermsType) Get() *PaymentTermsType {
	return v.value
}

func (v *NullablePaymentTermsType) Set(val *PaymentTermsType) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentTermsType) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentTermsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentTermsType(val *PaymentTermsType) *NullablePaymentTermsType {
	return &NullablePaymentTermsType{value: val, isSet: true}
}

func (v NullablePaymentTermsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentTermsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

