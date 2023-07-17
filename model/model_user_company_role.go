/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.29
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
)

// UserCompanyRole User company role
type UserCompanyRole string

// List of UserCompanyRole
var UserCompanyRoles = struct {
	MASTER UserCompanyRole
	SUBACCOUNT UserCompanyRole
	EMPLOYEE UserCompanyRole
} {
	MASTER: "master",
	SUBACCOUNT: "subaccount",
	EMPLOYEE: "employee",
}

// All allowed values of UserCompanyRole enum
var AllowedUserCompanyRoleEnumValues = []UserCompanyRole{
	"master",
	"subaccount",
	"employee",
}

func (v *UserCompanyRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserCompanyRole(value)
	for _, existing := range AllowedUserCompanyRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserCompanyRole", value)
}

// NewUserCompanyRoleFromValue returns a pointer to a valid UserCompanyRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserCompanyRoleFromValue(v string) (*UserCompanyRole, error) {
	ev := UserCompanyRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserCompanyRole: valid values are %v", v, AllowedUserCompanyRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserCompanyRole) IsValid() bool {
	for _, existing := range AllowedUserCompanyRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserCompanyRole value
func (v UserCompanyRole) Ptr() *UserCompanyRole {
	return &v
}

type NullableUserCompanyRole struct {
	value *UserCompanyRole
	isSet bool
}

func (v NullableUserCompanyRole) Get() *UserCompanyRole {
	return v.value
}

func (v *NullableUserCompanyRole) Set(val *UserCompanyRole) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCompanyRole) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCompanyRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCompanyRole(val *UserCompanyRole) *NullableUserCompanyRole {
	return &NullableUserCompanyRole{value: val, isSet: true}
}

func (v NullableUserCompanyRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCompanyRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

