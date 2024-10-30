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

// checks if the ModifyWebhooksSubscriptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyWebhooksSubscriptionRequest{}

// ModifyWebhooksSubscriptionRequest struct for ModifyWebhooksSubscriptionRequest
type ModifyWebhooksSubscriptionRequest struct {
Data *WebhooksSubscription `json:"data,omitempty"`
}

// NewModifyWebhooksSubscriptionRequest instantiates a new ModifyWebhooksSubscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyWebhooksSubscriptionRequest() *ModifyWebhooksSubscriptionRequest {
	this := ModifyWebhooksSubscriptionRequest{}
	return &this
}

// NewModifyWebhooksSubscriptionRequestWithDefaults instantiates a new ModifyWebhooksSubscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyWebhooksSubscriptionRequestWithDefaults() *ModifyWebhooksSubscriptionRequest {
	this := ModifyWebhooksSubscriptionRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModifyWebhooksSubscriptionRequest) GetData() WebhooksSubscription {
	if o == nil || IsNil(o.Data) {
		var ret WebhooksSubscription
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyWebhooksSubscriptionRequest) GetDataOk() (*WebhooksSubscription, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModifyWebhooksSubscriptionRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given WebhooksSubscription and assigns it to the Data field.
func (o *ModifyWebhooksSubscriptionRequest) SetData(v WebhooksSubscription) *ModifyWebhooksSubscriptionRequest {
	o.Data = &v
        return o
}

func (o ModifyWebhooksSubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyWebhooksSubscriptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableModifyWebhooksSubscriptionRequest struct {
	value *ModifyWebhooksSubscriptionRequest
	isSet bool
}

func (v NullableModifyWebhooksSubscriptionRequest) Get() *ModifyWebhooksSubscriptionRequest {
	return v.value
}

func (v *NullableModifyWebhooksSubscriptionRequest) Set(val *ModifyWebhooksSubscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyWebhooksSubscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyWebhooksSubscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyWebhooksSubscriptionRequest(val *ModifyWebhooksSubscriptionRequest) *NullableModifyWebhooksSubscriptionRequest {
	return &NullableModifyWebhooksSubscriptionRequest{value: val, isSet: true}
}

func (v NullableModifyWebhooksSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyWebhooksSubscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


