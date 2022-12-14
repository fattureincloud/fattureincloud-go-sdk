/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.24
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
)

// EmailRecipientStatus Email recipient status.
type EmailRecipientStatus string

// List of EmailRecipientStatus
var EmailRecipientStatuses = struct {
	UNKNOWN EmailRecipientStatus
	DOCUMENT_OPENED EmailRecipientStatus
	EMAIL_OPENED EmailRecipientStatus
} {
	UNKNOWN: "unknown",
	DOCUMENT_OPENED: "document_opened",
	EMAIL_OPENED: "email_opened",
}

// All allowed values of EmailRecipientStatus enum
var AllowedEmailRecipientStatusEnumValues = []EmailRecipientStatus{
	"unknown",
	"document_opened",
	"email_opened",
}

func (v *EmailRecipientStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmailRecipientStatus(value)
	for _, existing := range AllowedEmailRecipientStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EmailRecipientStatus", value)
}

// NewEmailRecipientStatusFromValue returns a pointer to a valid EmailRecipientStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmailRecipientStatusFromValue(v string) (*EmailRecipientStatus, error) {
	ev := EmailRecipientStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EmailRecipientStatus: valid values are %v", v, AllowedEmailRecipientStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmailRecipientStatus) IsValid() bool {
	for _, existing := range AllowedEmailRecipientStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmailRecipientStatus value
func (v EmailRecipientStatus) Ptr() *EmailRecipientStatus {
	return &v
}

type NullableEmailRecipientStatus struct {
	value *EmailRecipientStatus
	isSet bool
}

func (v NullableEmailRecipientStatus) Get() *EmailRecipientStatus {
	return v.value
}

func (v *NullableEmailRecipientStatus) Set(val *EmailRecipientStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailRecipientStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailRecipientStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailRecipientStatus(val *EmailRecipientStatus) *NullableEmailRecipientStatus {
	return &NullableEmailRecipientStatus{value: val, isSet: true}
}

func (v NullableEmailRecipientStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailRecipientStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

