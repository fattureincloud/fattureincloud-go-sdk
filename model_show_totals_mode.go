/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fattureincloud

import (
	"encoding/json"
	"fmt"
)

// ShowTotalsMode Totals mode.
type ShowTotalsMode string

// List of ShowTotalsMode
var ShowTotalsModes = struct {
	NONE ShowTotalsMode
	NETS ShowTotalsMode
	ALL ShowTotalsMode
} {
	NONE: "none",
	NETS: "nets",
	ALL: "all",
}

// All allowed values of ShowTotalsMode enum
var AllowedShowTotalsModeEnumValues = []ShowTotalsMode{
	"none",
	"nets",
	"all",
}

func (v *ShowTotalsMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ShowTotalsMode(value)
	for _, existing := range AllowedShowTotalsModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ShowTotalsMode", value)
}

// NewShowTotalsModeFromValue returns a pointer to a valid ShowTotalsMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShowTotalsModeFromValue(v string) (*ShowTotalsMode, error) {
	ev := ShowTotalsMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShowTotalsMode: valid values are %v", v, AllowedShowTotalsModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShowTotalsMode) IsValid() bool {
	for _, existing := range AllowedShowTotalsModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ShowTotalsMode value
func (v ShowTotalsMode) Ptr() *ShowTotalsMode {
	return &v
}

type NullableShowTotalsMode struct {
	value *ShowTotalsMode
	isSet bool
}

func (v NullableShowTotalsMode) Get() *ShowTotalsMode {
	return v.value
}

func (v *NullableShowTotalsMode) Set(val *ShowTotalsMode) {
	v.value = val
	v.isSet = true
}

func (v NullableShowTotalsMode) IsSet() bool {
	return v.isSet
}

func (v *NullableShowTotalsMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShowTotalsMode(val *ShowTotalsMode) *NullableShowTotalsMode {
	return &NullableShowTotalsMode{value: val, isSet: true}
}

func (v NullableShowTotalsMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShowTotalsMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

