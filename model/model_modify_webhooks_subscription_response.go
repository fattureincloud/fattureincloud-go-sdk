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

// checks if the ModifyWebhooksSubscriptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyWebhooksSubscriptionResponse{}

// ModifyWebhooksSubscriptionResponse struct for ModifyWebhooksSubscriptionResponse
type ModifyWebhooksSubscriptionResponse struct {
Data *WebhooksSubscription `json:"data,omitempty"`
	// Webhooks registration warnings
Warnings []string `json:"warnings,omitempty"`
}

// NewModifyWebhooksSubscriptionResponse instantiates a new ModifyWebhooksSubscriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyWebhooksSubscriptionResponse() *ModifyWebhooksSubscriptionResponse {
	this := ModifyWebhooksSubscriptionResponse{}
	return &this
}

// NewModifyWebhooksSubscriptionResponseWithDefaults instantiates a new ModifyWebhooksSubscriptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyWebhooksSubscriptionResponseWithDefaults() *ModifyWebhooksSubscriptionResponse {
	this := ModifyWebhooksSubscriptionResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModifyWebhooksSubscriptionResponse) GetData() WebhooksSubscription {
	if o == nil || IsNil(o.Data) {
		var ret WebhooksSubscription
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyWebhooksSubscriptionResponse) GetDataOk() (*WebhooksSubscription, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModifyWebhooksSubscriptionResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given WebhooksSubscription and assigns it to the Data field.
func (o *ModifyWebhooksSubscriptionResponse) SetData(v WebhooksSubscription) *ModifyWebhooksSubscriptionResponse {
	o.Data = &v
        return o
}

// GetWarnings returns the Warnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModifyWebhooksSubscriptionResponse) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModifyWebhooksSubscriptionResponse) GetWarningsOk() ([]string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModifyWebhooksSubscriptionResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *ModifyWebhooksSubscriptionResponse) SetWarnings(v []string) *ModifyWebhooksSubscriptionResponse {
	o.Warnings = v
        return o
}

func (o ModifyWebhooksSubscriptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyWebhooksSubscriptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableModifyWebhooksSubscriptionResponse struct {
	value *ModifyWebhooksSubscriptionResponse
	isSet bool
}

func (v NullableModifyWebhooksSubscriptionResponse) Get() *ModifyWebhooksSubscriptionResponse {
	return v.value
}

func (v *NullableModifyWebhooksSubscriptionResponse) Set(val *ModifyWebhooksSubscriptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyWebhooksSubscriptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyWebhooksSubscriptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyWebhooksSubscriptionResponse(val *ModifyWebhooksSubscriptionResponse) *NullableModifyWebhooksSubscriptionResponse {
	return &NullableModifyWebhooksSubscriptionResponse{value: val, isSet: true}
}

func (v NullableModifyWebhooksSubscriptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyWebhooksSubscriptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


