/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.2
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the CreateWebhooksSubscriptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWebhooksSubscriptionRequest{}

// CreateWebhooksSubscriptionRequest struct for CreateWebhooksSubscriptionRequest
type CreateWebhooksSubscriptionRequest struct {
	Data *WebhooksSubscription `json:"data,omitempty"`
}

// NewCreateWebhooksSubscriptionRequest instantiates a new CreateWebhooksSubscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebhooksSubscriptionRequest() *CreateWebhooksSubscriptionRequest {
	this := CreateWebhooksSubscriptionRequest{}
	return &this
}

// NewCreateWebhooksSubscriptionRequestWithDefaults instantiates a new CreateWebhooksSubscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebhooksSubscriptionRequestWithDefaults() *CreateWebhooksSubscriptionRequest {
	this := CreateWebhooksSubscriptionRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateWebhooksSubscriptionRequest) GetData() WebhooksSubscription {
	if o == nil || IsNil(o.Data) {
		var ret WebhooksSubscription
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhooksSubscriptionRequest) GetDataOk() (*WebhooksSubscription, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateWebhooksSubscriptionRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given WebhooksSubscription and assigns it to the Data field.
func (o *CreateWebhooksSubscriptionRequest) SetData(v WebhooksSubscription) *CreateWebhooksSubscriptionRequest {
	o.Data = &v
	return o
}

func (o CreateWebhooksSubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWebhooksSubscriptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateWebhooksSubscriptionRequest struct {
	value *CreateWebhooksSubscriptionRequest
	isSet bool
}

func (v NullableCreateWebhooksSubscriptionRequest) Get() *CreateWebhooksSubscriptionRequest {
	return v.value
}

func (v *NullableCreateWebhooksSubscriptionRequest) Set(val *CreateWebhooksSubscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebhooksSubscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebhooksSubscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebhooksSubscriptionRequest(val *CreateWebhooksSubscriptionRequest) *NullableCreateWebhooksSubscriptionRequest {
	return &NullableCreateWebhooksSubscriptionRequest{value: val, isSet: true}
}

func (v NullableCreateWebhooksSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebhooksSubscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


