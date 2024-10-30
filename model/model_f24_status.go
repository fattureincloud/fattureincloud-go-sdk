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

// F24Status F24 status
type F24Status string

// List of F24Status
var F24Statuses = struct {
    PAID F24Status
    NOT_PAID F24Status
    REVERSED F24Status
} {
            PAID: "paid",
            NOT_PAID: "not_paid",
            REVERSED: "reversed",
}

// All allowed values of F24Status enum
var AllowedF24StatusEnumValues = []F24Status{
	"paid",
	"not_paid",
	"reversed",
}

func (v *F24Status) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := F24Status(value)
	for _, existing := range AllowedF24StatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid F24Status", value)
}

// NewF24StatusFromValue returns a pointer to a valid F24Status
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewF24StatusFromValue(v string) (*F24Status, error) {
	ev := F24Status(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for F24Status: valid values are %v", v, AllowedF24StatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v F24Status) IsValid() bool {
	for _, existing := range AllowedF24StatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to F24Status value
func (v F24Status) Ptr() *F24Status {
	return &v
}

type NullableF24Status struct {
	value *F24Status
	isSet bool
}

func (v NullableF24Status) Get() *F24Status {
	return v.value
}

func (v *NullableF24Status) Set(val *F24Status) {
	v.value = val
	v.isSet = true
}

func (v NullableF24Status) IsSet() bool {
	return v.isSet
}

func (v *NullableF24Status) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableF24Status(val *F24Status) *NullableF24Status {
	return &NullableF24Status{value: val, isSet: true}
}

func (v NullableF24Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableF24Status) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

