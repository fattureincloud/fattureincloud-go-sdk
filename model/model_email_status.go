/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.31
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
)

// EmailStatus Email status
type EmailStatus string

// List of EmailStatus
var EmailStatuses = struct {
	SENDING EmailStatus
	PENDING EmailStatus
	SENT EmailStatus
} {
	SENDING: "sending",
	PENDING: "pending",
	SENT: "sent",
}

// All allowed values of EmailStatus enum
var AllowedEmailStatusEnumValues = []EmailStatus{
	"sending",
	"pending",
	"sent",
}

func (v *EmailStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmailStatus(value)
	for _, existing := range AllowedEmailStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EmailStatus", value)
}

// NewEmailStatusFromValue returns a pointer to a valid EmailStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmailStatusFromValue(v string) (*EmailStatus, error) {
	ev := EmailStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EmailStatus: valid values are %v", v, AllowedEmailStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmailStatus) IsValid() bool {
	for _, existing := range AllowedEmailStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmailStatus value
func (v EmailStatus) Ptr() *EmailStatus {
	return &v
}

type NullableEmailStatus struct {
	value *EmailStatus
	isSet bool
}

func (v NullableEmailStatus) Get() *EmailStatus {
	return v.value
}

func (v *NullableEmailStatus) Set(val *EmailStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailStatus(val *EmailStatus) *NullableEmailStatus {
	return &NullableEmailStatus{value: val, isSet: true}
}

func (v NullableEmailStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

