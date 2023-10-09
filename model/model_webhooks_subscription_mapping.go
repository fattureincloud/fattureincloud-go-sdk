/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.30
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
)

// WebhooksSubscriptionMapping Webhooks Subscription mapping
type WebhooksSubscriptionMapping string

// List of WebhooksSubscriptionMapping
var WebhooksSubscriptionMappings = struct {
	BINARY WebhooksSubscriptionMapping
	STRUCTURED WebhooksSubscriptionMapping
} {
	BINARY: "binary",
	STRUCTURED: "structured",
}

// All allowed values of WebhooksSubscriptionMapping enum
var AllowedWebhooksSubscriptionMappingEnumValues = []WebhooksSubscriptionMapping{
	"binary",
	"structured",
}

func (v *WebhooksSubscriptionMapping) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhooksSubscriptionMapping(value)
	for _, existing := range AllowedWebhooksSubscriptionMappingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhooksSubscriptionMapping", value)
}

// NewWebhooksSubscriptionMappingFromValue returns a pointer to a valid WebhooksSubscriptionMapping
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhooksSubscriptionMappingFromValue(v string) (*WebhooksSubscriptionMapping, error) {
	ev := WebhooksSubscriptionMapping(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhooksSubscriptionMapping: valid values are %v", v, AllowedWebhooksSubscriptionMappingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhooksSubscriptionMapping) IsValid() bool {
	for _, existing := range AllowedWebhooksSubscriptionMappingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhooksSubscriptionMapping value
func (v WebhooksSubscriptionMapping) Ptr() *WebhooksSubscriptionMapping {
	return &v
}

type NullableWebhooksSubscriptionMapping struct {
	value *WebhooksSubscriptionMapping
	isSet bool
}

func (v NullableWebhooksSubscriptionMapping) Get() *WebhooksSubscriptionMapping {
	return v.value
}

func (v *NullableWebhooksSubscriptionMapping) Set(val *WebhooksSubscriptionMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhooksSubscriptionMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhooksSubscriptionMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhooksSubscriptionMapping(val *WebhooksSubscriptionMapping) *NullableWebhooksSubscriptionMapping {
	return &NullableWebhooksSubscriptionMapping{value: val, isSet: true}
}

func (v NullableWebhooksSubscriptionMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhooksSubscriptionMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

