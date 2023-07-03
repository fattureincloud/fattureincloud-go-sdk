/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.28
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the WebhooksSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhooksSubscription{}

// WebhooksSubscription struct for WebhooksSubscription
type WebhooksSubscription struct {
	// Webhooks subscription id
	Id NullableString `json:"id,omitempty"`
	// Webhooks callback uri.
	Sink NullableString `json:"sink,omitempty"`
	// [Read Only] True if the webhooks subscription has been verified.
	Verified NullableBool `json:"verified,omitempty"`
	// Webhooks events types.
	Types []EventType `json:"types,omitempty"`
}

// NewWebhooksSubscription instantiates a new WebhooksSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhooksSubscription() *WebhooksSubscription {
	this := WebhooksSubscription{}
	return &this
}

// NewWebhooksSubscriptionWithDefaults instantiates a new WebhooksSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhooksSubscriptionWithDefaults() *WebhooksSubscription {
	this := WebhooksSubscription{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksSubscription) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksSubscription) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *WebhooksSubscription) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *WebhooksSubscription) SetId(v string) *WebhooksSubscription {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *WebhooksSubscription) SetIdNil() *WebhooksSubscription {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *WebhooksSubscription) UnsetId() {
	o.Id.Unset()
}

// GetSink returns the Sink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksSubscription) GetSink() string {
	if o == nil || IsNil(o.Sink.Get()) {
		var ret string
		return ret
	}
	return *o.Sink.Get()
}

// GetSinkOk returns a tuple with the Sink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksSubscription) GetSinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sink.Get(), o.Sink.IsSet()
}

// HasSink returns a boolean if a field has been set.
func (o *WebhooksSubscription) HasSink() bool {
	if o != nil && o.Sink.IsSet() {
		return true
	}

	return false
}

// SetSink gets a reference to the given NullableString and assigns it to the Sink field.
func (o *WebhooksSubscription) SetSink(v string) *WebhooksSubscription {
	o.Sink.Set(&v)
	return o
}
// SetSinkNil sets the value for Sink to be an explicit nil
func (o *WebhooksSubscription) SetSinkNil() *WebhooksSubscription {
	o.Sink.Set(nil)
	return o
}

// UnsetSink ensures that no value is present for Sink, not even an explicit nil
func (o *WebhooksSubscription) UnsetSink() {
	o.Sink.Unset()
}

// GetVerified returns the Verified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksSubscription) GetVerified() bool {
	if o == nil || IsNil(o.Verified.Get()) {
		var ret bool
		return ret
	}
	return *o.Verified.Get()
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksSubscription) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verified.Get(), o.Verified.IsSet()
}

// HasVerified returns a boolean if a field has been set.
func (o *WebhooksSubscription) HasVerified() bool {
	if o != nil && o.Verified.IsSet() {
		return true
	}

	return false
}

// SetVerified gets a reference to the given NullableBool and assigns it to the Verified field.
func (o *WebhooksSubscription) SetVerified(v bool) *WebhooksSubscription {
	o.Verified.Set(&v)
	return o
}
// SetVerifiedNil sets the value for Verified to be an explicit nil
func (o *WebhooksSubscription) SetVerifiedNil() *WebhooksSubscription {
	o.Verified.Set(nil)
	return o
}

// UnsetVerified ensures that no value is present for Verified, not even an explicit nil
func (o *WebhooksSubscription) UnsetVerified() {
	o.Verified.Unset()
}

// GetTypes returns the Types field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksSubscription) GetTypes() []EventType {
	if o == nil {
		var ret []EventType
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksSubscription) GetTypesOk() ([]EventType, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *WebhooksSubscription) HasTypes() bool {
	if o != nil && IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []EventType and assigns it to the Types field.
func (o *WebhooksSubscription) SetTypes(v []EventType) *WebhooksSubscription {
	o.Types = v
	return o
}

func (o WebhooksSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhooksSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Sink.IsSet() {
		toSerialize["sink"] = o.Sink.Get()
	}
	if o.Verified.IsSet() {
		toSerialize["verified"] = o.Verified.Get()
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	return toSerialize, nil
}

type NullableWebhooksSubscription struct {
	value *WebhooksSubscription
	isSet bool
}

func (v NullableWebhooksSubscription) Get() *WebhooksSubscription {
	return v.value
}

func (v *NullableWebhooksSubscription) Set(val *WebhooksSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhooksSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhooksSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhooksSubscription(val *WebhooksSubscription) *NullableWebhooksSubscription {
	return &NullableWebhooksSubscription{value: val, isSet: true}
}

func (v NullableWebhooksSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhooksSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


