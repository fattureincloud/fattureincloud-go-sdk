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
)

// checks if the WebhooksSubscriptionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhooksSubscriptionConfig{}

// WebhooksSubscriptionConfig struct for WebhooksSubscriptionConfig
type WebhooksSubscriptionConfig struct {
Mapping *WebhooksSubscriptionMapping `json:"mapping,omitempty"`
}

// NewWebhooksSubscriptionConfig instantiates a new WebhooksSubscriptionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhooksSubscriptionConfig() *WebhooksSubscriptionConfig {
	this := WebhooksSubscriptionConfig{}
	return &this
}

// NewWebhooksSubscriptionConfigWithDefaults instantiates a new WebhooksSubscriptionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhooksSubscriptionConfigWithDefaults() *WebhooksSubscriptionConfig {
	this := WebhooksSubscriptionConfig{}
	return &this
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *WebhooksSubscriptionConfig) GetMapping() WebhooksSubscriptionMapping {
	if o == nil || IsNil(o.Mapping) {
		var ret WebhooksSubscriptionMapping
		return ret
	}
	return *o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksSubscriptionConfig) GetMappingOk() (*WebhooksSubscriptionMapping, bool) {
	if o == nil || IsNil(o.Mapping) {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *WebhooksSubscriptionConfig) HasMapping() bool {
	if o != nil && !IsNil(o.Mapping) {
		return true
	}

	return false
}

// SetMapping gets a reference to the given WebhooksSubscriptionMapping and assigns it to the Mapping field.
func (o *WebhooksSubscriptionConfig) SetMapping(v WebhooksSubscriptionMapping) *WebhooksSubscriptionConfig {
	o.Mapping = &v
        return o
}

func (o WebhooksSubscriptionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhooksSubscriptionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mapping) {
		toSerialize["mapping"] = o.Mapping
	}
	return toSerialize, nil
}

type NullableWebhooksSubscriptionConfig struct {
	value *WebhooksSubscriptionConfig
	isSet bool
}

func (v NullableWebhooksSubscriptionConfig) Get() *WebhooksSubscriptionConfig {
	return v.value
}

func (v *NullableWebhooksSubscriptionConfig) Set(val *WebhooksSubscriptionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhooksSubscriptionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhooksSubscriptionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhooksSubscriptionConfig(val *WebhooksSubscriptionConfig) *NullableWebhooksSubscriptionConfig {
	return &NullableWebhooksSubscriptionConfig{value: val, isSet: true}
}

func (v NullableWebhooksSubscriptionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhooksSubscriptionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


